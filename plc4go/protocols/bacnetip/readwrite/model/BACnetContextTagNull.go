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

// BACnetContextTagNull is the data-structure of this message
type BACnetContextTagNull struct {
	*BACnetContextTag

	// Arguments.
	TagNumberArgument uint8
}

// IBACnetContextTagNull is the corresponding interface of BACnetContextTagNull
type IBACnetContextTagNull interface {
	IBACnetContextTag
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

func (m *BACnetContextTagNull) GetDataType() BACnetDataType {
	return BACnetDataType_NULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagNull) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagNull) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

// NewBACnetContextTagNull factory function for BACnetContextTagNull
func NewBACnetContextTagNull(header *BACnetTagHeader, tagNumberArgument uint8) *BACnetContextTagNull {
	_result := &BACnetContextTagNull{
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetContextTagNull(structType interface{}) *BACnetContextTagNull {
	if casted, ok := structType.(BACnetContextTagNull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagNull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagNull(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagNull(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagNull) GetTypeName() string {
	return "BACnetContextTagNull"
}

func (m *BACnetContextTagNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetContextTagNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagNullParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, header *BACnetTagHeader) (*BACnetContextTagNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((header.GetActualLength()) == (0))) {
		return nil, errors.WithStack(utils.ParseValidationError{"length field should be 0"})
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagNull")
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagNull{
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagNull")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
