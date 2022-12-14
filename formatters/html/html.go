package html

import (
	"fmt"
	"html"
	"io"
	"sort"
	"strings"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

// Option sets an option of the HTML formatter.
type Option func(f *Formatter)

// Standalone configures the HTML formatter for generating a standalone HTML document.
func Standalone(b bool) Option { return func(f *Formatter) { f.standalone = b } }

// ClassPrefix sets the CSS class prefix.
func ClassPrefix(prefix string) Option { return func(f *Formatter) { f.prefix = prefix } }

// WithClasses emits HTML using CSS classes, rather than inline styles.
func WithClasses(b bool) Option { return func(f *Formatter) { f.Classes = b } }

// WithAllClasses disables an optimisation that omits redundant CSS classes.
func WithAllClasses(b bool) Option { return func(f *Formatter) { f.allClasses = b } }

// TabWidth sets the number of characters for a tab. Defaults to 8.
func TabWidth(width int) Option { return func(f *Formatter) { f.tabWidth = width } }

// PreventSurroundingPre prevents the surrounding pre tags around the generated code.
func PreventSurroundingPre(b bool) Option {
	return func(f *Formatter) {
		if b {
			f.preWrapper = nopPreWrapper
		} else {
			f.preWrapper = defaultPreWrapper
		}
	}
}

// WithPreWrapper allows control of the surrounding pre tags.
func WithPreWrapper(wrapper PreWrapper) Option {
	return func(f *Formatter) {
		f.preWrapper = wrapper
	}
}

// WrapLongLines wraps long lines.
func WrapLongLines(b bool) Option {
	return func(f *Formatter) {
		f.wrapLongLines = b
	}
}

// WithLineNumbers formats output with line numbers.
func WithLineNumbers(b bool) Option {
	return func(f *Formatter) {
		f.lineNumbers = b
	}
}

// LineNumbersInTable will, when combined with WithLineNumbers, separate the line numbers
// and code in table td's, which make them copy-and-paste friendly.
func LineNumbersInTable(b bool) Option {
	return func(f *Formatter) {
		f.lineNumbersInTable = b
	}
}

// LinkableLineNumbers decorates the line numbers HTML elements with an "id"
// attribute so they can be linked.
func LinkableLineNumbers(b bool, prefix string) Option {
	return func(f *Formatter) {
		f.linkableLineNumbers = b
		f.lineNumbersIDPrefix = prefix
	}
}

// HighlightLines higlights the given line ranges with the Highlight style.
//
// A range is the beginning and ending of a range as 1-based line numbers, inclusive.
func HighlightLines(ranges [][2]int) Option {
	return func(f *Formatter) {
		f.highlightRanges = ranges
		sort.Sort(f.highlightRanges)
	}
}

// BaseLineNumber sets the initial number to start line numbering at. Defaults to 1.
func BaseLineNumber(n int) Option {
	return func(f *Formatter) {
		f.baseLineNumber = n
	}
}

// New HTML formatter.
func New(options ...Option) *Formatter {
	f := &Formatter{
		baseLineNumber: 1,
		preWrapper:     defaultPreWrapper,
	}
	for _, option := range options {
		option(f)
	}
	return f
}

// PreWrapper defines the operations supported in WithPreWrapper.
type PreWrapper interface {
	// Start is called to write a start <pre> element.
	// The code flag tells whether this block surrounds
	// highlighted code. This will be false when surrounding
	// line numbers.
	Start(code bool, styleAttr string) string

	// End is called to write the end </pre> element.
	End(code bool) string
}

type preWrapper struct {
	start func(code bool, styleAttr string) string
	end   func(code bool) string
}

func (p preWrapper) Start(code bool, styleAttr string) string {
	return p.start(code, styleAttr)
}

func (p preWrapper) End(code bool) string {
	return p.end(code)
}

var (
	nopPreWrapper = preWrapper{
		start: func(code bool, styleAttr string) string { return "" },
		end:   func(code bool) string { return "" },
	}
	defaultPreWrapper = preWrapper{
		start: func(code bool, styleAttr string) string {
			if code {
				return fmt.Sprintf(`<pre tabindex="0"%s><code>`, styleAttr)
			}

			return fmt.Sprintf(`<pre tabindex="0"%s>`, styleAttr)
		},
		end: func(code bool) string {
			if code {
				return `</code></pre>`
			}

			return `</pre>`
		},
	}
)

// Formatter that generates HTML.
type Formatter struct {
	standalone          bool
	prefix              string
	Classes             bool // Exported field to detect when classes are being used
	allClasses          bool
	preWrapper          PreWrapper
	tabWidth            int
	wrapLongLines       bool
	lineNumbers         bool
	lineNumbersInTable  bool
	linkableLineNumbers bool
	lineNumbersIDPrefix string
	highlightRanges     highlightRanges
	baseLineNumber      int
}

type highlightRanges [][2]int

func (h highlightRanges) Len() int           { return len(h) }
func (h highlightRanges) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h highlightRanges) Less(i, j int) bool { return h[i][0] < h[j][0] }

