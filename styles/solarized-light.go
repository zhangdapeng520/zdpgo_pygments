package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// SolarizedLight style.
var SolarizedLight = Register(zdpgo_pygments.MustNewStyle("solarized-light", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:             "bg: #eee8d5 #586e75",
	zdpgo_pygments.Keyword:          "#859900",
	zdpgo_pygments.KeywordConstant:  "bold",
	zdpgo_pygments.KeywordNamespace: "#dc322f bold",
	zdpgo_pygments.KeywordType:      "bold",
	zdpgo_pygments.Name:             "#268bd2",
	zdpgo_pygments.NameBuiltin:      "#cb4b16",
	zdpgo_pygments.NameClass:        "#cb4b16",
	zdpgo_pygments.NameTag:          "bold",
	zdpgo_pygments.Literal:          "#2aa198",
	zdpgo_pygments.LiteralNumber:    "bold",
	zdpgo_pygments.OperatorWord:     "#859900",
	zdpgo_pygments.Comment:          "#93a1a1 italic",
	zdpgo_pygments.Generic:          "#d33682",
	zdpgo_pygments.Background:       " bg:#eee8d5",
}))
