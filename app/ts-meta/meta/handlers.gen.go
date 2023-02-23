// Code generated by tmpl; DO NOT EDIT.
// https://github.com/benbjohnson/tmpl
//
// Source: handlers.gen.go.tmpl

/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package meta

import (
	"github.com/openGemini/openGemini/app/ts-meta/meta/message"
	"github.com/openGemini/openGemini/engine/executor"
	"github.com/openGemini/openGemini/engine/executor/spdy/transport"
)

func New(typ uint8) RPCHandler {
	switch typ {
	case message.PingRequestMessage:
		return &Ping{}
	case message.PeersRequestMessage:
		return &Peers{}
	case message.CreateNodeRequestMessage:
		return &CreateNode{}
	case message.SnapshotRequestMessage:
		return &Snapshot{}
	case message.ExecuteRequestMessage:
		return &Execute{}
	case message.UpdateRequestMessage:
		return &Update{}
	case message.ReportRequestMessage:
		return &Report{}
	case message.GetShardInfoRequestMessage:
		return &GetShardInfo{}
	case message.GetDownSampleInfoRequestMessage:
		return &GetDownSampleInfo{}
	case message.GetRpMstInfosRequestMessage:
		return &GetRpMstInfos{}
	case message.GetUserInfoRequestMessage:
		return &GetUserInfo{}
	case message.GetStreamInfoRequestMessage:
		return &GetStreamInfo{}
	case message.GetMeasurementInfoRequestMessage:
		return &GetMeasurementInfo{}
	default:
		return nil
	}
}

type Ping struct {
	BaseHandler

	req *message.PingRequest
}

func (h *Ping) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.PingRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.PingRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Ping) Instance() RPCHandler {
	return &Ping{}
}

type Peers struct {
	BaseHandler

	req *message.PeersRequest
}

func (h *Peers) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.PeersRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.PeersRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Peers) Instance() RPCHandler {
	return &Peers{}
}

type CreateNode struct {
	BaseHandler

	req *message.CreateNodeRequest
}

func (h *CreateNode) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.CreateNodeRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.CreateNodeRequest", data)
	}
	h.req = msg
	return nil
}

func (h *CreateNode) Instance() RPCHandler {
	return &CreateNode{}
}

type Snapshot struct {
	BaseHandler

	req *message.SnapshotRequest
}

func (h *Snapshot) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.SnapshotRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.SnapshotRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Snapshot) Instance() RPCHandler {
	return &Snapshot{}
}

type Execute struct {
	BaseHandler

	req *message.ExecuteRequest
}

func (h *Execute) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.ExecuteRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.ExecuteRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Execute) Instance() RPCHandler {
	return &Execute{}
}

type Update struct {
	BaseHandler

	req *message.UpdateRequest
}

func (h *Update) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.UpdateRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.UpdateRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Update) Instance() RPCHandler {
	return &Update{}
}

type Report struct {
	BaseHandler

	req *message.ReportRequest
}

func (h *Report) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.ReportRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.ReportRequest", data)
	}
	h.req = msg
	return nil
}

func (h *Report) Instance() RPCHandler {
	return &Report{}
}

type GetShardInfo struct {
	BaseHandler

	req *message.GetShardInfoRequest
}

func (h *GetShardInfo) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetShardInfoRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetShardInfoRequest", data)
	}
	h.req = msg
	return nil
}

func (h *GetShardInfo) Instance() RPCHandler {
	return &GetShardInfo{}
}

type GetDownSampleInfo struct {
	BaseHandler
	req *message.GetDownSampleInfoRequest
}

func (h *GetDownSampleInfo) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetDownSampleInfoRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetDownSampleInfoRequest", data)
	}
	h.req = msg
	return nil
}

func (h *GetDownSampleInfo) Instance() RPCHandler {
	return &GetDownSampleInfo{}
}

type GetRpMstInfos struct {
	BaseHandler
	req *message.GetRpMstInfosRequest
}

func (h *GetRpMstInfos) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetRpMstInfosRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetRpMstInfosRequest", data)
	}
	h.req = msg
	return nil
}
func (h *GetRpMstInfos) Instance() RPCHandler {
	return &GetRpMstInfos{}
}

type GetUserInfo struct {
	BaseHandler

	req *message.GetUserInfoRequest
}

func (h *GetUserInfo) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetUserInfoRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetUserInfoRequest", data)
	}
	h.req = msg
	return nil
}

func (h *GetUserInfo) Instance() RPCHandler {
	return &GetUserInfo{}
}

type GetStreamInfo struct {
	BaseHandler

	req *message.GetStreamInfoRequest
}

func (h *GetStreamInfo) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetStreamInfoRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetStreamInfoRequest", data)
	}
	h.req = msg
	return nil
}

func (h *GetStreamInfo) Instance() RPCHandler {
	return &GetStreamInfo{}
}

type GetMeasurementInfo struct {
	BaseHandler

	req *message.GetMeasurementInfoRequest
}

func (h *GetMeasurementInfo) SetRequestMsg(data transport.Codec) error {
	msg, ok := data.(*message.GetMeasurementInfoRequest)
	if !ok {
		return executor.NewInvalidTypeError("*message.GetMeasurementInfoRequest", data)
	}
	h.req = msg
	return nil
}

func (h *GetMeasurementInfo) Instance() RPCHandler {
	return &GetMeasurementInfo{}
}
