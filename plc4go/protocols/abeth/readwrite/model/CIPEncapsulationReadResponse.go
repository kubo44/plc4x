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

// CIPEncapsulationReadResponse is the data-structure of this message
type CIPEncapsulationReadResponse struct {
	*CIPEncapsulationPacket
	Response *DF1ResponseMessage

	// Arguments.
	Len uint16
}

// ICIPEncapsulationReadResponse is the corresponding interface of CIPEncapsulationReadResponse
type ICIPEncapsulationReadResponse interface {
	ICIPEncapsulationPacket
	// GetResponse returns Response (property field)
	GetResponse() *DF1ResponseMessage
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

func (m *CIPEncapsulationReadResponse) GetCommandType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CIPEncapsulationReadResponse) InitializeParent(parent *CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.CIPEncapsulationPacket.SessionHandle = sessionHandle
	m.CIPEncapsulationPacket.Status = status
	m.CIPEncapsulationPacket.SenderContext = senderContext
	m.CIPEncapsulationPacket.Options = options
}

func (m *CIPEncapsulationReadResponse) GetParent() *CIPEncapsulationPacket {
	return m.CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CIPEncapsulationReadResponse) GetResponse() *DF1ResponseMessage {
	return m.Response
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationReadResponse factory function for CIPEncapsulationReadResponse
func NewCIPEncapsulationReadResponse(response *DF1ResponseMessage, sessionHandle uint32, status uint32, senderContext []uint8, options uint32, len uint16) *CIPEncapsulationReadResponse {
	_result := &CIPEncapsulationReadResponse{
		Response:               response,
		CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result.Child = _result
	return _result
}

func CastCIPEncapsulationReadResponse(structType interface{}) *CIPEncapsulationReadResponse {
	if casted, ok := structType.(CIPEncapsulationReadResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(CIPEncapsulationPacket); ok {
		return CastCIPEncapsulationReadResponse(casted.Child)
	}
	if casted, ok := structType.(*CIPEncapsulationPacket); ok {
		return CastCIPEncapsulationReadResponse(casted.Child)
	}
	return nil
}

func (m *CIPEncapsulationReadResponse) GetTypeName() string {
	return "CIPEncapsulationReadResponse"
}

func (m *CIPEncapsulationReadResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CIPEncapsulationReadResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (response)
	lengthInBits += m.Response.GetLengthInBits()

	return lengthInBits
}

func (m *CIPEncapsulationReadResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPEncapsulationReadResponseParse(readBuffer utils.ReadBuffer, len uint16) (*CIPEncapsulationReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (response)
	if pullErr := readBuffer.PullContext("response"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for response")
	}
	_response, _responseErr := DF1ResponseMessageParse(readBuffer, uint16(len))
	if _responseErr != nil {
		return nil, errors.Wrap(_responseErr, "Error parsing 'response' field")
	}
	response := CastDF1ResponseMessage(_response)
	if closeErr := readBuffer.CloseContext("response"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for response")
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadResponse")
	}

	// Create a partially initialized instance
	_child := &CIPEncapsulationReadResponse{
		Response:               CastDF1ResponseMessage(response),
		CIPEncapsulationPacket: &CIPEncapsulationPacket{},
	}
	_child.CIPEncapsulationPacket.Child = _child
	return _child, nil
}

func (m *CIPEncapsulationReadResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadResponse")
		}

		// Simple Field (response)
		if pushErr := writeBuffer.PushContext("response"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for response")
		}
		_responseErr := writeBuffer.WriteSerializable(m.Response)
		if popErr := writeBuffer.PopContext("response"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for response")
		}
		if _responseErr != nil {
			return errors.Wrap(_responseErr, "Error serializing 'response' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CIPEncapsulationReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
