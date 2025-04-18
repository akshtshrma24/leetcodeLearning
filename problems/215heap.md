# 215 

real end game

```py

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        n = [-s for s in nums]
        print(n)

        heapq.heapify(n)
        p = n[0]
        for i in range(k):
            p = heapq.heappop(n)
        
        return -p

```

kth largest, just negate them, use a heap to do get them in O(n)

then pop from beginning to get the largest