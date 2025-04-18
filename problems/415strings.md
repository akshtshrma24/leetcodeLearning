# 415 

real end game 

```py
class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        res = []
        carry = 0
        i, j = len(num1) - 1, len(num2) - 1
        ordzero = ord('0')
        while i >= 0 or j >= 0 or carry: 
            d1 = ord(num1[i]) - ordzero if i >= 0 else 0
            d2 = ord(num2[j]) - ordzero if j >= 0 else 0

            total = d1 + d2 + carry

            carry = total // 10
            print(total % 10)
            res.append(chr(total % 10 + ordzero))
            i -= 1
            j -= 1
        return ''.join(res[::-1])
```

just get the integer by subtarcting the char by ord zero if you can
add the two numbers together 
and you get the carry by just dividing by 10 kinda simple
then add the modded 10 because you dont want to add 2 integets