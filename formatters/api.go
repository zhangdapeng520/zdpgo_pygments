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

// Registry 格式化的存储对象
var Registry = map[string]zdpgo_pygments.Formatter{}

// Names 获取所有的格式化名称
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get 根据名字获取格式化对象
func Get(name string) zdpgo_pygments.Formatter {
	if f, ok := Registry[name]; ok {
		return f
	}
	return Fallback
}

// Register 注册一个格式化对象
func Register(name string, formatter zdpgo_pygments.Formatter) zdpgo_pygments.Formatter {
	Registry[name] = formatter
	return formatter
}
