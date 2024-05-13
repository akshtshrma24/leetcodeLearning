# 200

```py
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        # if its null just return 
        if not grid: return 0
        # get the rows and cols
        rows, cols = len(grid), len(grid[0])
        # keep track of the ones that you visited
        visit = set()
        islands = 0
        # initialize islands 
        # breadth first search so like you are using a queue
        def bfs(r, c): 
            # get the queue
            q = collections.deque()
            # add the row and column to visited
            visit.add((r, c))
            # add it to the queue
            q.append((r,c))
            while q:
                # while the queeu is not empty pop
                row, col = q.popleft()
                # directions you could go through
                directions = [[1, 0], [-1,0], [0,1], [0,-1]]
                # for x y in directions
                for dr, dc in directions:
                    # get updated r c
                    r, c = row + dr, col + dc 
                    # check if in bounds, check if 1 and check if not in visited
                    # if it isnt, add it to visited and append it to queue
                    if (r in range(rows) and
                        c in range(cols) and 
                        grid[r][c] == "1" and
                        (r, c) not in visit):
                        q.append((r, c))
                        visit.add((r, c))
            # just for the remaining ones pop it
            while q:
                row, col = q.popleft()
        # run bfs, increment islands
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == "1" and (r,c) not in visit:
                    bfs(r, c)
                    islands += 1
        return islands
```

Just a loop with bfs attached to it, bfs uses queue.
