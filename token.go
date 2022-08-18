package zdpgo_pygments

import (
	"bytes"
	"strings"
)

var (
	removeOperator = map[string]bool{";": true, ",": true, ".": true}
)

// GetToken 使用指定的词法分析器和代码内容获取token
func GetToken(lexer Lexer, content string) (string, error) {
	// 词法分析
	tokenise, err := lexer.Tokenise(nil, content)
	if err != nil {
		return "", err
	}

	// 处理token
	var buf bytes.Buffer

	for _, token := range tokenise.Tokens() {
		// 获取类型字符串
		typeStr := token.Type.String()
		valueStr := strings.TrimSpace(token.Value)

		// 忽略空格和注释
		if typeStr == "CommentSingle" || // 忽略单行注释
			typeStr == "CommentMultiline" || // 忽略多行注释
			typeStr == "Text" || // 忽略无法解析的文本
			typeStr == "NameNamespace" || // 忽略名称空间
			typeStr == "KeywordNamespace" || // Java名称空间关键字
			valueStr == "" { // 忽略空字符串
			continue
		} else if typeStr == "NameDecorator" { // 注解
			buf.WriteString("@Z")
		} else if typeStr == "Keyword" ||
			typeStr == "KeywordDeclaration" { // 关键字
			buf.WriteString("K")
		} else if typeStr == "KeywordType" { // Java类型关键字
			buf.WriteString("T")
		} else if typeStr == "Operator" { // 运算符
			if _, ok := removeOperator[valueStr]; !ok {
				buf.WriteString(valueStr)
			}
		} else if typeStr == "NameClass" { // Java：所有的类名改为C
			buf.WriteString("C")
		} else if typeStr == "NameOther" { // 所有其他变量名称为“O”
			buf.WriteString("O")
		} else if typeStr == "Name" ||
			typeStr == "NameVariable" { // 所有变量名称为“N”
			buf.WriteString("N")
		} else if typeStr == "NameAttribute" { // Java中的对象属性，全部转换为A
			buf.WriteString("A")
		} else if typeStr == "LiteralString" || // 所有双引号都变成D
			typeStr == "LiteralStringDouble" {
			buf.WriteString("D")
		} else if typeStr == "LiteralStringSingle" { // 所有单引号变成S
			buf.WriteString("S")
		} else if typeStr == "NameFunction" { // 用户定义的函数名为“F”
			buf.WriteString("F")
		} else if typeStr == "Punctuation" { // 运算符
			buf.WriteString(valueStr)
		} else if typeStr == "LiteralNumberInteger" { // 将所有的数字都替换为1，需要确认是否影响准确度
			buf.WriteString("1")
		} else if typeStr == "LiteralNumberFloat" { // 将所有的浮点数都替换为2，需要确认是否影响准确度
			buf.WriteString("2")
		} else {
			buf.WriteString(valueStr)
		}
	}
	return buf.String(), nil
}

// GetTokenArr 获取token数组
func GetTokenArr(lexer Lexer, contents []string) ([]string, error) {
	var results []string
	for _, content := range contents {
		// 处理PHP，解决无法获取token的问题
		isHandlePHP := lexer.Config().Name == "PHTML" && !strings.Contains(content, "<?php")
		if isHandlePHP {
			content = "<?php " + content + " ?>"
		}
		token, err := GetToken(lexer, content)
		if err != nil {
			return nil, err
		}

		// 处理token
		if isHandlePHP {
			token = strings.Replace(token, "<?php", "", 1)
			token = strings.Replace(token, "?>", "", 1)
		}

		// token不为空，才追加
		if token != "" {
			results = append(results, token)
		}
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
