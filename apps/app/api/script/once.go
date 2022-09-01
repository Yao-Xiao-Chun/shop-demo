package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
	"time"
)

//模拟一共有五个请求，其中只有一个请求会走到数据库 ，剩余相同的4个请求会等到结果，进行返回
func main() {
	const round = 5
	var wg sync.WaitGroup
	barrier := syncx.NewSingleFlight()

	wg.Add(round)
	for i := 0; i < round; i++ { //同时有五个线程执行 由于使用了限制，只会允许一个
		// 多个线程同时执行
		go func() {
			defer wg.Done()
			// 可以看到，多个线程在同一个key上去请求资源，获取资源的实际函数只会被调用一次
			val, err := barrier.Do("once", func() (interface{}, error) {
				// sleep 1秒，为了让多个线程同时取once这个key上的数据
				time.Sleep(time.Second)
				// 生成了一个随机的id
				return stringx.RandId(), nil
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}

	wg.Wait()
}
