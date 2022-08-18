package main

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_pygments"
)

func main() {
	suffixes := []string{
		//".java",
		//".py",
		".php",
		//".c",
		//".cpp",
	}
	for _, suffix := range suffixes {
		filePath := "examples/test_data/level1_1" + suffix
		result, err := zdpgo_clearcode.ClearCode(filePath)
		if err != nil {
			panic(err)
		}

		lexer := zdpgo_lexers.Match(filePath)
		contents := strings.Split(string(result), "\n")
		tokenArr, err := zdpgo_pygments.GetTokenArr(lexer, contents)

		if err != nil {
			fmt.Println("获取token失败：", err)
			return
		}

		tokenS := strings.Join(tokenArr, "")
		fmt.Println(tokenS)
	}
}
