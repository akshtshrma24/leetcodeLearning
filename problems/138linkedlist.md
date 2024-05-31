# 138 

```py
"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        old_to_copy = { None: None}

        cur = head
        while cur:
            new_node = Node(cur.val)
            old_to_copy[cur] = new_node
            cur = cur.next
        
        cur = head
        while cur: 
            temp = old_to_copy[cur]
            temp.next = old_to_copy[cur.next]
            temp.random = old_to_copy[cur.random]
            cur = cur.next
        return old_to_copy[head]

```
So first we just loop through entire array and we put the new nodes at the key of the old node
this works because when we go in the second loop trying to set the next to the next variable we 
set it to the cur.next inside of the old copy map this will grab the right node thanks to getting 
the key. 