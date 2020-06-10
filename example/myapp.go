package main

import (
	"os"
	"sync"
	"time"

	"github.com/jiyu93/golog"
)

func main() {
	t1 := time.Now()
	rt := golog.NewRotater("app1.log", 1, 5, true)
	golog.SetDefaultOutput(rt)
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(x int) {
			golog.Info("xxx")
			wg.Done()
		}(i)
	}
	logger1 := golog.NewLogger(golog.NewRotater("app2.log", 1, 5, true), golog.LevelDebug)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(x int) {
			logger1.Info("yyy")
			wg.Done()
		}(i)
	}
	wg.Wait()
	golog.SetDefaultOutput(os.Stdout)
	golog.Info("time used:", time.Now().Sub(t1).Milliseconds(), "ms")
}
