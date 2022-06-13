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

// BACnetConstructedDataPassbackTimeout is the data-structure of this message
type BACnetConstructedDataPassbackTimeout struct {
	*BACnetConstructedData
	PassbackTimeout *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPassbackTimeout is the corresponding interface of BACnetConstructedDataPassbackTimeout
type IBACnetConstructedDataPassbackTimeout interface {
	IBACnetConstructedData
	// GetPassbackTimeout returns PassbackTimeout (property field)
	GetPassbackTimeout() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataPassbackTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPassbackTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPassbackTimeout) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPassbackTimeout) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPassbackTimeout) GetPassbackTimeout() *BACnetApplicationTagUnsignedInteger {
	return m.PassbackTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassbackTimeout factory function for BACnetConstructedDataPassbackTimeout
func NewBACnetConstructedDataPassbackTimeout(passbackTimeout *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPassbackTimeout {
	_result := &BACnetConstructedDataPassbackTimeout{
		PassbackTimeout:       passbackTimeout,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPassbackTimeout(structType interface{}) *BACnetConstructedDataPassbackTimeout {
	if casted, ok := structType.(BACnetConstructedDataPassbackTimeout); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackTimeout); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassbackTimeout(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPassbackTimeout(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPassbackTimeout) GetTypeName() string {
	return "BACnetConstructedDataPassbackTimeout"
}

func (m *BACnetConstructedDataPassbackTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPassbackTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passbackTimeout)
	lengthInBits += m.PassbackTimeout.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPassbackTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassbackTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPassbackTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passbackTimeout)
	if pullErr := readBuffer.PullContext("passbackTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passbackTimeout")
	}
	_passbackTimeout, _passbackTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _passbackTimeoutErr != nil {
		return nil, errors.Wrap(_passbackTimeoutErr, "Error parsing 'passbackTimeout' field")
	}
	passbackTimeout := CastBACnetApplicationTagUnsignedInteger(_passbackTimeout)
	if closeErr := readBuffer.CloseContext("passbackTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passbackTimeout")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackTimeout")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPassbackTimeout{
		PassbackTimeout:       CastBACnetApplicationTagUnsignedInteger(passbackTimeout),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPassbackTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackTimeout")
		}

		// Simple Field (passbackTimeout)
		if pushErr := writeBuffer.PushContext("passbackTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passbackTimeout")
		}
		_passbackTimeoutErr := writeBuffer.WriteSerializable(m.PassbackTimeout)
		if popErr := writeBuffer.PopContext("passbackTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passbackTimeout")
		}
		if _passbackTimeoutErr != nil {
			return errors.Wrap(_passbackTimeoutErr, "Error serializing 'passbackTimeout' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPassbackTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
