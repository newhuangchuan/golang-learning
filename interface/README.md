# 大纲
- 函数 defer和panic
- 结构体（重点）
- 面向对象
- 接口
- 错误处理
- io操作
- 反射

## defer遇到panic
- 遇见panic，强行defer出栈。并且每个defer均收到异常，不捕获则转移到下一个
- 知道defer捕获异常或者上层显示异常
> defer最大的功能就是panic之后依然有效。
> 所以defer可以保证你的一些资源一定会被关闭，从而避免一些异常的出现问题。
> 遇到panic会执行已经压栈的defer。

> 不用recovery捕获异常
- panic之后的defer还没来得及压栈
> 使用recovery捕获异常
- 如果遇到recover，返回recover处会继续往下执行，不会抛出panic异常信息
- 需要注意的是recovery处会向下执行，而不会返回再执行。

## defer中含有panic
- panic仅有最后一个可以被recover捕获。
- 这个异常将会覆盖掉main中的异常`panic("panic"")`，最后这个异常被第二个执行的defer捕获到了。
```go
package main

import "fmt"

func main() {
	panicIndefer()


}

func panicIndefer()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("hahahha")
		}
	}()

	defer func() {
		panic("defer 内部触发的一个panic")
	}()

	//外部触发一个panic
	panic("外部触发的panic")
}

```
## defer中函数参数包含子函数
- 先压栈1，再压栈2。
- 子函数在压栈时要执行获取返回值作为参数。
- 因为压栈时需要联通函数地址，函数形参一同进栈。
- 所以调用顺序为，压栈f1，调用f3，压栈f2，调用f4，调用f2，f1。
```go

package main

import "fmt"

func main() {
	defer function(1,function(3, 0))
	defer function(1,function(4, 0))

}

func function(index int , value int) int {
	fmt.Println(index)
	return index
}

```
## 结构体知识点
- 结构体的定义
- 结构体声明
- 结构体初始化
- new函数介绍
- 属性访问与修改
- 结构体命名嵌入
- 结构体匿名嵌入
- 结构体执行类型嵌入
- 结构体指针类型嵌入
- 结构体可见性
- 结构体深浅拷贝

### 结构体的特点
- Go语言中数组可以存储同意类型的数据，但是在结构体中我们可以为不同项定义不同的数据类型。
- 结构体的目的就是把数据聚集在一起，一遍能够更加便捷的操作这些数据。
### 结构体和类的概念
- Go里面没有类，Go用一种特殊的方式，把结构体本身看作一个类。
- 一个成熟的类，具备成员变量和成员函数，结构体本身就是成员变量，在给他绑定上成员函数，就可以了

## 01 结构体的定义
### 定义方式
- 在Golang中最常用的方法就是使用关键字type和struct来定义一个结构体
```go
type Person struct {
	Name string
	Age int
	Label map[string]string
}
```
### 字段说明
- 结构体中的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。

#### 使用结构体类型
- 使用结构体，举例： 配置文件不同段，global段、mysql段、redis段。
- 使用全局的配置变量可以拿到下属的所有段信息。

#### 字段标记信息
- 在定义结构体时可以为字段指定一个标记信息，这些标记信息通过反射接口可见，冰参与结构体的类型标识。
- 常用来做yaml解析、json解析、xorm/gorm连接mysql字段标记，距离如公有云ecs结构体

## 02 结构体声明和初始化
### 方法一： 使用var 声明变量并初始化
- 关键字var可以使用结构体类型声明变量，并初始化为零值，
- 关键字var创建了类型为Person且名为p的变量，p被成为类型Person的一个实例（interface）
```go
package main

import "fmt"

type Person struct {
	Name string
	Age int
	Label map[string]string
}

func main() {
	var p Person
	fmt.Println(p)
}
```
#### var 等号： 使用自定义数据初始化
- 当声明变量时，这个变量对应的值总是会被初始化。
- 使用var关键字用零值初始化，对于数值类型来说，零值是0；对于字符串来说，零值是空字符串；对于布尔值来说，零值是false。
- var + NewXXX 初始化和全局变量
```go
package main

import (
	"log"
)

type Person struct {
	Name string
	Age int
	Label map[string]string
}

func main() {
	var p = Person{Name: "root"}
	log.Printf("[%v  %+v]\n",p, p)
}

```
#### var特点，可以在函数外部使用，可以声明初始化全局变量
- 举例使用var=XXX初始化全局缓存

### 方法二： 使用短变量声明操作符（:=）初始化
#### 特点1 使用自定义值初始化
- 写法1： 明确写出字段的名字以及对应的值，可以不关心定义时的字段顺序
- 写法2： 不屑字段名称，只写值，这时字段的顺序必须和定义时一致`name := Person{"root", 18, "123456789@qq.com""}`