func (f *Formatter) Format(w io.Writer, style *zdpgo_pygments.Style, iterator zdpgo_pygments.Iterator) (err error) {
	return f.writeHTML(w, style, iterator.Tokens())
}

// We deliberately don't use html/template here because it is two orders of magnitude slower (benchmarked).
//
// OTOH we need to be super careful about correct escaping...
func (f *Formatter) writeHTML(w io.Writer, style *zdpgo_pygments.Style, tokens []zdpgo_pygments.Token) (err error) { // nolint: gocyclo
	css := f.styleToCSS(style)
	if !f.Classes {
		for t, style := range css {
			css[t] = compressStyle(style)
		}
	}
	if f.standalone {
		fmt.Fprint(w, "<html>\n")
		if f.Classes {
			fmt.Fprint(w, "<style type=\"text/css\">\n")
			err = f.WriteCSS(w, style)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "body { %s; }\n", css[zdpgo_pygments.Background])
			fmt.Fprint(w, "</style>")
		}
		fmt.Fprintf(w, "<body%s>\n", f.styleAttr(css, zdpgo_pygments.Background))
	}

	wrapInTable := f.lineNumbers && f.lineNumbersInTable

	lines := zdpgo_pygments.SplitTokensIntoLines(tokens)
	lineDigits := len(fmt.Sprintf("%d", f.baseLineNumber+len(lines)-1))
	highlightIndex := 0

	if wrapInTable {
		// List line numbers in its own <td>
		fmt.Fprintf(w, "<div%s>\n", f.styleAttr(css, zdpgo_pygments.PreWrapper))
		fmt.Fprintf(w, "<table%s><tr>", f.styleAttr(css, zdpgo_pygments.LineTable))
		fmt.Fprintf(w, "<td%s>\n", f.styleAttr(css, zdpgo_pygments.LineTableTD))
		fmt.Fprintf(w, f.preWrapper.Start(false, f.styleAttr(css, zdpgo_pygments.PreWrapper)))
		for index := range lines {
			line := f.baseLineNumber + index
			highlight, next := f.shouldHighlight(highlightIndex, line)
			if next {
				highlightIndex++
			}
			if highlight {
				fmt.Fprintf(w, "<span%s>", f.styleAttr(css, zdpgo_pygments.LineHighlight))
			}

			fmt.Fprintf(w, "<span%s%s>%s\n</span>", f.styleAttr(css, zdpgo_pygments.LineNumbersTable), f.lineIDAttribute(line), f.lineTitleWithLinkIfNeeded(lineDigits, line))

			if highlight {
				fmt.Fprintf(w, "</span>")
			}
		}
		fmt.Fprint(w, f.preWrapper.End(false))
		fmt.Fprint(w, "</td>\n")
		fmt.Fprintf(w, "<td%s>\n", f.styleAttr(css, zdpgo_pygments.LineTableTD, "width:100%"))
	}

	fmt.Fprintf(w, f.preWrapper.Start(true, f.styleAttr(css, zdpgo_pygments.PreWrapper)))

	highlightIndex = 0
	for index, tokens := range lines {
		// 1-based line number.
		line := f.baseLineNumber + index
		highlight, next := f.shouldHighlight(highlightIndex, line)
		if next {
			highlightIndex++
		}

		// Start of Line
		fmt.Fprint(w, `<span`)
		if highlight {
			// Line + LineHighlight
			if f.Classes {
				fmt.Fprintf(w, ` class="%s %s"`, f.class(zdpgo_pygments.Line), f.class(zdpgo_pygments.LineHighlight))
			} else {
				fmt.Fprintf(w, ` style="%s %s"`, css[zdpgo_pygments.Line], css[zdpgo_pygments.LineHighlight])
			}
			fmt.Fprint(w, `>`)
		} else {
			fmt.Fprintf(w, "%s>", f.styleAttr(css, zdpgo_pygments.Line))
		}

		// Line number
		if f.lineNumbers && !wrapInTable {
			fmt.Fprintf(w, "<span%s%s>%s</span>", f.styleAttr(css, zdpgo_pygments.LineNumbers), f.lineIDAttribute(line), f.lineTitleWithLinkIfNeeded(lineDigits, line))
		}

		fmt.Fprintf(w, `<span%s>`, f.styleAttr(css, zdpgo_pygments.CodeLine))

		for _, token := range tokens {
			html := html.EscapeString(token.String())
			attr := f.styleAttr(css, token.Type)
			if attr != "" {
				html = fmt.Sprintf("<span%s>%s</span>", attr, html)
			}
			fmt.Fprint(w, html)
		}

		fmt.Fprint(w, `</span>`) // End of CodeLine

		fmt.Fprint(w, `</span>`) // End of Line
	}

	fmt.Fprintf(w, f.preWrapper.End(true))

	if wrapInTable {
		fmt.Fprint(w, "</td></tr></table>\n")
		fmt.Fprint(w, "</div>\n")
	}

	if f.standalone {
		fmt.Fprint(w, "\n</body>\n")
		fmt.Fprint(w, "</html>\n")
	}

	return nil
}

