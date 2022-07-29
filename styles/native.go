package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Native style.
var Native = Register(zdpgo_pygments.MustNewStyle("native", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Background:         "#d0d0d0 bg:#202020",
	zdpgo_pygments.TextWhitespace:     "#666666",
	zdpgo_pygments.Comment:            "italic #999999",
	zdpgo_pygments.CommentPreproc:     "noitalic bold #cd2828",
	zdpgo_pygments.CommentSpecial:     "noitalic bold #e50808 bg:#520000",
	zdpgo_pygments.Keyword:            "bold #6ab825",
	zdpgo_pygments.KeywordPseudo:      "nobold",
	zdpgo_pygments.OperatorWord:       "bold #6ab825",
	zdpgo_pygments.LiteralString:      "#ed9d13",
	zdpgo_pygments.LiteralStringOther: "#ffa500",
	zdpgo_pygments.LiteralNumber:      "#3677a9",
	zdpgo_pygments.NameBuiltin:        "#24909d",
	zdpgo_pygments.NameVariable:       "#40ffff",
	zdpgo_pygments.NameConstant:       "#40ffff",
	zdpgo_pygments.NameClass:          "underline #447fcf",
	zdpgo_pygments.NameFunction:       "#447fcf",
	zdpgo_pygments.NameNamespace:      "underline #447fcf",
	zdpgo_pygments.NameException:      "#bbbbbb",
	zdpgo_pygments.NameTag:            "bold #6ab825",
	zdpgo_pygments.NameAttribute:      "#bbbbbb",
	zdpgo_pygments.NameDecorator:      "#ffa500",
	zdpgo_pygments.GenericHeading:     "bold #ffffff",
	zdpgo_pygments.GenericSubheading:  "underline #ffffff",
	zdpgo_pygments.GenericDeleted:     "#d22323",
	zdpgo_pygments.GenericInserted:    "#589819",
	zdpgo_pygments.GenericError:       "#d22323",
	zdpgo_pygments.GenericEmph:        "italic",
	zdpgo_pygments.GenericStrong:      "bold",
	zdpgo_pygments.GenericPrompt:      "#aaaaaa",
	zdpgo_pygments.GenericOutput:      "#cccccc",
	zdpgo_pygments.GenericTraceback:   "#d22323",
	zdpgo_pygments.GenericUnderline:   "underline",
	zdpgo_pygments.Error:              "bg:#e3d2d2 #a61717",
}))
