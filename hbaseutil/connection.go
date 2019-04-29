package hbaseutil

import (
	"github.com/tsuna/gohbase"
	"fmt"
	"sync"
)

var (
	Client gohbase.Client
	once sync.Once
)

func init() {
	once.Do(func () {
		fmt.Println("创建连接")
		Client = gohbase.NewClient("116.62.156.102")
	})
}