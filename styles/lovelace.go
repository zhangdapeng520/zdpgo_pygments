package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Lovelace style.
var Lovelace = Register(zdpgo_pygments.MustNewStyle("lovelace", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:         "#a89028",
	zdpgo_pygments.Comment:                "italic #888888",
	zdpgo_pygments.CommentHashbang:        "#287088",
	zdpgo_pygments.CommentMultiline:       "#888888",
	zdpgo_pygments.CommentPreproc:         "noitalic #289870",
	zdpgo_pygments.Keyword:                "#2838b0",
	zdpgo_pygments.KeywordConstant:        "italic #444444",
	zdpgo_pygments.KeywordDeclaration:     "italic",
	zdpgo_pygments.KeywordType:            "italic",
	zdpgo_pygments.Operator:               "#666666",
	zdpgo_pygments.OperatorWord:           "#a848a8",
	zdpgo_pygments.Punctuation:            "#888888",
	zdpgo_pygments.NameAttribute:          "#388038",
	zdpgo_pygments.NameBuiltin:            "#388038",
	zdpgo_pygments.NameBuiltinPseudo:      "italic",
	zdpgo_pygments.NameClass:              "#287088",
	zdpgo_pygments.NameConstant:           "#b85820",
	zdpgo_pygments.NameDecorator:          "#287088",
	zdpgo_pygments.NameEntity:             "#709030",
	zdpgo_pygments.NameException:          "#908828",
	zdpgo_pygments.NameFunction:           "#785840",
	zdpgo_pygments.NameFunctionMagic:      "#b85820",
	zdpgo_pygments.NameLabel:              "#289870",
	zdpgo_pygments.NameNamespace:          "#289870",
	zdpgo_pygments.NameTag:                "#2838b0",
	zdpgo_pygments.NameVariable:           "#b04040",
	zdpgo_pygments.NameVariableGlobal:     "#908828",
	zdpgo_pygments.NameVariableMagic:      "#b85820",
	zdpgo_pygments.LiteralString:          "#b83838",
	zdpgo_pygments.LiteralStringAffix:     "#444444",
	zdpgo_pygments.LiteralStringChar:      "#a848a8",
	zdpgo_pygments.LiteralStringDelimiter: "#b85820",
	zdpgo_pygments.LiteralStringDoc:       "italic #b85820",
	zdpgo_pygments.LiteralStringEscape:    "#709030",
	zdpgo_pygments.LiteralStringInterpol:  "underline",
	zdpgo_pygments.LiteralStringOther:     "#a848a8",
	zdpgo_pygments.LiteralStringRegex:     "#a848a8",
	zdpgo_pygments.LiteralNumber:          "#444444",
	zdpgo_pygments.GenericDeleted:         "#c02828",
	zdpgo_pygments.GenericEmph:            "italic",
	zdpgo_pygments.GenericError:           "#c02828",
	zdpgo_pygments.GenericHeading:         "#666666",
	zdpgo_pygments.GenericSubheading:      "#444444",
	zdpgo_pygments.GenericInserted:        "#388038",
	zdpgo_pygments.GenericOutput:          "#666666",
	zdpgo_pygments.GenericPrompt:          "#444444",
	zdpgo_pygments.GenericStrong:          "bold",
	zdpgo_pygments.GenericTraceback:       "#2838b0",
	zdpgo_pygments.GenericUnderline:       "underline",
	zdpgo_pygments.Error:                  "bg:#a848a8",
	zdpgo_pygments.Background:             " bg:#ffffff",
}))
