package zdpgo_pygments

import (
	"strings"
)

// GetToken 使用指定的词法分析器和代码内容获取token
func GetToken(lexer Lexer, content string) (string, error) {
	var (
		result      string
		singleCount int // 单引号数量
		doubleCount int // 双引号数量
	)

	// 词法分析
	tokenise, err := lexer.Tokenise(nil, content)
	if err != nil {
		return "", err
	}

	for _, token := range tokenise.Tokens() {
		// 获取类型字符串
		typeStr := token.Type.String()
		valueStr := strings.TrimSpace(token.Value)

		// 忽略空格和注释
		if typeStr == "Text" || typeStr == "CommentSingle" || typeStr == "CommentMultiline" { // 忽略空格和注释
			result += ""
		} else if typeStr == "Keyword" { // 关键字
			result += valueStr
		} else if typeStr == "NameClass" { // Java：所有的类名改为C
			result += "C"
		} else if typeStr == "Name" { // 所有变量名称为“N”
			if valueStr != "" {
				result += "N"
			}
		} else if typeStr == "NameAttribute" { // Java中的对象属性，全部转换为A
			if valueStr != "" {
				result += "A"
			}
		} else if typeStr == "NameVariable" { // PHP中的变量类型
			if valueStr != "" {
				result += "N"
			}
		} else if typeStr == "LiteralString" { // 所有双引号都变成S
			// 双引号会连续出现3个：" 内容 "
			if doubleCount%3 == 0 {
				if valueStr != "" {
					result += "S"
				}
			}
			doubleCount++
		} else if typeStr == "LiteralStringDouble" { // PHP/Python的双引号
			if valueStr != "" {
				if lexer.Config().Name == "Python" {
					// Python双引号会连续出现3个：" 内容 "
					if doubleCount%3 == 0 {
						result += "S"
						doubleCount = 0 // 重置，防止该数过大
					}
					doubleCount++
				} else {
					result += "S"
				}
			}
		} else if typeStr == "LiteralStringSingle" { // 所有单引号变成G
			// 在Python中这里会连续出现3个
			if singleCount%3 == 0 {
				if valueStr != "" {
					result += "G"
				}
			}
			singleCount++
		} else if typeStr == "NameFunction" { // 用户定义的函数名为“F”
			if valueStr != "" {
				result += "F"
			}
		} else if typeStr == "Punctuation" { // 运算符
			result += valueStr
		} else if typeStr == "LiteralNumberInteger" { // 将所有的数字都替换为1，需要确认是否影响准确度
			result += "1"
		} else if typeStr == "LiteralNumberFloat" { // 将所有的浮点数都替换为1.0，需要确认是否影响准确度
			result += "1.0"
		} else {
			result += valueStr
		}
	}
	return result, nil
}

// GetTokenArr 获取token数组
func GetTokenArr(lexer Lexer, contents []string) ([]string, error) {
	var results []string
	for _, content := range contents {
		token, err := GetToken(lexer, content)
		if err != nil {
			return nil, err
		}
		results = append(results, token)
	}
	return results, nil
}

// 将token数组按照指定的数量展开
func GetSpreadTokenArr(tokens []string, lines int) []string {
	var result []string

	// 特殊情况1：行数大于或等于token总数
	if lines >= len(tokens) {
		token := strings.Join(tokens, "")
		result = append(result, token)
		return result
	}

	// 特殊情况2：行数小于或等于0
	if lines <= 0 {
		return tokens
	}

	// 按照指定行数展开token，并合并
	for i := 0; i <= len(tokens)-lines; i++ {
		token := strings.Join(tokens[i:i+lines], "")
		result = append(result, token)
	}
	return result
}
