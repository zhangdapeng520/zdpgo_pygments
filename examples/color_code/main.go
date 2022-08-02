package main

import (
	"bytes"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pygments/formatters"
	"github.com/zhangdapeng520/zdpgo_pygments/quick"
	"github.com/zhangdapeng520/zdpgo_pygments/styles"
	"log"
)

func main() {
	// 源码
	s := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
`
	res := colorCode(s) // example 返回，高亮字符
	fmt.Println(res)
}

// 接 下面的 example-*
func colorCode(s string) string {
	res := new(bytes.Buffer) // 放 高亮的字符

	style := styles.Get("dark") // 主题设置
	if style == nil {
		fmt.Println(style)
		style = styles.Fallback // 默认主题
	}

	attr := "terminal"                // 终端
	formatter := formatters.Get(attr) // 环境的变量，默认为`html`
	if formatter == nil {
		formatter = formatters.Fallback // 默认格式化
	}

	err := quick.Highlight(res, s, "go", attr, style.Name) // 直接高亮
	// 1. 放哪，2. 源码，3. 语言，4. 格式环境，5. 主题名
	if err != nil {
		log.Fatalln(err)
	}
	return res.String()
}
