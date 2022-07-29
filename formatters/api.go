package formatters

import (
	"io"
	"sort"

	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/formatters/html"
	"github.com/zhangdapeng520/zdpgo_pygments/formatters/svg"
)

var (
	// NoOp formatter.
	NoOp = Register("noop", zdpgo_pygments.FormatterFunc(func(w io.Writer, s *zdpgo_pygments.Style, iterator zdpgo_pygments.Iterator) error {
		for t := iterator(); t != zdpgo_pygments.EOF; t = iterator() {
			if _, err := io.WriteString(w, t.Value); err != nil {
				return err
			}
		}
		return nil
	}))
	// Default HTML formatter outputs self-contained HTML.
	htmlFull = Register("html", html.New(html.Standalone(true), html.WithClasses(true))) // nolint
	SVG      = Register("svg", svg.New(svg.EmbedFont("Liberation Mono", svg.FontLiberationMono, svg.WOFF)))
)

// Fallback formatter.
var Fallback = NoOp

// Registry of Formatters.
var Registry = map[string]zdpgo_pygments.Formatter{}

// Names of registered formatters.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get formatter by name.
//
// If the given formatter is not found, the Fallback formatter will be returned.
func Get(name string) zdpgo_pygments.Formatter {
	if f, ok := Registry[name]; ok {
		return f
	}
	return Fallback
}

// Register a named formatter.
func Register(name string, formatter zdpgo_pygments.Formatter) zdpgo_pygments.Formatter {
	Registry[name] = formatter
	return formatter
}
