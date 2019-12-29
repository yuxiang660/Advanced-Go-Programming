# 第1章 语言基础
## 1.2 Hello, World的革命
* [hello.go](./code/hello/hello.go)
* [http-hello.go](./code/http-hello/http-hello.go)

## 1.3 数组、字符串和切片
* 赋值方式<br>
除了闭包函数以引用方式访问外，其他赋值或函数传参都是传值。

### 1.3.1 数组
* 值语义：赋值或函数传参时是整体复制（包括底层内存）
* 不同长度代表不同类型，因此不能互相赋值
* 定义：[array-defination.go](./code/array/defination.go)
* 指针与赋值：[array-pointer.go](./code/array/pointer.go)
* 二维数组：[array-two-dimension.go](./code/array/two-dimension.go)

### 1.3.2 字符串
* 只读
* 赋值时只复制了地址和长度
* 底层头部结构`reflect.StringHeader`: [string-len.go](./code/string/len.go)
```go
type StringHeader struct {
    Data uintptr
    Len int
}
```
* 字符串以UTF8编码：[string-utf8.go](./code/string/utf8.go)
* `rune`：[string-rune.go](./code/string/rune.go)
    - `rune` means the same as "code point" which refers to the item represented by a single value. For example, the code point U+2318, with hexadecimal value 2318, represents the symbol ⌘. So,
    > `⌘` is `rune` with integer value 0x2318.
* 单引号、双引号、反引号：[string-single-double-back-quote.go](./code/string/single-double-back-quote.go)
    - 单引号表示`rune`类型
    - 双引号表示`string`类型
    - 反引号用来创建原生的字符串字面量，不支持任何转义序列
* `string`、`[]byte`和`[]rune`之间的转换：[string-bytes-transition](./code/string/string-bytes-runes-transition.go)

### 1.3.3 切片
* 头部含有底层数据指针和容量信息
* 赋值或传参时只赋值头部信息（注意这也时按值（头部的值）传递）
* 长度不是类型的组成部分，是简化版的动态数组
* 底层头部结构`reflect.SliceHeader`:
```go
type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
}
```
* 定义：[slice-defination](./code/slice/defination.go)
* 追加切片元素：[slice-append](./code/slice/append.go)

## 1.4 函数、方法和接口
* 函数
    - 具名函数<br>
    是匿名函数的一种特例
    - 匿名函数<br>
    匿名如果引用了外部作用域中的变量就变成了闭包函数
* 方法
    - 绑定到一个具体类型的特殊函数
    - 编译时静态绑定
* 接口
    - 方法的集合
    - 运行时动态绑定
    - 通过隐式接口机制实现了鸭子面向对象模型
* Go程序初始化入口：`main.main`函数
    1. 导入其他包，执行每个包的`init`函数
    2. `main`包的所有包级常量、变量被创建和初始化，并执行`main`包的`init`函数

### 1.4.1 函数
* 函数是第一类的对象
* 具名函数
```go
func Add(a, b int) int {
    return a+b
}
```
* 匿名函数
```go
var Add = func(a, b int) int {
    return a+b
}
```
* 多返回函数
```go
func Swap(a, b int) (int, int) {
    return b, a
}
```
* 可变参数；[func-variable-parameter](./code/func/variable-parameter.go)
    - `func foo(parm ...int)`打包，将多个传入`foo`的`int`类型打包成一个`[]int`
    - `foo(parm...)`拆包，将`parm`拆成多个`int`类型的数据，传给`foo`
* `defer`关键字
    - 延时被修饰的函数到当前函数结束的时候执行
    - `defer`多个函数的按照先进后出的顺序执行
    - `defer`一个匿名函数<br>
    需要注意闭包捕获是按**引用**方式传递的，如果多个`defer`匿名函数捕获外部变量，会是同一个引用，参见：[func-defer](./code/func/defer.go)
* Go语言的栈
    - Go编译器会处理栈和堆的问题，不需要程序员考虑，参见：[func-stack-heap](./code/func/stack-heap.go)
    - 不要假设变量在内存中的位置是固定不变的，因为Go采用动态栈实现函数栈大小的调整。每个`goroutine`刚启动时只会分配很小的栈，当有需求增大栈空间时，新内存空间被开辟，为了保持栈的连续，之前的数据会移动到新开辟的内存空间，这会导致之前中全部变量的地址发生变化。

### 1.4.2 方法
* Go方法不支持重载，每个方法的名字必须时唯一的
* 将方法还原为普通类型的函数
```go
// func (f *File) Close() error {...}
// to
// func CloseFile(f *File) error
var CloseFile = (*File).Close
```
* Go语言的继承
    - 通过结构体内置匿名成员来继承匿名成员的内部成员和对应方法：[method-inherit](./code/method/inherit.go)
    - `childType.parentVar`会在编译时被展开为`childType.parentType.parentVar`
    - `C++`子类方法是在运行时动态绑定到对象的，因此基类实现的某些方法看到的`this`可能不是基类类型对应的对象，这个特性会导致基类方法运行的不确定性。Go语言通过嵌入匿名的成员来“继承”基类的方法，`this`就是实现该方法的类型的对象。但是这样的方法并不能实现`C++`中虚函数的多态特性。如果需要虚函数的多态特性，我们需要借助Go语言接口来实现。

### 1.4.3 接口
* 接口类型实现了对鸭子类型的支持
* 利用`fmt.Printf`的接口参数，实现大写输出：<br>
    - [interface-upper-writer](./code/interface/upper-writer.go)
    - [interface-upper-string](./code/interface/upper-string.go)
* 接口的显式转换与隐式转换：[interface-conversion](./code/interface/conversion.go)
    - 非接口类型（如`int`型）不支持隐式转换，必须显示转换，如：
    > `var numInt64 = (int64)(numInt)`
    - 接口和对象，接口和接口之间的转换支持隐式转换
* 内置接口
```go
type error interface {
    //只要实现了Error()函数，返回值为String的都实现了error接口
    Error() String
}
```
* 接口“继承”接口
```go
type runtime.Error interface {
    error
    // 既需要实现上面的Error()函数，也需要实现RuntimeError()函数，才实现了runtime.Error接口
    RuntimeError()
}
```

## 1.5 面向并发的内存模型
* 与Erlang不同，Go语言的Goroutine之间是共享内存的
### 1.5.1 Goroutine和系统线程
* goroutine和系统线程不是等价的
### 1.5.2 原子操作
* `sync.Mutex`：[lock-mutex](./code/lock/mutex.go)
* `atomic`: [lock-atomic](./code/lock/atomic.go)
* `sync.Once`：[lock-singleton](./code/lock/singleton.go)
### 1.5.3 顺序一致性内存模型
* 两个Goroutine之间，一个Goroutine不能知道另一个Goroutine的内部执行顺序，甚至一直不能看到对共享资源的改动（可能始终在寄存器中），直到另一个Goroutine结束。
### 1.5.6 基于Channel的通信
* 对Channel的发送和接收通常发生在不同的Goroutine上，在同一个Goroutine上执行2个操作很容易导致死锁。
* 无缓冲的Channel上的发送操作总在对应的接收操作完成前发生：[unbuffered-channel](./code/channel/unbuffered-channel.go)
* 对于带缓冲的Channel，第K个接收完成操作发生在第K+C个发送完成之前，其中C是Channel的缓存大小。对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送**完成**之前。对于从无缓冲Channel进行的接收**完成**，发生在对该Channel进行的发送完成之后。
* `select{}`是一个空的管道选择语句，会导致线程阻塞。