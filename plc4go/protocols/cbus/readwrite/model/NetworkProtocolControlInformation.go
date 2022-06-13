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
)

// Code generated by code-generation. DO NOT EDIT.

// NetworkProtocolControlInformation is the data-structure of this message
type NetworkProtocolControlInformation struct {
	StackCounter uint8
	StackDepth   uint8
}

// INetworkProtocolControlInformation is the corresponding interface of NetworkProtocolControlInformation
type INetworkProtocolControlInformation interface {
	// GetStackCounter returns StackCounter (property field)
	GetStackCounter() uint8
	// GetStackDepth returns StackDepth (property field)
	GetStackDepth() uint8
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

func (m *NetworkProtocolControlInformation) GetStackCounter() uint8 {
	return m.StackCounter
}

func (m *NetworkProtocolControlInformation) GetStackDepth() uint8 {
	return m.StackDepth
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkProtocolControlInformation factory function for NetworkProtocolControlInformation
func NewNetworkProtocolControlInformation(stackCounter uint8, stackDepth uint8) *NetworkProtocolControlInformation {
	return &NetworkProtocolControlInformation{StackCounter: stackCounter, StackDepth: stackDepth}
}

func CastNetworkProtocolControlInformation(structType interface{}) *NetworkProtocolControlInformation {
	if casted, ok := structType.(NetworkProtocolControlInformation); ok {
		return &casted
	}
	if casted, ok := structType.(*NetworkProtocolControlInformation); ok {
		return casted
	}
	return nil
}

func (m *NetworkProtocolControlInformation) GetTypeName() string {
	return "NetworkProtocolControlInformation"
}

func (m *NetworkProtocolControlInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NetworkProtocolControlInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (stackCounter)
	lengthInBits += 3

	// Simple field (stackDepth)
	lengthInBits += 3

	return lengthInBits
}

func (m *NetworkProtocolControlInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NetworkProtocolControlInformationParse(readBuffer utils.ReadBuffer) (*NetworkProtocolControlInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkProtocolControlInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkProtocolControlInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (stackCounter)
	_stackCounter, _stackCounterErr := readBuffer.ReadUint8("stackCounter", 3)
	if _stackCounterErr != nil {
		return nil, errors.Wrap(_stackCounterErr, "Error parsing 'stackCounter' field")
	}
	stackCounter := _stackCounter

	// Simple Field (stackDepth)
	_stackDepth, _stackDepthErr := readBuffer.ReadUint8("stackDepth", 3)
	if _stackDepthErr != nil {
		return nil, errors.Wrap(_stackDepthErr, "Error parsing 'stackDepth' field")
	}
	stackDepth := _stackDepth

	if closeErr := readBuffer.CloseContext("NetworkProtocolControlInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkProtocolControlInformation")
	}

	// Create the instance
	return NewNetworkProtocolControlInformation(stackCounter, stackDepth), nil
}

func (m *NetworkProtocolControlInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NetworkProtocolControlInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkProtocolControlInformation")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 2, uint8(0x0))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (stackCounter)
	stackCounter := uint8(m.StackCounter)
	_stackCounterErr := writeBuffer.WriteUint8("stackCounter", 3, (stackCounter))
	if _stackCounterErr != nil {
		return errors.Wrap(_stackCounterErr, "Error serializing 'stackCounter' field")
	}

	// Simple Field (stackDepth)
	stackDepth := uint8(m.StackDepth)
	_stackDepthErr := writeBuffer.WriteUint8("stackDepth", 3, (stackDepth))
	if _stackDepthErr != nil {
		return errors.Wrap(_stackDepthErr, "Error serializing 'stackDepth' field")
	}

	if popErr := writeBuffer.PopContext("NetworkProtocolControlInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkProtocolControlInformation")
	}
	return nil
}

func (m *NetworkProtocolControlInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
