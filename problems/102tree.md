# 102 

end game 

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
    vector<vector<int>> levelOrder(TreeNode* root) {
        queue<TreeNode*> q;
        vector<vector<int>> res;
        if(!root) return res;
        q.push(root);
        while(!q.empty()){
            int len = q.size();
            vector<int> temp; 
            for(int i = 0; i < len; i++){
                TreeNode* curr = q.front();
                q.pop();
                if(curr != NULL){
                    temp.push_back(curr->val);
                    q.push(curr->left);
                    q.push(curr->right);
                }
            }
            if(temp.size() > 0) res.push_back(temp);
        }
        return res;
    }
};
```
need to remember how to do bfs