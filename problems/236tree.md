# 236 

real end game

```py
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':

        def dfs(root, p, q):
            if not root: 
                return None
            
            if root == p or root == q:
                return root
        
            l = dfs(root.left, p, q)
            r = dfs(root.right, p, q)

            if l and r: 
                return root
            else: 
                return l or r
            
        return dfs(root, p, q)

# T: O(N) worst case have to go all the way in the tree 
# S: O(1) not allocating extra space, but recursive solution, so if not counting recursive stack frames
    # otherwise O(N) to count for stack
```

so you are basically doing dfs right
you obviously check if its root is null if it is cant do antthign jsut return None 
BUT if root is p or root is q
then you want to pop up say hey I am one of them search in the other sie then you search in the other side 
if you find in other side you know the one you are on is the lowest, if it isnt in the other side then you just return l or r
becasue which ever one popped up from underneath that said hey im p or q is l or r so just return l or r