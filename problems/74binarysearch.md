# 74 

```py
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        r, c = len(matrix), len(matrix[0])
        top, bot = 0, r - 1
        while top <= bot:
            row = (top + bot) // 2
            if target > matrix[row][-1]:
                top = row + 1
            elif target < matrix[row][0]:
                bot = row - 1
            else: break
        if not top <= bot: return False
        row = (top + bot) // 2
        l,r = 0, len(matrix[0]) - 1
        while(l <= r):
            m = (l + r) // 2
            if matrix[row][m] < target:
                l = m + 1
            elif matrix[row][m] > target:
                r = m - 1
            else:
                return True
        return False
```

This is just binary saerch but with another dimension, just search the ends, of it.
just 2 binary searches 