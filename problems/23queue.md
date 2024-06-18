# 23 

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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<int> pq;
        if(lists.size() == 0) return NULL;
        for(auto& i : lists){
            while(i){
                pq.push(i->val * -1);
                i = i->next;
            }
        }
        if(pq.size() == 0) return NULL;
        ListNode* temp = new ListNode();
        ListNode* temp_next = new ListNode();
        temp->next = temp_next;
        while(!pq.empty()){
            temp_next->val = pq.top() * -1;
            pq.pop();
            if(pq.size() > 0){
                ListNode *temp_next_next = new ListNode();
                temp_next->next = temp_next_next;
                temp_next = temp_next->next;
            }
        }
        return temp->next;
    }
};

```

just min heaps for sorting;