# 528 

real end game 

```py
class Solution:

    def __init__(self, weights: List[int]):
        total = 0
        self.arr = []
        for w in weights: 
            total += w
            self.arr.append(total)

    def pickIndex(self) -> int:
        target = random.uniform(0, self.arr[-1])
        l, r = 0, len(self.arr)

        while l < r: 
            m = (l + r) // 2

            if(self.arr[m] > target):
                r = m
            elif(self.arr[m] < target):
                l = m + 1
            else:
                l = m
        
        return l
# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()
```
this one is kind of hard to remmeber the binary search if we find it we want to continue looking left to see fi theres any other one that 
is less that we can find 
so do that