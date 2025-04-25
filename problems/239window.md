# 239 

real end game 

```py
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        res = []
        q = collections.deque()
        l = r = 0

        while r < len(nums):
            while q and nums[q[-1]] < nums[r]:
                q.pop()
            q.append(r)

            if l > q[0]:
                q.popleft()

            if (r + 1) >= k: 
                res.append(nums[q[0]])
                l += 1
            r += 1

        return res
```

this one is easy actually 
all you do is a monotonic queu 
so while r is less than nums 
    you do a while loop when q is non empty and the one at the end is less than the nums right
    then you just append the index of r

    if the left most index is less than left then u obviously have to pop left

    if the window is greater than or requal to k you append the beginning (which should be taht windows max)
        increment l 
    then increment r

then just return it 