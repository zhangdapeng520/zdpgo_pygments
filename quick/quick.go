package quick

import (
	"io"

	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/formatters"
	"github.com/zhangdapeng520/zdpgo_pygments/styles"
)

// Highlight some text.
//
// Lexer, formatter and style may be empty, in which case a best-effort is made.
func Highlight(w io.Writer, source, lexer, formatter, style string) error {
	// Determine lexer.
	l := zdpgo_lexers.Get(lexer)
	if l == nil {
		l = zdpgo_lexers.Analyse(source)
	}
	if l == nil {
		l = zdpgo_lexers.Fallback
	}
	l = zdpgo_pygments.Coalesce(l)

	// Determine formatter.
	f := formatters.Get(formatter)
	if f == nil {
		f = formatters.Fallback
	}

	// Determine style.
	s := styles.Get(style)
	if s == nil {
		s = styles.Fallback
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return f.Format(w, s, it)
}
