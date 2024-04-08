// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdu

import (
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutlv"
)

// PDU Types.
const (
	GenericNACKID         ID = 0x80000000
	BindReceiverID        ID = 0x00000001
	BindReceiverRespID    ID = 0x80000001
	BindTransmitterID     ID = 0x00000002
	BindTransmitterRespID ID = 0x80000002
	QuerySMID             ID = 0x00000003
	QuerySMRespID         ID = 0x80000003
	SubmitSMID            ID = 0x00000004
	SubmitSMRespID        ID = 0x80000004
	DeliverSMID           ID = 0x00000005
	DeliverSMRespID       ID = 0x80000005
	UnbindID              ID = 0x00000006
	UnbindRespID          ID = 0x80000006
	ReplaceSMID           ID = 0x00000007
	ReplaceSMRespID       ID = 0x80000007
	CancelSMID            ID = 0x00000008
	CancelSMRespID        ID = 0x80000008
	BindTransceiverID     ID = 0x00000009
	BindTransceiverRespID ID = 0x80000009
	OutbindID             ID = 0x0000000B
	EnquireLinkID         ID = 0x00000015
	EnquireLinkRespID     ID = 0x80000015
	SubmitMultiID         ID = 0x00000021
	SubmitMultiRespID     ID = 0x80000021
	AlertNotificationID   ID = 0x00000102
	DataSMID              ID = 0x00000103
	DataSMRespID          ID = 0x80000103
)

// GenericNACK PDU.
type GenericNACK struct{ *Codec }

func newGenericNACK(hdr *Header) *Codec {
	return &Codec{h: hdr}
}

// NewGenericNACK creates and initializes a GenericNACK PDU.
func NewGenericNACK() Body {
	b := newGenericNACK(&Header{ID: GenericNACKID})
	b.Init()
	return b
}

// Bind PDU.
type Bind struct{ *Codec }

func newBind(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.SystemID,
			pdufield.Password,
			pdufield.SystemType,
			pdufield.InterfaceVersion,
			pdufield.AddrTON,
			pdufield.AddrNPI,
			pdufield.AddressRange,
		},
		r: raw}
}

// NewBindReceiver creates a new Bind PDU.
func NewBindReceiver() Body {
	b := newBind(&Header{ID: BindReceiverID}, nil)
	b.Init()
	return b
}

// NewBindTransceiver creates a new Bind PDU.
func NewBindTransceiver() Body {
	b := newBind(&Header{ID: BindTransceiverID}, nil)
	b.Init()
	return b
}

// NewBindTransmitter creates a new Bind PDU.
func NewBindTransmitter() Body {
	b := newBind(&Header{ID: BindTransmitterID}, nil)
	b.Init()
	return b
}

// BindResp PDU.
type BindResp struct{ *Codec }

func newBindResp(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{pdufield.SystemID},
		r: raw,
	}
}

// NewBindReceiverResp creates and initializes a new BindResp PDU.
func NewBindReceiverResp() Body {
	b := newBindResp(&Header{ID: BindReceiverRespID}, nil)
	b.Init()
	return b
}

// NewBindTransceiverResp creates and initializes a new BindResp PDU.
func NewBindTransceiverResp() Body {
	b := newBindResp(&Header{ID: BindTransceiverRespID}, nil)
	b.Init()
	return b
}

// NewBindTransmitterResp creates and initializes a new BindResp PDU.
func NewBindTransmitterResp() Body {
	b := newBindResp(&Header{ID: BindTransmitterRespID}, nil)
	b.Init()
	return b
}

// QuerySM PDU.
type QuerySM struct{ *Codec }

func newQuerySM(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.MessageID,
			pdufield.SourceAddrTON,
			pdufield.SourceAddrNPI,
			pdufield.SourceAddr,
		},
		r: raw,
	}
}

// NewQuerySM creates and initializes a new QuerySM PDU.
func NewQuerySM() Body {
	b := newQuerySM(&Header{ID: QuerySMID}, nil)
	b.Init()
	return b
}

// QuerySMResp PDU.
type QuerySMResp struct{ *Codec }

func newQuerySMResp(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.MessageID,
			pdufield.FinalDate,
			pdufield.MessageState,
			pdufield.ErrorCode,
		},
		r: raw,
	}
}

