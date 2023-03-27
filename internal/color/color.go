/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-27 14:16:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-27 14:16:57
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/color/color.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package color

import "fmt"

// 前景色
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color 表示文本的颜色
type Color uint8

// Add 将颜色添加到给定的字符串中
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}
