# 1539 

```py
class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        counter = 0
        for i in range(arr[-1]):
            if i not in arr:
                counter += 1
            if counter == k + 1:
                return i

        return arr[-1] + (k + 1 - counter)
```
easy 