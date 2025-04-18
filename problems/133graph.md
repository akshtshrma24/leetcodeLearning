# 133 

real end game

```py

"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

from typing import Optional
class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if not node:
            return
            
        oldToNew = {}

        def dfs(node):
            if node in oldToNew:
                return oldToNew[node]
            copy = Node(node.val)
            oldToNew[node] = copy

            for nei in node.neighbors:
                copy.neighbors.append(dfs(nei))
            return copy
    
        return dfs(node)
```
this is kinda easy, so first just check if node exists or else just return 
make a hashmap 

in the dfs function if you have seen the node in the map before then you just return the seen node because think of it like this
1 <-> 3 

we make a copy of 1, then the oldToNew[node] you set that as its copy
then for its neighbors you loop so you go to 3, its not in there you make the copy and add to map so currently map os {1:1, 3:3}
then you loop through 3s neighbros and you see that node is already in oldToNew otherwise youd infintely hop back between the neighbros
then you just return the 1 instead of doing all the copy stuff to avoid loop, then you set 3 as neighbros as 1, then you return copy 
and then 3 is returned, well the copy of 3 is returned, then you are back at line 30 at 1s isntance and append it into neihgbors, thats the only neighbor so u go to next line which is just retunr copy