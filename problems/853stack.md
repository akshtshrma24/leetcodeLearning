# 853

```py
class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        tuples = []
        fleet = 0
        arrival_times = []
        for i in range(len(position)):
            tuples.append((position[i], speed[i]))
        tuples = sorted(tuples, key=lambda x: (-x[0], x[1]))
        for i in tuples:
            arrival = (target - i[0]) / i[1]
            arrival_times.append(arrival)

        bottle_neck = -1
        for i in range(len(arrival_times)):
            bottle_neck = max(arrival_times[i], bottle_neck)
            arrival_times[i] = bottle_neck

        while arrival_times:
            popped = arrival_times.pop()
            fleet += 1
            if(arrival_times and popped <= arrival_times[-1]):
                while(arrival_times and popped <= arrival_times[-1]):
                    popped = arrival_times.pop()
        return fleet
```
this one was kinda tricky, you have to get htem as tuples, sort by descending order, calculate the 
arrival time, then apply th ebottle neck as so
1, 1, 7, 3, 12, 
becomes
1, 1, 7, 7, 12
because 3 cant pass 7 so it gets put as 7 
then you pop from stack 
if popped > peek value fleet + 1
else loop through til its not,  