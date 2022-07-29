// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint
import (
	"github.com/zhangdapeng520/zdpgo_pygments"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/a"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/b"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/c"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/circular"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/d"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/e"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/f"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/g"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/h"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/i"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers/internal"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/j"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/k"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/l"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/m"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/n"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/o"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/p"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/q"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/r"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/s"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/t"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/v"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/w"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/x"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/y"
	_ "github.com/zhangdapeng520/zdpgo_pygments/lexers/z"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) zdpgo_pygments.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) zdpgo_pygments.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) zdpgo_pygments.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) zdpgo_pygments.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer zdpgo_pygments.Lexer) zdpgo_pygments.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