// NewQuerySMResp creates and initializes a new QuerySMResp PDU.
func NewQuerySMResp() Body {
	b := newQuerySMResp(&Header{ID: QuerySMRespID}, nil)
	b.Init()
	return b
}

// SubmitSM PDU.
type SubmitSM struct{ *Codec }

func newSubmitSM(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.ServiceType,
			pdufield.SourceAddrTON,
			pdufield.SourceAddrNPI,
			pdufield.SourceAddr,
			pdufield.DestAddrTON,
			pdufield.DestAddrNPI,
			pdufield.DestinationAddr,
			pdufield.ESMClass,
			pdufield.ProtocolID,
			pdufield.PriorityFlag,
			pdufield.ScheduleDeliveryTime,
			pdufield.ValidityPeriod,
			pdufield.RegisteredDelivery,
			pdufield.ReplaceIfPresentFlag,
			pdufield.DataCoding,
			pdufield.SMDefaultMsgID,
			pdufield.SMLength,
			pdufield.ShortMessage,
		},
		r: raw,
	}
}

// NewSubmitSM creates and initializes a new SubmitSM PDU.
func NewSubmitSM(fields pdutlv.Fields) Body {
	b := newSubmitSM(&Header{ID: SubmitSMID}, nil)
	b.Init()
	for tag, value := range fields {
		b.t.Set(tag, value)
	}
	return b
}

// SubmitSMResp PDU.
type SubmitSMResp struct{ *Codec }

func newSubmitSMResp(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.MessageID,
		},
		r: raw,
	}
}

// NewSubmitSMResp creates and initializes a new SubmitSMResp PDU.
func NewSubmitSMResp() Body {
	b := newSubmitSMResp(&Header{ID: SubmitSMRespID}, nil)
	b.Init()
	return b
}

// SubmitMulti PDU.
type SubmitMulti struct{ *Codec }

func newSubmitMulti(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.ServiceType,
			pdufield.SourceAddrTON,
			pdufield.SourceAddrNPI,
			pdufield.SourceAddr,
			pdufield.NumberDests,
			pdufield.DestinationList, // contains DestFlag, DestAddrTON and DestAddrNPI for each address
			pdufield.ESMClass,
			pdufield.ProtocolID,
			pdufield.PriorityFlag,
			pdufield.ScheduleDeliveryTime,
			pdufield.ValidityPeriod,
			pdufield.RegisteredDelivery,
			pdufield.ReplaceIfPresentFlag,
			pdufield.DataCoding,
			pdufield.SMDefaultMsgID,
			pdufield.SMLength,
			pdufield.ShortMessage,
		},
		r: raw,
	}
}

// NewSubmitMulti creates and initializes a new SubmitMulti PDU.
func NewSubmitMulti(fields pdutlv.Fields) Body {
	b := newSubmitMulti(&Header{ID: SubmitMultiID}, nil)
	b.Init()
	for tag, value := range fields {
		b.t.Set(tag, value)
	}
	return b
}

// SubmitMultiResp PDU.
type SubmitMultiResp struct{ *Codec }

func newSubmitMultiResp(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.MessageID,
			pdufield.NoUnsuccess,
			pdufield.UnsuccessSme,
		},
		r: raw,
	}
}

// NewSubmitMultiResp creates and initializes a new SubmitMultiResp PDU.
func NewSubmitMultiResp() Body {
	b := newSubmitMultiResp(&Header{ID: SubmitMultiRespID}, nil)
	b.Init()
	return b
}

// DeliverSM PDU.
type DeliverSM struct{ *Codec }

func newDeliverSM(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.ServiceType,
			pdufield.SourceAddrTON,
			pdufield.SourceAddrNPI,
			pdufield.SourceAddr,
			pdufield.DestAddrTON,
			pdufield.DestAddrNPI,
			pdufield.DestinationAddr,
			pdufield.ESMClass,
			pdufield.ProtocolID,
			pdufield.PriorityFlag,
			pdufield.ScheduleDeliveryTime,
			pdufield.ValidityPeriod,
			pdufield.RegisteredDelivery,
			pdufield.ReplaceIfPresentFlag,
			pdufield.DataCoding,
			pdufield.SMDefaultMsgID,
			pdufield.SMLength,
			pdufield.ShortMessage,
		},
		r: raw,
	}
}

