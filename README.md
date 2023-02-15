# Go语言编程范例

*[English ]()* | *简体中文*

## 数据结构

切片作为队列

切片作为栈

`map`作为`set`

## 面向对象

Go语言面向对象编程的特点：
- `struct`可继承多个`struct`
- `struct`可实现多个`interface`
- `interface`可以继承`interface`
- `struct`继承`interface`需继承实现`interface`的`struct`

示例代码[./animal](./animal)中，各个interface的继承关系如下：
- [Animal](./animal/animal.go)
  - [Mammal](./animal/mammal.go)
    - [Cat](./animal/cat.go)
    - [Dog](./animal/dog.go)

Go语言推崇”组合优于继承“的编程思想。Go语言接口的实现无需显式地声明具体实现了哪些接口，只要实现了某个接口所有的方法，这个`struct`就是实现了该接口。
还有，`struct`的继承不像Java那样有`extends`和`implements`的区别，并且是声明为属性的形式，所以被叫做“组合”。

本小节举例的是多层继承，这是不符合“组合优于继承”的，但现实应用中，特别是在大型、复杂的项目中，我们无法避免地需要用到多层继承，本小节意在指出Go语言实现多层继承的方法。
