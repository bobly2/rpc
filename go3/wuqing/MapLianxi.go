package wuqing

//import  "fmt"

func TestMap() {
	//Map是一种数据结构，是一个集合，用于存储一系列无序的键值对 Map存储的是无序的键值对集合。
	//每次迭代Map的时候，打印的Key和Value是无序的

	//创建1
	//dict := make(map[string]int)
	//dict["张三"] = 43
	//创建2
	//dict := map[string]int{"张三": 43}
	//dict := map[string]int{"张三": 43, "李四": 50}
	////当然我们可以不指定任何键值对，也就是一个空map。
	//dict := map[string]int{}
	//
	////创建一个nil的map，未初始化，需要make一下
	//var dict map[string]int
	//dict = make(map[string]int)
	//dict["张三"] = 43
	//fmt.Println(dict)
	////如果键张三存在，则对其值修改，如果不存在，则新增这个键值对。
	//
	////还可以检测是否存在   boolean类型
	//age, exists := dict["李四"]
	//delete(dict, "张三")
	//
	//for k, v := range dict {
	//	fmt.Println(k, v)
	//}
	//这是无序的，若想有序，需要先排序

	//在函数间传递Map
	//函数间传递Map是不会拷贝一个该Map的副本的，也就是说如果一个Map传递给一个函数，该函数对这个Map做了修改，
	//那么这个Map的所有引用，都会感知到这个修改。
	//也就是说，引用map 会修改原本的值

	//go 类型
	//数值类型、浮点类型、字符类型以及布尔类型，基本类型，作为参数引用，不会改变本身

	//引用类型
	//引用类型有切片、map、接口、函数类型以及chan。
	//它的修改可以影响到任何引用到它的变量

	//引用类型之所以可以引用，是因为我们创建引用类型的变量，其实是一个标头值，标头值里包含一个指针，指向底层的数据结构，
	//当我们在函数中传递引用类型时，其实传递的是这个标头值的副本，它所指向的底层结构并没有被复制传递，这也是引用类型传递高效的原因。

	//结构类型
}
