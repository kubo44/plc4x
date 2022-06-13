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

// BACnetPropertyStatesAccessEvent is the data-structure of this message
type BACnetPropertyStatesAccessEvent struct {
	*BACnetPropertyStates
	AccessEvent *BACnetAccessEventTagged
}

// IBACnetPropertyStatesAccessEvent is the corresponding interface of BACnetPropertyStatesAccessEvent
type IBACnetPropertyStatesAccessEvent interface {
	IBACnetPropertyStates
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() *BACnetAccessEventTagged
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

func (m *BACnetPropertyStatesAccessEvent) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesAccessEvent) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesAccessEvent) GetAccessEvent() *BACnetAccessEventTagged {
	return m.AccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAccessEvent factory function for BACnetPropertyStatesAccessEvent
func NewBACnetPropertyStatesAccessEvent(accessEvent *BACnetAccessEventTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesAccessEvent {
	_result := &BACnetPropertyStatesAccessEvent{
		AccessEvent:          accessEvent,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesAccessEvent(structType interface{}) *BACnetPropertyStatesAccessEvent {
	if casted, ok := structType.(BACnetPropertyStatesAccessEvent); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessEvent(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesAccessEvent(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesAccessEvent) GetTypeName() string {
	return "BACnetPropertyStatesAccessEvent"
}

func (m *BACnetPropertyStatesAccessEvent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesAccessEvent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesAccessEvent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesAccessEventParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEvent)
	if pullErr := readBuffer.PullContext("accessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEvent")
	}
	_accessEvent, _accessEventErr := BACnetAccessEventTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accessEventErr != nil {
		return nil, errors.Wrap(_accessEventErr, "Error parsing 'accessEvent' field")
	}
	accessEvent := CastBACnetAccessEventTagged(_accessEvent)
	if closeErr := readBuffer.CloseContext("accessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEvent")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAccessEvent")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesAccessEvent{
		AccessEvent:          CastBACnetAccessEventTagged(accessEvent),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesAccessEvent) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAccessEvent")
		}

		// Simple Field (accessEvent)
		if pushErr := writeBuffer.PushContext("accessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEvent")
		}
		_accessEventErr := writeBuffer.WriteSerializable(m.AccessEvent)
		if popErr := writeBuffer.PopContext("accessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEvent")
		}
		if _accessEventErr != nil {
			return errors.Wrap(_accessEventErr, "Error serializing 'accessEvent' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
