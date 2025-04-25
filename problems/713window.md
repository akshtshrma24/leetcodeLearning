# 713 

real end game 

```py

class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        res = 0
        l = 0
        product = 1
        for r in range(len(nums)):
            product *= nums[r]
            while l <= r and product >= k: 
                product = product // nums[l]
                l += 1
            res += (r - l + 1)

        return res

```

easy once you know how it works, just bring the windows left up by 1 everytime the 
the main trick is 
            res += (r - l + 1)
once you know that its easy 
so lets say our sub array is 
