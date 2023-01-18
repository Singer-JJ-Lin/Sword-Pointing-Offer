# 二叉树的最近公共祖先
DFS对二叉树进行先序遍历，遇到p和q返回，从底至顶回溯，当节点p, q在节点root的异侧时，节点root即为最近公共祖先，则向上返回root 。
## 递归解析
### 终止条件
1. 当root,p,q有一个为空，返回空
2. 当root等于p,q， 则直接返回root

## 递归遍历左右子树,返回值分别为left和right
1. 如果当 left和righ同时为空，说明 root的左右子树中都不包含p和q,返回null
2. 当 left和right同时不为空,说明 p, q分列在root的异侧(分别在左右子树)，因此root为最近公共祖先，返回root
3. 当 left为空 ，right不为空，p,q都不在root的左子树中，直接返回right
4. 当 left不为空，right为空返回left 


