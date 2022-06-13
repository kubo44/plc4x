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

// BACnetConstructedDataLargeAnalogValueFaultHighLimit is the data-structure of this message
type BACnetConstructedDataLargeAnalogValueFaultHighLimit struct {
	*BACnetConstructedData
	FaultHighLimit *BACnetApplicationTagDouble

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLargeAnalogValueFaultHighLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueFaultHighLimit
type IBACnetConstructedDataLargeAnalogValueFaultHighLimit interface {
	IBACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() *BACnetApplicationTagDouble
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

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetFaultHighLimit() *BACnetApplicationTagDouble {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLargeAnalogValueFaultHighLimit factory function for BACnetConstructedDataLargeAnalogValueFaultHighLimit
func NewBACnetConstructedDataLargeAnalogValueFaultHighLimit(faultHighLimit *BACnetApplicationTagDouble, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	_result := &BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		FaultHighLimit:        faultHighLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLargeAnalogValueFaultHighLimit(structType interface{}) *BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueFaultHighLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLargeAnalogValueFaultHighLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueFaultHighLimit"
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLargeAnalogValueFaultHighLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLargeAnalogValueFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultHighLimit)
	if pullErr := readBuffer.PullContext("faultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultHighLimit")
	}
	_faultHighLimit, _faultHighLimitErr := BACnetApplicationTagParse(readBuffer)
	if _faultHighLimitErr != nil {
		return nil, errors.Wrap(_faultHighLimitErr, "Error parsing 'faultHighLimit' field")
	}
	faultHighLimit := CastBACnetApplicationTagDouble(_faultHighLimit)
	if closeErr := readBuffer.CloseContext("faultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultHighLimit")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		FaultHighLimit:        CastBACnetApplicationTagDouble(faultHighLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}

		// Simple Field (faultHighLimit)
		if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
		}
		_faultHighLimitErr := writeBuffer.WriteSerializable(m.FaultHighLimit)
		if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultHighLimit")
		}
		if _faultHighLimitErr != nil {
			return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLargeAnalogValueFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
