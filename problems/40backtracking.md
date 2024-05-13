# 40 

```py
class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        # sort the candidates so that we can loop through out of the 
        res = []
        # this is dfs
        def backtrack(cur, pos, target):
            # if the target is 8 then you can append the copy
            if target == 0:
                res.append(cur.copy())
            # but if it is less than 0 means we went over exit out
            if target <= 0: return
            prev = -1
            # prev should be set to -1 because we need to initialise it to something
            for i in range(pos, len(candidates)):
                # if duplicate continue
                if(candidates[i] == prev): continue
                # after duplicate then we can append to current and then backtrack 
                cur.append(candidates[i])
                backtrack(cur, i + 1, target - candidates[i])
                cur.pop()
                prev = candidates[i]
        backtrack([], 0, target)
        return res
```
I think this will just be like the one that we had to get the permutations except duplicates
but we check where it is equal to the target and if it is stop recursing on the branch
like dfs
