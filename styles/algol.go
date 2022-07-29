package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Algol style.
var Algol = Register(zdpgo_pygments.MustNewStyle("algol", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Comment:            "italic #888",
	zdpgo_pygments.CommentPreproc:     "bold noitalic #888",
	zdpgo_pygments.CommentSpecial:     "bold noitalic #888",
	zdpgo_pygments.Keyword:            "underline bold",
	zdpgo_pygments.KeywordDeclaration: "italic",
	zdpgo_pygments.NameBuiltin:        "bold italic",
	zdpgo_pygments.NameBuiltinPseudo:  "bold italic",
	zdpgo_pygments.NameNamespace:      "bold italic #666",
	zdpgo_pygments.NameClass:          "bold italic #666",
	zdpgo_pygments.NameFunction:       "bold italic #666",
	zdpgo_pygments.NameVariable:       "bold italic #666",
	zdpgo_pygments.NameConstant:       "bold italic #666",
	zdpgo_pygments.OperatorWord:       "bold",
	zdpgo_pygments.LiteralString:      "italic #666",
	zdpgo_pygments.Error:              "border:#FF0000",
	zdpgo_pygments.Background:         " bg:#ffffff",
}))
