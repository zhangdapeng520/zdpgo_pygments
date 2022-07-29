package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Theme based on HackerRank Dark Editor theme
var HrDark = Register(zdpgo_pygments.MustNewStyle("hrdark", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:      "italic #828b96",
	zdpgo_pygments.Keyword:      "#ff636f",
	zdpgo_pygments.OperatorWord: "#ff636f",
	zdpgo_pygments.Name:         "#58a1dd",
	zdpgo_pygments.Literal:      "#a6be9d",
	zdpgo_pygments.Operator:     "#ff636f",
	zdpgo_pygments.Background:   "#1d2432",
	zdpgo_pygments.Other:        "#fff",
}))
