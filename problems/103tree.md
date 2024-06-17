# 103 

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
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        queue<TreeNode*> q;
        vector<vector<int>> res;
        if(!root) return res;
        q.push(root);
        while(!q.empty()){
            int len = q.size();
            vector<int> temp_list;
            for(int i = 0; i < len; i++){
                TreeNode* temp = q.front();
                q.pop();
                if(temp){
                    temp_list.push_back(temp->val);
                    q.push(temp->left);
                    q.push(temp->right);
                }
            }
            if(temp_list.size() > 0) {
                if(res.size() % 2 == 0) res.push_back(temp_list);
                else {
                    reverse(begin(temp_list), end(temp_list));
                    res.push_back(temp_list);
                }
            }
        }
        return res;
    }
};
```
this is just the same as normal traversal but since we want zig zag we just make sure to reverse it every other go around 
