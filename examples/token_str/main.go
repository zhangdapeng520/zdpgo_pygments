package main

import (
	"fmt"
	"io/ioutil"

	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_pygments"
)

func main() {
	suffixes := []string{
		//".java",
		// ".py",
		".php",
		// ".c",
		// ".cpp",
	}
	for _, suffix := range suffixes {
		filePath := "examples/test_data/level1_1" + suffix
		result, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		lexer := zdpgo_lexers.Match(filePath)
		token, err := zdpgo_pygments.GetToken(lexer, string(result))

		s := "<?php $name = $email = $gender = $comment = $website = ''; ?>"
		token, err = zdpgo_pygments.GetToken(lexer, s)
		if err != nil {
			fmt.Println("获取token失败：", err)
			return
		}
		fmt.Println(token)
	}
}
