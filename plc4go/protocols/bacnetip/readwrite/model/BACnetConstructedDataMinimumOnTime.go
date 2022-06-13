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

// BACnetConstructedDataMinimumOnTime is the data-structure of this message
type BACnetConstructedDataMinimumOnTime struct {
	*BACnetConstructedData
	MinimumOnTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinimumOnTime is the corresponding interface of BACnetConstructedDataMinimumOnTime
type IBACnetConstructedDataMinimumOnTime interface {
	IBACnetConstructedData
	// GetMinimumOnTime returns MinimumOnTime (property field)
	GetMinimumOnTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMinimumOnTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinimumOnTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_ON_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinimumOnTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinimumOnTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinimumOnTime) GetMinimumOnTime() *BACnetApplicationTagUnsignedInteger {
	return m.MinimumOnTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumOnTime factory function for BACnetConstructedDataMinimumOnTime
func NewBACnetConstructedDataMinimumOnTime(minimumOnTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinimumOnTime {
	_result := &BACnetConstructedDataMinimumOnTime{
		MinimumOnTime:         minimumOnTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinimumOnTime(structType interface{}) *BACnetConstructedDataMinimumOnTime {
	if casted, ok := structType.(BACnetConstructedDataMinimumOnTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOnTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOnTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOnTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinimumOnTime) GetTypeName() string {
	return "BACnetConstructedDataMinimumOnTime"
}

func (m *BACnetConstructedDataMinimumOnTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinimumOnTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumOnTime)
	lengthInBits += m.MinimumOnTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinimumOnTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumOnTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinimumOnTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOnTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOnTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumOnTime)
	if pullErr := readBuffer.PullContext("minimumOnTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumOnTime")
	}
	_minimumOnTime, _minimumOnTimeErr := BACnetApplicationTagParse(readBuffer)
	if _minimumOnTimeErr != nil {
		return nil, errors.Wrap(_minimumOnTimeErr, "Error parsing 'minimumOnTime' field")
	}
	minimumOnTime := CastBACnetApplicationTagUnsignedInteger(_minimumOnTime)
	if closeErr := readBuffer.CloseContext("minimumOnTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumOnTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOnTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOnTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinimumOnTime{
		MinimumOnTime:         CastBACnetApplicationTagUnsignedInteger(minimumOnTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinimumOnTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOnTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOnTime")
		}

		// Simple Field (minimumOnTime)
		if pushErr := writeBuffer.PushContext("minimumOnTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumOnTime")
		}
		_minimumOnTimeErr := writeBuffer.WriteSerializable(m.MinimumOnTime)
		if popErr := writeBuffer.PopContext("minimumOnTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumOnTime")
		}
		if _minimumOnTimeErr != nil {
			return errors.Wrap(_minimumOnTimeErr, "Error serializing 'minimumOnTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOnTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOnTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinimumOnTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
