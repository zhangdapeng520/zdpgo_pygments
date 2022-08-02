package zdpgo_pygments

// Coalesce is a Lexer interceptor that collapses runs of common types into a single token.
func Coalesce(lexer Lexer) Lexer { return &coalescer{lexer} }

type coalescer struct{ Lexer }

func (d *coalescer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	var prev Token
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	return func() Token {
		for token := it(); token != (EOF); token = it() {
			if len(token.Value) == 0 {
				continue
			}
			if prev == EOF {
				prev = token
			} else {
				if prev.Type == token.Type && len(prev.Value) < 8192 {
					prev.Value += token.Value
				} else {
					out := prev
					prev = token
					return out
				}
			}
		}
		out := prev
		prev = EOF
		return out
	}, nil
}

//
//// Tokenize 将文件内容token化，得到一个hash元组
//func Tokenize(filename string) (string, int, int, error) {
//	// 读取文件内容
//	text, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return "", 0, 0, err
//	}
//
//	// 词法分析
//	lexer := lexers.Match(filename)
//
//	// 获取token
//	tokens, err := lexer.Tokenise(nil, string(text))
//	if err != nil {
//		return "", 0, 0, err
//	}
//
//	// 处理tokens并返回
//	return handle_tokens(tokens)
//}
//
//// 处理tokens
//func HandleTokens(tokens []Token)(string, int, int){
//	tokenLength := len(tokens)
//	count1 := 0  // 标记来存储原始代码文件中每个元素的对应位置
//	count2 := 0 // 标记来存储清理后的代码文本中每个元素的位置
//
//	// 这些标记用于标记原始代码文件中抄袭的内容。
//	for i := 0; i < tokenLength; i++ {
//		if (tokens[i][0]==Name)
//		if tokens[i][0] == token.Name and not i == token_length - 1 and not tokens[i + 1][1] == '(':
//		result.append(('N', count1, count2))  # 所有变量名称为“N”
//		count2 += 1
//		elif tokens[i][0] in token.Literal.String:
//		result.append(('S', count1, count2))  # 所有字符串为'S'
//		count2 += 1
//		elif tokens[i][0] in token.Name.Function:
//		result.append(('F', count1, count2))  # 用户定义的函数名为“F”
//		count2 += 1
//		elif tokens[i][0] == token.Text or tokens[i][0] in token.Comment:
//		pass  # 忽略了空白和注释
//		else:
//		result.append((tokens[i][1], count1, count2))
//		# 结果元组(每个元素，例如'def'，它在原始代码文件中的位置，在清理后的代码/文本中的位置)
//		count2 += len(tokens[i][1])
//		count1 += len(tokens[i][1])
//	}
//
//	return result
//}
