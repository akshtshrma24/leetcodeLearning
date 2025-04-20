# 199 

real end game

```py
while(q):
    q_len = len(q)
    curr = []
    for i in range(q_len):
        curr.append(q[i])
        q.append(q[i].left) if q[i].left else q
        q.append(q[i].right) if q[i].right else q

    q = q[q_len:]
    
    res.append(curr)
```

this one is easy once you know how level order traversal works

we go while the q is there 

take a temporary array and take the current length of q, we take the current one because we are appending to it so we dont want to go infinite 

then we append the current node, and then if q[i] has a left append it same thing with the right
then we take teh q and cut off the part we already traversed, like if the tree is below 
   1
 2  3 
3 2  3

we weill go like this q len is 1 at the moment so append 1 to it, it has a left it has a right, so append that so our q at end of loop is 
[1,2,3] then we cut off the traversed 
[2,3]
then we go to the length of the q so 2 
it has 2 has a left and a right so add those to the q, 3 only has a right so append that 
our q is [2,3,3,2,3]
remove q_len
[3,2,3]

do you see now? 
