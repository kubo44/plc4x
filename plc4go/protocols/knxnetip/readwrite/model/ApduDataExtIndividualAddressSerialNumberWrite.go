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

// ApduDataExtIndividualAddressSerialNumberWrite is the data-structure of this message
type ApduDataExtIndividualAddressSerialNumberWrite struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtIndividualAddressSerialNumberWrite is the corresponding interface of ApduDataExtIndividualAddressSerialNumberWrite
type IApduDataExtIndividualAddressSerialNumberWrite interface {
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

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetExtApciType() uint8 {
	return 0x1E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtIndividualAddressSerialNumberWrite) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtIndividualAddressSerialNumberWrite factory function for ApduDataExtIndividualAddressSerialNumberWrite
func NewApduDataExtIndividualAddressSerialNumberWrite(length uint8) *ApduDataExtIndividualAddressSerialNumberWrite {
	_result := &ApduDataExtIndividualAddressSerialNumberWrite{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtIndividualAddressSerialNumberWrite(structType interface{}) *ApduDataExtIndividualAddressSerialNumberWrite {
	if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtIndividualAddressSerialNumberWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtIndividualAddressSerialNumberWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberWrite"
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtIndividualAddressSerialNumberWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberWrite")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtIndividualAddressSerialNumberWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtIndividualAddressSerialNumberWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
