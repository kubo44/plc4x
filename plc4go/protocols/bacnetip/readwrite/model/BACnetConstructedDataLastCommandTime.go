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

// BACnetConstructedDataLastCommandTime is the data-structure of this message
type BACnetConstructedDataLastCommandTime struct {
	*BACnetConstructedData
	LastCommandTime *BACnetTimeStamp

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLastCommandTime is the corresponding interface of BACnetConstructedDataLastCommandTime
type IBACnetConstructedDataLastCommandTime interface {
	IBACnetConstructedData
	// GetLastCommandTime returns LastCommandTime (property field)
	GetLastCommandTime() *BACnetTimeStamp
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

func (m *BACnetConstructedDataLastCommandTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLastCommandTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_COMMAND_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLastCommandTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLastCommandTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLastCommandTime) GetLastCommandTime() *BACnetTimeStamp {
	return m.LastCommandTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCommandTime factory function for BACnetConstructedDataLastCommandTime
func NewBACnetConstructedDataLastCommandTime(lastCommandTime *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLastCommandTime {
	_result := &BACnetConstructedDataLastCommandTime{
		LastCommandTime:       lastCommandTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLastCommandTime(structType interface{}) *BACnetConstructedDataLastCommandTime {
	if casted, ok := structType.(BACnetConstructedDataLastCommandTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCommandTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCommandTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCommandTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLastCommandTime) GetTypeName() string {
	return "BACnetConstructedDataLastCommandTime"
}

func (m *BACnetConstructedDataLastCommandTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLastCommandTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastCommandTime)
	lengthInBits += m.LastCommandTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLastCommandTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastCommandTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLastCommandTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCommandTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCommandTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCommandTime)
	if pullErr := readBuffer.PullContext("lastCommandTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastCommandTime")
	}
	_lastCommandTime, _lastCommandTimeErr := BACnetTimeStampParse(readBuffer)
	if _lastCommandTimeErr != nil {
		return nil, errors.Wrap(_lastCommandTimeErr, "Error parsing 'lastCommandTime' field")
	}
	lastCommandTime := CastBACnetTimeStamp(_lastCommandTime)
	if closeErr := readBuffer.CloseContext("lastCommandTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastCommandTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCommandTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCommandTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLastCommandTime{
		LastCommandTime:       CastBACnetTimeStamp(lastCommandTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLastCommandTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCommandTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCommandTime")
		}

		// Simple Field (lastCommandTime)
		if pushErr := writeBuffer.PushContext("lastCommandTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastCommandTime")
		}
		_lastCommandTimeErr := writeBuffer.WriteSerializable(m.LastCommandTime)
		if popErr := writeBuffer.PopContext("lastCommandTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastCommandTime")
		}
		if _lastCommandTimeErr != nil {
			return errors.Wrap(_lastCommandTimeErr, "Error serializing 'lastCommandTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCommandTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCommandTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLastCommandTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
