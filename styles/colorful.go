package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Colorful style.
var Colorful = Register(zdpgo_pygments.MustNewStyle("colorful", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:        "#bbbbbb",
	zdpgo_pygments.Comment:               "#888",
	zdpgo_pygments.CommentPreproc:        "#579",
	zdpgo_pygments.CommentSpecial:        "bold #cc0000",
	zdpgo_pygments.Keyword:               "bold #080",
	zdpgo_pygments.KeywordPseudo:         "#038",
	zdpgo_pygments.KeywordType:           "#339",
	zdpgo_pygments.Operator:              "#333",
	zdpgo_pygments.OperatorWord:          "bold #000",
	zdpgo_pygments.NameBuiltin:           "#007020",
	zdpgo_pygments.NameFunction:          "bold #06B",
	zdpgo_pygments.NameClass:             "bold #B06",
	zdpgo_pygments.NameNamespace:         "bold #0e84b5",
	zdpgo_pygments.NameException:         "bold #F00",
	zdpgo_pygments.NameVariable:          "#963",
	zdpgo_pygments.NameVariableInstance:  "#33B",
	zdpgo_pygments.NameVariableClass:     "#369",
	zdpgo_pygments.NameVariableGlobal:    "bold #d70",
	zdpgo_pygments.NameConstant:          "bold #036",
	zdpgo_pygments.NameLabel:             "bold #970",
	zdpgo_pygments.NameEntity:            "bold #800",
	zdpgo_pygments.NameAttribute:         "#00C",
	zdpgo_pygments.NameTag:               "#070",
	zdpgo_pygments.NameDecorator:         "bold #555",
	zdpgo_pygments.LiteralString:         "bg:#fff0f0",
	zdpgo_pygments.LiteralStringChar:     "#04D bg:",
	zdpgo_pygments.LiteralStringDoc:      "#D42 bg:",
	zdpgo_pygments.LiteralStringInterpol: "bg:#eee",
	zdpgo_pygments.LiteralStringEscape:   "bold #666",
	zdpgo_pygments.LiteralStringRegex:    "bg:#fff0ff #000",
	zdpgo_pygments.LiteralStringSymbol:   "#A60 bg:",
	zdpgo_pygments.LiteralStringOther:    "#D20",
	zdpgo_pygments.LiteralNumber:         "bold #60E",
	zdpgo_pygments.LiteralNumberInteger:  "bold #00D",
	zdpgo_pygments.LiteralNumberFloat:    "bold #60E",
	zdpgo_pygments.LiteralNumberHex:      "bold #058",
	zdpgo_pygments.LiteralNumberOct:      "bold #40E",
	zdpgo_pygments.GenericHeading:        "bold #000080",
	zdpgo_pygments.GenericSubheading:     "bold #800080",
	zdpgo_pygments.GenericDeleted:        "#A00000",
	zdpgo_pygments.GenericInserted:       "#00A000",
	zdpgo_pygments.GenericError:          "#FF0000",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericPrompt:         "bold #c65d09",
	zdpgo_pygments.GenericOutput:         "#888",
	zdpgo_pygments.GenericTraceback:      "#04D",
	zdpgo_pygments.GenericUnderline:      "underline",
	zdpgo_pygments.Error:                 "#F00 bg:#FAA",
	zdpgo_pygments.Background:            " bg:#ffffff",
}))
