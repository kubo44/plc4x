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

// BACnetConstructedDataNetworkNumber is the data-structure of this message
type BACnetConstructedDataNetworkNumber struct {
	*BACnetConstructedData
	NetworkNumber *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataNetworkNumber is the corresponding interface of BACnetConstructedDataNetworkNumber
type IBACnetConstructedDataNetworkNumber interface {
	IBACnetConstructedData
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataNetworkNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataNetworkNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNetworkNumber) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNetworkNumber) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNetworkNumber) GetNetworkNumber() *BACnetApplicationTagUnsignedInteger {
	return m.NetworkNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkNumber factory function for BACnetConstructedDataNetworkNumber
func NewBACnetConstructedDataNetworkNumber(networkNumber *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataNetworkNumber {
	_result := &BACnetConstructedDataNetworkNumber{
		NetworkNumber:         networkNumber,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNetworkNumber(structType interface{}) *BACnetConstructedDataNetworkNumber {
	if casted, ok := structType.(BACnetConstructedDataNetworkNumber); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkNumber); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkNumber(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkNumber(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNetworkNumber) GetTypeName() string {
	return "BACnetConstructedDataNetworkNumber"
}

func (m *BACnetConstructedDataNetworkNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNetworkNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkNumber)
	lengthInBits += m.NetworkNumber.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNetworkNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkNumberParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataNetworkNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumber)
	if pullErr := readBuffer.PullContext("networkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkNumber")
	}
	_networkNumber, _networkNumberErr := BACnetApplicationTagParse(readBuffer)
	if _networkNumberErr != nil {
		return nil, errors.Wrap(_networkNumberErr, "Error parsing 'networkNumber' field")
	}
	networkNumber := CastBACnetApplicationTagUnsignedInteger(_networkNumber)
	if closeErr := readBuffer.CloseContext("networkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkNumber")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkNumber")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNetworkNumber{
		NetworkNumber:         CastBACnetApplicationTagUnsignedInteger(networkNumber),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNetworkNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkNumber")
		}

		// Simple Field (networkNumber)
		if pushErr := writeBuffer.PushContext("networkNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkNumber")
		}
		_networkNumberErr := writeBuffer.WriteSerializable(m.NetworkNumber)
		if popErr := writeBuffer.PopContext("networkNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkNumber")
		}
		if _networkNumberErr != nil {
			return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNetworkNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
