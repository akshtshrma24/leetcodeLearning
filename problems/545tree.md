# 545 

real end game 

```py
from typing import (
    List,
)
from lintcode import (
    TreeNode,
)

class Solution:
    def boundary_of_binary_tree(self, root: TreeNode) -> List[int]:
        # write your code here
        ans = []
        def dfs(node: TreeNode, left_bound: bool, right_bound: bool) -> None:
            if not node:
                return

            if left_bound and (node.left or node.right):
                ans.append(node.val)
            
            if not node.left and not node.right:
                ans.append(node.val)
                return

            dfs(node.left, left_bound and node.left, right_bound and not left_bound and not node.right)
            dfs(node.right, left_bound and not right_bound and not node.left, right_bound and node.right)

            if right_bound and not left_bound:
                ans.append(node.val)

        dfs(root, True, True)
        return ans
```

so what we do here is we run dfs, if we are left bound and there is a node iether at node left or node right, add the value 
