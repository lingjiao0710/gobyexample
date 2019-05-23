## [Go by Example](https://gobyexample.com/): Structs

Go语言结构体数据类是将各个类型的变量定义的集合，通常用来表示记录。

[structs.go](<../src/structs.go>)

```go
package main

import "fmt"

type Address struct{
	city, state string
}

// 这个person结构体有name和age成员
type person struct {
	name string
	age  int
	Address
}

//匿名字段
type Student struct{
	string
	int
}

func main() {

	// 这个语法创建一个新结构体变量
	fmt.Println(person{"Bob", 20, Address{"NewYork", "sanda"}})

	// 可以使用"成员:值"的方式来初始化结构体变量
	fmt.Println(person{name: "Alice", age: 30})

	// 未显式赋值的成员初始值为零值
	fmt.Println(person{name: "Fred"})

	// 可以使用&来获取结构体变量的地址
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用点号(.)来访问结构体成员
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 结构体指针也可以使用点号(.)来访问结构体成员
	// Go语言会自动识别出来
	sp := &s
	fmt.Println(sp.age)

	// 结构体成员变量的值是可以改变的
	sp.age = 51
	fmt.Println(sp.age)

	//结构的零值
	var emp person
	fmt.Println("epm :", emp)


	//匿名字段
	p := Student{"Lilei", 100}
	fmt.Println(p)
	p.string = "Hanmeimei"
	p.int = 50
	fmt.Println(p)

	//嵌套结构,并且嵌套的结构属于匿名字段一样，可以直接通过外部结构访问
	emp.name = "Navven"
	emp.age = 40
	emp.Address = Address{
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println(emp)
	fmt.Println(emp.city)

	//当结构中的字段以大写字母开头时，该字段称为导出字段，可以被其他包访问


	//结构是值类型，如果每个字段都可以比较，那结构就是可以比较的
	stu1 := Student{"abc", 11}
	stu2 := Student{"abc", 11}

	if(stu1 == stu2){
		fmt.Println("stu1 and stu2 are equal")
	}


	//如果结构中包含不可比较的字段，那结构就不能比较
	/*
	type image struct{
		data map[int]string
	}

	image1 := image{
		data: map[int]string{5: "test"},
	}
	image2 := image{
		data: map[int]string{5: "test"},
	}

	//invalid operation: image1 == image2 (struct containing map[int]string cannot be compared)
	if(image1 == image2){
		fmt.Println("image1 and image1 are equal")
	}*/
}
```

```bash
$go run structs.go
{Bob 20 {NewYork sanda}}
	{Alice 30 { }}
	{Fred 0 { }}
	&{Ann 40 { }}
	Sean
	50
	51
	epm : { 0 { }}
	{Lilei 100}
	{Hanmeimei 50}
	{Navven 40 {Chicago Illinois}}
	Chicago
	stu1 and stu2 are equal
```

下一篇:[方法 Methods](methods.md)