# 134 

real end game

```py
class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        if sum(gas) < sum(cost):
            return -1 
        
        total = 0
        res = 0
        start = 0

        for i in range(len(gas)):
            total += (gas[i] - cost[i])

            if total < 0: 
                total = 0
                start = i + 1
        
        return start
```

this one is easy 
first verify a solution exists you do that by making suure sum of gas is less than the cost

then you just go through the list
and if total ever goes below - reset total and make start = i + 1

then just return start

