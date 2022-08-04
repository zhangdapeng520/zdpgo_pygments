package zdpgo_pygments

import (
	"encoding/json"
	"fmt"
)

// TokenType 要高亮的token的类型
// 同时也是一个 Emitter 发射器, 发射一个token字符串信号，这个token字符串就是它自己
type TokenType int

// MarshalJSON 实现JSON序列化功能
func (t TokenType) MarshalJSON() ([]byte, error) { return json.Marshal(t.String()) }

// UnmarshalJSON 实现JSON反序列化功能
func (t *TokenType) UnmarshalJSON(data []byte) error {
	key := ""
	err := json.Unmarshal(data, &key)
	if err != nil {
		return err
	}
	for tt, text := range _TokenType_map {
		if text == key {
			*t = tt
			return nil
		}
	}
	return fmt.Errorf("unknown TokenType %q", data)
}

// token类型集合
//
// Categories of types are grouped in ranges of 1000, while sub-categories are in ranges of 100. For
// example, the literal category is in the range 3000-3999. The sub-category for literal strings is
// in the range 3100-3199.

// token元类型
const (
	Background       TokenType = -1 - iota // 默认背景样式
	PreWrapper                             // 预包装样式
	Line                                   // 行样式
	LineNumbers                            // 行号
	LineNumbersTable                       // 表格行号
	LineHighlight                          // 行高亮样式
	LineTable                              // 表格行包裹样式
	LineTableTD                            // 表格行TD样式
	CodeLine                               // 代码行样式
	Error                                  // 错误
	Other                                  // 其他
	None                                   // 不高亮
	EOFType          TokenType = 0         // 结束
)

// 关键字
const (
	Keyword TokenType = 1000 + iota
	KeywordConstant
	KeywordDeclaration
	KeywordNamespace
	KeywordPseudo
	KeywordReserved
	KeywordType
)

// 名称
const (
	Name TokenType = 2000 + iota
	NameAttribute
	NameBuiltin
	NameBuiltinPseudo
	NameClass
	NameConstant
	NameDecorator
	NameEntity
	NameException
	NameFunction
	NameFunctionMagic
	NameKeyword
	NameLabel
	NameNamespace
	NameOperator
	NameOther
	NamePseudo
	NameProperty
	NameTag
	NameVariable
	NameVariableAnonymous
	NameVariableClass
	NameVariableGlobal
	NameVariableInstance
	NameVariableMagic
)

// Literals.
const (
	Literal TokenType = 3000 + iota
	LiteralDate
	LiteralOther
)

// Strings.
const (
	LiteralString TokenType = 3100 + iota
	LiteralStringAffix
	LiteralStringAtom
	LiteralStringBacktick
	LiteralStringBoolean
	LiteralStringChar
	LiteralStringDelimiter
	LiteralStringDoc
	LiteralStringDouble
	LiteralStringEscape
	LiteralStringHeredoc
	LiteralStringInterpol
	LiteralStringName
	LiteralStringOther
	LiteralStringRegex
	LiteralStringSingle
	LiteralStringSymbol
)

// Literals.
const (
	LiteralNumber TokenType = 3200 + iota
	LiteralNumberBin
	LiteralNumberFloat
	LiteralNumberHex
	LiteralNumberInteger
	LiteralNumberIntegerLong
	LiteralNumberOct
)

// Operators.
const (
	Operator TokenType = 4000 + iota
	OperatorWord
)

// Punctuation.
const (
	Punctuation TokenType = 5000 + iota
)

// Comments.
const (
	Comment TokenType = 6000 + iota
	CommentHashbang
	CommentMultiline
	CommentSingle
	CommentSpecial
)

// Preprocessor "comments".
const (
	CommentPreproc TokenType = 6100 + iota
	CommentPreprocFile
)

// Generic tokens.
const (
	Generic TokenType = 7000 + iota
	GenericDeleted
	GenericEmph
	GenericError
	GenericHeading
	GenericInserted
	GenericOutput
	GenericPrompt
	GenericStrong
	GenericSubheading
	GenericTraceback
	GenericUnderline
)

// Text.
const (
	Text TokenType = 8000 + iota
	TextWhitespace
	TextSymbol
	TextPunctuation
)

