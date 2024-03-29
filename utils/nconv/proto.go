/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 19:43:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:44:41
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoMsgToMapE 将 protobuf 消息转换为 Map 类型
func ProtoMsgToMapE(msg proto.Message) (m map[string]any, err error) {
	options := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	var jsonBuf []byte
	if jsonBuf, err = options.Marshal(msg); err != nil {
		return
	}

	m = make(map[string]any)
	err = json.Unmarshal(jsonBuf, &m)
	return
}
