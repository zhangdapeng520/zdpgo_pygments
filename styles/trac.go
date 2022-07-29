package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Trac style.
var Trac = Register(zdpgo_pygments.MustNewStyle("trac", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:     "#bbbbbb",
	zdpgo_pygments.Comment:            "italic #999988",
	zdpgo_pygments.CommentPreproc:     "bold noitalic #999999",
	zdpgo_pygments.CommentSpecial:     "bold #999999",
	zdpgo_pygments.Operator:           "bold",
	zdpgo_pygments.LiteralString:      "#bb8844",
	zdpgo_pygments.LiteralStringRegex: "#808000",
	zdpgo_pygments.LiteralNumber:      "#009999",
	zdpgo_pygments.Keyword:            "bold",
	zdpgo_pygments.KeywordType:        "#445588",
	zdpgo_pygments.NameBuiltin:        "#999999",
	zdpgo_pygments.NameFunction:       "bold #990000",
	zdpgo_pygments.NameClass:          "bold #445588",
	zdpgo_pygments.NameException:      "bold #990000",
	zdpgo_pygments.NameNamespace:      "#555555",
	zdpgo_pygments.NameVariable:       "#008080",
	zdpgo_pygments.NameConstant:       "#008080",
	zdpgo_pygments.NameTag:            "#000080",
	zdpgo_pygments.NameAttribute:      "#008080",
	zdpgo_pygments.NameEntity:         "#800080",
	zdpgo_pygments.GenericHeading:     "#999999",
	zdpgo_pygments.GenericSubheading:  "#aaaaaa",
	zdpgo_pygments.GenericDeleted:     "bg:#ffdddd #000000",
	zdpgo_pygments.GenericInserted:    "bg:#ddffdd #000000",
	zdpgo_pygments.GenericError:       "#aa0000",
	zdpgo_pygments.GenericEmph:        "italic",
	zdpgo_pygments.GenericStrong:      "bold",
	zdpgo_pygments.GenericPrompt:      "#555555",
	zdpgo_pygments.GenericOutput:      "#888888",
	zdpgo_pygments.GenericTraceback:   "#aa0000",
	zdpgo_pygments.GenericUnderline:   "underline",
	zdpgo_pygments.Error:              "bg:#e3d2d2 #a61717",
	zdpgo_pygments.Background:         " bg:#ffffff",
}))
