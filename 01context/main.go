package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个耗时任务（比如下载文件）
func download(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 收到取消信号
			fmt.Println("任务被取消，停止下载")
			return
		default:
			fmt.Println("正在下载...")
			time.Sleep(1 * time.Second) // 模拟下载耗时
		}
	}
}

// 1、任务取消
//func main() {
//	// 创建一个可取消的 context，返回 ctx 和取消函数 cancel
//	ctx, cancel := context.WithCancel(context.Background())
//
//	// 启动下载任务（并发执行）
//	go download(ctx)
//
//	// 3秒后取消任务
//	time.Sleep(3 * time.Second)
//	cancel() // 发送取消信号
//
//	// 等待一下，看任务是否停止
//	time.Sleep(1 * time.Second)
//}

// 2、超时控制timeout
//func main() {
//	// 创建一个10秒后自动取消的 context
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel() // 确保最终会释放资源
//
//	go download(ctx) // 启动任务
//
//	// 等待任务结束（或超时）
//	select {
//	case <-ctx.Done():
//		fmt.Println("任务结束：", ctx.Err()) // 会打印超时原因
//	}
//}

// 3、传递共享信息
func main() {
	// 创建一个带键值对的 context
	ctx := context.WithValue(context.Background(), "userID", 1225)

	// 在另一个函数中获取值
	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	// 取出传递的 userID
	userID := ctx.Value("userID").(int) // 需要类型转换
	fmt.Println("当前用户ID：", userID)      // 输出：123
}
