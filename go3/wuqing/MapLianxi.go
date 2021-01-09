package wuqing

import "fmt"

func TestMap() {
	//Map是一种数据结构，是一个集合，用于存储一系列无序的键值对 Map存储的是无序的键值对集合。
	//每次迭代Map的时候，打印的Key和Value是无序的

	//创建1
	//dict := make(map[string]int)
	//dict["张三"] = 43
	//创建2
	dict := map[string]int{"张三": 43}
	dict := map[string]int{"张三":43,"李四":50}
	//当然我们可以不指定任何键值对，也就是一个空map。
	dict := map[string]int{}

	//创建一个nil的map，未初始化，需要make一下
	var dict map[string]int
	dict = make(map[string]int)
	dict["张三"] = 43
	fmt.Println(dict)
	//如果键张三存在，则对其值修改，如果不存在，则新增这个键值对。
}
