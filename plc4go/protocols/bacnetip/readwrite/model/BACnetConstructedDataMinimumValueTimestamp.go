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

// BACnetConstructedDataMinimumValueTimestamp is the data-structure of this message
type BACnetConstructedDataMinimumValueTimestamp struct {
	*BACnetConstructedData
	MinimumValueTimestamp *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinimumValueTimestamp is the corresponding interface of BACnetConstructedDataMinimumValueTimestamp
type IBACnetConstructedDataMinimumValueTimestamp interface {
	IBACnetConstructedData
	// GetMinimumValueTimestamp returns MinimumValueTimestamp (property field)
	GetMinimumValueTimestamp() *BACnetDateTime
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

func (m *BACnetConstructedDataMinimumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinimumValueTimestamp) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinimumValueTimestamp) GetMinimumValueTimestamp() *BACnetDateTime {
	return m.MinimumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumValueTimestamp factory function for BACnetConstructedDataMinimumValueTimestamp
func NewBACnetConstructedDataMinimumValueTimestamp(minimumValueTimestamp *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinimumValueTimestamp {
	_result := &BACnetConstructedDataMinimumValueTimestamp{
		MinimumValueTimestamp: minimumValueTimestamp,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinimumValueTimestamp(structType interface{}) *BACnetConstructedDataMinimumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMinimumValueTimestamp); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumValueTimestamp(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumValueTimestamp(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMinimumValueTimestamp"
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumValueTimestamp)
	lengthInBits += m.MinimumValueTimestamp.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinimumValueTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumValueTimestampParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinimumValueTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumValueTimestamp)
	if pullErr := readBuffer.PullContext("minimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumValueTimestamp")
	}
	_minimumValueTimestamp, _minimumValueTimestampErr := BACnetDateTimeParse(readBuffer)
	if _minimumValueTimestampErr != nil {
		return nil, errors.Wrap(_minimumValueTimestampErr, "Error parsing 'minimumValueTimestamp' field")
	}
	minimumValueTimestamp := CastBACnetDateTime(_minimumValueTimestamp)
	if closeErr := readBuffer.CloseContext("minimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumValueTimestamp")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumValueTimestamp")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinimumValueTimestamp{
		MinimumValueTimestamp: CastBACnetDateTime(minimumValueTimestamp),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinimumValueTimestamp) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumValueTimestamp")
		}

		// Simple Field (minimumValueTimestamp)
		if pushErr := writeBuffer.PushContext("minimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumValueTimestamp")
		}
		_minimumValueTimestampErr := writeBuffer.WriteSerializable(m.MinimumValueTimestamp)
		if popErr := writeBuffer.PopContext("minimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumValueTimestamp")
		}
		if _minimumValueTimestampErr != nil {
			return errors.Wrap(_minimumValueTimestampErr, "Error serializing 'minimumValueTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumValueTimestamp")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinimumValueTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
