package channel

import (
	"fmt"
	"time"
)

// readandwrite :=make(chan int)
// onlyread := make(<-chan int)  //创建只读channel
// onlywrite := make(chan<- int) //创建只写channel
// 这段代码执行时会出现一个错误：fatal error: all goroutines are asleep - deadlock!
// 我们创建了一个无缓冲的channel，然后给这个channel赋值了，程序就是在赋值完成后陷入了死锁。
// 因为我们的channel是无缓冲的，即同步的，赋值完成后来不及读取channel，程序就已经阻塞了。
// 这里介绍一个非常重要的概念：channel的机制是先进先出，如果你给channel赋值了，那么必须要读取它的值，
// 不然就会造成阻塞，当然这个只对无缓冲的channel有效。对于有缓冲的channel，
// 发送方会一直阻塞直到数据被拷贝到缓冲区；如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复
func sample1() {
	// 读写channel 无缓冲
	ch := make(chan int)
	ch <- 1 // 在这边进入了死锁
	go func() {
		<-ch
		fmt.Println("1")
	}()
	fmt.Println("2")
}

func sample2() {
	// 读写channel 给channel增加缓冲区，然后在程序的最后让主线程休眠一秒
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("2")
}

func sample3() {
	// 读写channel 给channel增加缓冲区，然后在程序的最后让主线程休眠一秒
	ch := make(chan int)
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	ch <- 1
	fmt.Println("2")
}

// produce只写
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

// consumer只读
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

// server 不带缓冲的例子
// 这段代码中channel是没有缓冲的，所以当生产者给channel赋值后，生产者这个线程会阻塞，
// 直到消费者线程将channel中的数据取出。消费者第一次将数据取出后，进行下一次循环时，消费者的线程也会阻塞，
// 因为生产者还没有将数据存入，这时程序会去执行生产者的线程。程序就这样在消费者和生产者两个线程间不断切换，
// 直到循环结束
func server() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

// 带缓冲的例子
// 在这个程序中，缓冲区可以存储10个int类型的整数，
// 在执行生产者线程的时候，线程就不会阻塞，一次性将10个整数存入channel，在读取的时候，也是一次性读取
func server2() {
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

func sample4() {
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("Ping!")
	fmt.Println(<-messages) // 取消息
}

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	<-complete // 执行完毕了，发个消息
}

func sample5() {
	go loop()
	complete <- 0 // 直到线程跑完, 取到消息. main在此阻塞住
}

func sample6() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	for v := range ch {
		fmt.Println(v)
		if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
			break
		}
	}
}

func sample7() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// 显式地关闭信道
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// 这是一个模型，开出很多小goroutine, 它们各自跑各自的，最后跑完了向主线报告。
// 方案1：只使用单个无缓冲信道阻塞主线
var quit chan int // 只开一个信道

func foo(id int) {
	fmt.Println(id)
	quit <- 0 // ok, finished
}

func sample8() {
	count := 100
	quit = make(chan int) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		<-quit
	}
}
// 方案2 把信道换成缓冲1000的
func sample9() {
	count := 10
	quit = make(chan int,10)

	for i := 0; i < count; i++ {
		go foo(i)
	}
	for i := 0; i < count; i++ {
		<- quit
	}
}