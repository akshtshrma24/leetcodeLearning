# 84 

real end game


```py
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        maxArea = 0 
        stack = []  

        for i, h in enumerate(heights):
            start = i
            while stack and stack[-1][1] > h: 
                index, height = stack.pop()
                maxArea = max(maxArea, height * (i - index))
                start = index
            stack.append((start, h))
        
        for i, h in stack:
            maxArea = max(maxArea, h * (len(heights) - i))
        return maxArea
```

for this one you keep a stack of heights 

you keep track of the indexes and the heights, 
while the current value is less than the top of the stack, 
you want to compute the area, by taking the current index, and multiplying the top of the stacks index
then you keep that index and append that to stack

otherwise u just append it to the stack

then u have to go through it once more because what if you havent got the max year
this one you just take the height and multiply it by the length of the heights minus i because if we have 1, 3 right the 3 can be at most 3 
but the 1 can be 2 because it can take a section of the 2, so it would be h which is 3 * 2 - 1 which is 3 * 1, then when u get to 1
it will be 1 * (2 - 0) which is 2 so this one would be 3 but then u want to compare it to max area in case it was already found