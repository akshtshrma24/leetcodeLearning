# 26 

real end game

```py
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        l = 1

        for r in range(1, len(nums)):
            if nums[r] != nums[r - 1]:
                nums[l] = nums[r]
                l += 1
        
        return l
```

this one is also really easy once u know the trick,

loop r through the array 
if nums of r is not the same as the nums of r before it, you set the number to nums [l]
and incrememnt nums of l 

[0,1,1,2,2]
if it is this 
l starts by pointing at first 1 so does r
then r sees its the same so im going to keep going and im going to leave l as the same 
it keeps going til r is at index 3, 
then it sees 2 is not same as 1 the one before it so what it does is 
[0,2,1,2,2]
     l 

left stays at this indexed and you just return that as it basically counts how many unique ones there are 