func (f *Formatter) lineIDAttribute(line int) string {
	if !f.linkableLineNumbers {
		return ""
	}
	return fmt.Sprintf(" id=\"%s\"", f.lineID(line))
}

func (f *Formatter) lineTitleWithLinkIfNeeded(lineDigits, line int) string {
	title := fmt.Sprintf("%*d", lineDigits, line)
	if !f.linkableLineNumbers {
		return title
	}
	return fmt.Sprintf("<a style=\"outline: none; text-decoration:none; color:inherit\" href=\"#%s\">%s</a>", f.lineID(line), title)
}

func (f *Formatter) lineID(line int) string {
	return fmt.Sprintf("%s%d", f.lineNumbersIDPrefix, line)
}

func (f *Formatter) shouldHighlight(highlightIndex, line int) (bool, bool) {
	next := false
	for highlightIndex < len(f.highlightRanges) && line > f.highlightRanges[highlightIndex][1] {
		highlightIndex++
		next = true
	}
	if highlightIndex < len(f.highlightRanges) {
		hrange := f.highlightRanges[highlightIndex]
		if line >= hrange[0] && line <= hrange[1] {
			return true, next
		}
	}
	return false, next
}

func (f *Formatter) class(t zdpgo_pygments.TokenType) string {
	for t != 0 {
		if cls, ok := zdpgo_pygments.StandardTypes[t]; ok {
			if cls != "" {
				return f.prefix + cls
			}
			return ""
		}
		t = t.Parent()
	}
	if cls := zdpgo_pygments.StandardTypes[t]; cls != "" {
		return f.prefix + cls
	}
	return ""
}

func (f *Formatter) styleAttr(styles map[zdpgo_pygments.TokenType]string, tt zdpgo_pygments.TokenType, extraCSS ...string) string {
	if f.Classes {
		cls := f.class(tt)
		if cls == "" {
			return ""
		}
		return fmt.Sprintf(` class="%s"`, cls)
	}
	if _, ok := styles[tt]; !ok {
		tt = tt.SubCategory()
		if _, ok := styles[tt]; !ok {
			tt = tt.Category()
			if _, ok := styles[tt]; !ok {
				return ""
			}
		}
	}
	css := []string{styles[tt]}
	css = append(css, extraCSS...)
	return fmt.Sprintf(` style="%s"`, strings.Join(css, ";"))
}

func (f *Formatter) tabWidthStyle() string {
	if f.tabWidth != 0 && f.tabWidth != 8 {
		return fmt.Sprintf("; -moz-tab-size: %[1]d; -o-tab-size: %[1]d; tab-size: %[1]d", f.tabWidth)
	}
	return ""
}

