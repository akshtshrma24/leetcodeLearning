# 523 

real end game

```py
class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        remainder = { 0: -1 }
        total = 0

        for i, n in enumerate(nums):
            total += n
            r = total % k

            if r not in remainder: 
                remainder[r] = i
            elif i - remainder[r] > 1:
                return True
        return False
```

this is easy once you know the trick, the trick is that if you find the remainder again, that means something had to be added to the sum that was divisible 
by k

so just loop through adding and then u also want to make sure the difference is greater than 1 bc you need 2 elements in the array 
