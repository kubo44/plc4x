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

// S7PayloadReadVarResponse is the data-structure of this message
type S7PayloadReadVarResponse struct {
	*S7Payload
	Items []*S7VarPayloadDataItem

	// Arguments.
	Parameter *S7Parameter
}

// IS7PayloadReadVarResponse is the corresponding interface of S7PayloadReadVarResponse
type IS7PayloadReadVarResponse interface {
	IS7Payload
	// GetItems returns Items (property field)
	GetItems() []*S7VarPayloadDataItem
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

func (m *S7PayloadReadVarResponse) GetParameterParameterType() uint8 {
	return 0x04
}

func (m *S7PayloadReadVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadReadVarResponse) InitializeParent(parent *S7Payload) {}

func (m *S7PayloadReadVarResponse) GetParent() *S7Payload {
	return m.S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadReadVarResponse) GetItems() []*S7VarPayloadDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadReadVarResponse factory function for S7PayloadReadVarResponse
func NewS7PayloadReadVarResponse(items []*S7VarPayloadDataItem, parameter *S7Parameter) *S7PayloadReadVarResponse {
	_result := &S7PayloadReadVarResponse{
		Items:     items,
		S7Payload: NewS7Payload(parameter),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadReadVarResponse(structType interface{}) *S7PayloadReadVarResponse {
	if casted, ok := structType.(S7PayloadReadVarResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadReadVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7Payload); ok {
		return CastS7PayloadReadVarResponse(casted.Child)
	}
	if casted, ok := structType.(*S7Payload); ok {
		return CastS7PayloadReadVarResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadReadVarResponse) GetTypeName() string {
	return "S7PayloadReadVarResponse"
}

func (m *S7PayloadReadVarResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadReadVarResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadReadVarResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadReadVarResponseParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7PayloadReadVarResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadReadVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadReadVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]*S7VarPayloadDataItem, CastS7ParameterReadVarResponse(parameter).GetNumItems())
	{
		for curItem := uint16(0); curItem < uint16(CastS7ParameterReadVarResponse(parameter).GetNumItems()); curItem++ {
			_item, _err := S7VarPayloadDataItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastS7VarPayloadDataItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadReadVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadReadVarResponse")
	}

	// Create a partially initialized instance
	_child := &S7PayloadReadVarResponse{
		Items:     items,
		S7Payload: &S7Payload{},
	}
	_child.S7Payload.Child = _child
	return _child, nil
}

func (m *S7PayloadReadVarResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadReadVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadReadVarResponse")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for items")
			}
			for _, _element := range m.Items {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for items")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadReadVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadReadVarResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadReadVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
