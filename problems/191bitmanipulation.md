# 191 

```py
class Solution:
    def hammingWeight(self, n: int) -> int:
        res = 0
        while n:
            n &= (n - 1)
            res += 1
        return res     
```
This works y runnign the loop for as many 1s there are in the bits 
if you remove a bit then & it by itself you effectively are just removing 
only one 1 bit. do that for as many times as you can while n is greater than 0 