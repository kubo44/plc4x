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

// BACnetFaultParameterFaultExtendedParametersEntryNull is the data-structure of this message
type BACnetFaultParameterFaultExtendedParametersEntryNull struct {
	*BACnetFaultParameterFaultExtendedParametersEntry
	NullValue *BACnetApplicationTagNull
}

// IBACnetFaultParameterFaultExtendedParametersEntryNull is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryNull
type IBACnetFaultParameterFaultExtendedParametersEntryNull interface {
	IBACnetFaultParameterFaultExtendedParametersEntry
	// GetNullValue returns NullValue (property field)
	GetNullValue() *BACnetApplicationTagNull
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

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) InitializeParent(parent *BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameterFaultExtendedParametersEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetParent() *BACnetFaultParameterFaultExtendedParametersEntry {
	return m.BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetNullValue() *BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryNull factory function for BACnetFaultParameterFaultExtendedParametersEntryNull
func NewBACnetFaultParameterFaultExtendedParametersEntryNull(nullValue *BACnetApplicationTagNull, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultExtendedParametersEntryNull {
	_result := &BACnetFaultParameterFaultExtendedParametersEntryNull{
		NullValue: nullValue,
		BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultExtendedParametersEntryNull(structType interface{}) *BACnetFaultParameterFaultExtendedParametersEntryNull {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryNull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryNull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryNull(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryNull(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryNull"
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryNullParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultExtendedParametersEntryNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParse(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field")
	}
	nullValue := CastBACnetApplicationTagNull(_nullValue)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryNull")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultExtendedParametersEntryNull{
		NullValue: CastBACnetApplicationTagNull(nullValue),
		BACnetFaultParameterFaultExtendedParametersEntry: &BACnetFaultParameterFaultExtendedParametersEntry{},
	}
	_child.BACnetFaultParameterFaultExtendedParametersEntry.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := writeBuffer.WriteSerializable(m.NullValue)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
