# 128 

```py
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) == 0: return 0
        nums = set(nums)
        res = -1
        for i in nums: 
            if(i - 1 not in nums): 
                temp = []
                j = i
                while(j in nums):
                    temp.append(j)
                    j += 1
                res = max(len(temp), res)
        return res
```

This one is really easy, just check if n - 1 is in nums
if it is means its not start of sequence, so skip that 
then when it is loop through +1 til its not in set 
return longest