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

// BACnetConstructedDataScheduleDefault is the data-structure of this message
type BACnetConstructedDataScheduleDefault struct {
	*BACnetConstructedData
	ScheduleDefault *BACnetConstructedDataElement

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataScheduleDefault is the corresponding interface of BACnetConstructedDataScheduleDefault
type IBACnetConstructedDataScheduleDefault interface {
	IBACnetConstructedData
	// GetScheduleDefault returns ScheduleDefault (property field)
	GetScheduleDefault() *BACnetConstructedDataElement
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

func (m *BACnetConstructedDataScheduleDefault) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataScheduleDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SCHEDULE_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataScheduleDefault) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataScheduleDefault) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataScheduleDefault) GetScheduleDefault() *BACnetConstructedDataElement {
	return m.ScheduleDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataScheduleDefault factory function for BACnetConstructedDataScheduleDefault
func NewBACnetConstructedDataScheduleDefault(scheduleDefault *BACnetConstructedDataElement, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataScheduleDefault {
	_result := &BACnetConstructedDataScheduleDefault{
		ScheduleDefault:       scheduleDefault,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataScheduleDefault(structType interface{}) *BACnetConstructedDataScheduleDefault {
	if casted, ok := structType.(BACnetConstructedDataScheduleDefault); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScheduleDefault); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataScheduleDefault(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataScheduleDefault(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataScheduleDefault) GetTypeName() string {
	return "BACnetConstructedDataScheduleDefault"
}

func (m *BACnetConstructedDataScheduleDefault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataScheduleDefault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (scheduleDefault)
	lengthInBits += m.ScheduleDefault.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataScheduleDefault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataScheduleDefaultParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataScheduleDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScheduleDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScheduleDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (scheduleDefault)
	if pullErr := readBuffer.PullContext("scheduleDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for scheduleDefault")
	}
	_scheduleDefault, _scheduleDefaultErr := BACnetConstructedDataElementParse(readBuffer, BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(propertyIdentifierArgument), nil)
	if _scheduleDefaultErr != nil {
		return nil, errors.Wrap(_scheduleDefaultErr, "Error parsing 'scheduleDefault' field")
	}
	scheduleDefault := CastBACnetConstructedDataElement(_scheduleDefault)
	if closeErr := readBuffer.CloseContext("scheduleDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for scheduleDefault")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScheduleDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScheduleDefault")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataScheduleDefault{
		ScheduleDefault:       CastBACnetConstructedDataElement(scheduleDefault),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataScheduleDefault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScheduleDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScheduleDefault")
		}

		// Simple Field (scheduleDefault)
		if pushErr := writeBuffer.PushContext("scheduleDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for scheduleDefault")
		}
		_scheduleDefaultErr := writeBuffer.WriteSerializable(m.ScheduleDefault)
		if popErr := writeBuffer.PopContext("scheduleDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for scheduleDefault")
		}
		if _scheduleDefaultErr != nil {
			return errors.Wrap(_scheduleDefaultErr, "Error serializing 'scheduleDefault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScheduleDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScheduleDefault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataScheduleDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
