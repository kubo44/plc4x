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

// SerialInterfaceAddress is the data-structure of this message
type SerialInterfaceAddress struct {
	Address byte
}

// ISerialInterfaceAddress is the corresponding interface of SerialInterfaceAddress
type ISerialInterfaceAddress interface {
	// GetAddress returns Address (property field)
	GetAddress() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SerialInterfaceAddress) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSerialInterfaceAddress factory function for SerialInterfaceAddress
func NewSerialInterfaceAddress(address byte) *SerialInterfaceAddress {
	return &SerialInterfaceAddress{Address: address}
}

func CastSerialInterfaceAddress(structType interface{}) *SerialInterfaceAddress {
	if casted, ok := structType.(SerialInterfaceAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*SerialInterfaceAddress); ok {
		return casted
	}
	return nil
}

func (m *SerialInterfaceAddress) GetTypeName() string {
	return "SerialInterfaceAddress"
}

func (m *SerialInterfaceAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SerialInterfaceAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	return lengthInBits
}

func (m *SerialInterfaceAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SerialInterfaceAddressParse(readBuffer utils.ReadBuffer) (*SerialInterfaceAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SerialInterfaceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SerialInterfaceAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadByte("address")
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	if closeErr := readBuffer.CloseContext("SerialInterfaceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SerialInterfaceAddress")
	}

	// Create the instance
	return NewSerialInterfaceAddress(address), nil
}

func (m *SerialInterfaceAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SerialInterfaceAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SerialInterfaceAddress")
	}

	// Simple Field (address)
	address := byte(m.Address)
	_addressErr := writeBuffer.WriteByte("address", (address))
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}

	if popErr := writeBuffer.PopContext("SerialInterfaceAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SerialInterfaceAddress")
	}
	return nil
}

func (m *SerialInterfaceAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
