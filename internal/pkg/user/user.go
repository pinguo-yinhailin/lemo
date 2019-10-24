package user

import (
	"context"
	"fmt"
	"time"

	"github.com/pinguo-yinhailin/lemo/pkg/name"
	"github.com/spf13/viper"
)

type User interface {
	Run(ctx context.Context, config *viper.Viper)
	Create(ctx context.Context, quantity int) error
	C() <-chan *Person
}

type user struct {
	config         *viper.Viper
	streamTodo     chan<- *Person
	streamFinished <-chan *Person
}

func New() User {
	return &user{}
}

func (u *user) Run(ctx context.Context, config *viper.Viper) {
	u.config = config

	stream := make(chan *Person)
	u.streamTodo = stream

	u.streamFinished = u.createDocument(ctx, u.createAppearance(ctx, stream))

	go func() {
		defer close(stream)
		<-ctx.Done()
	}()
}

func (u *user) C() <-chan *Person {
	return u.streamFinished
}

func (u *user) Create(ctx context.Context, quantity int) error {
	if u.streamTodo == nil || u.streamFinished == nil {
		return fmt.Errorf("function Run should be executed before Create")
	}

	names := u.createName(quantity)
	for _, v := range names {
		p := &Person{Name: v}

		select {
		case <-ctx.Done():
			return nil
		case u.streamTodo <- p:
		}
	}
	return nil
}

func (u *user) createName(quantity int) []string {
	var res []string

	n, err := name.New(u.config.GetString("nameFile"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < quantity; i++ {
		res = append(res, n.Generate())
	}
	return res
}

// createAppearance 运行外观创建服务
// 这里假设相关的程序最大并发量为3
func (u *user) createAppearance(ctx context.Context, person <-chan *Person) <-chan *Person {
	streamRes := make(chan *Person)

	limitFunc := func(limit <-chan bool) {
		defer func() {
			<-limit
		}()

		select {
		case <-ctx.Done():
			return
		case p := <-person:
			fmt.Printf("[1]生成外观：%s\n", p.Name)
			p.HasAppearance = true
			time.Sleep(time.Second * 3)

			select {
			case <-ctx.Done():
				return
			case streamRes <- p:
			}
		}
	}

	streamLimit := make(chan bool, 3)
	go func() {
		defer close(streamLimit)
		defer close(streamRes)

		for {
			streamLimit <- true
			go limitFunc(streamLimit)
		}
	}()

	return streamRes
}

// createDocument 运行档案创建服务
// 这里假设相关的程序只允许串行执行
func (u *user) createDocument(ctx context.Context, person <-chan *Person) <-chan *Person {
	stream := make(chan *Person)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case p := <-person:
				fmt.Printf("[2]创建档案：%s\n", p.Name)
				p.HasDocument = true
				time.Sleep(time.Second)

				select {
				case <-ctx.Done():
					return
				case stream <- p:
				}
			}
		}
	}()

	return stream
}

type Person struct {
	Name          string
	HasAppearance bool
	HasDocument   bool
}
