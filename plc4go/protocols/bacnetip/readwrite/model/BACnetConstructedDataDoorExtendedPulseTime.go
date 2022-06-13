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

// BACnetConstructedDataDoorExtendedPulseTime is the data-structure of this message
type BACnetConstructedDataDoorExtendedPulseTime struct {
	*BACnetConstructedData
	DoorExtendedPulseTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDoorExtendedPulseTime is the corresponding interface of BACnetConstructedDataDoorExtendedPulseTime
type IBACnetConstructedDataDoorExtendedPulseTime interface {
	IBACnetConstructedData
	// GetDoorExtendedPulseTime returns DoorExtendedPulseTime (property field)
	GetDoorExtendedPulseTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_EXTENDED_PULSE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDoorExtendedPulseTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetDoorExtendedPulseTime() *BACnetApplicationTagUnsignedInteger {
	return m.DoorExtendedPulseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorExtendedPulseTime factory function for BACnetConstructedDataDoorExtendedPulseTime
func NewBACnetConstructedDataDoorExtendedPulseTime(doorExtendedPulseTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDoorExtendedPulseTime {
	_result := &BACnetConstructedDataDoorExtendedPulseTime{
		DoorExtendedPulseTime: doorExtendedPulseTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDoorExtendedPulseTime(structType interface{}) *BACnetConstructedDataDoorExtendedPulseTime {
	if casted, ok := structType.(BACnetConstructedDataDoorExtendedPulseTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorExtendedPulseTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorExtendedPulseTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorExtendedPulseTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetTypeName() string {
	return "BACnetConstructedDataDoorExtendedPulseTime"
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorExtendedPulseTime)
	lengthInBits += m.DoorExtendedPulseTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorExtendedPulseTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDoorExtendedPulseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorExtendedPulseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorExtendedPulseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorExtendedPulseTime)
	if pullErr := readBuffer.PullContext("doorExtendedPulseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorExtendedPulseTime")
	}
	_doorExtendedPulseTime, _doorExtendedPulseTimeErr := BACnetApplicationTagParse(readBuffer)
	if _doorExtendedPulseTimeErr != nil {
		return nil, errors.Wrap(_doorExtendedPulseTimeErr, "Error parsing 'doorExtendedPulseTime' field")
	}
	doorExtendedPulseTime := CastBACnetApplicationTagUnsignedInteger(_doorExtendedPulseTime)
	if closeErr := readBuffer.CloseContext("doorExtendedPulseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorExtendedPulseTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorExtendedPulseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorExtendedPulseTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDoorExtendedPulseTime{
		DoorExtendedPulseTime: CastBACnetApplicationTagUnsignedInteger(doorExtendedPulseTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorExtendedPulseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorExtendedPulseTime")
		}

		// Simple Field (doorExtendedPulseTime)
		if pushErr := writeBuffer.PushContext("doorExtendedPulseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorExtendedPulseTime")
		}
		_doorExtendedPulseTimeErr := writeBuffer.WriteSerializable(m.DoorExtendedPulseTime)
		if popErr := writeBuffer.PopContext("doorExtendedPulseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorExtendedPulseTime")
		}
		if _doorExtendedPulseTimeErr != nil {
			return errors.Wrap(_doorExtendedPulseTimeErr, "Error serializing 'doorExtendedPulseTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorExtendedPulseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorExtendedPulseTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDoorExtendedPulseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
