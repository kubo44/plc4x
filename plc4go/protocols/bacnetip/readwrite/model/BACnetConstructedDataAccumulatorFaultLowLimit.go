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

// BACnetConstructedDataAccumulatorFaultLowLimit is the data-structure of this message
type BACnetConstructedDataAccumulatorFaultLowLimit struct {
	*BACnetConstructedData
	FaultLowLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAccumulatorFaultLowLimit is the corresponding interface of BACnetConstructedDataAccumulatorFaultLowLimit
type IBACnetConstructedDataAccumulatorFaultLowLimit interface {
	IBACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetFaultLowLimit() *BACnetApplicationTagUnsignedInteger {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorFaultLowLimit factory function for BACnetConstructedDataAccumulatorFaultLowLimit
func NewBACnetConstructedDataAccumulatorFaultLowLimit(faultLowLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAccumulatorFaultLowLimit {
	_result := &BACnetConstructedDataAccumulatorFaultLowLimit{
		FaultLowLimit:         faultLowLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccumulatorFaultLowLimit(structType interface{}) *BACnetConstructedDataAccumulatorFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorFaultLowLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorFaultLowLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccumulatorFaultLowLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorFaultLowLimit"
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccumulatorFaultLowLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAccumulatorFaultLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultLowLimit)
	if pullErr := readBuffer.PullContext("faultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultLowLimit")
	}
	_faultLowLimit, _faultLowLimitErr := BACnetApplicationTagParse(readBuffer)
	if _faultLowLimitErr != nil {
		return nil, errors.Wrap(_faultLowLimitErr, "Error parsing 'faultLowLimit' field")
	}
	faultLowLimit := CastBACnetApplicationTagUnsignedInteger(_faultLowLimit)
	if closeErr := readBuffer.CloseContext("faultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultLowLimit")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorFaultLowLimit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccumulatorFaultLowLimit{
		FaultLowLimit:         CastBACnetApplicationTagUnsignedInteger(faultLowLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorFaultLowLimit")
		}

		// Simple Field (faultLowLimit)
		if pushErr := writeBuffer.PushContext("faultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultLowLimit")
		}
		_faultLowLimitErr := writeBuffer.WriteSerializable(m.FaultLowLimit)
		if popErr := writeBuffer.PopContext("faultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultLowLimit")
		}
		if _faultLowLimitErr != nil {
			return errors.Wrap(_faultLowLimitErr, "Error serializing 'faultLowLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorFaultLowLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccumulatorFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
