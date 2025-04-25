# 19 

real end game 

```py
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(val=69420, next=head)

        right = head

        for _ in range(n):
            right = right.next
        
        left = dummy

        while right:
            left = left.next 
            right = right.next
        
        left.next = left.next.next

        return dummy.next
```

this one is easy 
just make a dummy ode
set right as the head go until n
then set left as dummy
then go until right is null
this puts left at the exact right place to be then jsut delete the node
and return dummmy.next