# 62 

```py
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        row = [1] * n 
        for i in range(m - 1):
            # initialize a new row 1 will always be on the right side.
            new_row = [1] * n
            for j in range(n - 2, -1, -1): # this means 2nd to last pos
            # until the end, reverse order
                # new_row[j] is the element to the right [j + 1] + row[j]
                new_row[j] = new_row[j + 1] + row[j]
            # set the row to the new row
            row = new_row
        return row[0]
```

This sounds like a dfs/backtracking one, we just need to go through all of the possible combinations 
but rather than calculating for each one check if [-1] or [][-1] has been calculated 
That was right for a dfs cacheing solution but there is also a 2dp solution 
the 2dp solution is that since the bot can only go right and down the right most coloumn and bottom most
column will always be 1 since no more moves, so then you just look at the right element and below element
add them together and thats how many moves you have you keep doing that 