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

// NLMWhoIsRouterToNetwork is the data-structure of this message
type NLMWhoIsRouterToNetwork struct {
	*NLM
	DestinationNetworkAddress []uint16

	// Arguments.
	ApduLength uint16
}

// INLMWhoIsRouterToNetwork is the corresponding interface of NLMWhoIsRouterToNetwork
type INLMWhoIsRouterToNetwork interface {
	INLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() []uint16
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

func (m *NLMWhoIsRouterToNetwork) GetMessageType() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NLMWhoIsRouterToNetwork) InitializeParent(parent *NLM, vendorId *BACnetVendorId) {
	m.NLM.VendorId = vendorId
}

func (m *NLMWhoIsRouterToNetwork) GetParent() *NLM {
	return m.NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NLMWhoIsRouterToNetwork) GetDestinationNetworkAddress() []uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMWhoIsRouterToNetwork factory function for NLMWhoIsRouterToNetwork
func NewNLMWhoIsRouterToNetwork(destinationNetworkAddress []uint16, vendorId *BACnetVendorId, apduLength uint16) *NLMWhoIsRouterToNetwork {
	_result := &NLMWhoIsRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       NewNLM(vendorId, apduLength),
	}
	_result.Child = _result
	return _result
}

func CastNLMWhoIsRouterToNetwork(structType interface{}) *NLMWhoIsRouterToNetwork {
	if casted, ok := structType.(NLMWhoIsRouterToNetwork); ok {
		return &casted
	}
	if casted, ok := structType.(*NLMWhoIsRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(NLM); ok {
		return CastNLMWhoIsRouterToNetwork(casted.Child)
	}
	if casted, ok := structType.(*NLM); ok {
		return CastNLMWhoIsRouterToNetwork(casted.Child)
	}
	return nil
}

func (m *NLMWhoIsRouterToNetwork) GetTypeName() string {
	return "NLMWhoIsRouterToNetwork"
}

func (m *NLMWhoIsRouterToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMWhoIsRouterToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.DestinationNetworkAddress) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddress))
	}

	return lengthInBits
}

func (m *NLMWhoIsRouterToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMWhoIsRouterToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLMWhoIsRouterToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMWhoIsRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMWhoIsRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (destinationNetworkAddress)
	if pullErr := readBuffer.PullContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destinationNetworkAddress")
	}
	// Length array
	destinationNetworkAddress := make([]uint16, 0)
	{
		_destinationNetworkAddressLength := uint16(apduLength) - uint16(uint16(utils.InlineIf(bool(bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255))))), func() interface{} { return uint16(uint16(3)) }, func() interface{} { return uint16(uint16(1)) }).(uint16)))
		_destinationNetworkAddressEndPos := positionAware.GetPos() + uint16(_destinationNetworkAddressLength)
		for positionAware.GetPos() < _destinationNetworkAddressEndPos {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field")
			}
			destinationNetworkAddress = append(destinationNetworkAddress, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("destinationNetworkAddress", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destinationNetworkAddress")
	}

	if closeErr := readBuffer.CloseContext("NLMWhoIsRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMWhoIsRouterToNetwork")
	}

	// Create a partially initialized instance
	_child := &NLMWhoIsRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child, nil
}

func (m *NLMWhoIsRouterToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMWhoIsRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMWhoIsRouterToNetwork")
		}

		// Array Field (destinationNetworkAddress)
		if m.DestinationNetworkAddress != nil {
			if pushErr := writeBuffer.PushContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for destinationNetworkAddress")
			}
			for _, _element := range m.DestinationNetworkAddress {
				_elementErr := writeBuffer.WriteUint16("", 16, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'destinationNetworkAddress' field")
				}
			}
			if popErr := writeBuffer.PopContext("destinationNetworkAddress", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for destinationNetworkAddress")
			}
		}

		if popErr := writeBuffer.PopContext("NLMWhoIsRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMWhoIsRouterToNetwork")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMWhoIsRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
