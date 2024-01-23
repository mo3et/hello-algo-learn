---
markmap:
---

# 二叉树

## 定义

- 非线性数据结构，代表祖先与后代直接的派生关系
- 与链表类似，二叉树的基本单元是「节点 Node」

## 常见术语

- 「根节点 root node」：位于二叉树顶层的节点，没有父节点。
- 「叶节点 leaf node」：没有子节点的节点，其两个指针均指向 。
- 「边 edge」：连接两个节点的线段，即节点引用（指针）。
- 节点所在的「层 level」：从顶至底递增，根节点所在层为 1 。
- 节点的「度 degree」：节点的子节点的数量。在二叉树中，度的取值范围是 0、1、2 。
- 二叉树的「高度 height」：从根节点到最远叶节点所经过的边的数量。
- 节点的「深度 depth」：从根节点到该节点所经过的边的数量。
- 节点的「高度 height」：从距离该节点最远的叶节点到该节点所经过的边的数量。
- ![](https://www.hello-algo.com/chapter_tree/binary_tree.assets/binary_tree_terminology.png)

## 基本操作

### 初始化二叉树

- 与链表类似，首先初始化节点，然后构建引用（指针）
- ```go
  /* 初始化二叉树 */
  // 初始化节点
  n1:=NewTreeNode(1)
  n2:=NewTreeNode(2)
  n3:=NewTreeNode(3)
  n4:=NewTreeNode(4)
  n5:=NewTreeNode(5)

  // 构建引用指向(指针)
  n1.Left = n2
  n1.Right = n3
  n2.Left = n4
  n2.Right = n5
  ```

### 插入与删除节点

- 在二叉树插入与删除节点可以通过修改指针来实现。
- ![](https://www.hello-algo.com/chapter_tree/binary_tree.assets/binary_tree_add_remove.png)
- ```go
    /* 插入与删除节点 */
  // 在n1 -> n2 中间插入节点 P
  p:= NewTreeNode(0)
  n1.Left=p
  p.Left = n2
  // 删除节点 P
  n1.Left = n2
  ```
- [Attention]需要注意的是，插入节点可能会改变二叉树的原有逻辑结构，而删除节点通常意味着删除该节点及其所有子树。因此，在二叉树中，插入与删除操作通常是由一套操作配合完成的，以实现有实际意义的操作。

### 常见二叉树类型
