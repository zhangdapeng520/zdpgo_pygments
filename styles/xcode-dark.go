package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

var (
	// Inspired by Apple's Xcode "Default (Dark)" Theme
	background                  = "#1F1F24"
	plainText                   = "#FFFFFF"
	comments                    = "#6C7986"
	strings                     = "#FC6A5D"
	numbers                     = "#D0BF69"
	keywords                    = "#FC5FA3"
	preprocessorStatements      = "#FD8F3F"
	typeDeclarations            = "#5DD8FF"
	otherDeclarations           = "#41A1C0"
	otherFunctionAndMethodNames = "#A167E6"
	otherTypeNames              = "#D0A8FF"
)

// Xcode dark style
var XcodeDark = Register(zdpgo_pygments.MustNewStyle("xcode-dark", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Background: plainText + " bg:" + background,

	zdpgo_pygments.Comment:          comments,
	zdpgo_pygments.CommentMultiline: comments,
	zdpgo_pygments.CommentPreproc:   preprocessorStatements,
	zdpgo_pygments.CommentSingle:    comments,
	zdpgo_pygments.CommentSpecial:   comments + " italic",

	zdpgo_pygments.Error: "#960050",

	zdpgo_pygments.Keyword:            keywords,
	zdpgo_pygments.KeywordConstant:    keywords,
	zdpgo_pygments.KeywordDeclaration: keywords,
	zdpgo_pygments.KeywordReserved:    keywords,

	zdpgo_pygments.LiteralNumber:        numbers,
	zdpgo_pygments.LiteralNumberBin:     numbers,
	zdpgo_pygments.LiteralNumberFloat:   numbers,
	zdpgo_pygments.LiteralNumberHex:     numbers,
	zdpgo_pygments.LiteralNumberInteger: numbers,
	zdpgo_pygments.LiteralNumberOct:     numbers,

	zdpgo_pygments.LiteralString:         strings,
	zdpgo_pygments.LiteralStringEscape:   strings,
	zdpgo_pygments.LiteralStringInterpol: plainText,

	zdpgo_pygments.Name:              plainText,
	zdpgo_pygments.NameBuiltin:       otherTypeNames,
	zdpgo_pygments.NameBuiltinPseudo: otherFunctionAndMethodNames,
	zdpgo_pygments.NameClass:         typeDeclarations,
	zdpgo_pygments.NameFunction:      otherDeclarations,
	zdpgo_pygments.NameVariable:      otherDeclarations,

	zdpgo_pygments.Operator: plainText,

	zdpgo_pygments.Punctuation: plainText,

	zdpgo_pygments.Text: plainText,
}))
