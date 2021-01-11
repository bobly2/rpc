package wuqing

import "fmt"

//使用通道，在多个goroutine发送和接受共享的数据，达到数据同步的目的。
//在两个routine之间架设的管道，一个goroutine可以往这个管道里塞数据，另外一个可以从这个管道里取数据，有点类似于我们说的队列
func MainTongDao() {
	//声明一个通道
	ch := make(chan int)
	ch <- 2   //发送数值2给这个通道
	x := <-ch //从通道里读取值，并把读取的值赋值给x变量
	<-ch      //从通道里读取值，然后忽略

	//关闭。关闭后，不能推送数据，但还能接收数据
	close(ch)

	//第二个参数，指定通道的大小，默认没有第二个参数的时候，通道的大小为0，这种通道也被成为无缓冲通道。
	ch0 := make(chan int, 0)
	ch2 := make(chan int, 2)

	//无缓冲的通道，发送goroutine和接收gouroutine必须是同步的，同时准备后，
	//如果没有同时准备好的话，先执行的操作就会阻塞等待，直到另一个相对应的操作准备好为止。这种无缓冲的通道我们也称之为同步通道

}

func MainTestC() {
	ch := make(chan int)

	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()

	fmt.Println(<-ch)

}

//在计算sum和的goroutine没有执行完，把值赋给ch通道之前，fmt.Println(<-ch)会一直等待，
//所以main主goroutine就不会终止，只有当计算和的goroutine完成后，并且发送到ch通道的操作准备好后，
//同时<-ch就会接收计算好的值，然后打印出来。

func mainAA() {
	one := make(chan int)
	two := make(chan int)

	go func() {
		one <- 100
	}()

	go func() {
		v := <-one
		two <- v
	}()

	fmt.Println(<-two)

}

//有缓冲通道，其实是一个队列
//当队列满的时候，发送操作会阻塞；当队列空的时候，接受操作会阻塞
//cap(ch)返回通道的最大容量
//len(ch)返回现在通道里有几个元素

func mirroredQuery() string {

	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}
func request(hostname string) (response string) { /* ... */ return }

//我们定义了一个容量为3的通道responses，然后同时发起3个并发goroutine向这三个镜像获取数据，
//获取到的数据发送到通道responses中，最后我们使用return <-responses返回获取到的第一个数据，
//也就是最快返回的那个镜像的数据。

//单向通道
var send chan<- int    //只能发送
var receive <-chan int //只能接收


