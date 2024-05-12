# 90

```py
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()
        def backtrack(i, subset):
            if(i == len(nums)): 
                # gone through entire array
                res.append(subset[::]) # you take a copy of it because
                # when you add it to the res you dont want to overwrite the 
                # previous entry you still want it there 
                return
            #  All subsets that contain nums[i]
            subset.append(nums[i])
            backtrack(i + 1, subset)
            subset.pop()

            # ALl subsets that dont contain nums[i]
            while(i + 1 < len(nums) and nums[i] == nums[i + 1]): 
                # while there are duplicates
                i += 1
            backtrack(i + 1, subset)
        backtrack(0, [])
        return res
```
I think this is exactly like 46, just adding it to a set rather than an array so you dont do duplicares
No its not, its kinda like it but you eliminate the duplicates so you dont expand on every single branch
but rather kinda every other branch every branch other than duplicates so a little smaller than 
n*2^n but worst case still that bc it could not have any duplicates