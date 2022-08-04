package main

import (
	"fmt"
	"io/ioutil"

	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
)

func main() {
	suffixes := []string{
		".java",
		".py",
		".php",
		".c",
		".cpp",
	}
	for _, suffix := range suffixes {
		filePath := "examples/test_data/level1_1" + suffix
		result, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		lexer := lexers.Match(filePath)
		token, err := zdpgo_pygments.GetToken(lexer, string(result))
		if err != nil {
			fmt.Println("获取token失败：", err)
			return
		}
		fmt.Println(token)
	}
}
