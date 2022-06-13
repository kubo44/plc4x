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

// BACnetChannelValueDouble is the data-structure of this message
type BACnetChannelValueDouble struct {
	*BACnetChannelValue
	DoubleValue *BACnetApplicationTagDouble
}

// IBACnetChannelValueDouble is the corresponding interface of BACnetChannelValueDouble
type IBACnetChannelValueDouble interface {
	IBACnetChannelValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() *BACnetApplicationTagDouble
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

func (m *BACnetChannelValueDouble) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueDouble) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueDouble) GetDoubleValue() *BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueDouble factory function for BACnetChannelValueDouble
func NewBACnetChannelValueDouble(doubleValue *BACnetApplicationTagDouble, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueDouble {
	_result := &BACnetChannelValueDouble{
		DoubleValue:        doubleValue,
		BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueDouble(structType interface{}) *BACnetChannelValueDouble {
	if casted, ok := structType.(BACnetChannelValueDouble); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueDouble(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueDouble(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueDouble) GetTypeName() string {
	return "BACnetChannelValueDouble"
}

func (m *BACnetChannelValueDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueDoubleParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doubleValue")
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParse(readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field")
	}
	doubleValue := CastBACnetApplicationTagDouble(_doubleValue)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doubleValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueDouble")
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueDouble{
		DoubleValue:        CastBACnetApplicationTagDouble(doubleValue),
		BACnetChannelValue: &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueDouble")
		}

		// Simple Field (doubleValue)
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doubleValue")
		}
		_doubleValueErr := writeBuffer.WriteSerializable(m.DoubleValue)
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doubleValue")
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueDouble")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
