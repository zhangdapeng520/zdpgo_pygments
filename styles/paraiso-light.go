package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// ParaisoLight style.
var ParaisoLight = Register(zdpgo_pygments.MustNewStyle("paraiso-light", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:                  "#2f1e2e",
	zdpgo_pygments.Error:                 "#ef6155",
	zdpgo_pygments.Comment:               "#8d8687",
	zdpgo_pygments.Keyword:               "#815ba4",
	zdpgo_pygments.KeywordNamespace:      "#5bc4bf",
	zdpgo_pygments.KeywordType:           "#fec418",
	zdpgo_pygments.Operator:              "#5bc4bf",
	zdpgo_pygments.Punctuation:           "#2f1e2e",
	zdpgo_pygments.Name:                  "#2f1e2e",
	zdpgo_pygments.NameAttribute:         "#06b6ef",
	zdpgo_pygments.NameClass:             "#fec418",
	zdpgo_pygments.NameConstant:          "#ef6155",
	zdpgo_pygments.NameDecorator:         "#5bc4bf",
	zdpgo_pygments.NameException:         "#ef6155",
	zdpgo_pygments.NameFunction:          "#06b6ef",
	zdpgo_pygments.NameNamespace:         "#fec418",
	zdpgo_pygments.NameOther:             "#06b6ef",
	zdpgo_pygments.NameTag:               "#5bc4bf",
	zdpgo_pygments.NameVariable:          "#ef6155",
	zdpgo_pygments.LiteralNumber:         "#f99b15",
	zdpgo_pygments.Literal:               "#f99b15",
	zdpgo_pygments.LiteralDate:           "#48b685",
	zdpgo_pygments.LiteralString:         "#48b685",
	zdpgo_pygments.LiteralStringChar:     "#2f1e2e",
	zdpgo_pygments.LiteralStringDoc:      "#8d8687",
	zdpgo_pygments.LiteralStringEscape:   "#f99b15",
	zdpgo_pygments.LiteralStringInterpol: "#f99b15",
	zdpgo_pygments.GenericDeleted:        "#ef6155",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericHeading:        "bold #2f1e2e",
	zdpgo_pygments.GenericInserted:       "#48b685",
	zdpgo_pygments.GenericPrompt:         "bold #8d8687",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericSubheading:     "bold #5bc4bf",
	zdpgo_pygments.Background:            "bg:#e7e9db",
}))
