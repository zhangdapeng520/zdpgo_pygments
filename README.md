# zdpgo_pygments

词法分析，类似于 Python 的 pygments，专用于各类编程语言的词法分析

## 版本历史

- v0.1.0 2022/08/02 新增：token 化源码
- v0.1.1 2022/08/02 优化：解决 Java 对象属性无法清洗的问题
- v0.1.2 2022/08/03 优化：解决 Python 的双引号连续出现 3 个 S
- v0.1.3 2022/08/04 新增：获取 token 数组和展开 token 数组
- v0.1.4 2022/08/04 优化：获取词法分析器
- v0.1.5 2022/08/04 新增：词法分析支持去除注释
- v0.1.6 2022/08/04 优化：抽离 lexers
- v0.1.7 2022/08/04 升级：词法分析器升级
- v0.1.8 2022/08/08 BUG修复：修复无法解析PHP数组token的BUG
- v0.1.9 2022/08/09 BUG修复：修复PHP词法分析单引号无法正常解析的BUG
- v1.0.0 2022/08/19 优化：移除token相关代码
