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

// BACnetConstructedDataDoorAlarmState is the data-structure of this message
type BACnetConstructedDataDoorAlarmState struct {
	*BACnetConstructedData
	DoorAlarmState *BACnetDoorAlarmStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDoorAlarmState is the corresponding interface of BACnetConstructedDataDoorAlarmState
type IBACnetConstructedDataDoorAlarmState interface {
	IBACnetConstructedData
	// GetDoorAlarmState returns DoorAlarmState (property field)
	GetDoorAlarmState() *BACnetDoorAlarmStateTagged
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

func (m *BACnetConstructedDataDoorAlarmState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDoorAlarmState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_ALARM_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDoorAlarmState) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDoorAlarmState) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDoorAlarmState) GetDoorAlarmState() *BACnetDoorAlarmStateTagged {
	return m.DoorAlarmState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorAlarmState factory function for BACnetConstructedDataDoorAlarmState
func NewBACnetConstructedDataDoorAlarmState(doorAlarmState *BACnetDoorAlarmStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDoorAlarmState {
	_result := &BACnetConstructedDataDoorAlarmState{
		DoorAlarmState:        doorAlarmState,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDoorAlarmState(structType interface{}) *BACnetConstructedDataDoorAlarmState {
	if casted, ok := structType.(BACnetConstructedDataDoorAlarmState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorAlarmState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorAlarmState(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDoorAlarmState(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDoorAlarmState) GetTypeName() string {
	return "BACnetConstructedDataDoorAlarmState"
}

func (m *BACnetConstructedDataDoorAlarmState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDoorAlarmState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorAlarmState)
	lengthInBits += m.DoorAlarmState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDoorAlarmState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorAlarmStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDoorAlarmState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorAlarmState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorAlarmState)
	if pullErr := readBuffer.PullContext("doorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorAlarmState")
	}
	_doorAlarmState, _doorAlarmStateErr := BACnetDoorAlarmStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _doorAlarmStateErr != nil {
		return nil, errors.Wrap(_doorAlarmStateErr, "Error parsing 'doorAlarmState' field")
	}
	doorAlarmState := CastBACnetDoorAlarmStateTagged(_doorAlarmState)
	if closeErr := readBuffer.CloseContext("doorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorAlarmState")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorAlarmState")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDoorAlarmState{
		DoorAlarmState:        CastBACnetDoorAlarmStateTagged(doorAlarmState),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDoorAlarmState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorAlarmState")
		}

		// Simple Field (doorAlarmState)
		if pushErr := writeBuffer.PushContext("doorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorAlarmState")
		}
		_doorAlarmStateErr := writeBuffer.WriteSerializable(m.DoorAlarmState)
		if popErr := writeBuffer.PopContext("doorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorAlarmState")
		}
		if _doorAlarmStateErr != nil {
			return errors.Wrap(_doorAlarmStateErr, "Error serializing 'doorAlarmState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorAlarmState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDoorAlarmState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
