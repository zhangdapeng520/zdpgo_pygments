package formatters

import (
	"fmt"
	"io"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", zdpgo_pygments.FormatterFunc(func(w io.Writer, s *zdpgo_pygments.Style, it zdpgo_pygments.Iterator) error {
	for t := it(); t != zdpgo_pygments.EOF; t = it() {
		if _, err := fmt.Fprintln(w, t.GoString()); err != nil {
			return err
		}
	}
	return nil
}))
