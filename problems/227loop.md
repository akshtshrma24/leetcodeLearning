# 227 

real end game 

```py
class Solution:
    def calculate(self, s: str) -> int:
        i = 0
        curr = prev = res = 0
        curr_operation = '+'

        while i < len(s):
            curr_char = s[i]
            if curr_char.isdigit():
                while i < len(s) and s[i].isdigit():
                    curr = curr * 10 + int(s[i])
                    i += 1
                i -= 1

                if curr_operation == '+':
                    res += curr
                    prev = curr
                elif curr_operation == '-':
                    res -= curr
                    prev = -curr
                elif curr_operation == '*':
                    res -= prev
                    res += prev * curr

                    prev *= curr
                else:
                    res -= prev
                    res += int(prev / curr)
                    
                    prev = int(prev / curr)
                curr = 0
            elif curr_char != ' ':
                curr_operation = curr_char
            i += 1
        return res
```
this one is not that easy, teh whole idea behind it is if its addition, you store the current positive value as previous and if its negative you set 
the value as the negaitve verison 

and when its multiplication you remove the previous value from the result becuase you are basically undoing that operation, then you multiply the previous to the current 
then you set previous as the previous times current, same for divisions