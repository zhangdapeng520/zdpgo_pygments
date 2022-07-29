package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// 1S:Designer color palette
var OnesEnterprise = Register(zdpgo_pygments.MustNewStyle("onesenterprise", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:           "#000000",
	zdpgo_pygments.Comment:        "#008000",
	zdpgo_pygments.CommentPreproc: "#963200",
	zdpgo_pygments.Operator:       "#FF0000",
	zdpgo_pygments.Keyword:        "#FF0000",
	zdpgo_pygments.Punctuation:    "#FF0000",
	zdpgo_pygments.LiteralString:  "#000000",
	zdpgo_pygments.Name:           "#0000FF",
}))
