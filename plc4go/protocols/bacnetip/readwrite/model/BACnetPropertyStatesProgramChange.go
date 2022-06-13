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

// BACnetPropertyStatesProgramChange is the data-structure of this message
type BACnetPropertyStatesProgramChange struct {
	*BACnetPropertyStates
	ProgramState *BACnetProgramStateTagged
}

// IBACnetPropertyStatesProgramChange is the corresponding interface of BACnetPropertyStatesProgramChange
type IBACnetPropertyStatesProgramChange interface {
	IBACnetPropertyStates
	// GetProgramState returns ProgramState (property field)
	GetProgramState() *BACnetProgramStateTagged
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

func (m *BACnetPropertyStatesProgramChange) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesProgramChange) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesProgramChange) GetProgramState() *BACnetProgramStateTagged {
	return m.ProgramState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesProgramChange factory function for BACnetPropertyStatesProgramChange
func NewBACnetPropertyStatesProgramChange(programState *BACnetProgramStateTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesProgramChange {
	_result := &BACnetPropertyStatesProgramChange{
		ProgramState:         programState,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesProgramChange(structType interface{}) *BACnetPropertyStatesProgramChange {
	if casted, ok := structType.(BACnetPropertyStatesProgramChange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesProgramChange(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesProgramChange(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesProgramChange) GetTypeName() string {
	return "BACnetPropertyStatesProgramChange"
}

func (m *BACnetPropertyStatesProgramChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesProgramChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programState)
	lengthInBits += m.ProgramState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesProgramChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesProgramChangeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesProgramChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programState)
	if pullErr := readBuffer.PullContext("programState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programState")
	}
	_programState, _programStateErr := BACnetProgramStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _programStateErr != nil {
		return nil, errors.Wrap(_programStateErr, "Error parsing 'programState' field")
	}
	programState := CastBACnetProgramStateTagged(_programState)
	if closeErr := readBuffer.CloseContext("programState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProgramChange")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesProgramChange{
		ProgramState:         CastBACnetProgramStateTagged(programState),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesProgramChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProgramChange")
		}

		// Simple Field (programState)
		if pushErr := writeBuffer.PushContext("programState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programState")
		}
		_programStateErr := writeBuffer.WriteSerializable(m.ProgramState)
		if popErr := writeBuffer.PopContext("programState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programState")
		}
		if _programStateErr != nil {
			return errors.Wrap(_programStateErr, "Error serializing 'programState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProgramChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