#### 特点2 不可以在函数外部使用，不可以声明初始化全局变量
```go
package main

import (
	"log"
)

type Person struct {
	Name string
	Age int
	Label map[string]string
}

func main() {
	var p = Person{
		Name: "root",
		Age: 18,
		Label: map[string]string{},
	}
	p2 := Person{
		Name: "administrator",
		Age: 99,
		Label: map[string]string{},
	}
	log.Printf("[%v  %+v]\n",p, p)
	log.Printf("[%+v]\n", p2)
}
```
## 03 new和make
### new只能把内存初始化为零值并返回其指针
- 举例来对比new和var返回的类型
```go
	p3 := new(Person)
	p3 = &Person{"admin", 18, map[string]string{}}
	log.Printf("[%+v]\n", p3)
```


### new和make对比
- 简单来说new只分配内存，make用于slice，map和channel的初始化
- new可以对所有的类型进行内存分配，返回指针，填充零值。
- make只能创建类型（slice map channel），返回引用，填充非零值
```go
	var p1 *Person = new(Person)
	p1.Age = 19
	p1.Name = "xiao"
	log.Println(*p1)
```

# 匿名结构体
## 匿名字段
- 顾名思义就是没有字段名的字段
```go
type test struct {
name string
age int
int
}
```
- 匿名字段和面向对象编程中的继承概念相似，可以用来模拟类似继承的行为

## 匿名结构体
- 又称为内嵌结构体
- 结构体可以包含一个或者多个匿名的字段，即这些字段没有显示的名字。
- 举例student和person
```go
package main

import (
	"log"
)

type Person1 struct {
	Name string
	Age  int
	Label map[string]string
}

type Student struct {
	StudentId int
	Person1  //匿名结构体的嵌套
}

func main() {
	p1 := Person1{
		Name: "abc",
		Age: 18,
		Label: nil,
	}
	//访问属性
	log.Printf("[p.name : %v][p.age : %v]\n", p1.Name, p1.Age)
	//修改属性
	p1.Age +=1
	log.Printf("[p.name : %v][p.age : %v]\n", p1.Name, p1.Age)

	s1  := Student{
		StudentId: 2292929,
		Person1: p1,
	}

	log.Printf("[s1.name : %v][s1.age : %v]\n", s1.Name, s1.Age)
}
```

# 结构体命名嵌入
- 结构体命名嵌入，就是嵌入的结构体又字段名
    - 使用字段名属性名访问
    
- 结构体匿名嵌入，嵌入的结构体无字段名，Go语言又一个特性让我们只声明一个成员对应的数据类型而不指定成员你的名字，这类的成员就叫做匿名成员
    - 使用属性直接访问
    - 需要注意小写属性导出问题
    
```go
package main

import "log"

type Person2 struct {
	Name string
	Age int
	Label map[string]string
}

type Student1 struct {
	StudentId int
	Person2
}

type  Teacher struct {
	TeacherId int
	P Person2
}
func main() {
	 p1 := Person2{
	 	Name: "tty",
	 	Age: 99,
	 	Label: map[string]string{},
	 }
	//匿名结构体的嵌入
	 s1 := Student1{
	 	StudentId: 994399223,
	 	Person2: p1,
	 }

	 p2 := Person2{
	 	Name: "tty2",
	 	Age: 89,
	 	Label: map[string]string{},
	 }
	//命名结构体的嵌入
	 t1 := Teacher{
	 	TeacherId: 2287334443,
	 	P: p2,
	 }
	//匿名结构体可以直接访问继承的属性
	 log.Printf("[s1.name is %v,s1.age is %v, s1.studentid is %v]\n", s1.Name, s1.Age, s1.StudentId)
	 log.Printf("[s1.Person2.name is %v, s1.Person2.age is %v, s1.studentid is %v]\n", s1.Person2.Name, s1.Person2.Age,s1.StudentId)
	 log.Printf("[t1.P.name is %v,t1.P.age is %v, t1.studentid is %v]\n",t1.P.Name,t1.P.Age,t1.TeacherId )
}

```
> 结构体匿名嵌入的作用就是为了处理公共字段

# 结构体指针类型的嵌入
- 指针的嵌入有什么不同，零值不同
- 空指针不能直接访问属性进行操作
## 结构体单例绑定
- sayHello() 用了指针的方式进行绑定，相当于给结构体绑定了一个函数，这个结构体等价于对象。
- 唯一的不同点就是如果使用*绑定函数，那么这种对象就是单例，引用的是同一个结构体。

# 结构体可见性
> 结论 结构体名的大小写影响结构体本身的可访问性，首字母小写则包外不可见
## 02 结构体变量的成员不同包访问
> 结论 结构体字段名的大小写影响字段的可访问性，首字母小写则包外不可见
## 03 结构体变量的成员变量同包访问
> 结论，同包内，结构体变量的成员变量可以随时被访问
















