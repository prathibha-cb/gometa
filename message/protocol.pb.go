// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Proposal
	Accept
	Commit
	Vote
	FollowerInfo
	EpochAck
	LeaderInfo
	NewLeader
	NewLeaderAck
	LogEntry
	Request
	Abort
	Response
*/
package message

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Proposal struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Txnid            *uint64 `protobuf:"varint,2,req,name=txnid" json:"txnid,omitempty"`
	Fid              *string `protobuf:"bytes,3,req,name=fid" json:"fid,omitempty"`
	ReqId            *uint64 `protobuf:"varint,4,req,name=reqId" json:"reqId,omitempty"`
	OpCode           *uint32 `protobuf:"varint,5,req,name=opCode" json:"opCode,omitempty"`
	Key              *string `protobuf:"bytes,6,req,name=key" json:"key,omitempty"`
	Content          []byte  `protobuf:"bytes,7,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}

func (m *Proposal) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Proposal) GetTxnid() uint64 {
	if m != nil && m.Txnid != nil {
		return *m.Txnid
	}
	return 0
}

func (m *Proposal) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

func (m *Proposal) GetReqId() uint64 {
	if m != nil && m.ReqId != nil {
		return *m.ReqId
	}
	return 0
}

func (m *Proposal) GetOpCode() uint32 {
	if m != nil && m.OpCode != nil {
		return *m.OpCode
	}
	return 0
}

func (m *Proposal) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Proposal) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Accept struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Txnid            *uint64 `protobuf:"varint,2,req,name=txnid" json:"txnid,omitempty"`
	Fid              *string `protobuf:"bytes,3,req,name=fid" json:"fid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Accept) Reset()         { *m = Accept{} }
func (m *Accept) String() string { return proto.CompactTextString(m) }
func (*Accept) ProtoMessage()    {}

func (m *Accept) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Accept) GetTxnid() uint64 {
	if m != nil && m.Txnid != nil {
		return *m.Txnid
	}
	return 0
}

func (m *Accept) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

type Commit struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Txnid            *uint64 `protobuf:"varint,2,req,name=txnid" json:"txnid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}

func (m *Commit) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Commit) GetTxnid() uint64 {
	if m != nil && m.Txnid != nil {
		return *m.Txnid
	}
	return 0
}

type Vote struct {
	Version           *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Round             *uint64 `protobuf:"varint,2,req,name=round" json:"round,omitempty"`
	Status            *uint32 `protobuf:"varint,3,req,name=status" json:"status,omitempty"`
	Epoch             *uint32 `protobuf:"varint,4,req,name=epoch" json:"epoch,omitempty"`
	CndId             *string `protobuf:"bytes,5,req,name=cndId" json:"cndId,omitempty"`
	CndLoggedTxnId    *uint64 `protobuf:"varint,6,req,name=cndLoggedTxnId" json:"cndLoggedTxnId,omitempty"`
	CndCommittedTxnId *uint64 `protobuf:"varint,7,req,name=cndCommittedTxnId" json:"cndCommittedTxnId,omitempty"`
	Solicit           *bool   `protobuf:"varint,8,req,name=solicit" json:"solicit,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}

func (m *Vote) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Vote) GetRound() uint64 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *Vote) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *Vote) GetEpoch() uint32 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

func (m *Vote) GetCndId() string {
	if m != nil && m.CndId != nil {
		return *m.CndId
	}
	return ""
}

func (m *Vote) GetCndLoggedTxnId() uint64 {
	if m != nil && m.CndLoggedTxnId != nil {
		return *m.CndLoggedTxnId
	}
	return 0
}

func (m *Vote) GetCndCommittedTxnId() uint64 {
	if m != nil && m.CndCommittedTxnId != nil {
		return *m.CndCommittedTxnId
	}
	return 0
}

func (m *Vote) GetSolicit() bool {
	if m != nil && m.Solicit != nil {
		return *m.Solicit
	}
	return false
}

type FollowerInfo struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	AcceptedEpoch    *uint32 `protobuf:"varint,2,req,name=acceptedEpoch" json:"acceptedEpoch,omitempty"`
	Fid              *string `protobuf:"bytes,3,req,name=fid" json:"fid,omitempty"`
	Voting           *bool   `protobuf:"varint,4,req,name=voting" json:"voting,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FollowerInfo) Reset()         { *m = FollowerInfo{} }
func (m *FollowerInfo) String() string { return proto.CompactTextString(m) }
func (*FollowerInfo) ProtoMessage()    {}

func (m *FollowerInfo) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *FollowerInfo) GetAcceptedEpoch() uint32 {
	if m != nil && m.AcceptedEpoch != nil {
		return *m.AcceptedEpoch
	}
	return 0
}

func (m *FollowerInfo) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

func (m *FollowerInfo) GetVoting() bool {
	if m != nil && m.Voting != nil {
		return *m.Voting
	}
	return false
}

type EpochAck struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	LastLoggedTxid   *uint64 `protobuf:"varint,2,req,name=lastLoggedTxid" json:"lastLoggedTxid,omitempty"`
	CurrentEpoch     *uint32 `protobuf:"varint,3,req,name=currentEpoch" json:"currentEpoch,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EpochAck) Reset()         { *m = EpochAck{} }
