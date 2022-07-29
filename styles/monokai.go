package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Monokai style.
var Monokai = Register(zdpgo_pygments.MustNewStyle("monokai", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:                "#f8f8f2",
	zdpgo_pygments.Error:               "#960050 bg:#1e0010",
	zdpgo_pygments.Comment:             "#75715e",
	zdpgo_pygments.Keyword:             "#66d9ef",
	zdpgo_pygments.KeywordNamespace:    "#f92672",
	zdpgo_pygments.Operator:            "#f92672",
	zdpgo_pygments.Punctuation:         "#f8f8f2",
	zdpgo_pygments.Name:                "#f8f8f2",
	zdpgo_pygments.NameAttribute:       "#a6e22e",
	zdpgo_pygments.NameClass:           "#a6e22e",
	zdpgo_pygments.NameConstant:        "#66d9ef",
	zdpgo_pygments.NameDecorator:       "#a6e22e",
	zdpgo_pygments.NameException:       "#a6e22e",
	zdpgo_pygments.NameFunction:        "#a6e22e",
	zdpgo_pygments.NameOther:           "#a6e22e",
	zdpgo_pygments.NameTag:             "#f92672",
	zdpgo_pygments.LiteralNumber:       "#ae81ff",
	zdpgo_pygments.Literal:             "#ae81ff",
	zdpgo_pygments.LiteralDate:         "#e6db74",
	zdpgo_pygments.LiteralString:       "#e6db74",
	zdpgo_pygments.LiteralStringEscape: "#ae81ff",
	zdpgo_pygments.GenericDeleted:      "#f92672",
	zdpgo_pygments.GenericEmph:         "italic",
	zdpgo_pygments.GenericInserted:     "#a6e22e",
	zdpgo_pygments.GenericStrong:       "bold",
	zdpgo_pygments.GenericSubheading:   "#75715e",
	zdpgo_pygments.Background:          "bg:#272822",
}))
