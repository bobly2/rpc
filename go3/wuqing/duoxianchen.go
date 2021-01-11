package wuqing

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*进程	一个程序对应一个独立程序空间
线程	一个执行空间，一个进程可以有多个线程
逻辑处理器	执行创建的goroutine，绑定一个线程
调度器	Go运行时中的，分配goroutine给不同的逻辑处理器
全局运行队列	所有刚创建的goroutine都会放到这里
本地运行队列	逻辑处理器的goroutine队列*/

//当我们创建一个goroutine的后，会先存放在全局运行队列中，等待Go运行时的调度器进行调度，
//把他们分配给其中的一个逻辑处理器，并放到这个逻辑处理器对应的本地运行队列中，最终等着被逻辑处理器执行即可。

//这一套管理、调度、执行goroutine的方式称之为Go的并发

//么Go的并行是怎样的呢？其实答案非常简单，多创建一个逻辑处理器就好了，
//这样调度器就可以同时分配全局运行队列中的goroutine到不同的逻辑处理器上并行执行

func MainDXC() {
	//sync.WaitGroup其实是一个计数的信号量，使用它的目的是要main函数等待两个goroutine执行完成后再结束，
	//不然这两个goroutine还在运行的时候，程序就结束了

	//sync.WaitGroup的使用也非常简单，先是使用Add 方法设设置计算器为2，每一个goroutine的函数执行完之后，就调用Done方法减1。
	//Wait方法的意思是如果计数器大于0，就会阻塞，所以main 函数会一直等待2个goroutine完成后，再结束。
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}

//默认情况下，Go默认是给每个可用的物理处理器都分配一个逻辑处理器，因为我的电脑是4核的，所以上面的例子默认创建了4个逻辑处理器，
//所以这个例子中同时也有并行的调度，如果我们强制只使用一个逻辑处理器，我们再看看结果。

//如果需要设置的话，一般我们采用如下代码设置。
func MainDXC2() {
	//设置一个逻辑处理器
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10000; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 10000; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}

//结果依然是并发

//有并发，就有资源竞争,
//如果两个或者多个goroutine在没有相互同步的情况下，访问某个共享的资源，比如同时对该资源进行读写时，就会处于相互竞争的状态，
//这就是并发中的资源竞争。

var (
	count int32
	wg    sync.WaitGroup
)

func MainJZ() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		//runtime.Gosched()是让当前goroutine暂停的意思，退回执行队列，让其他等待的goroutine运行，目的是让我们演示资源竞争的结果更明显

		value++
		count = value
	}
}

//多运行几次这个程序，会发现结果可能是2，也可以是3，也可能是4

//们对于同一个资源的读写必须是原子化的，也就是说，同一时间只能有一个goroutine对共享资源进行读写操作。

//共享资源竞争的问题，非常复杂
//Go为我们提供了一个工具帮助我们检查，这个就是go build -race命令。我们在当前项目目录下执行这个命令，生成一个可以执行文件，然后再运行这个可执行文件，就可以看到打印出的检测信息。

func MainYuanzi() {
	wg.Add(2)
	go incCount2()
	go incCount2()
	wg.Wait()
	fmt.Println(count)
}

func incCount2() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}

//atomic.LoadInt32,  读取int32类型变量的值，
//atomic.StoreInt32   修改int32类型变量的值，这两个都是原子性的操作

func MainSuo() {
	wg.Add(2)
	go incCount3()
	go incCount3()
	wg.Wait()
	fmt.Println(count)
}

var (
	mutex sync.Mutex
)

func incCount3() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}
//被sync互斥锁控制的这段代码范围，被称之为临界区，临界区的代码，同一时间，只能又一个goroutine访问
//新声明了一个互斥锁mutex sync.Mutex，这个互斥锁有两个方法，一个是mutex.Lock(),一个是mutex.Unlock(),这两个之间的区域就是临界区，临界区的代码是安全的。