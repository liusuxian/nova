/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-26 15:33:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-26 16:19:48
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nnet

import (
	"github.com/pkg/errors"
	"net"
)

// IsPrivateIPv4 判断是否是私有 IPv4 地址
func IsPrivateIPv4(ip net.IP) (ok bool) {
	return ip != nil && (ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

// PrivateIPv4 获取私有 IPv4 地址
func PrivateIPv4() (ip net.IP, err error) {
	var as []net.Addr
	as, err = net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip = ipnet.IP.To4()
		if IsPrivateIPv4(ip) {
			return
		}
	}

	err = errors.New("no private ip address")
	return
}
