# 130 

```py
class Solution:
    def solve(self, board: List[List[int]]) -> List[List[int]]:
        ROWS, COLS = len(board), len(board[0])

        def dfs(r, c):
            if(r < 0 or r == ROWS or
                c < 0 or c == COLS or
                board[r][c] != 'O'): 
                return
            board[r][c] = 'T'
            dfs(r,c + 1)
            dfs(r,c - 1)
            dfs(r + 1,c)
            dfs(r - 1,c)


        for r in range(0, ROWS):
            if(board[r][0] == 'O'):
                dfs(r, 0)
            if(board[r][COLS-1] == 'O'):
                dfs(r, COLS - 1)

        for c in range(0, COLS):
            if(board[0][c] == 'O'):
                dfs(0, c)
            if(board[ROWS - 1][c] == 'O'):
                dfs(ROWS - 1, c)
        
        for r in range(ROWS):
            for c in range(COLS):
                if(board[r][c] == 'O'): board[r][c] = 'X'
        for r in range(ROWS):
            for c in range(COLS):
                if(board[r][c] == 'T'): board[r][c] = 'O'
```

This seems just like the rotting oranges one
It was actually like the pacific atlantic one, I need to also think of things as a reverse way 
