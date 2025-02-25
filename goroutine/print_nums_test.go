package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// 题目来源：https://www.v2ex.com/t/1113786
// 考察通道的掌握情况
func TestPrintNums(t *testing.T) {
	ch := make(chan int)
	group := sync.WaitGroup{}
	group.Add(2)

	go func(ch chan int) {
		defer group.Done()
		for i := range 5 {
			ch <- i
		}
		// 考察点：应该由生产者关闭通道，否则向已关闭的通道发送数据会引发panic
		close(ch)
	}(ch)
	go func(ch chan int) {
		// 考察点：用defer保证函数结束时调用
		defer group.Done()
		// 考察点：不能用fori读5次，因为已关闭的通道能继续读出零值
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)

	// 考察点：需等待协程执行完成才能退出程序
	group.Wait()
}
