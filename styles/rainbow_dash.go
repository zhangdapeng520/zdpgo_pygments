package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// RainbowDash style.
var RainbowDash = Register(zdpgo_pygments.MustNewStyle("rainbow_dash", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:             "italic #0080ff",
	zdpgo_pygments.CommentPreproc:      "noitalic",
	zdpgo_pygments.CommentSpecial:      "bold",
	zdpgo_pygments.Error:               "bg:#cc0000 #ffffff",
	zdpgo_pygments.GenericDeleted:      "border:#c5060b bg:#ffcccc",
	zdpgo_pygments.GenericEmph:         "italic",
	zdpgo_pygments.GenericError:        "#ff0000",
	zdpgo_pygments.GenericHeading:      "bold #2c5dcd",
	zdpgo_pygments.GenericInserted:     "border:#00cc00 bg:#ccffcc",
	zdpgo_pygments.GenericOutput:       "#aaaaaa",
	zdpgo_pygments.GenericPrompt:       "bold #2c5dcd",
	zdpgo_pygments.GenericStrong:       "bold",
	zdpgo_pygments.GenericSubheading:   "bold #2c5dcd",
	zdpgo_pygments.GenericTraceback:    "#c5060b",
	zdpgo_pygments.GenericUnderline:    "underline",
	zdpgo_pygments.Keyword:             "bold #2c5dcd",
	zdpgo_pygments.KeywordPseudo:       "nobold",
	zdpgo_pygments.KeywordType:         "#5918bb",
	zdpgo_pygments.NameAttribute:       "italic #2c5dcd",
	zdpgo_pygments.NameBuiltin:         "bold #5918bb",
	zdpgo_pygments.NameClass:           "underline",
	zdpgo_pygments.NameConstant:        "#318495",
	zdpgo_pygments.NameDecorator:       "bold #ff8000",
	zdpgo_pygments.NameEntity:          "bold #5918bb",
	zdpgo_pygments.NameException:       "bold #5918bb",
	zdpgo_pygments.NameFunction:        "bold #ff8000",
	zdpgo_pygments.NameTag:             "bold #2c5dcd",
	zdpgo_pygments.LiteralNumber:       "bold #5918bb",
	zdpgo_pygments.Operator:            "#2c5dcd",
	zdpgo_pygments.OperatorWord:        "bold",
	zdpgo_pygments.LiteralString:       "#00cc66",
	zdpgo_pygments.LiteralStringDoc:    "italic",
	zdpgo_pygments.LiteralStringEscape: "bold #c5060b",
	zdpgo_pygments.LiteralStringOther:  "#318495",
	zdpgo_pygments.LiteralStringSymbol: "bold #c5060b",
	zdpgo_pygments.Text:                "#4d4d4d",
	zdpgo_pygments.TextWhitespace:      "#cbcbcb",
	zdpgo_pygments.Background:          " bg:#ffffff",
}))
