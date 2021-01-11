package wuqing

import (
	"fmt"
	"math"
)

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

type welcome string

//从java理解，相当于   welcome  是一个类， 这个类实现了接口的方法  ()相当于 impl
func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}
func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}



func Main2() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26
	var w welcome = "大家好"
	//Each(persons, w)
	//想要改变do实现方法的名字
	//w.selfInfo("bob",27)
	//Each(persons, HandlerFunc(w.selfInfo))
	EachFunc(persons, w.selfInfo)
}

//我们定义了一个新的类型HandlerFunc，它是一个func(k, v interface{})类型
//新的HandlerFunc实现了Handler接口，Do方法的实现是调用HandlerFunc本身，因为HandlerFunc类型的变量就是一个方法。
//HandlerFunc(w.selfInfo)不是方法的调用，而是转型，因为selfInfo和HandlerFunc是同一种类型，所以可以强制转型



// 声明一个函数类型  只知道参数，不知道内部具体，
type HandlerFunc func(k, v interface{})
//   HandlerFunc实现了 Do， 将 w.selfInfo("bob",27)  转型 为 HandlerFunc
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}


//3
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m,HandlerFunc(f))
}


//最终 去掉  自定义类welcome
func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}
func Main22() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26
	EachFunc(persons, selfInfo)
}



//定义接口
type Phone interface {
	call()
	call2()
}

type Phone1 struct {
	id            int
	name          string
	category_id   int
	category_name string
}

//第一个类的第一个回调函数
func (test Phone1) call() {
	fmt.Println("这是第一个类的第一个接口回调函数 结构体数据：", Phone1{id: 1, name: "浅笑"})
}

//第一个类的第二个回调函数
func (test Phone1) call2() {
	fmt.Println("这是一个类的第二个接口回调函数call2", Phone1{id: 1, name: "浅笑", category_id: 4, category_name: "分类名称"})
}

//第二个结构体的数据类型
type Phone2 struct {
	member_id       int
	member_balance  float32
	member_sex      bool
	member_nickname string
}

//第二个类的第一个回调函数
func (test2 Phone2) call() {
	fmt.Println("这是第二个类的第一个接口回调函数call", Phone2{member_id: 22, member_balance: 15.23, member_sex: false, member_nickname: "浅笑18"})
}

//第二个类的第二个回调函数
func (test2 Phone2) call2() {
	fmt.Println("这是第二个类的第二个接口回调函数call2", Phone2{member_id: 44, member_balance: 100, member_sex: true, member_nickname: "陈超"})
}

//开始运行
func Main3() {
	var phone Phone

	//先实例化第一个接口
	phone = new(Phone1)
	phone.call()
	phone.call2()

	//实例化第二个接口
	phone = new(Phone2)
	phone.call()
	phone.call2()
}



//函数作为实参
func Main33() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}

// 声明一个函数类型
type cb func(int) int

func Main32() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}
func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}


