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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CALReply_CR byte = 0x0D
const CALReply_LF byte = 0x0A

// CALReply is the data-structure of this message
type CALReply struct {
	CalType byte
	CalData *CALData
	Child   ICALReplyChild
}

// ICALReply is the corresponding interface of CALReply
type ICALReply interface {
	// GetCalType returns CalType (property field)
	GetCalType() byte
	// GetCalData returns CalData (property field)
	GetCalData() *CALData
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICALReplyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICALReply, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICALReplyChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CALReply, calType byte, calData *CALData)
	GetParent() *CALReply

	GetTypeName() string
	ICALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALReply) GetCalType() byte {
	return m.CalType
}

func (m *CALReply) GetCalData() *CALData {
	return m.CalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *CALReply) GetCr() byte {
	return CALReply_CR
}

func (m *CALReply) GetLf() byte {
	return CALReply_LF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReply factory function for CALReply
func NewCALReply(calType byte, calData *CALData) *CALReply {
	return &CALReply{CalType: calType, CalData: calData}
}

func CastCALReply(structType interface{}) *CALReply {
	if casted, ok := structType.(CALReply); ok {
		return &casted
	}
	if casted, ok := structType.(*CALReply); ok {
		return casted
	}
	if casted, ok := structType.(ICALReplyChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *CALReply) GetTypeName() string {
	return "CALReply"
}

func (m *CALReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CALReply) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits()

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyParse(readBuffer utils.ReadBuffer) (*CALReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (calType)
	currentPos = positionAware.GetPos()
	calType, _err := readBuffer.ReadByte("calType")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'calType' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CALReplyChild interface {
		InitializeParent(*CALReply, byte, *CALData)
		GetParent() *CALReply
	}
	var _child CALReplyChild
	var typeSwitchError error
	switch {
	case calType == 0x86: // CALReplyLong
		_child, typeSwitchError = CALReplyLongParse(readBuffer)
	case true: // CALReplyShort
		_child, typeSwitchError = CALReplyShortParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (calData)
	if pullErr := readBuffer.PullContext("calData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calData")
	}
	_calData, _calDataErr := CALDataParse(readBuffer)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field")
	}
	calData := CastCALData(_calData)
	if closeErr := readBuffer.CloseContext("calData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calData")
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CALReply_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CALReply_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field")
	}
	if lf != CALReply_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CALReply_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("CALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReply")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), calType, calData)
	return _child.GetParent(), nil
}

func (m *CALReply) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CALReply) SerializeParent(writeBuffer utils.WriteBuffer, child ICALReply, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CALReply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CALReply")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (calData)
	if pushErr := writeBuffer.PushContext("calData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for calData")
	}
	_calDataErr := writeBuffer.WriteSerializable(m.CalData)
	if popErr := writeBuffer.PopContext("calData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for calData")
	}
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("CALReply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CALReply")
	}
	return nil
}

func (m *CALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
