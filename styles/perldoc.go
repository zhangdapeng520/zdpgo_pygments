package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Perldoc style.
var Perldoc = Register(zdpgo_pygments.MustNewStyle("perldoc", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:       "#bbbbbb",
	zdpgo_pygments.Comment:              "#228B22",
	zdpgo_pygments.CommentPreproc:       "#1e889b",
	zdpgo_pygments.CommentSpecial:       "#8B008B bold",
	zdpgo_pygments.LiteralString:        "#CD5555",
	zdpgo_pygments.LiteralStringHeredoc: "#1c7e71 italic",
	zdpgo_pygments.LiteralStringRegex:   "#1c7e71",
	zdpgo_pygments.LiteralStringOther:   "#cb6c20",
	zdpgo_pygments.LiteralNumber:        "#B452CD",
	zdpgo_pygments.OperatorWord:         "#8B008B",
	zdpgo_pygments.Keyword:              "#8B008B bold",
	zdpgo_pygments.KeywordType:          "#00688B",
	zdpgo_pygments.NameClass:            "#008b45 bold",
	zdpgo_pygments.NameException:        "#008b45 bold",
	zdpgo_pygments.NameFunction:         "#008b45",
	zdpgo_pygments.NameNamespace:        "#008b45 underline",
	zdpgo_pygments.NameVariable:         "#00688B",
	zdpgo_pygments.NameConstant:         "#00688B",
	zdpgo_pygments.NameDecorator:        "#707a7c",
	zdpgo_pygments.NameTag:              "#8B008B bold",
	zdpgo_pygments.NameAttribute:        "#658b00",
	zdpgo_pygments.NameBuiltin:          "#658b00",
	zdpgo_pygments.GenericHeading:       "bold #000080",
	zdpgo_pygments.GenericSubheading:    "bold #800080",
	zdpgo_pygments.GenericDeleted:       "#aa0000",
	zdpgo_pygments.GenericInserted:      "#00aa00",
	zdpgo_pygments.GenericError:         "#aa0000",
	zdpgo_pygments.GenericEmph:          "italic",
	zdpgo_pygments.GenericStrong:        "bold",
	zdpgo_pygments.GenericPrompt:        "#555555",
	zdpgo_pygments.GenericOutput:        "#888888",
	zdpgo_pygments.GenericTraceback:     "#aa0000",
	zdpgo_pygments.GenericUnderline:     "underline",
	zdpgo_pygments.Error:                "bg:#e3d2d2 #a61717",
	zdpgo_pygments.Background:           " bg:#eeeedd",
}))
