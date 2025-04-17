# 543

real end game

```py
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        ans = 0
        
        def dfs(root):
            nonlocal ans
            
            if not root: 
                return 0

            left = dfs(root.left)
            right = dfs(root.right)

            ans = max(ans, left + right)

            return max(left, right) + 1
    
        dfs(root)
        return ans
```

this ones quite easy you take the left and right, diameter can be not just the longest path from left to right i nthe root node
so you take the maximum of the left + right or the max alread, then return the longest path + the current node