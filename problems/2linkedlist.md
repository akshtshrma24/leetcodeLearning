# 2

end game

```c++

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *res = new ListNode(-1);
        ListNode *to_use = new ListNode();
        res->next = to_use;
        int rem = 0;
        while(l1 != NULL || l2 != NULL){
            int digit1 = (l1 != NULL) ? l1->val : 0;
            int digit2 = (l2 != NULL) ? l2->val : 0;

            int sum = digit1 + digit2 + rem;
            int dig = sum % 10;
            rem = sum / 10;

            ListNode *tot = new ListNode(dig);
            to_use->next = tot;
            to_use = to_use->next;
            l1 = (l1 != NULL) ? l1->next : NULL;
            l2 = (l2 != NULL) ? l2->next : NULL;
        }
        if(rem > 0){
            ListNode *tot = new ListNode(rem);
            to_use->next = tot;
            to_use = to_use->next; 
        }
        res = res->next;
        return res->next;
    }
};
```


this one is also quite easy just need to make sure I pay attention to the way that you calculate the digit and the sum and learn the difference from it. 