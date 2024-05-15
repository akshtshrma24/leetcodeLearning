# 417

```py
class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        ROWS, COLS = len(heights), len(heights[0])
        pac, atl = set(), set()

        def dfs(r,c, prev_height, ocean):
            if((r,c) in ocean or r < 0 or c < 0 or r == ROWS or c == COLS
            or prev_height > heights[r][c]):
                return
            ocean.add((r,c))
            # run for each direction
            dfs(r - 1,c, heights[r][c], ocean)
            dfs(r + 1,c, heights[r][c], ocean)
            dfs(r,c + 1, heights[r][c], ocean)
            dfs(r,c - 1, heights[r][c], ocean)

        for c in range(0, COLS):
            dfs(0, c, heights[0][c], pac)
            dfs(ROWS - 1, c, heights[ROWS - 1][c], atl)
        
        for r in range(0, ROWS):
            dfs(r, 0, heights[r][0], pac)
            dfs(r, COLS - 1, heights[r][COLS - 1], atl)
        
        res = []
        for r in range(ROWS):
            for c in range(COLS):
                if((r,c) in pac and (r,c) in atl):
                    res.append([r,c])
        return res
```
looks like bfs we just check if it touches both sides pop from queue make sure touches left right/topbot
it is not bfs it is dfs, we do from the edges and go as far as we can, and add the ones that we can reach to 
a set for pac and atl, if we travers from the edges and we see that it is in both atl and pac, then we can 
safely say that it is what we are looking for, 
