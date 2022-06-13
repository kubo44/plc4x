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

// BACnetShedLevelAmount is the data-structure of this message
type BACnetShedLevelAmount struct {
	*BACnetShedLevel
	Amount *BACnetContextTagReal
}

// IBACnetShedLevelAmount is the corresponding interface of BACnetShedLevelAmount
type IBACnetShedLevelAmount interface {
	IBACnetShedLevel
	// GetAmount returns Amount (property field)
	GetAmount() *BACnetContextTagReal
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

func (m *BACnetShedLevelAmount) InitializeParent(parent *BACnetShedLevel, peekedTagHeader *BACnetTagHeader) {
	m.BACnetShedLevel.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetShedLevelAmount) GetParent() *BACnetShedLevel {
	return m.BACnetShedLevel
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetShedLevelAmount) GetAmount() *BACnetContextTagReal {
	return m.Amount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetShedLevelAmount factory function for BACnetShedLevelAmount
func NewBACnetShedLevelAmount(amount *BACnetContextTagReal, peekedTagHeader *BACnetTagHeader) *BACnetShedLevelAmount {
	_result := &BACnetShedLevelAmount{
		Amount:          amount,
		BACnetShedLevel: NewBACnetShedLevel(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetShedLevelAmount(structType interface{}) *BACnetShedLevelAmount {
	if casted, ok := structType.(BACnetShedLevelAmount); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetShedLevelAmount); ok {
		return casted
	}
	if casted, ok := structType.(BACnetShedLevel); ok {
		return CastBACnetShedLevelAmount(casted.Child)
	}
	if casted, ok := structType.(*BACnetShedLevel); ok {
		return CastBACnetShedLevelAmount(casted.Child)
	}
	return nil
}

func (m *BACnetShedLevelAmount) GetTypeName() string {
	return "BACnetShedLevelAmount"
}

func (m *BACnetShedLevelAmount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetShedLevelAmount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (amount)
	lengthInBits += m.Amount.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetShedLevelAmount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetShedLevelAmountParse(readBuffer utils.ReadBuffer) (*BACnetShedLevelAmount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelAmount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelAmount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (amount)
	if pullErr := readBuffer.PullContext("amount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for amount")
	}
	_amount, _amountErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _amountErr != nil {
		return nil, errors.Wrap(_amountErr, "Error parsing 'amount' field")
	}
	amount := CastBACnetContextTagReal(_amount)
	if closeErr := readBuffer.CloseContext("amount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for amount")
	}

	if closeErr := readBuffer.CloseContext("BACnetShedLevelAmount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelAmount")
	}

	// Create a partially initialized instance
	_child := &BACnetShedLevelAmount{
		Amount:          CastBACnetContextTagReal(amount),
		BACnetShedLevel: &BACnetShedLevel{},
	}
	_child.BACnetShedLevel.Child = _child
	return _child, nil
}

func (m *BACnetShedLevelAmount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelAmount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelAmount")
		}

		// Simple Field (amount)
		if pushErr := writeBuffer.PushContext("amount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for amount")
		}
		_amountErr := writeBuffer.WriteSerializable(m.Amount)
		if popErr := writeBuffer.PopContext("amount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for amount")
		}
		if _amountErr != nil {
			return errors.Wrap(_amountErr, "Error serializing 'amount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelAmount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelAmount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetShedLevelAmount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
