---
markmap:
---

# 哈希表

## 定义

- 建立 `key` 和 `value` 之间的映射
- 输入一个 key，可以在 **O(1)** 时间获取对应的 value

### 元素效率对比 **O(1)**

- 添加元素: **O(1)**
- 查询元素：**O(1)**
- 删除元素：**O(1)**

## Hash table 常用操作

- 初始化
- 查询操作
- 添加键值对
- 删除键值对

### 常用遍历方式

- 遍历键值对
- 遍历键
- 遍历值

## Hash table 简单实现

### 定义

- 数组中每个空位称为「桶」，每个桶存储一个 `KV Pair`。
- 查询操作：找到 `key` 对应的桶，获取桶中的`value`。

#### 如何定位桶

- 通过某种哈希算法`hash()`计算得到哈希值。
- 将哈希值对**桶数量**(数组长度)`capacity`取模，获取该`key`对应的数组索引`index`
- ```
  index = hash(key) % capacity
  ```
- 最后利用 index 在哈希表中访问对应桶，获取 `value`
- ![Alt text](https://www.hello-algo.com/chapter_hashing/hash_map.assets/hash_function.png)

## 哈希冲突与扩容

### 哈希函数定义

- 将所有 `key`构成的输入框架映射到数组所有索引构成的输出空间。
- (输入往往远大于输出)
- **理论上一定存在"多个输入对应相同输出"的情况**

### 哈希冲突

- ![Alt text](https://www.hello-algo.com/chapter_hashing/hash_map.assets/hash_collision.png)
- 多个输入对应同一个输出称为「哈希冲突 hash collision」
- 哈希表容量 `n` 越大，多个`key` 被分配到同一个桶中的概率就越低，冲突越少
- **可以通过扩容哈希表来减少哈希冲突**

### 哈希扩容

- 扩容类似数组扩容，需将所有键值对从原哈希表迁移至新哈希表，非常耗时。
- 哈希表容量 `capacity` 改变，我们需要通过哈希函数来重新计算所有键值对的存储位置
- 通常预留足够大的哈希表容量，避免频繁扩容。
- 「负载因子 load factor」

# 哈希冲突
## 定义
## 链式地址
