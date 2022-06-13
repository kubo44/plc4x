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

// BACnetEventParameterUnsignedRange is the data-structure of this message
type BACnetEventParameterUnsignedRange struct {
	*BACnetEventParameter
	OpeningTag *BACnetOpeningTag
	TimeDelay  *BACnetContextTagUnsignedInteger
	LowLimit   *BACnetContextTagUnsignedInteger
	HighLimit  *BACnetContextTagUnsignedInteger
	ClosingTag *BACnetClosingTag
}

// IBACnetEventParameterUnsignedRange is the corresponding interface of BACnetEventParameterUnsignedRange
type IBACnetEventParameterUnsignedRange interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() *BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() *BACnetContextTagUnsignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() *BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
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

func (m *BACnetEventParameterUnsignedRange) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterUnsignedRange) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterUnsignedRange) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterUnsignedRange) GetTimeDelay() *BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *BACnetEventParameterUnsignedRange) GetLowLimit() *BACnetContextTagUnsignedInteger {
	return m.LowLimit
}

func (m *BACnetEventParameterUnsignedRange) GetHighLimit() *BACnetContextTagUnsignedInteger {
	return m.HighLimit
}

func (m *BACnetEventParameterUnsignedRange) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterUnsignedRange factory function for BACnetEventParameterUnsignedRange
func NewBACnetEventParameterUnsignedRange(openingTag *BACnetOpeningTag, timeDelay *BACnetContextTagUnsignedInteger, lowLimit *BACnetContextTagUnsignedInteger, highLimit *BACnetContextTagUnsignedInteger, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterUnsignedRange {
	_result := &BACnetEventParameterUnsignedRange{
		OpeningTag:           openingTag,
		TimeDelay:            timeDelay,
		LowLimit:             lowLimit,
		HighLimit:            highLimit,
		ClosingTag:           closingTag,
		BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterUnsignedRange(structType interface{}) *BACnetEventParameterUnsignedRange {
	if casted, ok := structType.(BACnetEventParameterUnsignedRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterUnsignedRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterUnsignedRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterUnsignedRange(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterUnsignedRange) GetTypeName() string {
	return "BACnetEventParameterUnsignedRange"
}

func (m *BACnetEventParameterUnsignedRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterUnsignedRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits()

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventParameterUnsignedRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterUnsignedRangeParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterUnsignedRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterUnsignedRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterUnsignedRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(11)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := CastBACnetContextTagUnsignedInteger(_timeDelay)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowLimit")
	}
	_lowLimit, _lowLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field")
	}
	lowLimit := CastBACnetContextTagUnsignedInteger(_lowLimit)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowLimit")
	}

	// Simple Field (highLimit)
	if pullErr := readBuffer.PullContext("highLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highLimit")
	}
	_highLimit, _highLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _highLimitErr != nil {
		return nil, errors.Wrap(_highLimitErr, "Error parsing 'highLimit' field")
	}
	highLimit := CastBACnetContextTagUnsignedInteger(_highLimit)
	if closeErr := readBuffer.CloseContext("highLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highLimit")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(11)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterUnsignedRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterUnsignedRange")
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterUnsignedRange{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		TimeDelay:            CastBACnetContextTagUnsignedInteger(timeDelay),
		LowLimit:             CastBACnetContextTagUnsignedInteger(lowLimit),
		HighLimit:            CastBACnetContextTagUnsignedInteger(highLimit),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetEventParameter: &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterUnsignedRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterUnsignedRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterUnsignedRange")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(m.OpeningTag)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(m.TimeDelay)
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := writeBuffer.WriteSerializable(m.LowLimit)
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := writeBuffer.WriteSerializable(m.HighLimit)
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(m.ClosingTag)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterUnsignedRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterUnsignedRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterUnsignedRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
