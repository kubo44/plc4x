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

// ModbusPDUGetComEventCounterResponse is the data-structure of this message
type ModbusPDUGetComEventCounterResponse struct {
	*ModbusPDU
	Status     uint16
	EventCount uint16
}

// IModbusPDUGetComEventCounterResponse is the corresponding interface of ModbusPDUGetComEventCounterResponse
type IModbusPDUGetComEventCounterResponse interface {
	IModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
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

func (m *ModbusPDUGetComEventCounterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUGetComEventCounterResponse) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *ModbusPDUGetComEventCounterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUGetComEventCounterResponse) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUGetComEventCounterResponse) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUGetComEventCounterResponse) GetStatus() uint16 {
	return m.Status
}

func (m *ModbusPDUGetComEventCounterResponse) GetEventCount() uint16 {
	return m.EventCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUGetComEventCounterResponse factory function for ModbusPDUGetComEventCounterResponse
func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *ModbusPDUGetComEventCounterResponse {
	_result := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		ModbusPDU:  NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUGetComEventCounterResponse(structType interface{}) *ModbusPDUGetComEventCounterResponse {
	if casted, ok := structType.(ModbusPDUGetComEventCounterResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUGetComEventCounterResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUGetComEventCounterResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *ModbusPDUGetComEventCounterResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUGetComEventCounterResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUGetComEventCounterResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUGetComEventCounterResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUGetComEventCounterResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventCounterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (eventCount)
	_eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}
	eventCount := _eventCount

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventCounterResponse")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		ModbusPDU:  &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUGetComEventCounterResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventCounterResponse")
		}

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := writeBuffer.WriteUint16("status", 16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventCounterResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUGetComEventCounterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
