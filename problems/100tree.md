# 100 

END GAME

```c++
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool dfs(TreeNode *p, TreeNode* q) {
        if(!p && !q) return true;
        if(p && q && p->val == q->val){
            return dfs(p->left, q->left) && dfs(p->right, q->right);
        }
        return false;
    }
    bool isSameTree(TreeNode* p, TreeNode* q) {
        return dfs(p, q);
    }
};
```

this one is easy as well, just compare the left and right branches make sure they both get to null at the same time and
both are equivalent value wise otherwise return false, make sure both branhces return same thing