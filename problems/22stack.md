# 22 

```py
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        def dfs(left, right, s):
            if(len(s) == n * 2):
               res.append(s)
               return
            if(left < n):
                dfs(left + 1, right, s + '(') 
            if(right < left):
                dfs(left, right + 1, s + ')')
        dfs(0,0,'')
        return res
            
```
this one is also easy, i actually got this solution first try on paper
just need to remember how to implement dfs but its getting easier