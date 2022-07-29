package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// VisualStudio style.
var VisualStudio = Register(zdpgo_pygments.MustNewStyle("vs", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:           "#008000",
	zdpgo_pygments.CommentPreproc:    "#0000ff",
	zdpgo_pygments.Keyword:           "#0000ff",
	zdpgo_pygments.OperatorWord:      "#0000ff",
	zdpgo_pygments.KeywordType:       "#2b91af",
	zdpgo_pygments.NameClass:         "#2b91af",
	zdpgo_pygments.LiteralString:     "#a31515",
	zdpgo_pygments.GenericHeading:    "bold",
	zdpgo_pygments.GenericSubheading: "bold",
	zdpgo_pygments.GenericEmph:       "italic",
	zdpgo_pygments.GenericStrong:     "bold",
	zdpgo_pygments.GenericPrompt:     "bold",
	zdpgo_pygments.Error:             "border:#FF0000",
	zdpgo_pygments.Background:        " bg:#ffffff",
}))
