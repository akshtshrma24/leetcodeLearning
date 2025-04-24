# 316 

real end game 



```py

class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        last_pos = {char: i for i, char in enumerate(s)}
        stack = []
        seen = set()

        for i, char in enumerate(s):
            if char not in seen: 
                while stack and char < stack[-1] and i < last_pos[stack[-1]]:
                    seen.discard(stack.pop())
                seen.add(char)
                stack.append(char)
        
        return "".join(stack)
```

this one is kind of hard to visualise
you need a stack and a set
you make a dictionary with the last indexes of all of the positions 

then you go throught the string

if the character is not in the seen meaning you havent sene it before (duh)

you go to the while loop
while the stack is actually bigger than empty and the character is LESS than the one at the end AND the index is nto the last index aka less than it you 
can do a greedy thing 
you can just discard the stack.pop because you know that it is going to come again so then you discard the one from your set and add it to stack

the way this works is 
s = "cbacdcbc"
you make the dictionary so last_post = { a : 2, b: 6, c: 7, d: 4}
then you go through thte i and the char, if you havent seen the character before meaning it isnt a duplicate YET 
while there is a stack and the character is LESS than the end of stack and is the index is less than the index of the one at hte end of it (meaning we will see the end of the stack again) we discard the top of stack from set and add it to our stack after all that is done 

s = "cbacdcbc"
```py
for i, char in enumerate(s):
    if char not in seen: 
        while stack and char < stack[-1] and i < last_pos[stack[-1]]:
            seen.discard(stack.pop())
        seen.add(char)
        stack.append(char)
```

seen = ()
stack = []

c, 0

seen = (c)
stack = [c]

b, 1 
b < stack[1] True and 1 < 7 true
discard top of it 

i get it 