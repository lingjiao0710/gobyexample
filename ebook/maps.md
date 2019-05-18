## [Go by Example](https://gobyexample.com/): Maps

字典是Go语言内置的关联数据类型。因为数组是索引对应数组元素，而字典是键对应值。

[maps.go](<../src/maps.go>)

```go
package main

import "fmt"

func main() {

	// 创建一个字典可以使用内置函数make
	// "make(map[键类型]值类型)"
	m := make(map[string]int)

	// 使用经典的"name[key]=value"来为键设置值
	m["k1"] = 7
	m["k2"] = 13

	// 用Println输出字典，会输出所有的键值对
	fmt.Println("map:", m)

	// 获取一个键的值 "name[key]".
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 内置函数返回字典的元素个数
	fmt.Println("len:", len(m))

	// 内置函数delete从字典删除一个键对应的值
	delete(m, "k2")
	fmt.Println("map:", m)

	// 根据键来获取值有一个可选的返回值，这个返回值表示字典中是否
	// 存在该键，如果存在为true，返回对应值，否则为false，返回零值
	// 有的时候需要根据这个返回值来区分返回结果到底是存在的值还是零值
	// 比如字典不存在键x对应的整型值，返回零值就是0，但是恰好字典中有
	// 键y对应的值为0，这个时候需要那个可选返回值来判断是否零值。
	_, ok := m["k2"]
	fmt.Println("ok:", ok)

	// 你可以用 ":=" 同时定义和初始化一个字典
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//maps是引用类型，使用其他变量操作同一个map
	n1 := n
	n1["foo"] = 3
	n["test"] = 1000
	fmt.Println("map after changed ", n1)


	//使用for range 遍历maps
	fmt.Println("All items")
	for key, value := range n{
		fmt.Printf("[%s] = %d\n", key, value)
	}

	//maps不能使用==进行比较，比较maps只能对比maps中的每一个键值

}

```

```bash
$go run maps.go
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
ok: false
map: map[bar:2 foo:1]
map after changed  map[bar:2 foo:3 test:1000]
All items
[bar] = 2
[test] = 1000
[foo] = 3



```

下一篇:[Range](range.md)