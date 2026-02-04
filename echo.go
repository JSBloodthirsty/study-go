package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var path string
	path = "asdsad"
	fmt.Println(path)
	// 获取命令名（去掉路径）
	cmdName := path.Base(os.Args[0])
	fmt.Println("命令名:", cmdName)

	// 打印每个参数的索引和值
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("参数 %d: %s\n", i, os.Args[i])
	}

	// 测量潜在低效的版本和使用了strings.Join的版本的运行时间差异
	// 潜在低效的版本
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		s, sep := "", " "
		for _, arg := range os.Args[1:] {
			s += sep + arg
		}
	}
	fmt.Println("潜在低效版本运行时间:", time.Since(startTime))

	// 使用了strings.Join的版本
	startTime = time.Now()
	for i := 0; i < 100000; i++ {
		strings.Join(os.Args[1:], " ")
	}
	fmt.Println("使用strings.Join版本运行时间:", time.Since(startTime))
}
