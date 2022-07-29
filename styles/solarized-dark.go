package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// SolarizedDark style.
var SolarizedDark = Register(zdpgo_pygments.MustNewStyle("solarized-dark", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Keyword:               "#719e07",
	zdpgo_pygments.KeywordConstant:       "#CB4B16",
	zdpgo_pygments.KeywordDeclaration:    "#268BD2",
	zdpgo_pygments.KeywordReserved:       "#268BD2",
	zdpgo_pygments.KeywordType:           "#DC322F",
	zdpgo_pygments.NameAttribute:         "#93A1A1",
	zdpgo_pygments.NameBuiltin:           "#B58900",
	zdpgo_pygments.NameBuiltinPseudo:     "#268BD2",
	zdpgo_pygments.NameClass:             "#268BD2",
	zdpgo_pygments.NameConstant:          "#CB4B16",
	zdpgo_pygments.NameDecorator:         "#268BD2",
	zdpgo_pygments.NameEntity:            "#CB4B16",
	zdpgo_pygments.NameException:         "#CB4B16",
	zdpgo_pygments.NameFunction:          "#268BD2",
	zdpgo_pygments.NameTag:               "#268BD2",
	zdpgo_pygments.NameVariable:          "#268BD2",
	zdpgo_pygments.LiteralString:         "#2AA198",
	zdpgo_pygments.LiteralStringBacktick: "#586E75",
	zdpgo_pygments.LiteralStringChar:     "#2AA198",
	zdpgo_pygments.LiteralStringDoc:      "#93A1A1",
	zdpgo_pygments.LiteralStringEscape:   "#CB4B16",
	zdpgo_pygments.LiteralStringHeredoc:  "#93A1A1",
	zdpgo_pygments.LiteralStringRegex:    "#DC322F",
	zdpgo_pygments.LiteralNumber:         "#2AA198",
	zdpgo_pygments.Operator:              "#719e07",
	zdpgo_pygments.Comment:               "#586E75",
	zdpgo_pygments.CommentPreproc:        "#719e07",
	zdpgo_pygments.CommentSpecial:        "#719e07",
	zdpgo_pygments.GenericDeleted:        "#DC322F",
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericError:          "#DC322F bold",
	zdpgo_pygments.GenericHeading:        "#CB4B16",
	zdpgo_pygments.GenericInserted:       "#719e07",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericSubheading:     "#268BD2",
	zdpgo_pygments.Background:            "#93A1A1 bg:#002B36",
	zdpgo_pygments.Other:                 "#CB4B16",
}))
