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

// ApduDataExtMemoryBitWrite is the data-structure of this message
type ApduDataExtMemoryBitWrite struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtMemoryBitWrite is the corresponding interface of ApduDataExtMemoryBitWrite
type IApduDataExtMemoryBitWrite interface {
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

func (m *ApduDataExtMemoryBitWrite) GetExtApciType() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtMemoryBitWrite) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtMemoryBitWrite) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtMemoryBitWrite factory function for ApduDataExtMemoryBitWrite
func NewApduDataExtMemoryBitWrite(length uint8) *ApduDataExtMemoryBitWrite {
	_result := &ApduDataExtMemoryBitWrite{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtMemoryBitWrite(structType interface{}) *ApduDataExtMemoryBitWrite {
	if casted, ok := structType.(ApduDataExtMemoryBitWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtMemoryBitWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtMemoryBitWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtMemoryBitWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataExtMemoryBitWrite) GetTypeName() string {
	return "ApduDataExtMemoryBitWrite"
}

func (m *ApduDataExtMemoryBitWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtMemoryBitWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtMemoryBitWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtMemoryBitWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtMemoryBitWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtMemoryBitWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtMemoryBitWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtMemoryBitWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtMemoryBitWrite")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtMemoryBitWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtMemoryBitWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtMemoryBitWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtMemoryBitWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtMemoryBitWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtMemoryBitWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtMemoryBitWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
