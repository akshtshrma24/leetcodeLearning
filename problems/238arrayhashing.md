# 238 

```py
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [1] * len(nums)
        prefix = 1
        for i in range(len(nums)):
            res[i] = prefix
            prefix *= nums[i]
        postfix = 1
        for i in range(len(nums) - 1, -1, -1):
            res[i] *= postfix
            postfix *= nums[i]
        return res
```
3 arrays total
prefix, with all the multiplication before it 
postfix, with all multiplication after 
original array 
1, 2, 3, 4
solution: 24, 12, 8, 6
prefix array: 1 2 6, 24
postfi array: 24, 12, 4, 1
then you cross them so 
for the first position 
24, then you take 1 * 12 
then you take 2 * 4, then you take 6 * 1
