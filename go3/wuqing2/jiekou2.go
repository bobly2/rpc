package wuqing2

import "fmt"
//解决两点
//1.因为必须要实现Handler接口，Do这个方法名不能修改，不能定义一个更有意义的名字
//2.必须要新定义一个类型，才可以实现Handler接口，才能使用Each函数

//提供两种函数，既可以以接口的方式使用，也可以以方法的方式，对应我们例子中的Each和EachFunc这两个函数，灵活方便。
type Handler interface {
	Do(k, v interface{})
}
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}
//用一个函数类型
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func Main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)

}

