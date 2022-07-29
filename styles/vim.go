package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Vim style.
var Vim = Register(zdpgo_pygments.MustNewStyle("vim", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Background:         "#cccccc bg:#000000",
	zdpgo_pygments.Comment:            "#000080",
	zdpgo_pygments.CommentSpecial:     "bold #cd0000",
	zdpgo_pygments.Keyword:            "#cdcd00",
	zdpgo_pygments.KeywordDeclaration: "#00cd00",
	zdpgo_pygments.KeywordNamespace:   "#cd00cd",
	zdpgo_pygments.KeywordType:        "#00cd00",
	zdpgo_pygments.Operator:           "#3399cc",
	zdpgo_pygments.OperatorWord:       "#cdcd00",
	zdpgo_pygments.NameClass:          "#00cdcd",
	zdpgo_pygments.NameBuiltin:        "#cd00cd",
	zdpgo_pygments.NameException:      "bold #666699",
	zdpgo_pygments.NameVariable:       "#00cdcd",
	zdpgo_pygments.LiteralString:      "#cd0000",
	zdpgo_pygments.LiteralNumber:      "#cd00cd",
	zdpgo_pygments.GenericHeading:     "bold #000080",
	zdpgo_pygments.GenericSubheading:  "bold #800080",
	zdpgo_pygments.GenericDeleted:     "#cd0000",
	zdpgo_pygments.GenericInserted:    "#00cd00",
	zdpgo_pygments.GenericError:       "#FF0000",
	zdpgo_pygments.GenericEmph:        "italic",
	zdpgo_pygments.GenericStrong:      "bold",
	zdpgo_pygments.GenericPrompt:      "bold #000080",
	zdpgo_pygments.GenericOutput:      "#888",
	zdpgo_pygments.GenericTraceback:   "#04D",
	zdpgo_pygments.GenericUnderline:   "underline",
	zdpgo_pygments.Error:              "border:#FF0000",
}))
