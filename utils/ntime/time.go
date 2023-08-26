/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-08-26 10:51:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-08-26 11:20:24
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package ntime

import "time"

// RemainingTimeUntilTomorrow 获取距离次日0点0分0秒的剩余时间
func RemainingTimeUntilTomorrow(t time.Time) (remainingTime time.Duration) {
	// 获取明天的日期
	tomorrow := t.Add(24 * time.Hour)
	tomorrowDate := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
	// 计算剩余时间
	remainingTime = tomorrowDate.Sub(t)
	return
}