func (m *EpochAck) String() string { return proto.CompactTextString(m) }
func (*EpochAck) ProtoMessage()    {}

func (m *EpochAck) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *EpochAck) GetLastLoggedTxid() uint64 {
	if m != nil && m.LastLoggedTxid != nil {
		return *m.LastLoggedTxid
	}
	return 0
}

func (m *EpochAck) GetCurrentEpoch() uint32 {
	if m != nil && m.CurrentEpoch != nil {
		return *m.CurrentEpoch
	}
	return 0
}

type LeaderInfo struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	AcceptedEpoch    *uint32 `protobuf:"varint,2,req,name=acceptedEpoch" json:"acceptedEpoch,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LeaderInfo) Reset()         { *m = LeaderInfo{} }
func (m *LeaderInfo) String() string { return proto.CompactTextString(m) }
func (*LeaderInfo) ProtoMessage()    {}

func (m *LeaderInfo) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *LeaderInfo) GetAcceptedEpoch() uint32 {
	if m != nil && m.AcceptedEpoch != nil {
		return *m.AcceptedEpoch
	}
	return 0
}

type NewLeader struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	CurrentEpoch     *uint32 `protobuf:"varint,2,req,name=currentEpoch" json:"currentEpoch,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NewLeader) Reset()         { *m = NewLeader{} }
func (m *NewLeader) String() string { return proto.CompactTextString(m) }
func (*NewLeader) ProtoMessage()    {}

func (m *NewLeader) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *NewLeader) GetCurrentEpoch() uint32 {
	if m != nil && m.CurrentEpoch != nil {
		return *m.CurrentEpoch
	}
	return 0
}

type NewLeaderAck struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NewLeaderAck) Reset()         { *m = NewLeaderAck{} }
func (m *NewLeaderAck) String() string { return proto.CompactTextString(m) }
func (*NewLeaderAck) ProtoMessage()    {}

func (m *NewLeaderAck) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type LogEntry struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Txnid            *uint64 `protobuf:"varint,2,req,name=txnid" json:"txnid,omitempty"`
	OpCode           *uint32 `protobuf:"varint,3,req,name=opCode" json:"opCode,omitempty"`
	Key              *string `protobuf:"bytes,4,req,name=key" json:"key,omitempty"`
	Content          []byte  `protobuf:"bytes,5,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}

func (m *LogEntry) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *LogEntry) GetTxnid() uint64 {
	if m != nil && m.Txnid != nil {
		return *m.Txnid
	}
	return 0
}

func (m *LogEntry) GetOpCode() uint32 {
	if m != nil && m.OpCode != nil {
		return *m.OpCode
	}
	return 0
}

func (m *LogEntry) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *LogEntry) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Request struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	ReqId            *uint64 `protobuf:"varint,2,req,name=reqId" json:"reqId,omitempty"`
	OpCode           *uint32 `protobuf:"varint,3,req,name=opCode" json:"opCode,omitempty"`
	Key              *string `protobuf:"bytes,4,req,name=key" json:"key,omitempty"`
	Content          []byte  `protobuf:"bytes,5,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Request) GetReqId() uint64 {
	if m != nil && m.ReqId != nil {
		return *m.ReqId
	}
	return 0
}

func (m *Request) GetOpCode() uint32 {
	if m != nil && m.OpCode != nil {
		return *m.OpCode
	}
	return 0
}

func (m *Request) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Request) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Abort struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	ReqId            *uint64 `protobuf:"varint,2,req,name=reqId" json:"reqId,omitempty"`
	Fid              *string `protobuf:"bytes,3,req,name=fid" json:"fid,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Abort) Reset()         { *m = Abort{} }
func (m *Abort) String() string { return proto.CompactTextString(m) }
func (*Abort) ProtoMessage()    {}

func (m *Abort) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Abort) GetReqId() uint64 {
	if m != nil && m.ReqId != nil {
		return *m.ReqId
	}
	return 0
}

func (m *Abort) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

func (m *Abort) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type Response struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	ReqId            *uint64 `protobuf:"varint,2,req,name=reqId" json:"reqId,omitempty"`
	Fid              *string `protobuf:"bytes,3,req,name=fid" json:"fid,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	Content          []byte  `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Response) GetReqId() uint64 {
	if m != nil && m.ReqId != nil {
		return *m.ReqId
	}
	return 0
}

func (m *Response) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

func (m *Response) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *Response) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
}
