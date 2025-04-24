# 53 

real end game 

```py
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = float('-inf')
        cur_sum = -1
        for n in nums: 
            if cur_sum < 0: 
                cur_sum = 0
            cur_sum += n
            max_sum = max(cur_sum, max_sum)
        return max_sum

```