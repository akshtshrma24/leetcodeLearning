# 994

```py
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rotten_queue = deque()
        time, fresh = 0,0
        r,c = len(grid), len(grid[0])
        for row in range(r):
            for col in range(c):
                if(grid[row][col] == 1):
                    fresh += 1
                elif(grid[row][col] == 2):
                    rotten_queue.append([row,col])
        directions = [[1,0], [-1,0], [0,1], [0,-1]]
        while rotten_queue and fresh > 0:
            for i in range(len(rotten_queue)):
                r,c = rotten_queue.popleft()
                for dr, dc in directions:
                    row, col, = dr + r, dc + c
                    if(row < 0 or row == len(grid) or
                        col < 0 or col == len(grid[0]) or
                        grid[row][col] != 1
                        ): continue
                    print(row,col)
                    fresh -=1
                    grid[row][col] = 2
                    rotten_queue.append([row,col])
            time += 1
        return time if fresh == 0 else -1

```

I think this is just going to be bfs, but we add a visited one, and then we just check if its visited say its rotten
but so every 1 away the furthest orange is from the beginnign rotten one, should be the one
run multi sourced bfs
multi sourced bfs is just like bfs but you dont do it recursively, you use a loop
in the leet code problem its first getting the places where its rotten adding it to the queue, then it 
has another while loop that just makes sure that rotten queue and fresh are above 0 then it will go in the range
of the rotten queue then it will get the the rotten cord then it will go in the directions and mark the ones next to it
as rotten then it will keep going for all of them. time += 1 over there bc it does all the rotten queeue then it finishes but then
more stuff is added to the queue so it adds the time