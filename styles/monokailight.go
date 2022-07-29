package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// MonokaiLight style.
var MonokaiLight = Register(zdpgo_pygments.MustNewStyle("monokailight", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:             "#272822",
	zdpgo_pygments.Error:            "#960050 bg:#1e0010",
	zdpgo_pygments.Comment:          "#75715e",
	zdpgo_pygments.Keyword:          "#00a8c8",
	zdpgo_pygments.KeywordNamespace: "#f92672",
	zdpgo_pygments.Operator:         "#f92672",
	zdpgo_pygments.Punctuation:      "#111111",
	zdpgo_pygments.Name:             "#111111",
	zdpgo_pygments.NameAttribute:    "#75af00",
	zdpgo_pygments.NameClass:        "#75af00",
	zdpgo_pygments.NameConstant:     "#00a8c8",
	zdpgo_pygments.NameDecorator:    "#75af00",
	zdpgo_pygments.NameException:    "#75af00",
	zdpgo_pygments.NameFunction:     "#75af00",
	zdpgo_pygments.NameOther:        "#75af00",
	zdpgo_pygments.NameTag:          "#f92672",
	zdpgo_pygments.LiteralNumber:    "#ae81ff",
	zdpgo_pygments.Literal:          "#ae81ff",
	zdpgo_pygments.LiteralDate:      "#d88200",
	zdpgo_pygments.LiteralString:    "#d88200",
	zdpgo_pygments.LiteralStringEscape: "#8045FF",
	zdpgo_pygments.GenericEmph:         "italic",
	zdpgo_pygments.GenericStrong:       "bold",
	zdpgo_pygments.Background:          " bg:#fafafa",
}))
