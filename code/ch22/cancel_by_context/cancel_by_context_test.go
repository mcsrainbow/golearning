package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 判断上下文是否已取消
func isCancelled(ctx context.Context) bool {
	select {
	// 上下文已取消时返回 true
	case <-ctx.Done():
		return true
	// 上下文未取消时返回 false
	default:
		return false
	}
}

// 测试取消上下文
func TestCancel(t *testing.T) {
	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	// 启动 5 个 goroutine
	for i := 0; i < 5; i++ {
		// 启动异步操作
		go func(i int, ctx context.Context) {
			for {
				// 检查上下文是否已取消
				if isCancelled(ctx) {
					break
				}
				// 等待 5 毫秒后继续
				time.Sleep(time.Millisecond * 5)
			}
			// 打印 goroutine 取消信息
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	// 取消上下文
	cancel()
	// 等待 1 秒以确保所有 goroutine 打印完成
	time.Sleep(time.Second * 1)
}
