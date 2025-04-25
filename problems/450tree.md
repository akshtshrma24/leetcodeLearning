# 450 

real end game

```py

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        def dfs(root, key):
            if not root: 
                return root
            
            if key > root.val:
                root.right = dfs(root.right, key)
            elif key < root.val:
                root.left = dfs(root.left, key)
            else:
                if not root.left: 
                    return root.right
                elif not root.right: 
                    return root.left
                
                cur = root.right
                while cur.left: 
                    cur = cur.left
                root.val = cur.val 
                root.right = dfs(root.right, root.val)

            return root
        
        return dfs(root, key)
```
this one is not that easy, so you first just find the value 
when u find it if there is no root left or root right just return the other root kind of easy 
otherwise u set cur = root.right becuase you want the samllest value in the right tree

then go left as much as u can
st the value to roots value, then delete the right trees value recursively