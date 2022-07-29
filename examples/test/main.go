package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建字典
	counts := make(map[string]int)

	// 读取标准输入
	input := bufio.NewScanner(os.Stdin)

	// 遍历标准输入，添加到字典
	for input.Scan() {
		counts[input.Text()]++
	}

	// 遍历并打印每行，自动过滤重复行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
