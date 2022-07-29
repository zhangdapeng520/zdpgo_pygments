package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// SwapOff theme.
var SwapOff = Register(zdpgo_pygments.MustNewStyle("swapoff", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Background:        "#lightgray bg:#black",
	zdpgo_pygments.Number:            "bold #ansiyellow",
	zdpgo_pygments.Comment:           "#ansiteal",
	zdpgo_pygments.CommentPreproc:    "bold #ansigreen",
	zdpgo_pygments.String:            "bold #ansiturquoise",
	zdpgo_pygments.Keyword:           "bold #ansiwhite",
	zdpgo_pygments.NameKeyword:       "bold #ansiwhite",
	zdpgo_pygments.NameBuiltin:       "bold #ansiwhite",
	zdpgo_pygments.GenericHeading:    "bold",
	zdpgo_pygments.GenericSubheading: "bold",
	zdpgo_pygments.GenericStrong:     "bold",
	zdpgo_pygments.GenericUnderline:  "underline",
	zdpgo_pygments.NameTag:           "bold",
	zdpgo_pygments.NameAttribute:     "#ansiteal",
	zdpgo_pygments.Error:             "#ansired",
	zdpgo_pygments.LiteralDate:       "bold #ansiyellow",
}))
