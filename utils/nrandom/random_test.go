/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 00:19:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:24:27
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nrandom_test

import (
	"github.com/liusuxian/nova/utils/nrandom"
	"sort"
	"testing"
	"time"
)

func TestRandomWeight(t *testing.T) {
	weights := []int{4320, 984, 1201, 1060, 700, 400, 260, 200, 200, 390, 285}
	counts := map[int]int{}

	stime := time.Now() // 获取当前时间
	for i := 0; i < 1000000; i++ {
		index := nrandom.RandomWeight(weights)
		counts[index]++
	}
	elapsed := time.Since(stime)
	t.Logf("TestRandomWeight 执行完成耗时: %v\n", elapsed)

	keys := make([]int, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		t.Logf("TestRandomWeight index: %d count: %d\n", k, counts[k])
	}
}
