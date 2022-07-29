package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Theme based on HackerRank High Contrast Editor Theme
var HrHighContrast = Register(zdpgo_pygments.MustNewStyle("hr_high_contrast", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:              "#5a8349",
	zdpgo_pygments.Keyword:              "#467faf",
	zdpgo_pygments.OperatorWord:         "#467faf",
	zdpgo_pygments.Name:                 "#ffffff",
	zdpgo_pygments.LiteralString:        "#a87662",
	zdpgo_pygments.LiteralNumber:        "#fff",
	zdpgo_pygments.LiteralStringBoolean: "#467faf",
	zdpgo_pygments.Operator:             "#e4e400",
	zdpgo_pygments.Background:           "#000",
	zdpgo_pygments.Other:                "#d5d500",
}))
