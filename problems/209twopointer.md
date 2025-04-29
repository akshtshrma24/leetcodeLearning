# 209 

real end game

```py
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l, total = 0, 0
        res = float("inf")

        for r in range(len(nums)):
            total += nums[r]
            while total >= target: 
                res = min(r - l + 1, res)
                total -= nums[l]
                l += 1
                
        return 0 if res == float("inf") else res
```

standard sliding window
you increment r every time increment l only when the total >= target and remove total -= nums[l] because l is no longer in window
you do this because the its just a window bro u know
