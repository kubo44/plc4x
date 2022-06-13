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

// BACnetLandingCallStatusCommandDestination is the data-structure of this message
type BACnetLandingCallStatusCommandDestination struct {
	*BACnetLandingCallStatusCommand
	Destination *BACnetContextTagUnsignedInteger
}

// IBACnetLandingCallStatusCommandDestination is the corresponding interface of BACnetLandingCallStatusCommandDestination
type IBACnetLandingCallStatusCommandDestination interface {
	IBACnetLandingCallStatusCommand
	// GetDestination returns Destination (property field)
	GetDestination() *BACnetContextTagUnsignedInteger
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

func (m *BACnetLandingCallStatusCommandDestination) InitializeParent(parent *BACnetLandingCallStatusCommand, peekedTagHeader *BACnetTagHeader) {
	m.BACnetLandingCallStatusCommand.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetLandingCallStatusCommandDestination) GetParent() *BACnetLandingCallStatusCommand {
	return m.BACnetLandingCallStatusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLandingCallStatusCommandDestination) GetDestination() *BACnetContextTagUnsignedInteger {
	return m.Destination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingCallStatusCommandDestination factory function for BACnetLandingCallStatusCommandDestination
func NewBACnetLandingCallStatusCommandDestination(destination *BACnetContextTagUnsignedInteger, peekedTagHeader *BACnetTagHeader) *BACnetLandingCallStatusCommandDestination {
	_result := &BACnetLandingCallStatusCommandDestination{
		Destination:                    destination,
		BACnetLandingCallStatusCommand: NewBACnetLandingCallStatusCommand(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLandingCallStatusCommandDestination(structType interface{}) *BACnetLandingCallStatusCommandDestination {
	if casted, ok := structType.(BACnetLandingCallStatusCommandDestination); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLandingCallStatusCommandDestination); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLandingCallStatusCommand); ok {
		return CastBACnetLandingCallStatusCommandDestination(casted.Child)
	}
	if casted, ok := structType.(*BACnetLandingCallStatusCommand); ok {
		return CastBACnetLandingCallStatusCommandDestination(casted.Child)
	}
	return nil
}

func (m *BACnetLandingCallStatusCommandDestination) GetTypeName() string {
	return "BACnetLandingCallStatusCommandDestination"
}

func (m *BACnetLandingCallStatusCommandDestination) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLandingCallStatusCommandDestination) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destination)
	lengthInBits += m.Destination.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLandingCallStatusCommandDestination) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLandingCallStatusCommandDestinationParse(readBuffer utils.ReadBuffer) (*BACnetLandingCallStatusCommandDestination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingCallStatusCommandDestination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingCallStatusCommandDestination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destination)
	if pullErr := readBuffer.PullContext("destination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destination")
	}
	_destination, _destinationErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _destinationErr != nil {
		return nil, errors.Wrap(_destinationErr, "Error parsing 'destination' field")
	}
	destination := CastBACnetContextTagUnsignedInteger(_destination)
	if closeErr := readBuffer.CloseContext("destination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destination")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingCallStatusCommandDestination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingCallStatusCommandDestination")
	}

	// Create a partially initialized instance
	_child := &BACnetLandingCallStatusCommandDestination{
		Destination:                    CastBACnetContextTagUnsignedInteger(destination),
		BACnetLandingCallStatusCommand: &BACnetLandingCallStatusCommand{},
	}
	_child.BACnetLandingCallStatusCommand.Child = _child
	return _child, nil
}

func (m *BACnetLandingCallStatusCommandDestination) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLandingCallStatusCommandDestination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLandingCallStatusCommandDestination")
		}

		// Simple Field (destination)
		if pushErr := writeBuffer.PushContext("destination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for destination")
		}
		_destinationErr := writeBuffer.WriteSerializable(m.Destination)
		if popErr := writeBuffer.PopContext("destination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for destination")
		}
		if _destinationErr != nil {
			return errors.Wrap(_destinationErr, "Error serializing 'destination' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLandingCallStatusCommandDestination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLandingCallStatusCommandDestination")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLandingCallStatusCommandDestination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
