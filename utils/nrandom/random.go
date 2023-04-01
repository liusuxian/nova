/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 23:27:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-02 01:14:51
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nrandom/random.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nrandom

import (
	"math/rand"
	"sort"
	"time"
)

// RandomWeight 随机权重
func RandomWeight(weights []int) (index int) {
	// 计算前缀和
	length := len(weights)
	if length == 0 {
		return
	}
	prefixSum := make([]int, length)
	prefixSum[0] = weights[0]
	for i := 1; i < length; i++ {
		prefixSum[i] = prefixSum[i-1] + weights[i]
	}
	// 创建一个新的随机数生成器源
	source := rand.NewSource(time.Now().UnixNano())
	// 创建一个 PRNG
	rng := rand.New(source)
	// 生成一个随机权重值
	randomWeight := rng.Intn(prefixSum[length-1])
	// 使用二分查找算法找到随机权重值对应的下标
	return sort.SearchInts(prefixSum, randomWeight)
}
