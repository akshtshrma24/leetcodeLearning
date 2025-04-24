# 1004 

real end game

```py
class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        res = 0
        l = r = 0
        counter = 0

        while r < len(nums):
            if nums[r] == 0:
                counter += 1

            while(counter > k):
                if(nums[l] == 0):
                    counter -= 1
                l += 1

            if counter <= k: 
                res = max(res, r - l + 1)

            r += 1
        return res
```


realy easy just loop through it, create a window where the counter <= k and if counter is less than k just update res

while its bigger update left side of teh window decrementing counter for any 0s