package wuqing

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Go 并发示例-Runner 执行者，它可以在后台执行任何任务

//一个执行者，可以执行任何任务，但是这些任务是限制完成的，
//该执行者可以通过发送终止信号终止它
type Runner struct {
	tasks     []func(int)      //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务在多久内完成
	interrupt chan os.Signal   //可以控制强制终止的信号
}

//complete是一个无缓冲通道，也就是同步通道,因为我们要使用它来控制我们整个程序是否终止，
//所以它必须是同步通道，要让main routine等待，一致要任务完成或者被强制终止。

//interrupt是一个有缓冲的通道，这样做是因为，我们可以至少接收到一个操作系统的中断信息，
//这样Go runtime在发送这个信号的时候不会被阻塞，如果是无缓冲的通道就会阻塞了。

//定义一个工厂函数New,用于返回我们需要的Runner
func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

//将需要执行的任务，添加到Runner里
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

//执行任务，执行的过程中接收到中断信号时，返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) run() error {
	//新增的run方法也很简单，会使用for循环，不停的运行任务，在运行的每个任务之前，都会检测是否收到了中断信号，
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//检查是否接收到了中断信号
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}


func MainRunner() {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := common.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err:=r.Start();err!=nil{
		switch err {
		case common.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case common.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

//这里注意isInterrupt函数，它在实现的时候，使用了基于select的多路复用，select和switch很像，
//只不过它的每个case都是一个通信操作。那么到底选择哪个case块执行呢？
//原则就是哪个case的通信操作可以执行就执行哪个，如果同时有多个可以执行的case，那么就随机选择一个执行





