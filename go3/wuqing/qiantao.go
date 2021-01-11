package wuqing

import "fmt"

//嵌入类型 在其他语言中，有继承可以做同样的事情，在Go语言中，没有继承的概念
type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func MainTest() {



	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	fmt.Println("可以直接调用,名字为：", ad.name)
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)
	ad.user.sayHello()
	ad.sayHello()
}
func (u user) sayHello() {
	fmt.Println("Hello，i am a user")
}

func (a admin) sayHello() {
	fmt.Println("Hello，i am a admin")
}


//改变
func MainTest2() {
	ad:=admin{user{"张三","zhangsan@flysnow.org"},"管理员"}
	sayHello(ad.user)//使用user作为参数
	sayHello(ad)//使用admin作为参数
}
type Hello interface {
	hello()
}

func (u user) hello(){
	fmt.Println("Hello，i am a user")
}

func sayHello(h Hello){
	h.hello()
}

//这里就可以说明admin实现了接口Hello,但是我们又没有显示的声明类型 admin实现，所以这个实现是通过内部类型user实现的，
//因为admin包含了user所有的方法函数，所以也就实现了接口Hello。

//同名重写。同名继承