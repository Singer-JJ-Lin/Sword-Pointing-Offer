# 重建二叉树
1. 前序遍历序列的第一个节点为根节点，在中序遍历中找到这个节点，左边部分为左子树，右边部分为右子树
2. 根据中序遍历中左子树和右子树的长度，找到左子树和右子树在前序遍历的范围
3. 递归构建