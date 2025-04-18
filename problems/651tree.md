# 651 

real end game 

```py
from typing import (
    List,
)
from lintcode import (
    TreeNode,
)

"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left, self.right = None, None
"""

class Solution:
    """
    @param root: the root of tree
    @return: the vertical order traversal
    """
    def vertical_order(self, root: TreeNode) -> List[List[int]]:
        # write your code here
        if not root: 
            return []
        
        column_items = collections.defaultdict(list)

        min_x = float('inf')
        max_x = float('-inf')
        
        q = collections.deque([(0, root)])

        res = []

        while q: 
            x, n = q.popleft()
            column_items[x].append(n.val)

            min_x = min(min_x, x)
            max_x = max(max_x, x)

            if n.right:
                q.append((x + 1, n.right))
            if n.left:
                q.append((x - 1, n.left))

        for level in range(min_x, max_x + 1):
            res.append(column_items[level])
        
        return res
```

basically you do bfs but you keep track of the x axis. so the more left you go the more you subtract from x, and vice versa 