package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// ParaisoDark style.
var ParaisoDark = Register(zdpgo_pygments.MustNewStyle("paraiso-dark", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:                  "#e7e9db",
	zdpgo_pygments.Error:                 "#ef6155",
	zdpgo_pygments.Comment:               "#776e71",
	zdpgo_pygments.Keyword:               "#815ba4",
	zdpgo_pygments.KeywordNamespace:      "#5bc4bf",
	zdpgo_pygments.KeywordType:           "#fec418",
	zdpgo_pygments.Operator:              "#5bc4bf",
	zdpgo_pygments.Punctuation:           "#e7e9db",
	zdpgo_pygments.Name:                  "#e7e9db",
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
	zdpgo_pygments.LiteralStringChar:     "#e7e9db",
	zdpgo_pygments.LiteralStringDoc:      "#776e71",
	zdpgo_pygments.LiteralStringEscape:   "#f99b15",
	zdpgo_pygments.LiteralStringInterpol: "#f99b15",
	zdpgo_pygments.GenericDeleted:        "#ef6155",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericHeading:        "bold #e7e9db",
	zdpgo_pygments.GenericInserted:       "#48b685",
	zdpgo_pygments.GenericPrompt:         "bold #776e71",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericSubheading:     "bold #5bc4bf",
	zdpgo_pygments.Background:            "bg:#2f1e2e",
}))
