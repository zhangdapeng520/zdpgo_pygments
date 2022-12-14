package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Doom One 2 style. Inspired by Atom One and Doom Emacs's Atom One theme
var DoomOne2 = Register(zdpgo_pygments.MustNewStyle("doom-one2", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:                 "#b0c4de",
	zdpgo_pygments.Error:                "#b0c4de",
	zdpgo_pygments.Comment:              "italic #8a93a5",
	zdpgo_pygments.CommentHashbang:      "bold",
	zdpgo_pygments.Keyword:              "#76a9f9",
	zdpgo_pygments.KeywordConstant:      "#e5c07b",
	zdpgo_pygments.KeywordType:          "#e5c07b",
	zdpgo_pygments.Operator:             "#54b1c7",
	zdpgo_pygments.OperatorWord:         "bold #b756ff",
	zdpgo_pygments.Punctuation:          "#abb2bf",
	zdpgo_pygments.Name:                 "#aa89ea",
	zdpgo_pygments.NameAttribute:        "#cebc3a",
	zdpgo_pygments.NameBuiltin:          "#e5c07b",
	zdpgo_pygments.NameClass:            "#ca72ff",
	zdpgo_pygments.NameConstant:         "bold",
	zdpgo_pygments.NameDecorator:        "#e5c07b",
	zdpgo_pygments.NameEntity:           "#bda26f",
	zdpgo_pygments.NameException:        "bold #fd7474",
	zdpgo_pygments.NameFunction:         "#00b1f7",
	zdpgo_pygments.NameProperty:         "#cebc3a",
	zdpgo_pygments.NameLabel:            "#f5a40d",
	zdpgo_pygments.NameNamespace:        "#ca72ff",
	zdpgo_pygments.NameTag:              "#76a9f9",
	zdpgo_pygments.NameVariable:         "#DCAEEA",
	zdpgo_pygments.NameVariableClass:    "#DCAEEA",
	zdpgo_pygments.NameVariableGlobal:   "bold #DCAEEA",
	zdpgo_pygments.NameVariableInstance: "#e06c75",
	zdpgo_pygments.NameVariableMagic:    "#DCAEEA",
	zdpgo_pygments.Literal:              "#98c379",
	zdpgo_pygments.LiteralDate:          "#98c379",
	zdpgo_pygments.Number:               "#d19a66",
	zdpgo_pygments.NumberBin:            "#d19a66",
	zdpgo_pygments.NumberFloat:          "#d19a66",
	zdpgo_pygments.NumberHex:            "#d19a66",
	zdpgo_pygments.NumberInteger:        "#d19a66",
	zdpgo_pygments.NumberIntegerLong:    "#d19a66",
	zdpgo_pygments.NumberOct:            "#d19a66",
	zdpgo_pygments.String:               "#98c379",
	zdpgo_pygments.StringAffix:          "#98c379",
	zdpgo_pygments.StringBacktick:       "#98c379",
	zdpgo_pygments.StringDelimiter:      "#98c379",
	zdpgo_pygments.StringDoc:            "#7e97c3",
	zdpgo_pygments.StringDouble:         "#63c381",
	zdpgo_pygments.StringEscape:         "bold #d26464",
	zdpgo_pygments.StringHeredoc:        "#98c379",
	zdpgo_pygments.StringInterpol:       "#98c379",
	zdpgo_pygments.StringOther:          "#70b33f",
	zdpgo_pygments.StringRegex:          "#56b6c2",
	zdpgo_pygments.StringSingle:         "#98c379",
	zdpgo_pygments.StringSymbol:         "#56b6c2",
	zdpgo_pygments.Generic:              "#b0c4de",
	zdpgo_pygments.GenericDeleted:       "#b0c4de",
	zdpgo_pygments.GenericEmph:          "italic",
	zdpgo_pygments.GenericHeading:       "bold #a2cbff",
	zdpgo_pygments.GenericInserted:      "#a6e22e",
	zdpgo_pygments.GenericOutput:        "#a6e22e",
	zdpgo_pygments.GenericUnderline:     "underline",
	zdpgo_pygments.GenericPrompt:        "#a6e22e",
	zdpgo_pygments.GenericStrong:        "bold",
	zdpgo_pygments.GenericSubheading:    "#a2cbff",
	zdpgo_pygments.GenericTraceback:     "#a2cbff",
	zdpgo_pygments.Background:           "#b0c4de bg:#282c34",
}))
