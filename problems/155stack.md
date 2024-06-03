# 155 

```py
import sys
class MinStack:
    min_num = []
    stack = []
    def __init__(self):
        self.min_num = []
        self.stack = []

    def push(self, val: int) -> None:
        if(not self.min_num): temp = val
        else:temp = min(self.min_num[0], val)
        self.min_num.insert(0, temp)
        self.stack.insert(0, val)

    def pop(self) -> None:
        self.min_num = self.min_num[1:]
        self.stack = self.stack[1:]

    def top(self) -> int:
        return self.stack[0]

    def getMin(self) -> int:
        return self.min_num[0]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
```
really easy, all you have to do is make a stack
but for the min, just push the minimum value that way it works like this 
[3,2,1,1,1] minstack
[3,2,1,2,2] stack
that way when you are at 2 in the stack 1 is still the minium but when you pop 1
2 is the minum