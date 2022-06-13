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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AlarmMessageAckObjectPushType_VARIABLESPEC uint8 = 0x12

// AlarmMessageAckObjectPushType is the data-structure of this message
type AlarmMessageAckObjectPushType struct {
	LengthSpec     uint8
	SyntaxId       SyntaxIdType
	NumberOfValues uint8
	EventId        uint32
	AckStateGoing  *State
	AckStateComing *State
}

// IAlarmMessageAckObjectPushType is the corresponding interface of AlarmMessageAckObjectPushType
type IAlarmMessageAckObjectPushType interface {
	// GetLengthSpec returns LengthSpec (property field)
	GetLengthSpec() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetNumberOfValues returns NumberOfValues (property field)
	GetNumberOfValues() uint8
	// GetEventId returns EventId (property field)
	GetEventId() uint32
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() *State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() *State
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AlarmMessageAckObjectPushType) GetLengthSpec() uint8 {
	return m.LengthSpec
}

func (m *AlarmMessageAckObjectPushType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *AlarmMessageAckObjectPushType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *AlarmMessageAckObjectPushType) GetEventId() uint32 {
	return m.EventId
}

func (m *AlarmMessageAckObjectPushType) GetAckStateGoing() *State {
	return m.AckStateGoing
}

func (m *AlarmMessageAckObjectPushType) GetAckStateComing() *State {
	return m.AckStateComing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *AlarmMessageAckObjectPushType) GetVariableSpec() uint8 {
	return AlarmMessageAckObjectPushType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageAckObjectPushType factory function for AlarmMessageAckObjectPushType
func NewAlarmMessageAckObjectPushType(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, ackStateGoing *State, ackStateComing *State) *AlarmMessageAckObjectPushType {
	return &AlarmMessageAckObjectPushType{LengthSpec: lengthSpec, SyntaxId: syntaxId, NumberOfValues: numberOfValues, EventId: eventId, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing}
}

func CastAlarmMessageAckObjectPushType(structType interface{}) *AlarmMessageAckObjectPushType {
	if casted, ok := structType.(AlarmMessageAckObjectPushType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageAckObjectPushType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageAckObjectPushType) GetTypeName() string {
	return "AlarmMessageAckObjectPushType"
}

func (m *AlarmMessageAckObjectPushType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageAckObjectPushType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (lengthSpec)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8

	// Simple field (eventId)
	lengthInBits += 32

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits()

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits()

	return lengthInBits
}

func (m *AlarmMessageAckObjectPushType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageAckObjectPushTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageAckObjectPushType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckObjectPushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckObjectPushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field")
	}
	if variableSpec != AlarmMessageAckObjectPushType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageAckObjectPushType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Simple Field (lengthSpec)
	_lengthSpec, _lengthSpecErr := readBuffer.ReadUint8("lengthSpec", 8)
	if _lengthSpecErr != nil {
		return nil, errors.Wrap(_lengthSpecErr, "Error parsing 'lengthSpec' field")
	}
	lengthSpec := _lengthSpec

	// Simple Field (syntaxId)
	if pullErr := readBuffer.PullContext("syntaxId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for syntaxId")
	}
	_syntaxId, _syntaxIdErr := SyntaxIdTypeParse(readBuffer)
	if _syntaxIdErr != nil {
		return nil, errors.Wrap(_syntaxIdErr, "Error parsing 'syntaxId' field")
	}
	syntaxId := _syntaxId
	if closeErr := readBuffer.CloseContext("syntaxId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for syntaxId")
	}

	// Simple Field (numberOfValues)
	_numberOfValues, _numberOfValuesErr := readBuffer.ReadUint8("numberOfValues", 8)
	if _numberOfValuesErr != nil {
		return nil, errors.Wrap(_numberOfValuesErr, "Error parsing 'numberOfValues' field")
	}
	numberOfValues := _numberOfValues

	// Simple Field (eventId)
	_eventId, _eventIdErr := readBuffer.ReadUint32("eventId", 32)
	if _eventIdErr != nil {
		return nil, errors.Wrap(_eventIdErr, "Error parsing 'eventId' field")
	}
	eventId := _eventId

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateGoing")
	}
	_ackStateGoing, _ackStateGoingErr := StateParse(readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field")
	}
	ackStateGoing := CastState(_ackStateGoing)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateGoing")
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateComing")
	}
	_ackStateComing, _ackStateComingErr := StateParse(readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field")
	}
	ackStateComing := CastState(_ackStateComing)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateComing")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckObjectPushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckObjectPushType")
	}

	// Create the instance
	return NewAlarmMessageAckObjectPushType(lengthSpec, syntaxId, numberOfValues, eventId, ackStateGoing, ackStateComing), nil
}

func (m *AlarmMessageAckObjectPushType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AlarmMessageAckObjectPushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckObjectPushType")
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Simple Field (lengthSpec)
	lengthSpec := uint8(m.LengthSpec)
	_lengthSpecErr := writeBuffer.WriteUint8("lengthSpec", 8, (lengthSpec))
	if _lengthSpecErr != nil {
		return errors.Wrap(_lengthSpecErr, "Error serializing 'lengthSpec' field")
	}

	// Simple Field (syntaxId)
	if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for syntaxId")
	}
	_syntaxIdErr := writeBuffer.WriteSerializable(m.SyntaxId)
	if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for syntaxId")
	}
	if _syntaxIdErr != nil {
		return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
	}

	// Simple Field (numberOfValues)
	numberOfValues := uint8(m.NumberOfValues)
	_numberOfValuesErr := writeBuffer.WriteUint8("numberOfValues", 8, (numberOfValues))
	if _numberOfValuesErr != nil {
		return errors.Wrap(_numberOfValuesErr, "Error serializing 'numberOfValues' field")
	}

	// Simple Field (eventId)
	eventId := uint32(m.EventId)
	_eventIdErr := writeBuffer.WriteUint32("eventId", 32, (eventId))
	if _eventIdErr != nil {
		return errors.Wrap(_eventIdErr, "Error serializing 'eventId' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateGoing")
	}
	_ackStateGoingErr := writeBuffer.WriteSerializable(m.AckStateGoing)
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateGoing")
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateComing")
	}
	_ackStateComingErr := writeBuffer.WriteSerializable(m.AckStateComing)
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateComing")
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckObjectPushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckObjectPushType")
	}
	return nil
}

func (m *AlarmMessageAckObjectPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
