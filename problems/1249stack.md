# 1249 

real end game 

```py
class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        res = []

        cnt = 0 # extra ( parantheses 
        
        for c in s: 
            if c == '(':
                res.append(c)
                cnt += 1
            elif c == ')' and cnt > 0:
                res.append(c)
                cnt -= 1
            elif c != ')':
                res.append(c)
        
        filtered = []

        for c in res[::-1]:
            if(c == '(' and cnt > 0):
                cnt -= 1
            else:
                filtered.append(c)
        
        return "".join(filtered[::-1])
```

increase the count if it opens decrease if it is ")" and can add it only if the count is greater than 0 if it is 0 uk its bad 
then you go through it one more time because it could also have a count greater than 0 and add ( so then you decrease count and dont add it 
then just return filtered