# 5 

```py
class Solution:
    def longestPalindrome(self, s: str) -> str:
        res = ""
        res_len = 0
        for i in range(len(s)):
            # for the odd ones
            l, r = i, i
            while(l >= 0 and r < len(s) and s[l] == s[r]):
                if((r - l + 1) > res_len):
                    res = s[l: r + 1]
                    res_len = r - l + 1
                r += 1
                l -= 1

            # even length is 

            l, r = i, i + 1
            while(l >= 0 and r < len(s) and s[l] == s[r]):
                if((r - l + 1) > res_len):
                    res = s[l: r + 1]
                    res_len = r - l + 1
                r += 1
                l -= 1

        return res
```
kinda jsut like the O(n^3)  but you start from the middle instead and expand out 
so you start from 0 then you increase the windo til 1 of them is outta bounds
then you start at 1
then 2
and so on