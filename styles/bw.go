package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// BlackWhite style.
var BlackWhite = Register(zdpgo_pygments.MustNewStyle("bw", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:               "italic",
	zdpgo_pygments.CommentPreproc:        "noitalic",
	zdpgo_pygments.Keyword:               "bold",
	zdpgo_pygments.KeywordPseudo:         "nobold",
	zdpgo_pygments.KeywordType:           "nobold",
	zdpgo_pygments.OperatorWord:          "bold",
	zdpgo_pygments.NameClass:             "bold",
	zdpgo_pygments.NameNamespace:         "bold",
	zdpgo_pygments.NameException:         "bold",
	zdpgo_pygments.NameEntity:            "bold",
	zdpgo_pygments.NameTag:               "bold",
	zdpgo_pygments.LiteralString:         "italic",
	zdpgo_pygments.LiteralStringInterpol: "bold",
	zdpgo_pygments.LiteralStringEscape:   "bold",
	zdpgo_pygments.GenericHeading:        "bold",
	zdpgo_pygments.GenericSubheading:     "bold",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericPrompt:         "bold",
	zdpgo_pygments.Error:                 "border:#FF0000",
	zdpgo_pygments.Background:            " bg:#ffffff",
}))
