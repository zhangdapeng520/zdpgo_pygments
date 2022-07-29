package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
)

/*
@Time : 2022/7/28 17:58
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	filePath := "examples/test_data/level1_1.py"
	result, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		panic(err)
	}

	//lexer := lexers.Get("py")
	lexer := lexers.Match(filePath)
	tokenise, err := lexer.Tokenise(nil, result)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenise.Tokens())
}
