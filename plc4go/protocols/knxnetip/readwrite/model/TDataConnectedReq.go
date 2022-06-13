/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TDataConnectedReq is the data-structure of this message
type TDataConnectedReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// ITDataConnectedReq is the corresponding interface of TDataConnectedReq
type ITDataConnectedReq interface {
	ICEMI
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *TDataConnectedReq) GetMessageCode() uint8 {
	return 0x41
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *TDataConnectedReq) InitializeParent(parent *CEMI) {}

func (m *TDataConnectedReq) GetParent() *CEMI {
	return m.CEMI
}

// NewTDataConnectedReq factory function for TDataConnectedReq
func NewTDataConnectedReq(size uint16) *TDataConnectedReq {
	_result := &TDataConnectedReq{
		CEMI: NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastTDataConnectedReq(structType interface{}) *TDataConnectedReq {
	if casted, ok := structType.(TDataConnectedReq); ok {
		return &casted
	}
	if casted, ok := structType.(*TDataConnectedReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastTDataConnectedReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastTDataConnectedReq(casted.Child)
	}
	return nil
}

func (m *TDataConnectedReq) GetTypeName() string {
	return "TDataConnectedReq"
}

func (m *TDataConnectedReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *TDataConnectedReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *TDataConnectedReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TDataConnectedReqParse(readBuffer utils.ReadBuffer, size uint16) (*TDataConnectedReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataConnectedReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataConnectedReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataConnectedReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataConnectedReq")
	}

	// Create a partially initialized instance
	_child := &TDataConnectedReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *TDataConnectedReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataConnectedReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataConnectedReq")
		}

		if popErr := writeBuffer.PopContext("TDataConnectedReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataConnectedReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *TDataConnectedReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
