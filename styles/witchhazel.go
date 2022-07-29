// Copyright 2018 Alethea Katherine Flowers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package styles

import (
	"github.com/zhangdapeng520/zdpgo_pygments"
)

// WitchHazel Style
var WitchHazel = Register(zdpgo_pygments.MustNewStyle("witchhazel", zdpgo_pygments.StyleEntries{
	zdpgo_pygments.Text:              "#F8F8F2",
	zdpgo_pygments.Whitespace:        "#A8757B",
	zdpgo_pygments.Error:             "#960050 bg:#1e0010",
	zdpgo_pygments.Comment:           "#b0bec5",
	zdpgo_pygments.Keyword:           "#C2FFDF",
	zdpgo_pygments.KeywordNamespace:  "#FFB8D1",
	zdpgo_pygments.Operator:          "#FFB8D1",
	zdpgo_pygments.Punctuation:       "#F8F8F2",
	zdpgo_pygments.Name:              "#F8F8F2",
	zdpgo_pygments.NameAttribute:     "#ceb1ff",
	zdpgo_pygments.NameBuiltinPseudo: "#80cbc4",
	zdpgo_pygments.NameClass:         "#ceb1ff",
	zdpgo_pygments.NameConstant:      "#C5A3FF",
	zdpgo_pygments.NameDecorator:     "#ceb1ff",
	zdpgo_pygments.NameException:     "#ceb1ff",
	zdpgo_pygments.NameFunction:      "#ceb1ff",
	zdpgo_pygments.NameProperty:      "#F8F8F2",
	zdpgo_pygments.NameTag:           "#FFB8D1",
	zdpgo_pygments.NameVariable:      "#F8F8F2",
	zdpgo_pygments.Number:            "#C5A3FF",
	zdpgo_pygments.Literal:           "#ae81ff",
	zdpgo_pygments.LiteralDate:       "#e6db74",
	zdpgo_pygments.String:            "#1bc5e0",
	zdpgo_pygments.GenericDeleted:    "#f92672",
	zdpgo_pygments.GenericEmph:       "italic",
	zdpgo_pygments.GenericInserted:   "#a6e22e",
	zdpgo_pygments.GenericStrong:     "bold",
	zdpgo_pygments.GenericSubheading: "#75715e",
	zdpgo_pygments.Background:        " bg:#433e56",
}))
