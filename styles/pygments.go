package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Pygments default theme.
var Pygments = Register(zdpgo_pygments.MustNewStyle("pygments", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Whitespace:     "#bbbbbb",
	zdpgo_pygments.Comment:        "italic #408080",
	zdpgo_pygments.CommentPreproc: "noitalic #BC7A00",

	zdpgo_pygments.Keyword:       "bold #008000",
	zdpgo_pygments.KeywordPseudo: "nobold",
	zdpgo_pygments.KeywordType:   "nobold #B00040",

	zdpgo_pygments.Operator:     "#666666",
	zdpgo_pygments.OperatorWord: "bold #AA22FF",

	zdpgo_pygments.NameBuiltin:   "#008000",
	zdpgo_pygments.NameFunction:  "#0000FF",
	zdpgo_pygments.NameClass:     "bold #0000FF",
	zdpgo_pygments.NameNamespace: "bold #0000FF",
	zdpgo_pygments.NameException: "bold #D2413A",
	zdpgo_pygments.NameVariable:  "#19177C",
	zdpgo_pygments.NameConstant:  "#880000",
	zdpgo_pygments.NameLabel:     "#A0A000",
	zdpgo_pygments.NameEntity:    "bold #999999",
	zdpgo_pygments.NameAttribute: "#7D9029",
	zdpgo_pygments.NameTag:       "bold #008000",
	zdpgo_pygments.NameDecorator: "#AA22FF",

	zdpgo_pygments.String:         "#BA2121",
	zdpgo_pygments.StringDoc:      "italic",
	zdpgo_pygments.StringInterpol: "bold #BB6688",
	zdpgo_pygments.StringEscape:   "bold #BB6622",
	zdpgo_pygments.StringRegex:    "#BB6688",
	zdpgo_pygments.StringSymbol:   "#19177C",
	zdpgo_pygments.StringOther:    "#008000",
	zdpgo_pygments.Number:         "#666666",

	zdpgo_pygments.GenericHeading:    "bold #000080",
	zdpgo_pygments.GenericSubheading: "bold #800080",
	zdpgo_pygments.GenericDeleted:    "#A00000",
	zdpgo_pygments.GenericInserted:   "#00A000",
	zdpgo_pygments.GenericError:      "#FF0000",
	zdpgo_pygments.GenericEmph:       "italic",
	zdpgo_pygments.GenericStrong:     "bold",
	zdpgo_pygments.GenericPrompt:     "bold #000080",
	zdpgo_pygments.GenericOutput:     "#888",
	zdpgo_pygments.GenericTraceback:  "#04D",
	zdpgo_pygments.GenericUnderline:  "underline",

	zdpgo_pygments.Error: "border:#FF0000",
}))
