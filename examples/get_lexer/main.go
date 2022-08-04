package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
)

func TestLexer(t *testing.T) {
	// 读取文件内容
	filename := "examples/test_data/level1_1.py"
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	// 词法分析
	lexer := lexers.Match(filename)

	// 获取token
	tokens, err := lexer.Tokenise(nil, string(text))
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	// 遍历
	for _, token := range tokens.Tokens() {
		fmt.Println(token)
	}
}
