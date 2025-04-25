# 34 

Bin search 

```py
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left = self.binSearch(nums, target, True)        
        right = self.binSearch(nums, target, False)        
        return [left, right]

    def binSearch(self, nums, target, leftBias):
        l, r = 0, len(nums) - 1
        i = -1
        while l <= r: 
            m = (l + r) // 2
            if target > nums[m]:
                l = m + 1
            elif target < nums[m]:
                r = m - 1
            else:
                i = m
                if leftBias: 
                    r = m - 1
                else:
                    l = m + 1
        return i
```


this is normal binary serach, unless I find it if i find it then in that case i need to check if there is a left bias, if there is 
search the left side 
if not search the right side 

return the index
