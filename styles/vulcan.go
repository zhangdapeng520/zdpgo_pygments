package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

var (
	// inspired by Doom Emacs's One Doom Theme
	black  = "#282C34"
	grey   = "#3E4460"
	grey2  = "#43454f"
	white  = "#C9C9C9"
	red    = "#CF5967"
	yellow = "#ECBE7B"
	green  = "#82CC6A"
	cyan   = "#56B6C2"
	blue   = "#7FBAF5"
	blue2  = "#57C7FF"
	purple = "#BC74C4"
)

var Vulcan = Register(zdpgo_pygments.MustNewStyle("vulcan", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:                  grey,
	zdpgo_pygments.CommentHashbang:          grey + " italic",
	zdpgo_pygments.CommentMultiline:         grey,
	zdpgo_pygments.CommentPreproc:           blue,
	zdpgo_pygments.CommentSingle:            grey,
	zdpgo_pygments.CommentSpecial:           purple + " italic",
	zdpgo_pygments.Generic:                  white,
	zdpgo_pygments.GenericDeleted:           red,
	zdpgo_pygments.GenericEmph:              white + " underline",
	zdpgo_pygments.GenericError:             red + " bold",
	zdpgo_pygments.GenericHeading:           yellow + " bold",
	zdpgo_pygments.GenericInserted:          yellow,
	zdpgo_pygments.GenericOutput:            grey2,
	zdpgo_pygments.GenericPrompt:            white,
	zdpgo_pygments.GenericStrong:            red + " bold",
	zdpgo_pygments.GenericSubheading:        red + " italic",
	zdpgo_pygments.GenericTraceback:         white,
	zdpgo_pygments.GenericUnderline:         "underline",
	zdpgo_pygments.Error:                    red,
	zdpgo_pygments.Keyword:                  blue,
	zdpgo_pygments.KeywordConstant:          red + " bg:" + grey2,
	zdpgo_pygments.KeywordDeclaration:       blue,
	zdpgo_pygments.KeywordNamespace:         purple,
	zdpgo_pygments.KeywordPseudo:            purple,
	zdpgo_pygments.KeywordReserved:          blue,
	zdpgo_pygments.KeywordType:              blue2 + " bold",
	zdpgo_pygments.Literal:                  white,
	zdpgo_pygments.LiteralDate:              blue2,
	zdpgo_pygments.Name:                     white,
	zdpgo_pygments.NameAttribute:            purple,
	zdpgo_pygments.NameBuiltin:              blue,
	zdpgo_pygments.NameBuiltinPseudo:        blue,
	zdpgo_pygments.NameClass:                yellow,
	zdpgo_pygments.NameConstant:             yellow,
	zdpgo_pygments.NameDecorator:            yellow,
	zdpgo_pygments.NameEntity:               white,
	zdpgo_pygments.NameException:            red,
	zdpgo_pygments.NameFunction:             blue2,
	zdpgo_pygments.NameLabel:                red,
	zdpgo_pygments.NameNamespace:            white,
	zdpgo_pygments.NameOther:                white,
	zdpgo_pygments.NameTag:                  purple,
	zdpgo_pygments.NameVariable:             purple + " italic",
	zdpgo_pygments.NameVariableClass:        blue2 + " bold",
	zdpgo_pygments.NameVariableGlobal:       yellow,
	zdpgo_pygments.NameVariableInstance:     blue2,
	zdpgo_pygments.LiteralNumber:            cyan,
	zdpgo_pygments.LiteralNumberBin:         blue2,
	zdpgo_pygments.LiteralNumberFloat:       cyan,
	zdpgo_pygments.LiteralNumberHex:         blue2,
	zdpgo_pygments.LiteralNumberInteger:     cyan,
	zdpgo_pygments.LiteralNumberIntegerLong: cyan,
	zdpgo_pygments.LiteralNumberOct:         blue2,
	zdpgo_pygments.Operator:                 purple,
	zdpgo_pygments.OperatorWord:             purple,
	zdpgo_pygments.Other:                    white,
	zdpgo_pygments.Punctuation:              cyan,
	zdpgo_pygments.LiteralString:            green,
	zdpgo_pygments.LiteralStringBacktick:    blue2,
	zdpgo_pygments.LiteralStringChar:        blue2,
	zdpgo_pygments.LiteralStringDoc:         green,
	zdpgo_pygments.LiteralStringDouble:      green,
	zdpgo_pygments.LiteralStringEscape:      cyan,
	zdpgo_pygments.LiteralStringHeredoc:     cyan,
	zdpgo_pygments.LiteralStringInterpol:    green,
	zdpgo_pygments.LiteralStringOther:       green,
	zdpgo_pygments.LiteralStringRegex:       blue2,
	zdpgo_pygments.LiteralStringSingle:      green,
	zdpgo_pygments.LiteralStringSymbol:      green,
	zdpgo_pygments.Text:                     white,
	zdpgo_pygments.TextWhitespace:           white,
	zdpgo_pygments.Background:               " bg: " + black,
}))
