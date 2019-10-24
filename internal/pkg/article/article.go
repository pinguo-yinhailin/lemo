package article

import (
	"context"
	"fmt"
	"sync"
)

type Article interface {
	Sync() error
	Push(ctx context.Context, streamMsg chan<- string, max int) error
}

type article struct {
}

func New() Article {
	return &article{}
}

func (a *article) Sync() error {
	// 执行文章同步的通用服务
	return nil
}

// Push 并发推送文章到用户设备
func (a *article) Push(ctx context.Context, streamMsg chan<- string, max int) error {
	var wait sync.WaitGroup

	for i := 1; i <= max; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			a.PushOne(streamMsg, i)
		}(i)
	}

	wait.Wait()
	return nil
}

func (a *article) PushOne(streamMsg chan<- string, id int) {
	streamMsg <- fmt.Sprintf("推送文章到设备 %v", id)
}
