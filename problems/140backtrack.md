# 140 

real end game

```py
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        cur = []
        res = []
        
        def backtrack(i):
            if i == len(s):
                res.append(" ".join(cur))
                return
            for j in range(i, len(s)):
                w = s[i: j + 1]
                if w in wordDict:
                    cur.append(w)
                    backtrack(j + 1)
                    cur.pop()

        backtrack(0)
        return res
```
this is just backtracking with the words and checkin f its in the dictionary thats it
