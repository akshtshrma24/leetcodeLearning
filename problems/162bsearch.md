# 162 

real end game

```py
class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        l, r = 0, len(nums) - 1

        while l < r:
            m = (l + r) // 2

            if(nums[m] < nums[m + 1]):
                l = m + 1
            else:
                r = m 
        
        return l
```
really easy just search for the peak like this any time youre searching for an element in an array its most likely binary search
