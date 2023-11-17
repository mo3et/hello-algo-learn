# Hash table

Key-Value

> 「哈希表 hash table」，又称「散列表」，其通过建立键 key 与值 value 之间的映射，实现高效的元素查询。具体而言，我们向哈希表输入一个键 key ，则可以在 时间内获取对应的值 value 。

![Alt text](https://www.hello-algo.com/chapter_hashing/hash_map.assets/hash_table_lookup.png)

哈希表与数组和链表的查询效率对比

- 添加元素：仅需将元素添加至数组（链表）的尾部即可，使用 O(1) 时间。
- 查询元素：由于数组（链表）是乱序的，因此需要遍历其中的所有元素，使用 O(n) 时间。
- 删除元素：需要先查询到元素，再从数组（链表）中删除，使用 O(n) 时间。

## 哈希表常用操作

```go
/* 初始化 hash */
hmap := make(map[int]string)

/* 添加操作 */
// 在 hash table 中添加 kv
hmap[12836] = "小哈"
hmap[15937] = "小啰"
hmap[16750] = "小算"
hmap[13276] = "小法"
hmap[10583] = "小鸭"

/* 查询操作 */
// 向 Hash table 输入 key, 得到 value
name := hmap[15937]

/* 删除操作 */
delete(hmap.10583)
```

哈希表有三种常用遍历方式：
**遍历键值对、遍历键、遍历值**

```go
/* 遍历哈希表 */
// 遍历键值对 key->value
for key , value := range hmap {
    fmt.Println(key,"->",value)
}
// 单独遍历键 key

for key := range hmap {
    fmt.Println(key)
}
// 单独遍历键 value (同时获取kv，然后忽略key)

for _ , value := range hmap {
    fmt.Println(value)
}
```

## 哈希表简单实现

最简单的情况，**仅用一个数组来实现哈希表**。在哈希表中，我们将数组中的每个空位称为 「桶 bucket」，每个桶可存储一个键值对。因此，**查询操作就是找到`key`对应的桶，并在桶中提取`value`.**

如何基于`key`来定位对应的桶呢？通过「哈希函数 hash Function」实现的。

![Alt text](https://www.hello-algo.com/chapter_hashing/hash_map.assets/hash_function.png)

**简单哈希表实现**

```go
/* 将 key 和 value 封装成一个类 Pair 表示键值对 */
type pair struct {
    key int
    val string
}

/* 基于数组简易实现的哈希表 */
type arrayHashMap struct {
    buckets []*pair
}

/* 初始化哈希表 */
func newArrayHashmap() *arrayHashMap {

}

```
