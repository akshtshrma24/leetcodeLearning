# 143 

```py
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        s, f = head, head.next
        while f and f.next: # making sure that loop terminates when its null
            s = s.next
            f = f.next.next
        second = s.next # second half
        prev = s.next = None

        # this section gets the first half of it, by setting a slow and fast pointer

        # then while the seocnd one is there, you just want to reverse the pointers

        while second:
            tmp = second.next # temporary variable
            second.next = prev # this is why you need the temprary variable 
            prev = second # set the previous one to what hte og one was
            second = tmp # then assign the temporary variable back in 

        # 1 2 3 
        # tmp = 2
        # 


        # then you just go and mix match like how it wants it to do 
        first, second = head, prev
        while second:
            temp1, temp2 = first.next, second.next
            first.next = second
            second.next = temp1
            first, second = temp1, temp2
            
```

Very easy no notes needed, just remember how to reverse linked list