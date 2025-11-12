# Changelog

## v2.3.0

go-logger-limit v1.3.0 功能限定版功能修改是不完全正确的，替换了太多，并且做了包装，不合适；按原来的逻辑强制了控制台颜色色配置。

* **增加了默的ConsoleColor函数** 
  - 使用较为可靠的"github.com/fatih/color"
  - 控制台输出根据日志级别显示不同颜色：DEBUG(hi蓝色)、INFO(绿色)、WARN(黄色)、ERROR(hi红色)、FATAL(玫红)

* **增加示例fullFunc.go新增了控制台输出根据日志级别显示不同颜色的功能**
  - 修改logger\option.go的AttrFormat struct 增加SetConsoleColorFmt专控制台着色
  - 🆕 控制台输出根据日志级别显示不同颜色（与默认一致）：DEBUG(hi蓝色)、INFO(绿色)、WARN(黄色)、ERROR(hi红色)、FATAL(玫红)
  - 💄 日志级别格式调整：将[DEBUG],[INFO],[WARN],[ERROR],[FATAL]更改为固定长度的[DBG],[INF],[WRN],[ERR],[FAT]
  - 🆕 日志时间精确到毫秒（如 01:23:23.123）
  - 🆕 使用 formatter := "[{level} {time}] {message}\n"作为日志格式。效果——> [INF 20:19:49.073] this is a info message

## v1.3.0
- 🆕 控制台输出根据日志级别显示不同颜色：DEBUG(蓝色)、INFO(绿色)、WARN(黄色)、ERROR(红色)、FATAL(红色背景)

## v1.2.0
- 💄 日志级别格式调整：将[DEBUG],[INFO],[WARN],[ERROR],[FATAL]更改为固定长度的[DBG],[INF],[WRN],[ERR],[FAT]

## v1.1.0
- 🆕 新增毫秒时间格式化选项 FORMAT_MILLISECONDS，支持日志时间精确到毫秒（如 01:23:23.123）

## v1.0.0
- 初始版本发布