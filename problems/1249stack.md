# 1249 

real end game 

```py
class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        res = []

        cnt = 0 # extra ( parantheses 
        
        for c in s: 
            if c == '(':
                cnt += 1
                res.append(c)
            elif c == ')' and cnt > 0:
                cnt -= 1
                res.append(c)
            elif(c != ')'):
                res.append(c)

        print(''.join(res))
        filtered = []

        for c in res[::-1]:
            if c == '(' and cnt > 0:
                cnt -= 1
            else:
                filtered.append(c)

        return ''.join(filtered[::-1])
```

increase the count if it opens decrease if it is ")" and can add it only if the count is greater than 0 if it is 0 uk its bad 
then you go through it one more time because it could also have a count greater than 0 and add ( so then you decrease count and dont add it 
then just return filtered

increase the count if its opening on loop 1 and add it to res, decrease it only if its a closing and count is greater than 0 
this would mean there has been a + 1 and a - 1 would not result in negative
otherwise just add it bc its a letter

then loop backwards and if its an opening and count is greater than 0 do not add it and decrease count, this would only happen if there 
are excess (. 
when you are looping backwards and the count is greater than 0 it means the closing brackets were not enough to decrease count down to 0
