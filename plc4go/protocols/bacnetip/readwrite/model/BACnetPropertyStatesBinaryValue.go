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

// BACnetPropertyStatesBinaryValue is the data-structure of this message
type BACnetPropertyStatesBinaryValue struct {
	*BACnetPropertyStates
	BinaryValue *BACnetBinaryPVTagged
}

// IBACnetPropertyStatesBinaryValue is the corresponding interface of BACnetPropertyStatesBinaryValue
type IBACnetPropertyStatesBinaryValue interface {
	IBACnetPropertyStates
	// GetBinaryValue returns BinaryValue (property field)
	GetBinaryValue() *BACnetBinaryPVTagged
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesBinaryValue) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesBinaryValue) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesBinaryValue) GetBinaryValue() *BACnetBinaryPVTagged {
	return m.BinaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBinaryValue factory function for BACnetPropertyStatesBinaryValue
func NewBACnetPropertyStatesBinaryValue(binaryValue *BACnetBinaryPVTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesBinaryValue {
	_result := &BACnetPropertyStatesBinaryValue{
		BinaryValue:          binaryValue,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesBinaryValue(structType interface{}) *BACnetPropertyStatesBinaryValue {
	if casted, ok := structType.(BACnetPropertyStatesBinaryValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesBinaryValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryValue"
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (binaryValue)
	lengthInBits += m.BinaryValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBinaryValueParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesBinaryValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBinaryValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryValue)
	if pullErr := readBuffer.PullContext("binaryValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryValue")
	}
	_binaryValue, _binaryValueErr := BACnetBinaryPVTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _binaryValueErr != nil {
		return nil, errors.Wrap(_binaryValueErr, "Error parsing 'binaryValue' field")
	}
	binaryValue := CastBACnetBinaryPVTagged(_binaryValue)
	if closeErr := readBuffer.CloseContext("binaryValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBinaryValue")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesBinaryValue{
		BinaryValue:          CastBACnetBinaryPVTagged(binaryValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesBinaryValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBinaryValue")
		}

		// Simple Field (binaryValue)
		if pushErr := writeBuffer.PushContext("binaryValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for binaryValue")
		}
		_binaryValueErr := writeBuffer.WriteSerializable(m.BinaryValue)
		if popErr := writeBuffer.PopContext("binaryValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for binaryValue")
		}
		if _binaryValueErr != nil {
			return errors.Wrap(_binaryValueErr, "Error serializing 'binaryValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBinaryValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesBinaryValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
