# 226

real end game

```py

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs(root):
            if not root: 
                return 
            
            root.left, root.right = root.right, root.left

            dfs(root.left)
            dfs(root.right)

            return root
        
        return dfs(root)
```

stupid easy, just dfs through it swap the right and left, dfs both and then return the root

