# 48 

```py
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        l, r = 0, len(matrix) - 1
        while l < r:
            for i in range(r - l):
                top, bottom = l,r
                
                topLeft = matrix[top][l + i]

                # this one will go from the btotom curr into top curr
                matrix[top][l + i] = matrix[bottom - i][l]

                #this one moves bottom right into bottom curr
                matrix[bottom - i][l] = matrix[bottom][r - i]

                # moves from top right to bottom right
                matrix[bottom][r - i] = matrix[top + i][r]

                matrix[top + i][r] = topLeft
            r -= 1
            l += 1
```

In this its super simple it just loops through rows and shifts them around its really easy