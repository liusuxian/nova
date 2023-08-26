/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-08-26 10:51:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-08-26 11:19:39
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package ntime_test

import (
	"github.com/liusuxian/nova/utils/ntime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRemainingTimeUntilTomorrow(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(time.Hour, ntime.RemainingTimeUntilTomorrow(time.Date(2022, 12, 30, 23, 0, 0, 0, time.UTC)))
	assert.Equal(time.Hour, ntime.RemainingTimeUntilTomorrow(time.Date(2022, 12, 31, 23, 0, 0, 0, time.UTC)))
	assert.Equal(time.Hour, ntime.RemainingTimeUntilTomorrow(time.Date(2023, 1, 1, 23, 0, 0, 0, time.UTC)))
}
