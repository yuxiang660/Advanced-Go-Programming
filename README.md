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
* Go语言通过结构体内置匿名成员来实现继承：[method-inherit](./code/method/inherit.go)
