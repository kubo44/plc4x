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

// BACnetTagPayloadOctetString is the data-structure of this message
type BACnetTagPayloadOctetString struct {
	Octets []byte

	// Arguments.
	ActualLength uint32
}

// IBACnetTagPayloadOctetString is the corresponding interface of BACnetTagPayloadOctetString
type IBACnetTagPayloadOctetString interface {
	// GetOctets returns Octets (property field)
	GetOctets() []byte
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

func (m *BACnetTagPayloadOctetString) GetOctets() []byte {
	return m.Octets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadOctetString factory function for BACnetTagPayloadOctetString
func NewBACnetTagPayloadOctetString(octets []byte, actualLength uint32) *BACnetTagPayloadOctetString {
	return &BACnetTagPayloadOctetString{Octets: octets, ActualLength: actualLength}
}

func CastBACnetTagPayloadOctetString(structType interface{}) *BACnetTagPayloadOctetString {
	if casted, ok := structType.(BACnetTagPayloadOctetString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadOctetString); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadOctetString) GetTypeName() string {
	return "BACnetTagPayloadOctetString"
}

func (m *BACnetTagPayloadOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Octets) > 0 {
		lengthInBits += 8 * uint16(len(m.Octets))
	}

	return lengthInBits
}

func (m *BACnetTagPayloadOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadOctetStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (octets)
	numberOfBytesoctets := int(actualLength)
	octets, _readArrayErr := readBuffer.ReadByteArray("octets", numberOfBytesoctets)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'octets' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadOctetString")
	}

	// Create the instance
	return NewBACnetTagPayloadOctetString(octets, actualLength), nil
}

func (m *BACnetTagPayloadOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadOctetString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadOctetString")
	}

	// Array Field (octets)
	if m.Octets != nil {
		// Byte Array field (octets)
		_writeArrayErr := writeBuffer.WriteByteArray("octets", m.Octets)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'octets' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadOctetString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadOctetString")
	}
	return nil
}

func (m *BACnetTagPayloadOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