// NewDeliverSM creates and initializes a new DeliverSM PDU.
func NewDeliverSM() Body {
	b := newDeliverSM(&Header{ID: DeliverSMID}, nil)
	b.Init()
	return b
}

// DeliverSMResp PDU.
type DeliverSMResp struct{ *Codec }

func newDeliverSMResp(hdr *Header, raw []byte) *Codec {
	return &Codec{
		h: hdr,
		l: pdufield.List{
			pdufield.MessageID,
		},
		r: raw,
	}
}

// NewDeliverSMResp creates and initializes a new DeliverSMResp PDU.
func NewDeliverSMResp() Body {
	b := newDeliverSMResp(&Header{ID: DeliverSMRespID}, nil)
	b.Init()
	return b
}

// NewDeliverSMRespSeq creates and initializes a new DeliverSMResp PDU for a specific seq.
func NewDeliverSMRespSeq(seq uint32) Body {
	b := newDeliverSMResp(&Header{ID: DeliverSMRespID, Seq: seq}, nil)
	b.Init()
	return b
}

// DataSM PDU.
type DataSM struct{ *Codec }

func newDataSM(hdr *Header, raw []byte) *Codec {
	return NewCodec(
		hdr,
		pdufield.List{
			pdufield.ServiceType,
			pdufield.SourceAddrTON,
			pdufield.SourceAddrNPI,
			pdufield.SourceAddr,
			pdufield.DestAddrTON,
			pdufield.DestAddrNPI,
			pdufield.DestinationAddr,
			pdufield.ESMClass,
			pdufield.RegisteredDelivery,
			pdufield.DataCoding,
		},
		raw,
	)
}

// DataSMResp PDU.
type DataSMResp struct{ *Codec }

func newDataSMResp(hdr *Header, raw []byte) *Codec {
	return NewCodec(
		hdr,
		pdufield.List{
			pdufield.MessageID,
		},
		raw,
	)
}

func NewDataSMResp(seq uint32, messageId string) Body {
	b := newDataSMResp(&Header{ID: DataSMRespID, Seq: seq}, nil)
	b.Init()
	b.TLVFields().Set(pdutlv.TagReceiptedMessageID, messageId)
	return b
}

// Unbind PDU.
type Unbind struct{ *Codec }

func newUnbind(hdr *Header, raw []byte) *Codec {
	return &Codec{h: hdr, r: raw}
}

// NewUnbind creates and initializes a Unbind PDU.
func NewUnbind() Body {
	b := newUnbind(&Header{ID: UnbindID}, nil)
	b.Init()
	return b
}

// UnbindResp PDU.
type UnbindResp struct{ *Codec }

func newUnbindResp(hdr *Header, raw []byte) *Codec {
	return &Codec{h: hdr, r: raw}
}

// NewUnbindResp creates and initializes a UnbindResp PDU.
func NewUnbindResp() Body {
	b := newUnbindResp(&Header{ID: UnbindRespID}, nil)
	b.Init()
	return b
}

// EnquireLink PDU.
type EnquireLink struct{ *Codec }

func newEnquireLink(hdr *Header, raw []byte) *Codec {
	return &Codec{h: hdr, r: raw}
}

// NewEnquireLink creates and initializes a EnquireLink PDU.
func NewEnquireLink() Body {
	b := newEnquireLink(&Header{ID: EnquireLinkID}, nil)
	b.Init()
	return b
}

// EnquireLinkResp PDU.
type EnquireLinkResp struct{ *Codec }

func newEnquireLinkResp(hdr *Header, raw []byte) *Codec {
	return &Codec{h: hdr, r: raw}
}

// NewEnquireLinkResp creates and initializes a EnquireLinkResp PDU.
func NewEnquireLinkResp() Body {
	b := newEnquireLinkResp(&Header{ID: EnquireLinkRespID}, nil)
	b.Init()
	return b
}

// NewEnquireLinkRespSeq creates and initializes a EnquireLinkResp PDU for a specific seq.
func NewEnquireLinkRespSeq(seq uint32) Body {
	b := newEnquireLinkResp(&Header{ID: EnquireLinkRespID, Seq: seq}, nil)
	b.Init()
	return b
}
