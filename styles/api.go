package styles

import (
	"sort"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Registry of Styles.
var Registry = map[string]*zdpgo_pygments.Style{}

// Fallback style. Reassign to change the default fallback style.
var Fallback = SwapOff

// Register a zdpgo_pygments.Style.
func Register(style *zdpgo_pygments.Style) *zdpgo_pygments.Style {
	Registry[style.Name] = style
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get named style, or Fallback.
func Get(name string) *zdpgo_pygments.Style {
	if style, ok := Registry[name]; ok {
		return style
	}
	return Fallback
}
