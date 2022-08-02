package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
	"strings"
)

/*
@Time : 2022/7/28 17:58
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	filePath := "examples/test_data/level1_1.java"
	result, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		panic(err)
	}

	//lexer := lexers.Get("py")
	lexer := lexers.Match(filePath)

	// 词法分析
	tokenise, err := lexer.Tokenise(nil, result)
	if err != nil {
		panic(err)
	}
	//fmt.Println(tokenise.Tokens())
	var (
		s           string
		singleCount int // 单引号数量
		doubleCount int // 双引号数量
	)
	for _, token := range tokenise.Tokens() {
		// 获取类型字符串
		typeStr := token.Type.String()
		valueStr := strings.Trim(token.Value, " ")
		//valueStr = strings.Trim(token.Value, "\r")
		//valueStr = strings.Trim(token.Value, "\n")

		// 忽略空格和注释
		if typeStr == "Text" { // 忽略空格和注释
			s += ""
		} else if typeStr == "Keyword" { // 关键字
			s += valueStr
		} else if typeStr == "Name" { // 所有变量名称为“N”
			if valueStr != "" {
				s += "N"
			}
		} else if typeStr == "NameVariable" { // PHP中的变量类型
			if valueStr != "" {
				s += "N"
			}
		} else if typeStr == "LiteralString" { // 所有双引号都变成S
			// 双引号会连续出现3个：" 内容 "
			if doubleCount%3 == 0 {
				if valueStr != "" {
					s += "S"
				}
			}
			doubleCount++
		} else if typeStr == "LiteralStringDouble" { // PHP的双引号
			if valueStr != "" {
				s += "S"
			}
		} else if typeStr == "LiteralStringSingle" { // 所有单引号变成G
			// 在Python中这里会连续出现3个
			if strings.HasSuffix(filePath, ".py") && singleCount%3 == 0 {
				if valueStr != "" {
					s += "G"
				}
			}
			singleCount++
		} else if typeStr == "NameFunction" { // 用户定义的函数名为“F”
			if valueStr != "" {
				s += "F"
			}
		} else if typeStr == "Punctuation" { // 运算符
			s += valueStr
		} else if typeStr == "LiteralNumberInteger" { // TODO：将所有的数字都替换为1，需要确认是否影响准确度
			s += "1"
		} else if typeStr == "LiteralNumberFloat" { // TODO：将所有的浮点数都替换为1.0，需要确认是否影响准确度
			s += "1.0"
		} else {
			s += token.Value
		}
		fmt.Println("type = ", token.Type)
		fmt.Println("type category = ", token.Type.Category())
		fmt.Println("type category string = ", token.Type.Category().String())
		fmt.Println("type string = ", token.Type.String())
		fmt.Println("value = ", token.Value)
		fmt.Println("token string = ", token.String())
		fmt.Println("token go string = ", token.GoString())
		fmt.Println("==============")
	}
	fmt.Println(s)
}
