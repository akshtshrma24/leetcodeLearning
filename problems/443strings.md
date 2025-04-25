# 443 

real end game 

```py
class Solution:
    def compress(self, chars: List[str]) -> int:
        r = i = 0

        while i < len(chars):
            curr_char = chars[i]
            curr_occ = 0

            while i < len(chars) and chars[i] == curr_char:
                curr_occ += 1
                i += 1
            chars[r] = curr_char
            r += 1
            if(curr_occ > 1):
                curr_occ_str = str(curr_occ)
                for j in range(len(curr_occ_str)):
                    chars[r] = curr_occ_str[j]
                    r += 1
        return r
```

its not that hard, literally just count the occurences put it right after, 
and return r bc it gets dropped off right after the last character
