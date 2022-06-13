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

// BACnetPropertyStatesDoorSecuredStatus is the data-structure of this message
type BACnetPropertyStatesDoorSecuredStatus struct {
	*BACnetPropertyStates
	DoorSecuredStatus *BACnetDoorSecuredStatusTagged
}

// IBACnetPropertyStatesDoorSecuredStatus is the corresponding interface of BACnetPropertyStatesDoorSecuredStatus
type IBACnetPropertyStatesDoorSecuredStatus interface {
	IBACnetPropertyStates
	// GetDoorSecuredStatus returns DoorSecuredStatus (property field)
	GetDoorSecuredStatus() *BACnetDoorSecuredStatusTagged
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

func (m *BACnetPropertyStatesDoorSecuredStatus) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesDoorSecuredStatus) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesDoorSecuredStatus) GetDoorSecuredStatus() *BACnetDoorSecuredStatusTagged {
	return m.DoorSecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorSecuredStatus factory function for BACnetPropertyStatesDoorSecuredStatus
func NewBACnetPropertyStatesDoorSecuredStatus(doorSecuredStatus *BACnetDoorSecuredStatusTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesDoorSecuredStatus {
	_result := &BACnetPropertyStatesDoorSecuredStatus{
		DoorSecuredStatus:    doorSecuredStatus,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesDoorSecuredStatus(structType interface{}) *BACnetPropertyStatesDoorSecuredStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorSecuredStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorSecuredStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorSecuredStatus(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesDoorSecuredStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorSecuredStatus"
}

func (m *BACnetPropertyStatesDoorSecuredStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesDoorSecuredStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorSecuredStatus)
	lengthInBits += m.DoorSecuredStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesDoorSecuredStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesDoorSecuredStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesDoorSecuredStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorSecuredStatus)
	if pullErr := readBuffer.PullContext("doorSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorSecuredStatus")
	}
	_doorSecuredStatus, _doorSecuredStatusErr := BACnetDoorSecuredStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorSecuredStatusErr != nil {
		return nil, errors.Wrap(_doorSecuredStatusErr, "Error parsing 'doorSecuredStatus' field")
	}
	doorSecuredStatus := CastBACnetDoorSecuredStatusTagged(_doorSecuredStatus)
	if closeErr := readBuffer.CloseContext("doorSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorSecuredStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorSecuredStatus")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesDoorSecuredStatus{
		DoorSecuredStatus:    CastBACnetDoorSecuredStatusTagged(doorSecuredStatus),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesDoorSecuredStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorSecuredStatus")
		}

		// Simple Field (doorSecuredStatus)
		if pushErr := writeBuffer.PushContext("doorSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorSecuredStatus")
		}
		_doorSecuredStatusErr := writeBuffer.WriteSerializable(m.DoorSecuredStatus)
		if popErr := writeBuffer.PopContext("doorSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorSecuredStatus")
		}
		if _doorSecuredStatusErr != nil {
			return errors.Wrap(_doorSecuredStatusErr, "Error serializing 'doorSecuredStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorSecuredStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesDoorSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