// WriteCSS writes CSS style definitions (without any surrounding HTML).
func (f *Formatter) WriteCSS(w io.Writer, style *zdpgo_pygments.Style) error {
	css := f.styleToCSS(style)
	// Special-case background as it is mapped to the outer ".zdpgo_pygments" class.
	if _, err := fmt.Fprintf(w, "/* %s */ .%sbg { %s }\n", zdpgo_pygments.Background, f.prefix, css[zdpgo_pygments.Background]); err != nil {
		return err
	}
	// Special-case PreWrapper as it is the ".zdpgo_pygments" class.
	if _, err := fmt.Fprintf(w, "/* %s */ .%szdpgo_pygments { %s }\n", zdpgo_pygments.PreWrapper, f.prefix, css[zdpgo_pygments.PreWrapper]); err != nil {
		return err
	}
	// Special-case code column of table to expand width.
	if f.lineNumbers && f.lineNumbersInTable {
		if _, err := fmt.Fprintf(w, "/* %s */ .%szdpgo_pygments .%s:last-child { width: 100%%; }",
			zdpgo_pygments.LineTableTD, f.prefix, f.class(zdpgo_pygments.LineTableTD)); err != nil {
			return err
		}
	}
	// Special-case line number highlighting when targeted.
	if f.lineNumbers || f.lineNumbersInTable {
		targetedLineCSS := StyleEntryToCSS(style.Get(zdpgo_pygments.LineHighlight))
		for _, tt := range []zdpgo_pygments.TokenType{zdpgo_pygments.LineNumbers, zdpgo_pygments.LineNumbersTable} {
			fmt.Fprintf(w, "/* %s targeted by URL anchor */ .%szdpgo_pygments .%s:target { %s }\n", tt, f.prefix, f.class(tt), targetedLineCSS)
		}
	}
	tts := []int{}
	for tt := range css {
		tts = append(tts, int(tt))
	}
	sort.Ints(tts)
	for _, ti := range tts {
		tt := zdpgo_pygments.TokenType(ti)
		switch tt {
		case zdpgo_pygments.Background, zdpgo_pygments.PreWrapper:
			continue
		}
		class := f.class(tt)
		if class == "" {
			continue
		}
		styles := css[tt]
		if _, err := fmt.Fprintf(w, "/* %s */ .%szdpgo_pygments .%s { %s }\n", tt, f.prefix, class, styles); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) styleToCSS(style *zdpgo_pygments.Style) map[zdpgo_pygments.TokenType]string {
	classes := map[zdpgo_pygments.TokenType]string{}
	bg := style.Get(zdpgo_pygments.Background)
	// Convert the style.
	for t := range zdpgo_pygments.StandardTypes {
		entry := style.Get(t)
		if t != zdpgo_pygments.Background {
			entry = entry.Sub(bg)
		}
		if !f.allClasses && entry.IsZero() {
			continue
		}
		classes[t] = StyleEntryToCSS(entry)
	}
	classes[zdpgo_pygments.Background] += f.tabWidthStyle()
	classes[zdpgo_pygments.PreWrapper] += classes[zdpgo_pygments.Background] + `;`
	// Make PreWrapper a grid to show highlight style with full width.
	if len(f.highlightRanges) > 0 {
		classes[zdpgo_pygments.PreWrapper] += `display: grid;`
	}
	// Make PreWrapper wrap long lines.
	if f.wrapLongLines {
		classes[zdpgo_pygments.PreWrapper] += `white-space: pre-wrap; word-break: break-word;`
	}
	lineNumbersStyle := `white-space: pre; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;`
	// All rules begin with default rules followed by user provided rules
	classes[zdpgo_pygments.Line] = `display: flex;` + classes[zdpgo_pygments.Line]
	classes[zdpgo_pygments.LineNumbers] = lineNumbersStyle + classes[zdpgo_pygments.LineNumbers]
	classes[zdpgo_pygments.LineNumbersTable] = lineNumbersStyle + classes[zdpgo_pygments.LineNumbersTable]
	classes[zdpgo_pygments.LineTable] = "border-spacing: 0; padding: 0; margin: 0; border: 0;" + classes[zdpgo_pygments.LineTable]
	classes[zdpgo_pygments.LineTableTD] = "vertical-align: top; padding: 0; margin: 0; border: 0;" + classes[zdpgo_pygments.LineTableTD]
	return classes
}

// StyleEntryToCSS converts a zdpgo_pygments.StyleEntry to CSS attributes.
func StyleEntryToCSS(e zdpgo_pygments.StyleEntry) string {
	styles := []string{}
	if e.Colour.IsSet() {
		styles = append(styles, "color: "+e.Colour.String())
	}
	if e.Background.IsSet() {
		styles = append(styles, "background-color: "+e.Background.String())
	}
	if e.Bold == zdpgo_pygments.Yes {
		styles = append(styles, "font-weight: bold")
	}
	if e.Italic == zdpgo_pygments.Yes {
		styles = append(styles, "font-style: italic")
	}
	if e.Underline == zdpgo_pygments.Yes {
		styles = append(styles, "text-decoration: underline")
	}
	return strings.Join(styles, "; ")
}

// Compress CSS attributes - remove spaces, transform 6-digit colours to 3.
func compressStyle(s string) string {
	parts := strings.Split(s, ";")
	out := []string{}
	for _, p := range parts {
		p = strings.Join(strings.Fields(p), " ")
		p = strings.Replace(p, ": ", ":", 1)
		if strings.Contains(p, "#") {
			c := p[len(p)-6:]
			if c[0] == c[1] && c[2] == c[3] && c[4] == c[5] {
				p = p[:len(p)-6] + c[0:1] + c[2:3] + c[4:5]
			}
		}
		out = append(out, p)
	}
	return strings.Join(out, ";")
}
