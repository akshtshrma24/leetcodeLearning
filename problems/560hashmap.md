# 560 

real end game

```py
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        res = 0 
        cur_sum = 0
        prefix_sums = defaultdict(int)
        prefix_sums[0] = 1

        for n in nums: 
            cur_sum += n
            diff = cur_sum - k 
            res += prefix_sums[diff]        
            prefix_sums[cur_sum] += 1

        return res
```

this one is easy once u know the trick you just loop through it ocunting how many times youve seen this sum before, and if you get to a sunm of 0
that means you found k again so you increament it and add it to res

so you start off with the totallign n and if you have seen the difference before it will be atleast 1, otherwise it will be 0 so you can add it to res
