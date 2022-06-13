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

// BACnetConstructedDataAPDUTimeout is the data-structure of this message
type BACnetConstructedDataAPDUTimeout struct {
	*BACnetConstructedData
	ApduTimeout *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAPDUTimeout is the corresponding interface of BACnetConstructedDataAPDUTimeout
type IBACnetConstructedDataAPDUTimeout interface {
	IBACnetConstructedData
	// GetApduTimeout returns ApduTimeout (property field)
	GetApduTimeout() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAPDUTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAPDUTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAPDUTimeout) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAPDUTimeout) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAPDUTimeout) GetApduTimeout() *BACnetApplicationTagUnsignedInteger {
	return m.ApduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDUTimeout factory function for BACnetConstructedDataAPDUTimeout
func NewBACnetConstructedDataAPDUTimeout(apduTimeout *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAPDUTimeout {
	_result := &BACnetConstructedDataAPDUTimeout{
		ApduTimeout:           apduTimeout,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAPDUTimeout(structType interface{}) *BACnetConstructedDataAPDUTimeout {
	if casted, ok := structType.(BACnetConstructedDataAPDUTimeout); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAPDUTimeout(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAPDUTimeout(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataAPDUTimeout"
}

func (m *BACnetConstructedDataAPDUTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAPDUTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (apduTimeout)
	lengthInBits += m.ApduTimeout.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAPDUTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAPDUTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAPDUTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (apduTimeout)
	if pullErr := readBuffer.PullContext("apduTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apduTimeout")
	}
	_apduTimeout, _apduTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _apduTimeoutErr != nil {
		return nil, errors.Wrap(_apduTimeoutErr, "Error parsing 'apduTimeout' field")
	}
	apduTimeout := CastBACnetApplicationTagUnsignedInteger(_apduTimeout)
	if closeErr := readBuffer.CloseContext("apduTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apduTimeout")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDUTimeout")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAPDUTimeout{
		ApduTimeout:           CastBACnetApplicationTagUnsignedInteger(apduTimeout),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAPDUTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDUTimeout")
		}

		// Simple Field (apduTimeout)
		if pushErr := writeBuffer.PushContext("apduTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for apduTimeout")
		}
		_apduTimeoutErr := writeBuffer.WriteSerializable(m.ApduTimeout)
		if popErr := writeBuffer.PopContext("apduTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for apduTimeout")
		}
		if _apduTimeoutErr != nil {
			return errors.Wrap(_apduTimeoutErr, "Error serializing 'apduTimeout' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDUTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
