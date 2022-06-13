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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataBitText is the data-structure of this message
type BACnetConstructedDataBitText struct {
	*BACnetConstructedData
	NumberOfDataElements *BACnetApplicationTagUnsignedInteger
	BitText              []*BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBitText is the corresponding interface of BACnetConstructedDataBitText
type IBACnetConstructedDataBitText interface {
	IBACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger
	// GetBitText returns BitText (property field)
	GetBitText() []*BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
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

func (m *BACnetConstructedDataBitText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBitText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BIT_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBitText) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBitText) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBitText) GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *BACnetConstructedDataBitText) GetBitText() []*BACnetApplicationTagCharacterString {
	return m.BitText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataBitText) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBitText factory function for BACnetConstructedDataBitText
func NewBACnetConstructedDataBitText(numberOfDataElements *BACnetApplicationTagUnsignedInteger, bitText []*BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBitText {
	_result := &BACnetConstructedDataBitText{
		NumberOfDataElements:  numberOfDataElements,
		BitText:               bitText,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBitText(structType interface{}) *BACnetConstructedDataBitText {
	if casted, ok := structType.(BACnetConstructedDataBitText); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitText); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitText(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitText(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBitText) GetTypeName() string {
	return "BACnetConstructedDataBitText"
}

func (m *BACnetConstructedDataBitText) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBitText) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += (*m.NumberOfDataElements).GetLengthInBits()
	}

	// Array field
	if len(m.BitText) > 0 {
		for _, element := range m.BitText {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataBitText) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBitTextParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBitText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBitText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = CastBACnetApplicationTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (bitText)
	if pullErr := readBuffer.PullContext("bitText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitText")
	}
	// Terminated array
	bitText := make([]*BACnetApplicationTagCharacterString, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'bitText' field")
			}
			bitText = append(bitText, CastBACnetApplicationTagCharacterString(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("bitText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBitText")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBitText{
		NumberOfDataElements:  CastBACnetApplicationTagUnsignedInteger(numberOfDataElements),
		BitText:               bitText,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBitText) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBitText")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
		if m.NumberOfDataElements != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.NumberOfDataElements
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (bitText)
		if m.BitText != nil {
			if pushErr := writeBuffer.PushContext("bitText", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bitText")
			}
			for _, _element := range m.BitText {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'bitText' field")
				}
			}
			if popErr := writeBuffer.PopContext("bitText", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bitText")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBitText")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBitText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
