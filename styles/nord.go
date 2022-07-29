package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

var (
	// colors and palettes based on https://www.nordtheme.com/docs/colors-and-palettes
	nord0  = "#2e3440"
	nord1  = "#3b4252" // nolint
	nord2  = "#434c5e" // nolint
	nord3  = "#4c566a"
	nord3b = "#616e87"

	nord4 = "#d8dee9"
	nord5 = "#e5e9f0" // nolint
	nord6 = "#eceff4"

	nord7  = "#8fbcbb"
	nord8  = "#88c0d0"
	nord9  = "#81a1c1"
	nord10 = "#5e81ac"

	nord11 = "#bf616a"
	nord12 = "#d08770"
	nord13 = "#ebcb8b"
	nord14 = "#a3be8c"
	nord15 = "#b48ead"
)

// Nord, an arctic, north-bluish color palette
var Nord = Register(zdpgo_pygments.MustNewStyle("nord", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.TextWhitespace:        nord4,
	zdpgo_pygments.Comment:               "italic " + nord3b,
	zdpgo_pygments.CommentPreproc:        nord10,
	zdpgo_pygments.Keyword:               "bold " + nord9,
	zdpgo_pygments.KeywordPseudo:         "nobold " + nord9,
	zdpgo_pygments.KeywordType:           "nobold " + nord9,
	zdpgo_pygments.Operator:              nord9,
	zdpgo_pygments.OperatorWord:          "bold " + nord9,
	zdpgo_pygments.Name:                  nord4,
	zdpgo_pygments.NameBuiltin:           nord9,
	zdpgo_pygments.NameFunction:          nord8,
	zdpgo_pygments.NameClass:             nord7,
	zdpgo_pygments.NameNamespace:         nord7,
	zdpgo_pygments.NameException:         nord11,
	zdpgo_pygments.NameVariable:          nord4,
	zdpgo_pygments.NameConstant:          nord7,
	zdpgo_pygments.NameLabel:             nord7,
	zdpgo_pygments.NameEntity:            nord12,
	zdpgo_pygments.NameAttribute:         nord7,
	zdpgo_pygments.NameTag:               nord9,
	zdpgo_pygments.NameDecorator:         nord12,
	zdpgo_pygments.Punctuation:           nord6,
	zdpgo_pygments.LiteralString:         nord14,
	zdpgo_pygments.LiteralStringDoc:      nord3b,
	zdpgo_pygments.LiteralStringInterpol: nord14,
	zdpgo_pygments.LiteralStringEscape:   nord13,
	zdpgo_pygments.LiteralStringRegex:    nord13,
	zdpgo_pygments.LiteralStringSymbol:   nord14,
	zdpgo_pygments.LiteralStringOther:    nord14,
	zdpgo_pygments.LiteralNumber:         nord15,
	zdpgo_pygments.GenericHeading:        "bold " + nord8,
	zdpgo_pygments.GenericSubheading:     "bold " + nord8,
	zdpgo_pygments.GenericDeleted:        nord11,
	zdpgo_pygments.GenericInserted:       nord14,
	zdpgo_pygments.GenericError:          nord11,
	zdpgo_pygments.GenericEmph:           "italic",
	zdpgo_pygments.GenericStrong:         "bold",
	zdpgo_pygments.GenericPrompt:         "bold " + nord3,
	zdpgo_pygments.GenericOutput:         nord4,
	zdpgo_pygments.GenericTraceback:      nord11,
	zdpgo_pygments.Error:                 nord11,
	zdpgo_pygments.Background:            nord4 + " bg:" + nord0,
}))
