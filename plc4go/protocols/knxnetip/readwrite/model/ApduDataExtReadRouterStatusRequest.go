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

// ApduDataExtReadRouterStatusRequest is the data-structure of this message
type ApduDataExtReadRouterStatusRequest struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtReadRouterStatusRequest is the corresponding interface of ApduDataExtReadRouterStatusRequest
type IApduDataExtReadRouterStatusRequest interface {
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

func (m *ApduDataExtReadRouterStatusRequest) GetExtApciType() uint8 {
	return 0x0D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtReadRouterStatusRequest) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtReadRouterStatusRequest) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtReadRouterStatusRequest factory function for ApduDataExtReadRouterStatusRequest
func NewApduDataExtReadRouterStatusRequest(length uint8) *ApduDataExtReadRouterStatusRequest {
	_result := &ApduDataExtReadRouterStatusRequest{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtReadRouterStatusRequest(structType interface{}) *ApduDataExtReadRouterStatusRequest {
	if casted, ok := structType.(ApduDataExtReadRouterStatusRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtReadRouterStatusRequest(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtReadRouterStatusRequest(casted.Child)
	}
	return nil
}

func (m *ApduDataExtReadRouterStatusRequest) GetTypeName() string {
	return "ApduDataExtReadRouterStatusRequest"
}

func (m *ApduDataExtReadRouterStatusRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtReadRouterStatusRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtReadRouterStatusRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtReadRouterStatusRequestParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtReadRouterStatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterStatusRequest")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtReadRouterStatusRequest{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtReadRouterStatusRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterStatusRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtReadRouterStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
