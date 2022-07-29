package formatters

import (
	"fmt"
	"io"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

// TTY16m is a true-colour terminal formatter.
var TTY16m = Register("terminal16m", zdpgo_pygments.FormatterFunc(trueColourFormatter))

func trueColourFormatter(w io.Writer, style *zdpgo_pygments.Style, it zdpgo_pygments.Iterator) error {
	style = clearBackground(style)
	for token := it(); token != zdpgo_pygments.EOF; token = it() {
		entry := style.Get(token.Type)
		if !entry.IsZero() {
			out := ""
			if entry.Bold == zdpgo_pygments.Yes {
				out += "\033[1m"
			}
			if entry.Underline == zdpgo_pygments.Yes {
				out += "\033[4m"
			}
			if entry.Italic == zdpgo_pygments.Yes {
				out += "\033[3m"
			}
			if entry.Colour.IsSet() {
				out += fmt.Sprintf("\033[38;2;%d;%d;%dm", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
			}
			if entry.Background.IsSet() {
				out += fmt.Sprintf("\033[48;2;%d;%d;%dm", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
			}
			fmt.Fprint(w, out)
		}
		fmt.Fprint(w, token.Value)
		if !entry.IsZero() {
			fmt.Fprint(w, "\033[0m")
		}
	}
	return nil
}
