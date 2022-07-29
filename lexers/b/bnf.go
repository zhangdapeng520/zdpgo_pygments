package b

import (
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
	"github.com/zhangdapeng520/zdpgo_pygments/lexers/internal"
)

// Bnf lexer.
var Bnf = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "BNF",
		Aliases:   []string{"bnf"},
		Filenames: []string{"*.bnf"},
		MimeTypes: []string{"text/x-bnf"},
	},
	bnfRules,
))

func bnfRules() Rules {
	return Rules{
		"root": {
			{`(<)([ -;=?-~]+)(>)`, ByGroups(Punctuation, NameClass, Punctuation), nil},
			{`::=`, Operator, nil},
			{`[^<>:]+`, Text, nil},
			{`.`, Text, nil},
		},
	}
}
