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

// SysexCommandSamplingInterval is the data-structure of this message
type SysexCommandSamplingInterval struct {
	*SysexCommand
}

// ISysexCommandSamplingInterval is the corresponding interface of SysexCommandSamplingInterval
type ISysexCommandSamplingInterval interface {
	ISysexCommand
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

func (m *SysexCommandSamplingInterval) GetCommandType() uint8 {
	return 0x7A
}

func (m *SysexCommandSamplingInterval) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandSamplingInterval) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandSamplingInterval) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandSamplingInterval factory function for SysexCommandSamplingInterval
func NewSysexCommandSamplingInterval() *SysexCommandSamplingInterval {
	_result := &SysexCommandSamplingInterval{
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandSamplingInterval(structType interface{}) *SysexCommandSamplingInterval {
	if casted, ok := structType.(SysexCommandSamplingInterval); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandSamplingInterval); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandSamplingInterval(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandSamplingInterval(casted.Child)
	}
	return nil
}

func (m *SysexCommandSamplingInterval) GetTypeName() string {
	return "SysexCommandSamplingInterval"
}

func (m *SysexCommandSamplingInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandSamplingInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandSamplingInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandSamplingIntervalParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandSamplingInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSamplingInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSamplingInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSamplingInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSamplingInterval")
	}

	// Create a partially initialized instance
	_child := &SysexCommandSamplingInterval{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandSamplingInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSamplingInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSamplingInterval")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSamplingInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSamplingInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandSamplingInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
