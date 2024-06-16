# 543 

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
    int track = 0;
    int dfs(TreeNode* root){
        if(!root) return 0;
        int l = dfs(root->left);
        int r = dfs(root->right);
        track = max(track, l + r);
        return max(l, r) + 1;
    }

    int diameterOfBinaryTree(TreeNode* root) {
        dfs(root);
        return track;
    }
};
```
dfs global track variable other wise if you dont track can get set to 2 when you have like 1 2 bc it goes 1 then increment 1 then goes to 2 then 2 goes to Null but 1 to 2 is 1 then it adds 1 more in the return line 
so track gets set to 2