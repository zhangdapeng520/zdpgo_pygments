package p

import (
	. "github.com/zhangdapeng520/zdpgo_pygments" // nolint
	"github.com/zhangdapeng520/zdpgo_pygments/lexers/internal"
)

var Plaintext = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
		Priority:  0.1,
	},
	internal.PlaintextRules,
))
