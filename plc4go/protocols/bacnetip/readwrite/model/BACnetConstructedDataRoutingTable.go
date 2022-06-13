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

// BACnetConstructedDataRoutingTable is the data-structure of this message
type BACnetConstructedDataRoutingTable struct {
	*BACnetConstructedData
	RoutingTable []*BACnetRouterEntry

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataRoutingTable is the corresponding interface of BACnetConstructedDataRoutingTable
type IBACnetConstructedDataRoutingTable interface {
	IBACnetConstructedData
	// GetRoutingTable returns RoutingTable (property field)
	GetRoutingTable() []*BACnetRouterEntry
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

func (m *BACnetConstructedDataRoutingTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRoutingTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ROUTING_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRoutingTable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRoutingTable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRoutingTable) GetRoutingTable() []*BACnetRouterEntry {
	return m.RoutingTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRoutingTable factory function for BACnetConstructedDataRoutingTable
func NewBACnetConstructedDataRoutingTable(routingTable []*BACnetRouterEntry, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataRoutingTable {
	_result := &BACnetConstructedDataRoutingTable{
		RoutingTable:          routingTable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRoutingTable(structType interface{}) *BACnetConstructedDataRoutingTable {
	if casted, ok := structType.(BACnetConstructedDataRoutingTable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRoutingTable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRoutingTable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRoutingTable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRoutingTable) GetTypeName() string {
	return "BACnetConstructedDataRoutingTable"
}

func (m *BACnetConstructedDataRoutingTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRoutingTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.RoutingTable) > 0 {
		for _, element := range m.RoutingTable {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataRoutingTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRoutingTableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataRoutingTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRoutingTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRoutingTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (routingTable)
	if pullErr := readBuffer.PullContext("routingTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for routingTable")
	}
	// Terminated array
	routingTable := make([]*BACnetRouterEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRouterEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'routingTable' field")
			}
			routingTable = append(routingTable, CastBACnetRouterEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("routingTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for routingTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRoutingTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRoutingTable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRoutingTable{
		RoutingTable:          routingTable,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRoutingTable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRoutingTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRoutingTable")
		}

		// Array Field (routingTable)
		if m.RoutingTable != nil {
			if pushErr := writeBuffer.PushContext("routingTable", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for routingTable")
			}
			for _, _element := range m.RoutingTable {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'routingTable' field")
				}
			}
			if popErr := writeBuffer.PopContext("routingTable", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for routingTable")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRoutingTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRoutingTable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRoutingTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
