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

// ApduDataExtDomainAddressSerialNumberResponse is the data-structure of this message
type ApduDataExtDomainAddressSerialNumberResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtDomainAddressSerialNumberResponse is the corresponding interface of ApduDataExtDomainAddressSerialNumberResponse
type IApduDataExtDomainAddressSerialNumberResponse interface {
	IApduDataExt
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

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetExtApciType() uint8 {
	return 0x2D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtDomainAddressSerialNumberResponse) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtDomainAddressSerialNumberResponse factory function for ApduDataExtDomainAddressSerialNumberResponse
func NewApduDataExtDomainAddressSerialNumberResponse(length uint8) *ApduDataExtDomainAddressSerialNumberResponse {
	_result := &ApduDataExtDomainAddressSerialNumberResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtDomainAddressSerialNumberResponse(structType interface{}) *ApduDataExtDomainAddressSerialNumberResponse {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtDomainAddressSerialNumberResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtDomainAddressSerialNumberResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberResponse"
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtDomainAddressSerialNumberResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtDomainAddressSerialNumberResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSerialNumberResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSerialNumberResponse")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtDomainAddressSerialNumberResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSerialNumberResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSerialNumberResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
