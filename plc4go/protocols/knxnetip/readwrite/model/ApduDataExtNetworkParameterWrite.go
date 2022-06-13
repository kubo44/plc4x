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

// ApduDataExtNetworkParameterWrite is the data-structure of this message
type ApduDataExtNetworkParameterWrite struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtNetworkParameterWrite is the corresponding interface of ApduDataExtNetworkParameterWrite
type IApduDataExtNetworkParameterWrite interface {
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

func (m *ApduDataExtNetworkParameterWrite) GetExtApciType() uint8 {
	return 0x24
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtNetworkParameterWrite) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtNetworkParameterWrite) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtNetworkParameterWrite factory function for ApduDataExtNetworkParameterWrite
func NewApduDataExtNetworkParameterWrite(length uint8) *ApduDataExtNetworkParameterWrite {
	_result := &ApduDataExtNetworkParameterWrite{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtNetworkParameterWrite(structType interface{}) *ApduDataExtNetworkParameterWrite {
	if casted, ok := structType.(ApduDataExtNetworkParameterWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtNetworkParameterWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtNetworkParameterWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataExtNetworkParameterWrite) GetTypeName() string {
	return "ApduDataExtNetworkParameterWrite"
}

func (m *ApduDataExtNetworkParameterWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtNetworkParameterWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtNetworkParameterWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtNetworkParameterWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtNetworkParameterWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterWrite")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtNetworkParameterWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtNetworkParameterWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtNetworkParameterWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
