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
* 底层结构`reflect.StringHeader`: [string-len.go](./code/string/len.go)
```go
type StringHeader struct {
    Data unitptr
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

### 1.3.3 切片
* 头部含有底层数据指针和容量信息
* 赋值或传参时只赋值头部信息（注意这也时按值（头部的值）传递）
