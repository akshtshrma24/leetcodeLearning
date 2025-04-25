# 1209 

real end game

```py
class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = []
        
        for c in s: 
            if stack and stack[-1][0] == c:
                stack[-1][1] += 1
            else:
                stack.append([c, 1])

            if stack[-1][1] >= k: 
                stack.pop()
        
        res = ''

        for c, cnt in stack: 
            res += c * cnt

        return res
```
i knew it was stack was just confused on how to keep track of the k values 

literally just keep a tuple of the character, count and if the top of the stack is the character add 1
otherwise just append the character, 1

then check if the top of the stack is greater than or equal to k 

return res