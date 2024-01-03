// context中要重点关注Done方法和cancelFunc方法，Done方法负责接受上游结束的通知，cancelFunc方法负责通知下游结束
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	// 创建截止时间
	d := time.Now().Add(shortDuration)
	// 创建有截止时间的context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	// 使用select监听1s和有截止时间的context哪个先结束
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
