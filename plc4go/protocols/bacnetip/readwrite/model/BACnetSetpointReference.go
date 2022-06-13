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

// BACnetSetpointReference is the data-structure of this message
type BACnetSetpointReference struct {
	SetPointReference *BACnetObjectPropertyReferenceEnclosed
}

// IBACnetSetpointReference is the corresponding interface of BACnetSetpointReference
type IBACnetSetpointReference interface {
	// GetSetPointReference returns SetPointReference (property field)
	GetSetPointReference() *BACnetObjectPropertyReferenceEnclosed
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

func (m *BACnetSetpointReference) GetSetPointReference() *BACnetObjectPropertyReferenceEnclosed {
	return m.SetPointReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSetpointReference factory function for BACnetSetpointReference
func NewBACnetSetpointReference(setPointReference *BACnetObjectPropertyReferenceEnclosed) *BACnetSetpointReference {
	return &BACnetSetpointReference{SetPointReference: setPointReference}
}

func CastBACnetSetpointReference(structType interface{}) *BACnetSetpointReference {
	if casted, ok := structType.(BACnetSetpointReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetSetpointReference); ok {
		return casted
	}
	return nil
}

func (m *BACnetSetpointReference) GetTypeName() string {
	return "BACnetSetpointReference"
}

func (m *BACnetSetpointReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetSetpointReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (setPointReference)
	if m.SetPointReference != nil {
		lengthInBits += (*m.SetPointReference).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetSetpointReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSetpointReferenceParse(readBuffer utils.ReadBuffer) (*BACnetSetpointReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSetpointReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSetpointReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (setPointReference) (Can be skipped, if a given expression evaluates to false)
	var setPointReference *BACnetObjectPropertyReferenceEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("setPointReference"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for setPointReference")
		}
		_val, _err := BACnetObjectPropertyReferenceEnclosedParse(readBuffer, uint8(0))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'setPointReference' field")
		default:
			setPointReference = CastBACnetObjectPropertyReferenceEnclosed(_val)
			if closeErr := readBuffer.CloseContext("setPointReference"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for setPointReference")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetSetpointReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSetpointReference")
	}

	// Create the instance
	return NewBACnetSetpointReference(setPointReference), nil
}

func (m *BACnetSetpointReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetSetpointReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSetpointReference")
	}

	// Optional Field (setPointReference) (Can be skipped, if the value is null)
	var setPointReference *BACnetObjectPropertyReferenceEnclosed = nil
	if m.SetPointReference != nil {
		if pushErr := writeBuffer.PushContext("setPointReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for setPointReference")
		}
		setPointReference = m.SetPointReference
		_setPointReferenceErr := writeBuffer.WriteSerializable(setPointReference)
		if popErr := writeBuffer.PopContext("setPointReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for setPointReference")
		}
		if _setPointReferenceErr != nil {
			return errors.Wrap(_setPointReferenceErr, "Error serializing 'setPointReference' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetSetpointReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSetpointReference")
	}
	return nil
}

func (m *BACnetSetpointReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