// Aliases.
const (
	Whitespace = TextWhitespace

	Date = LiteralDate

	String          = LiteralString
	StringAffix     = LiteralStringAffix
	StringBacktick  = LiteralStringBacktick
	StringChar      = LiteralStringChar
	StringDelimiter = LiteralStringDelimiter
	StringDoc       = LiteralStringDoc
	StringDouble    = LiteralStringDouble
	StringEscape    = LiteralStringEscape
	StringHeredoc   = LiteralStringHeredoc
	StringInterpol  = LiteralStringInterpol
	StringOther     = LiteralStringOther
	StringRegex     = LiteralStringRegex
	StringSingle    = LiteralStringSingle
	StringSymbol    = LiteralStringSymbol

	Number            = LiteralNumber
	NumberBin         = LiteralNumberBin
	NumberFloat       = LiteralNumberFloat
	NumberHex         = LiteralNumberHex
	NumberInteger     = LiteralNumberInteger
	NumberIntegerLong = LiteralNumberIntegerLong
	NumberOct         = LiteralNumberOct
)

var (
	StandardTypes = map[TokenType]string{
		Background:       "bg",
		PreWrapper:       "zdpgo_pygments",
		Line:             "line",
		LineNumbers:      "ln",
		LineNumbersTable: "lnt",
		LineHighlight:    "hl",
		LineTable:        "lntable",
		LineTableTD:      "lntd",
		CodeLine:         "cl",
		Text:             "",
		Whitespace:       "w",
		Error:            "err",
		Other:            "x",
		// I have no idea what this is used for...
		// Escape:     "esc",

		Keyword:            "k",
		KeywordConstant:    "kc",
		KeywordDeclaration: "kd",
		KeywordNamespace:   "kn",
		KeywordPseudo:      "kp",
		KeywordReserved:    "kr",
		KeywordType:        "kt",

		Name:                 "n",
		NameAttribute:        "na",
		NameBuiltin:          "nb",
		NameBuiltinPseudo:    "bp",
		NameClass:            "nc",
		NameConstant:         "no",
		NameDecorator:        "nd",
		NameEntity:           "ni",
		NameException:        "ne",
		NameFunction:         "nf",
		NameFunctionMagic:    "fm",
		NameProperty:         "py",
		NameLabel:            "nl",
		NameNamespace:        "nn",
		NameOther:            "nx",
		NameTag:              "nt",
		NameVariable:         "nv",
		NameVariableClass:    "vc",
		NameVariableGlobal:   "vg",
		NameVariableInstance: "vi",
		NameVariableMagic:    "vm",

		Literal:     "l",
		LiteralDate: "ld",

		String:          "s",
		StringAffix:     "sa",
		StringBacktick:  "sb",
		StringChar:      "sc",
		StringDelimiter: "dl",
		StringDoc:       "sd",
		StringDouble:    "s2",
		StringEscape:    "se",
		StringHeredoc:   "sh",
		StringInterpol:  "si",
		StringOther:     "sx",
		StringRegex:     "sr",
		StringSingle:    "s1",
		StringSymbol:    "ss",

		Number:            "m",
		NumberBin:         "mb",
		NumberFloat:       "mf",
		NumberHex:         "mh",
		NumberInteger:     "mi",
		NumberIntegerLong: "il",
		NumberOct:         "mo",

		Operator:     "o",
		OperatorWord: "ow",

		Punctuation: "p",

		Comment:            "c",
		CommentHashbang:    "ch",
		CommentMultiline:   "cm",
		CommentPreproc:     "cp",
		CommentPreprocFile: "cpf",
		CommentSingle:      "c1",
		CommentSpecial:     "cs",

		Generic:           "g",
		GenericDeleted:    "gd",
		GenericEmph:       "ge",
		GenericError:      "gr",
		GenericHeading:    "gh",
		GenericInserted:   "gi",
		GenericOutput:     "go",
		GenericPrompt:     "gp",
		GenericStrong:     "gs",
		GenericSubheading: "gu",
		GenericTraceback:  "gt",
		GenericUnderline:  "gl",
	}
)

func (t TokenType) Parent() TokenType {
	if t%100 != 0 {
		return t / 100 * 100
	}
	if t%1000 != 0 {
		return t / 1000 * 1000
	}
	return 0
}

func (t TokenType) Category() TokenType {
	return t / 1000 * 1000
}

func (t TokenType) SubCategory() TokenType {
	return t / 100 * 100
}

func (t TokenType) InCategory(other TokenType) bool {
	return t/1000 == other/1000
}

func (t TokenType) InSubCategory(other TokenType) bool {
	return t/100 == other/100
}

func (t TokenType) Emit(groups []string, _ *LexerState) Iterator {
	return Literator(Token{Type: t, Value: groups[0]})
}
