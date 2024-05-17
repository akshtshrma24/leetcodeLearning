# 190

```py
class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        for i in range(32):
            bit = (n >> i) & 1
            res = res | (bit << (31 - i))
        return res
```
This was really easy, just shift it by i because you want to keep shifting it more and more
til you get to 32 then and it by 1 to actually get the bit otherwise you just get the number
then just or it to the res by the amount of bits left it should be 
