# 栈

stack 是一个遵循先入后出(FILO)逻辑的**线性**数据结构。

举例，叠盘子（叠叠乐），需要从上到下依次取出。

我们把堆叠元素的顶部称为「栈顶」，底部称为「栈底」。把元素添加到栈顶叫做「入栈」，删除栈顶元素叫做「出栈」。

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/stack_operations.png)

## 栈常用操作

常见的为`push()`、`pop()`、`peek()`，其中 peek 为访问栈顶元素，三个操作的时间复杂度都为O(1).

通常情况下，都可以直接使用栈类。然而某些语言没有专门提供栈类(例如 Go)，此时可以将该语言的“数组”或“链表”视作栈来使用，并在程序逻辑上忽略与栈无关的操作。

```go
/* 初始化栈 */
// 在 Go 中，推荐将 Slice 当作栈来使用
var stack []int

/* 元素入栈 */
stack = append(stack,1)
stack = append(stack,2)
stack = append(stack,3)
stack = append(stack,5)
stack = append(stack,4)

/* 访问栈顶元素 */
peek := stack[len(stack)-1]

/* 元素出栈 */
pop := stack[len(stack)-1]
stack := stack[:len(stack)-1]

/* 获取栈的长度 */
size :=len(stack)

/*判断是否为空 */
isEmpty := len(stack) == 0
```

## 栈的实现

栈遵循先入后出的原则，因此我们只能在栈顶添加或删除元素。然而，数组和链表都可以在任意位置添加和删除元素，**因此，栈可以被视为一种受限制的数组或链表**。换句话说，可以“屏蔽”数组或链表的部分无关操作，使其对外表现的逻辑符合栈的特性。

### 1. 基于链表的实现

使用链表来实现栈时，我们可以将链表的头节点视为栈顶，尾节点视为栈底。

对于入栈操作，我们只需将元素插入链表头部，这种节点插入方法被称为「头插法」。

对于出栈操作，只需将头节点从链表中删除即可。

**LinkedListStack**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/linkedlist_stack.png)

**push()**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/linkedlist_stack_push.png)

**pop()**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/linkedlist_stack_pop.png)

```go
/* 鉴于链表实现的栈 */
type linkedListStack struct{
    // 使用内置包 list 来实现栈
    data *list.List
}

/* 初始化栈 */
func newLinkedListStack() *linkedListStack{
    return &linkedListStack{
    data: list.New(),
    }
}

/* 入栈 */
func (s *linkedListStack) push(value int){
    s.data.PushBack(value)
}

/* 出栈 */
func (s *linkedListStack) pop() any{
    if s.isEmpty(){
        return nil
    }
    e:= s.data.Back()
    s.data.Remove(e)
    return e.Value
}

/* 访问栈顶元素 */
func (s *linkedListStack) peek() any {
    if s.isEmpty(){
        return
    }
    e:= s.data.Back()
    return e.Value
}

/* 获取栈的长度 */
func (s *linkedListStack) size() int {
    return s.data.Len()
}

/* 判断栈是否为空 */
func (s *linkedListStack) isEmpty() bool {
    return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListStack) toList() *list.List {
    return s.data
}
```

### 2. 基于数组的实现

使用数组实现栈时，可以将**数组的尾部作为栈顶**。入栈和出栈操作分别在数组尾部添加元素与删除元素，时间复杂度都为O(1).

**ArrayStack**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/array_stack.png)

**push()**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/array_stack_push.png)

**pop()**

![](https://www.hello-algo.com/chapter_stack_and_queue/stack.assets/array_stack_pop.png)

由于入栈的元素会不断增加，因此我们可以使用动态数组，无需处理数组扩容问题。

```go
/* 基于数组实现的栈 */
type arrayStack struct {
    data []int // 数据
}

/* 初始化栈 */
func newArrayStack() *arrayStack {
    return &arrayStack{
        // 设置栈的长度为 0, 容量为 16
        data: make([]int, 0, 16)
    }
}

/* 栈的长度 */
func (s *arrayStack) size() int {
    return len(s.data)
}

/* 栈是否为空 */
func (s *arrayStack) isEmpty() bool{
    return s.size() ==0
}

/* 入栈 */
func (s *arrayStack) peek() any{
    if s.isEmpty(){
    return nil
    }
}

/* 获取栈顶元素 */
func (s *arrayStack) peek() any {
    if s.isEmpty() {
        return nil
    }
    val := s.data[len(s.data)-1]
    return val
}

/* 获取 Slice 用于打印 */
func (s *arrayStack) toSlice() []int {
    return s.data
}
```
