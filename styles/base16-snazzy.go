package styles

import "github.com/zhangdapeng520/zdpgo_pygments"

// Base16Snazzy style
var Base16Snazzy = Register(zdpgo_pygments.MustNewStyle("base16-snazzy", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:                  "#78787e",
	zdpgo_pygments.CommentHashbang:          "#78787e",
	zdpgo_pygments.CommentMultiline:         "#78787e",
	zdpgo_pygments.CommentPreproc:           "#78787e",
	zdpgo_pygments.CommentSingle:            "#78787e",
	zdpgo_pygments.CommentSpecial:           "#78787e",
	zdpgo_pygments.Generic:                  "#e2e4e5",
	zdpgo_pygments.GenericDeleted:           "#ff5c57",
	zdpgo_pygments.GenericEmph:              "#e2e4e5 underline",
	zdpgo_pygments.GenericError:             "#ff5c57",
	zdpgo_pygments.GenericHeading:           "#e2e4e5 bold",
	zdpgo_pygments.GenericInserted:          "#e2e4e5 bold",
	zdpgo_pygments.GenericOutput:            "#43454f",
	zdpgo_pygments.GenericPrompt:            "#e2e4e5",
	zdpgo_pygments.GenericStrong:            "#e2e4e5 italic",
	zdpgo_pygments.GenericSubheading:        "#e2e4e5 bold",
	zdpgo_pygments.GenericTraceback:         "#e2e4e5",
	zdpgo_pygments.GenericUnderline:         "underline",
	zdpgo_pygments.Error:                    "#ff5c57",
	zdpgo_pygments.Keyword:                  "#ff6ac1",
	zdpgo_pygments.KeywordConstant:          "#ff6ac1",
	zdpgo_pygments.KeywordDeclaration:       "#ff5c57",
	zdpgo_pygments.KeywordNamespace:         "#ff6ac1",
	zdpgo_pygments.KeywordPseudo:            "#ff6ac1",
	zdpgo_pygments.KeywordReserved:          "#ff6ac1",
	zdpgo_pygments.KeywordType:              "#9aedfe",
	zdpgo_pygments.Literal:                  "#e2e4e5",
	zdpgo_pygments.LiteralDate:              "#e2e4e5",
	zdpgo_pygments.Name:                     "#e2e4e5",
	zdpgo_pygments.NameAttribute:            "#57c7ff",
	zdpgo_pygments.NameBuiltin:              "#ff5c57",
	zdpgo_pygments.NameBuiltinPseudo:        "#e2e4e5",
	zdpgo_pygments.NameClass:                "#f3f99d",
	zdpgo_pygments.NameConstant:             "#ff9f43",
	zdpgo_pygments.NameDecorator:            "#ff9f43",
	zdpgo_pygments.NameEntity:               "#e2e4e5",
	zdpgo_pygments.NameException:            "#e2e4e5",
	zdpgo_pygments.NameFunction:             "#57c7ff",
	zdpgo_pygments.NameLabel:                "#ff5c57",
	zdpgo_pygments.NameNamespace:            "#e2e4e5",
	zdpgo_pygments.NameOther:                "#e2e4e5",
	zdpgo_pygments.NameTag:                  "#ff6ac1",
	zdpgo_pygments.NameVariable:             "#ff5c57",
	zdpgo_pygments.NameVariableClass:        "#ff5c57",
	zdpgo_pygments.NameVariableGlobal:       "#ff5c57",
	zdpgo_pygments.NameVariableInstance:     "#ff5c57",
	zdpgo_pygments.LiteralNumber:            "#ff9f43",
	zdpgo_pygments.LiteralNumberBin:         "#ff9f43",
	zdpgo_pygments.LiteralNumberFloat:       "#ff9f43",
	zdpgo_pygments.LiteralNumberHex:         "#ff9f43",
	zdpgo_pygments.LiteralNumberInteger:     "#ff9f43",
	zdpgo_pygments.LiteralNumberIntegerLong: "#ff9f43",
	zdpgo_pygments.LiteralNumberOct:         "#ff9f43",
	zdpgo_pygments.Operator:                 "#ff6ac1",
	zdpgo_pygments.OperatorWord:             "#ff6ac1",
	zdpgo_pygments.Other:                    "#e2e4e5",
	zdpgo_pygments.Punctuation:              "#e2e4e5",
	zdpgo_pygments.LiteralString:            "#5af78e",
	zdpgo_pygments.LiteralStringBacktick:    "#5af78e",
	zdpgo_pygments.LiteralStringChar:        "#5af78e",
	zdpgo_pygments.LiteralStringDoc:         "#5af78e",
	zdpgo_pygments.LiteralStringDouble:      "#5af78e",
	zdpgo_pygments.LiteralStringEscape:      "#5af78e",
	zdpgo_pygments.LiteralStringHeredoc:     "#5af78e",
	zdpgo_pygments.LiteralStringInterpol:    "#5af78e",
	zdpgo_pygments.LiteralStringOther:       "#5af78e",
	zdpgo_pygments.LiteralStringRegex:       "#5af78e",
	zdpgo_pygments.LiteralStringSingle:      "#5af78e",
	zdpgo_pygments.LiteralStringSymbol:      "#5af78e",
	zdpgo_pygments.Text:                     "#e2e4e5",
	zdpgo_pygments.TextWhitespace:           "#e2e4e5",
	zdpgo_pygments.Background:               " bg:#282a36",
}))
