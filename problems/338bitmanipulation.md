# 338

```py
class Solution:
    def countBits(self, n: int) -> List[int]:
        res = []
        for i in range(n + 1):
            count = 0
            temp = i
            while temp:
                temp &= (temp - 1)
                count += 1
            res.append(count)
        return res
```

This was super easy just use the way to count 1s and count it in a loop
