package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// SolarizedDark256 style.
var SolarizedDark256 = Register(zdpgo_pygments.MustNewStyle("solarized-dark256", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Keyword:               "#5f8700",
	zdpgo_pygments.KeywordConstant:       "#d75f00",
	zdpgo_pygments.KeywordDeclaration:    "#0087ff",
	zdpgo_pygments.KeywordNamespace:      "#d75f00",
	zdpgo_pygments.KeywordReserved:       "#0087ff",
	zdpgo_pygments.KeywordType:           "#af0000",
	zdpgo_pygments.NameAttribute:         "#8a8a8a",
	zdpgo_pygments.NameBuiltin:           "#0087ff",
	zdpgo_pygments.NameBuiltinPseudo:     "#0087ff",
	zdpgo_pygments.NameClass:             "#0087ff",
	zdpgo_pygments.NameConstant:          "#d75f00",
	zdpgo_pygments.NameDecorator:         "#0087ff",
	zdpgo_pygments.NameEntity:            "#d75f00",
	zdpgo_pygments.NameException:         "#af8700",
	zdpgo_pygments.NameFunction:          "#0087ff",
	zdpgo_pygments.NameTag:               "#0087ff",
	zdpgo_pygments.NameVariable:          "#0087ff",
	zdpgo_pygments.LiteralString:         "#00afaf",
	zdpgo_pygments.LiteralStringBacktick: "#4e4e4e",
	zdpgo_pygments.LiteralStringChar:     "#00afaf",
	zdpgo_pygments.LiteralStringDoc:      "#00afaf",
	zdpgo_pygments.LiteralStringEscape:   "#af0000",
	zdpgo_pygments.LiteralStringHeredoc:  "#00afaf",
	zdpgo_pygments.LiteralStringRegex:    "#af0000",
	zdpgo_pygments.LiteralNumber:         "#00afaf",
	zdpgo_pygments.Operator:              "#8a8a8a",
	zdpgo_pygments.OperatorWord:          "#5f8700",
	zdpgo_pygments.Comment:               "#4e4e4e",
	zdpgo_pygments.CommentPreproc:        "#5f8700",
	zdpgo_pygments.CommentSpecial:        "#5f8700",
	zdpgo_pygments.GenericDeleted:        "#af0000",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericError:          "#af0000 bold",
	zdpgo_pygments.GenericHeading:        "#d75f00",
	zdpgo_pygments.GenericInserted:       "#5f8700",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericSubheading:     "#0087ff",
	zdpgo_pygments.Background:            "#8a8a8a bg:#1c1c1c",
	zdpgo_pygments.Other:                 "#d75f00",
}))
