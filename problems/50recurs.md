# 50 

real end game 
```py
class Solution:
    def myPow(self, x: float, n: int) -> float:
        
        def calc_power(x, n):
            if x == 0:
                return 0
            if n == 0:
                return 1

            res = calc_power(x, n // 2)
            res *= res

            if n % 2 == 1:
                return res * x
            
            return res
        
        ans = calc_power(x, abs(n))

        if n >= 0:
            return ans

        return 1 / ans
```

this one is kind of simple to follow 
x is just the number obviously if its 0 just return 0 
if n is 1 return 1 because you are just going to multiple it by 1 

then recurse down to it because you are breaking the problem down by a low.
2^10 becomes 2^5 * 2^5 and so one and so on, so just keep breaking it doen
if n % 2 == 1 meaning if its odd then return the res * x because you cant break it down more 

then just retunr it 

and finally just make sure its posiitive and return it 
otehrwise return the inverse