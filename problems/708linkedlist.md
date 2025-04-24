# 708 

real end game 

```py
class Solution 
    def insert(self, head: 'Optional[Node]', insertVal: int):
        if not head:
            new_node = Node(insertVal)
            new_node.next = new_node
            return new_node

        
        cur = head

        while cur.next != head:
            if cur.val <= insertVal and insertVal <= cur.next.val:
                new_node = Node(insertVal, cur.next)
                cur.next = new_node
                return head
            elif cur.next.val < cur.val:
                if insertVal >= cur.val or insertVal <= cur.next.val:
                    new_node = Node(insertVal, cur.next)
                    cur.next = newNode
                    return head
            
            cur = cur.next
        new_node = Node(insertVal, cur.next)
        cur.next = new_node

        return head
```

this one is very easy literally just check for the 3 cases inserting is easy 

if there is no head just make a loop and return it 

make a copy of head
and while cur.next != head this would exit when you have looped once, i reckon you could do this in a set but that owuld be harder 
then you check if cur.val is less than insertVal and and insertVal is less than cur.next.val this would mean that it needs to go inbetween them 

so you make a new node with the value as insert value and the next value as the cur.next
then you set cur.next as the new node 

and return head 

then you go to the new case where the cur.next.val < cur.val this would mean that it is about to loop so you want to check if the insertval is greater
than cur.val or if insert.val is less than cur.next.val this would mean you insert it inebetween those 2 nodes 

if none of those go into it you dont return anything meaning that all the values are the same

so you just make a node and insert it anywehre 
