# 680 

real end game

```py
class Solution:
    def validPalindrome(self, s: str) -> bool:

        def palindrome(s):
            l, r = 0, len(s) - 1
            while l < r: 
                if s[l] != s[r]:
                    return (l,r)
                l += 1
                r -= 1
            return True

        def remove_element(s, index):
            s = list(s)
            s.pop(index)
            return "".join(s)

        temp = palindrome(s)
        if(temp != True):
            l = temp[0]
            r = temp[1]
            left = remove_element(s, l)
            right = remove_element(s, r)

            if(palindrome(left) != True and palindrome(right) != True):
                return False
        
        return True

```
really easy 