# 1047 

real end game 

```py
class Solution:
    def removeDuplicates(self, s: str) -> str:
        stack = []

        for c in s: 
            if stack and c == stack[-1]:
                stack.pop()
            else: stack.append(c)
        
        return ''.join(stack)
```

really easy just a stack