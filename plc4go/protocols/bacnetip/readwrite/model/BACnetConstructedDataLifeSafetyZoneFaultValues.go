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

// BACnetConstructedDataLifeSafetyZoneFaultValues is the data-structure of this message
type BACnetConstructedDataLifeSafetyZoneFaultValues struct {
	*BACnetConstructedData
	FaultValues []*BACnetLifeSafetyStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLifeSafetyZoneFaultValues is the corresponding interface of BACnetConstructedDataLifeSafetyZoneFaultValues
type IBACnetConstructedDataLifeSafetyZoneFaultValues interface {
	IBACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []*BACnetLifeSafetyStateTagged
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

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetFaultValues() []*BACnetLifeSafetyStateTagged {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyZoneFaultValues factory function for BACnetConstructedDataLifeSafetyZoneFaultValues
func NewBACnetConstructedDataLifeSafetyZoneFaultValues(faultValues []*BACnetLifeSafetyStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLifeSafetyZoneFaultValues {
	_result := &BACnetConstructedDataLifeSafetyZoneFaultValues{
		FaultValues:           faultValues,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLifeSafetyZoneFaultValues(structType interface{}) *BACnetConstructedDataLifeSafetyZoneFaultValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneFaultValues); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZoneFaultValues(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZoneFaultValues(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneFaultValues"
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyZoneFaultValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLifeSafetyZoneFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (faultValues)
	if pullErr := readBuffer.PullContext("faultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultValues")
	}
	// Terminated array
	faultValues := make([]*BACnetLifeSafetyStateTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'faultValues' field")
			}
			faultValues = append(faultValues, CastBACnetLifeSafetyStateTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("faultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneFaultValues")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLifeSafetyZoneFaultValues{
		FaultValues:           faultValues,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneFaultValues")
		}

		// Array Field (faultValues)
		if m.FaultValues != nil {
			if pushErr := writeBuffer.PushContext("faultValues", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for faultValues")
			}
			for _, _element := range m.FaultValues {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'faultValues' field")
				}
			}
			if popErr := writeBuffer.PopContext("faultValues", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for faultValues")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneFaultValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLifeSafetyZoneFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
