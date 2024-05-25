# 206 

```py
# iterative 
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, curr = None, head
        while curr:
            next_one = curr.next
            curr.next = prev
            prev = curr
            curr = next_one
        return prev

# recursive

```
im thinking you make a copy of it and loop through it 
yea i was right you have to use 2 pointers to do this 
recursive way is just to get to the end and then change the next next 
to point back so that you can reverse it
