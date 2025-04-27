# 895 

real end game

```py
class FreqStack:

    def __init__(self):
        self.cnt = {}
        self.maxCnt = 0
        self.stack = {}

    def push(self, val: int) -> None:
        valCnt = 1 + self.cnt.get(val, 0)
        self.cnt[val] = valCnt
        if valCnt > self.maxCnt:
            self.maxCnt = valCnt
            self.stack[valCnt] = []
        self.stack[valCnt].append(val)

    def pop(self) -> int:
        res = self.stack[self.maxCnt].pop()
        self.cnt[res] -= 1

        if self.stack[self.maxCnt] == []:
            self.maxCnt -= 1
        
        return res
        


# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()
```

stupid ass question for real
we keep 2 maps
1 is of the count of the elemnt and 1 of the stack
we basically put the value in its current amount used stack

if push 5 and then push a 5 again 
we have 2 stacks
1: [5]
2: [5]

and if we pop we see that our maxCount is 2 so we pop from 2 
if the 2s stack is empty it means our maxCount should be decremented so we decrement the maxCount, now when we pop again
we pop the maxCount and the end of that because we go in order so we pop the 1