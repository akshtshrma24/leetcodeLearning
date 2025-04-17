# 71

real end game


```py
class Solution:
    def simplifyPath(self, path: str) -> str:
        splat = path.split('/')

        res = ""
        s = []

        for diri in splat:
            if diri == '..' and len(s) > 0: 
                print(s)
                s.pop(-1)
                continue
            if diri == '..' and len(s) == 0:
                s.append('')
                continue
            elif diri == '.' or diri == '':
                continue
            s.append(diri)

        print(s, 'here')

        for diri in s:
            if diri == '':
                continue
            res += f"/{diri}"
        
        if res == '':
            return "/"
        
        return res
```

this one is also quite easy, just make a stack and use that stack to pop the elements back when you get the ../
and have a nother if statemnt for the if case that .. is first 
