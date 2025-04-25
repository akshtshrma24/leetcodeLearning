# 76 

real end game 

```py
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not s or not t or len(t) > len(s):
            return ""
        
        t_counts = Counter(t)

        l = r = matches = 0
        required = len(t_counts)

        window_counts = defaultdict(int)

        ans = (float('inf'), 0, 0)

        while r < len(s):
            cur_char = s[r]
            window_counts[cur_char] += 1

            if cur_char in t_counts and t_counts[cur_char] == window_counts[cur_char]:
                matches += 1
            
            while l <= r and matches == required:
                to_remove = s[l]

                if r - l + 1 < ans[0]:
                    ans = (r - l + 1, l, r)
                
                window_counts[to_remove] -= 1

                if to_remove in t_counts and t_counts[to_remove] > window_counts[to_remove]:
                    matches -= 1
                
                l += 1
            
            r += 1
        
        return s[ans[1]: ans[2] + 1] if ans[0] != float('inf') else ""
```

this one is long but the principle is simple
have 2 dictionaries 
1 dictionary is the values in s and 1 is t
if the values that are in t are the same in s, check for minimumness, remove the left one, and check again. 
super simple low key