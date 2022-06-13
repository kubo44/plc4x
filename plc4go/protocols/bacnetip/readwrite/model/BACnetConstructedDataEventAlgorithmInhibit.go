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

// BACnetConstructedDataEventAlgorithmInhibit is the data-structure of this message
type BACnetConstructedDataEventAlgorithmInhibit struct {
	*BACnetConstructedData
	EventAlgorithmInhibit *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventAlgorithmInhibit is the corresponding interface of BACnetConstructedDataEventAlgorithmInhibit
type IBACnetConstructedDataEventAlgorithmInhibit interface {
	IBACnetConstructedData
	// GetEventAlgorithmInhibit returns EventAlgorithmInhibit (property field)
	GetEventAlgorithmInhibit() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ALGORITHM_INHIBIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventAlgorithmInhibit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetEventAlgorithmInhibit() *BACnetApplicationTagBoolean {
	return m.EventAlgorithmInhibit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventAlgorithmInhibit factory function for BACnetConstructedDataEventAlgorithmInhibit
func NewBACnetConstructedDataEventAlgorithmInhibit(eventAlgorithmInhibit *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventAlgorithmInhibit {
	_result := &BACnetConstructedDataEventAlgorithmInhibit{
		EventAlgorithmInhibit: eventAlgorithmInhibit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventAlgorithmInhibit(structType interface{}) *BACnetConstructedDataEventAlgorithmInhibit {
	if casted, ok := structType.(BACnetConstructedDataEventAlgorithmInhibit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventAlgorithmInhibit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventAlgorithmInhibit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventAlgorithmInhibit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetTypeName() string {
	return "BACnetConstructedDataEventAlgorithmInhibit"
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventAlgorithmInhibit)
	lengthInBits += m.EventAlgorithmInhibit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventAlgorithmInhibitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventAlgorithmInhibit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventAlgorithmInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventAlgorithmInhibit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventAlgorithmInhibit)
	if pullErr := readBuffer.PullContext("eventAlgorithmInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventAlgorithmInhibit")
	}
	_eventAlgorithmInhibit, _eventAlgorithmInhibitErr := BACnetApplicationTagParse(readBuffer)
	if _eventAlgorithmInhibitErr != nil {
		return nil, errors.Wrap(_eventAlgorithmInhibitErr, "Error parsing 'eventAlgorithmInhibit' field")
	}
	eventAlgorithmInhibit := CastBACnetApplicationTagBoolean(_eventAlgorithmInhibit)
	if closeErr := readBuffer.CloseContext("eventAlgorithmInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventAlgorithmInhibit")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventAlgorithmInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventAlgorithmInhibit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventAlgorithmInhibit{
		EventAlgorithmInhibit: CastBACnetApplicationTagBoolean(eventAlgorithmInhibit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventAlgorithmInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventAlgorithmInhibit")
		}

		// Simple Field (eventAlgorithmInhibit)
		if pushErr := writeBuffer.PushContext("eventAlgorithmInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventAlgorithmInhibit")
		}
		_eventAlgorithmInhibitErr := writeBuffer.WriteSerializable(m.EventAlgorithmInhibit)
		if popErr := writeBuffer.PopContext("eventAlgorithmInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventAlgorithmInhibit")
		}
		if _eventAlgorithmInhibitErr != nil {
			return errors.Wrap(_eventAlgorithmInhibitErr, "Error serializing 'eventAlgorithmInhibit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventAlgorithmInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventAlgorithmInhibit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventAlgorithmInhibit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
