# 136 

```py
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for i in nums:
            res = res ^ i
        return res
```
unintuitive but this is how it works
4
2
2
if you XOR the bits of this 
100
010
010
you will get this 
100 ^ 010 = 100
100 ^ 010 = 100
any which way you go it is same