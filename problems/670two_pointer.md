# 670

real end game

```py
class Solution:
    def maximumSwap(self, num: int) -> int:
        num = list(str(num))

        max_digit = "0"
        max_i = -1

        swap_i, swap_j = -1, -1

        for i in range(len(num) - 1, -1, -1):
            if num[i] > max_digit:
                max_digit = num[i]
                max_i = i
            if num[i] < max_digit: 
                swap_i, swap_j = i, max_i
        num[swap_i], num[swap_j] = num[swap_j], num[swap_i]

        return int("".join(num))

```
this one is easy in hind sight 
so first imagine your number is 2736
you turn it in to a string so its easy to swap and then you turn it into a list so its even easier
then you initialise max_digit and max-i
these will keep track of your current maxes so that you know what to swap
and then you keep trakc of the index you want to swap
loop through it backwards becasue we are going in order of significance
if the number is a max
    set it to that number 
    and note down the index
if it is less than the max
    note down the indexes because thats the number youre swapping

after youre done just make the swap and join them and youre good
