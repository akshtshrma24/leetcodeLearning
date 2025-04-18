# 791 

real end game 

```py
class Solution:
    def customSortString(self, order: str, s: str) -> str:
        m = dict(Counter(s))

        res = ""

        for o in order:
            limit = m.get(o, 0)
            for i in range(limit):
                res += o
            if limit != 0:
                del m[o]
        
        for k, v in m.items():
            for i in range(v):
                res += k

        return res
```

easy 
