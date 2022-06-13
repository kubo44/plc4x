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

// BACnetConstructedDataPassengerAlarm is the data-structure of this message
type BACnetConstructedDataPassengerAlarm struct {
	*BACnetConstructedData
	PassengerAlarm *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPassengerAlarm is the corresponding interface of BACnetConstructedDataPassengerAlarm
type IBACnetConstructedDataPassengerAlarm interface {
	IBACnetConstructedData
	// GetPassengerAlarm returns PassengerAlarm (property field)
	GetPassengerAlarm() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataPassengerAlarm) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPassengerAlarm) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSENGER_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPassengerAlarm) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPassengerAlarm) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPassengerAlarm) GetPassengerAlarm() *BACnetApplicationTagBoolean {
	return m.PassengerAlarm
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassengerAlarm factory function for BACnetConstructedDataPassengerAlarm
func NewBACnetConstructedDataPassengerAlarm(passengerAlarm *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPassengerAlarm {
	_result := &BACnetConstructedDataPassengerAlarm{
		PassengerAlarm:        passengerAlarm,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPassengerAlarm(structType interface{}) *BACnetConstructedDataPassengerAlarm {
	if casted, ok := structType.(BACnetConstructedDataPassengerAlarm); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassengerAlarm); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassengerAlarm(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassengerAlarm(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPassengerAlarm) GetTypeName() string {
	return "BACnetConstructedDataPassengerAlarm"
}

func (m *BACnetConstructedDataPassengerAlarm) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPassengerAlarm) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passengerAlarm)
	lengthInBits += m.PassengerAlarm.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPassengerAlarm) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassengerAlarmParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPassengerAlarm, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassengerAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassengerAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passengerAlarm)
	if pullErr := readBuffer.PullContext("passengerAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passengerAlarm")
	}
	_passengerAlarm, _passengerAlarmErr := BACnetApplicationTagParse(readBuffer)
	if _passengerAlarmErr != nil {
		return nil, errors.Wrap(_passengerAlarmErr, "Error parsing 'passengerAlarm' field")
	}
	passengerAlarm := CastBACnetApplicationTagBoolean(_passengerAlarm)
	if closeErr := readBuffer.CloseContext("passengerAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passengerAlarm")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassengerAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassengerAlarm")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPassengerAlarm{
		PassengerAlarm:        CastBACnetApplicationTagBoolean(passengerAlarm),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPassengerAlarm) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassengerAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassengerAlarm")
		}

		// Simple Field (passengerAlarm)
		if pushErr := writeBuffer.PushContext("passengerAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passengerAlarm")
		}
		_passengerAlarmErr := writeBuffer.WriteSerializable(m.PassengerAlarm)
		if popErr := writeBuffer.PopContext("passengerAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passengerAlarm")
		}
		if _passengerAlarmErr != nil {
			return errors.Wrap(_passengerAlarmErr, "Error serializing 'passengerAlarm' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassengerAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassengerAlarm")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPassengerAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
