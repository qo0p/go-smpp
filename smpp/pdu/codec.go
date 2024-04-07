// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdu

import (
	"bytes"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutlv"
)

var nextSeq uint32

// codec is the base type of all PDUs.
// It implements the PDU interface and provides a generic encoder.
type Codec struct {
	h *Header
	l pdufield.List
	f pdufield.Map
	t pdutlv.Map
}

// init initializes the codec's list and maps and sets the header
// sequence number.
func (pdu *Codec) Init() {
	if pdu.l == nil {
		pdu.l = pdufield.List{}
	}
	pdu.f = make(pdufield.Map)
	pdu.t = make(pdutlv.Map)
	if pdu.h.Seq == 0 { // If Seq not set
		pdu.h.Seq = atomic.AddUint32(&nextSeq, 1)
	}
}

// setup replaces the codec's current maps with the given ones.
func (pdu *Codec) Setup(f pdufield.Map, t pdutlv.Map) {
	pdu.f, pdu.t = f, t
}

// Header implements the PDU interface.
func (pdu *Codec) Header() *Header {
	return pdu.h
}

// Len implements the PDU interface.
func (pdu *Codec) Len() int {
	l := HeaderLen
	for _, f := range pdu.f {
		l += f.Len()
	}
	for _, t := range pdu.t {
		l += t.Len()
	}
	return l
}

// FieldList implements the PDU interface.
func (pdu *Codec) FieldList() pdufield.List {
	return pdu.l
}

// Fields implement the PDU interface.
func (pdu *Codec) Fields() pdufield.Map {
	return pdu.f
}

// Fields implement the PDU interface.
func (pdu *Codec) TLVFields() pdutlv.Map {
	return pdu.t
}

// SerializeTo implements the PDU interface.
func (pdu *Codec) SerializeTo(w io.Writer) error {
	var b bytes.Buffer
	for _, k := range pdu.FieldList() {
		f, ok := pdu.f[k]
		if !ok {
			pdu.f.Set(k, nil)
			f = pdu.f[k]
		}
		if err := f.SerializeTo(&b); err != nil {
			return err
		}
	}
	for _, f := range pdu.TLVFields() {
		if err := f.SerializeTo(&b); err != nil {
			return err
		}
	}
	pdu.h.Len = uint32(pdu.Len())
	err := pdu.h.SerializeTo(w)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, &b)
	return err
}

// Decoder wraps a PDU (e.g. Bind) and the codec together and is
// used for initializing new PDUs with map data decoded off the wire.
type Decoder interface {
	Body
	Setup(f pdufield.Map, t pdutlv.Map)
}

func DecodeFields(pdu Decoder, b []byte) (Body, error) {
	l := pdu.FieldList()
	r := bytes.NewBuffer(b)
	f, err := l.Decode(r)
	if err != nil {
		return nil, err
	}
	t, err := pdutlv.DecodeTLV(r)
	if err != nil {
		return nil, err
	}
	pdu.Setup(f, t)
	return pdu, nil
}

// Decode decodes binary PDU data. It returns a new PDU object, e.g. Bind,
// with header and all fields decoded. The returned PDU can be modified
// and re-serialized to its binary form.
func Decode(r io.Reader) (decoded Body, header *Header, raw []byte, err error) {
	header, err = DecodeHeader(r)
	if err != nil {
		return
	}
	raw = make([]byte, header.Len-HeaderLen)
	_, err = io.ReadFull(r, raw)
	if err != nil {
		return
	}
	switch header.ID {
	case AlertNotificationID:
		// TODO(fiorix): Implement AlertNotification.
	case BindReceiverID, BindTransceiverID, BindTransmitterID:
		decoded, err = DecodeFields(newBind(header), raw)
		return
	case BindReceiverRespID, BindTransceiverRespID, BindTransmitterRespID:
		decoded, err = DecodeFields(newBindResp(header), raw)
		return
	case CancelSMID:
		// TODO(fiorix): Implement CancelSM.
	case CancelSMRespID:
		// TODO(fiorix): Implement CancelSMResp.
	case DataSMID:
		// TODO(fiorix): Implement DataSM.
	case DataSMRespID:
		// TODO(fiorix): Implement DataSMResp.
	case DeliverSMID:
		decoded, err = DecodeFields(newDeliverSM(header), raw)
		return
	case DeliverSMRespID:
		decoded, err = DecodeFields(newDeliverSMResp(header), raw)
		return
	case EnquireLinkID:
		decoded, err = DecodeFields(newEnquireLink(header), raw)
		return
	case EnquireLinkRespID:
		decoded, err = DecodeFields(newEnquireLinkResp(header), raw)
		return
	case GenericNACKID:
		decoded, err = DecodeFields(newGenericNACK(header), raw)
		return
	case OutbindID:
		// TODO(fiorix): Implement Outbind.
	case QuerySMID:
		decoded, err = DecodeFields(newQuerySM(header), raw)
		return
	case QuerySMRespID:
		decoded, err = DecodeFields(newQuerySMResp(header), raw)
		return
	case ReplaceSMID:
		// TODO(fiorix): Implement ReplaceSM.
	case ReplaceSMRespID:
		// TODO(fiorix): Implement ReplaceSMResp.
	case SubmitMultiID:
		decoded, err = DecodeFields(newSubmitMulti(header), raw)
		return
	case SubmitMultiRespID:
		decoded, err = DecodeFields(newSubmitMultiResp(header), raw)
		return
	case SubmitSMID:
		decoded, err = DecodeFields(newSubmitSM(header), raw)
		return
	case SubmitSMRespID:
		decoded, err = DecodeFields(newSubmitSMResp(header), raw)
		return
	case UnbindID:
		decoded, err = DecodeFields(newUnbind(header), raw)
		return
	case UnbindRespID:
		decoded, err = DecodeFields(newUnbindResp(header), raw)
		return
	default:
		err = fmt.Errorf("unknown PDU type: %#x", header.ID)
		return
	}
	err = fmt.Errorf("PDU not implemented: %#x", header.ID)
	return
}
