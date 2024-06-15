# 141 

end game
```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        if(head == NULL) return false;
        ListNode *l = head;
        ListNode *r = head->next;
        while(r != NULL){
            if(l == r) return true;
            l = l->next;
            r = r->next;
            if(r == NULL) return false;
            r = r->next;
        }
        return false;
    }
};
```

this one uses a fast pointer and a slow pointer...