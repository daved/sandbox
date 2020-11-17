package main

import (
	"context"
	"fmt"
	"time"

	"github.com/codemodus/sigmon/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan sigmon.Signal, 1)

	sm := sigmon.New(nil)
	sm.Set(func(s *sigmon.State) {
		defer close(c)
		cancel()
		c <- s.Signal()
		sm.Set(nil)
	})
	sm.Start()

	fmt.Println("bgn main")

	var h Handler = &end{tag: "endpoint"}
	h = outer(inner(h))

	sleep()
	fmt.Println("err:", run(h, ctx))

	sm.Stop()
	fmt.Println("end main")
	if sig, ok := <-c; ok {
		fmt.Printf("received sig %q\n", sig)
	}
}

type end struct {
	tag string
}

func (e *end) Do(ctx context.Context) error {
	fmt.Println("   bgn end.do")
	defer fmt.Println("   end end.do")

	select {
	case <-ctx.Done():
		fmt.Println("   canceled before end.do")
		return ctx.Err()
	default:
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	sm := sigmon.New(nil)
	sm.Set(func(s *sigmon.State) {
		cancel()
		fmt.Println("   signal in handler")
		sm.Set(nil)
	})
	sm.Start()
	defer sm.Stop()

	sleep()
	select {
	case <-ctx.Done():
		fmt.Println("   canceled in end.do")
		return ctx.Err()
	default:
	}

	fmt.Println("  ", e.tag)
	return nil
}

func inner(next Handler) Handler {
	return HandleFunc(func(ctx context.Context) error {
		fmt.Println("  bgn inner")
		defer fmt.Println("  end inner")

		select {
		case <-ctx.Done():
			fmt.Println("  canceled before inner")
			return ctx.Err()
		default:
		}

		sleep()
		return next.Do(ctx)
	})
}

func outer(next Handler) Handler {
	return HandleFunc(func(ctx context.Context) error {
		fmt.Println(" bgn outer")
		defer fmt.Println(" end outer")

		select {
		case <-ctx.Done():
			fmt.Println(" canceled before outer")
			return ctx.Err()
		default:
		}

		sleep()
		return next.Do(ctx)
	})
}

func sleep() {
	time.Sleep(time.Second * 6)
}
