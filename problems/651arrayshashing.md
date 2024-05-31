# 271 


```py
class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            res += f"{str(len(s))}#{s}"
        return res

    def decode(self, s: str) -> List[str]: 
        res = [] 
        i = 0
        while i < len(s):
            j = i
            while(s[j] != "#"): j += 1
            print(s[i:j])
            length = int(s[i:j])
            i = j + 1
            j = i + length
            print(j, 'j')
            print(i,j)
            res.append(s[i:j])
            i = j
        return res
```

use # and lenght of string to encode
then decode you can just use the lenth and the # to note when 
