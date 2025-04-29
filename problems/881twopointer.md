# 881 

real end game

```py
class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        
        boats = 0

        l, r = 0, len(people) - 1
        while l <= r: 
            remaining = limit - people[r]
            r -= 1
            boats += 1
            if l <= r and remaining >= people[l]:
                l += 1
        return boats
```
sort the peoples
initialise 2 points while l is smaller than right
check the remaining you do this to see if you can fit more than 1 person on the boat, if the remaining is greater than the left persons weight
put him in the boat and increase left by 1 
just increament boats then 

