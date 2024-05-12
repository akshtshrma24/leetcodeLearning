# 46 

```py
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        def dfs(curr, used):
            if len(curr) >= len(nums):
                # this means that current one is greater than lenght of nums so you have to stop
                #  otherwise you go forever
                res.append(curr.copy())
                return

            for i, letter in enumerate(nums):
                # skip used letters
                if used[i]: continue
                # add letter to permutation, mark letter as used
                curr.append(letter)
                used[i] = True
                dfs(curr, used)
                # remove letter from permutation, mark letter as unused
                curr.pop()
                used[i] = False 
        dfs([], [False] * len(nums))
        return res
```

so this one it uses dfs because this
[1,2,3]
if you chose a 1 first then next one can only be a 2 or a 3 hten if you pick 2 
next one can only be a 3 vice versa. thats where you get the dfs. 