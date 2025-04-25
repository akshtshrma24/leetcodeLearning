# 15 

real end game 

```py
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        res = []

        nums.sort()

        for i, a in enumerate(nums):
            if i > 0 and a == nums[i - 1]:
                continue
            
           
            l = i + 1
            r = len(nums) - 1
            while l < r:
                three_sum = nums[l] + nums[r] + a

                if three_sum > 0: 
                    r -= 1
                elif three_sum < 0: 
                    l += 1
                else:
                    res.append([nums[l], nums[r], a])
                    l += 1
                    while nums[l] == nums[l - 1] and l < r:
                        l += 1
        return res
```

so first waht you have to do is sort it 