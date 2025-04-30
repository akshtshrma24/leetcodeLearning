# 129 

real end game

```py
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        res = 0

        def dfs(root, curr_val):
            if not root:
                return 0 
            
            curr_val = 10 * curr_val + root.val

            if not root.right and not root.left:
                return curr_val

            return dfs(root.left, curr_val) + dfs(root.right, curr_val)

        temp = dfs(root, 0)
        return temp
```

this one is easy, just go down the list appending the current number to curr_val and at the end just sum 