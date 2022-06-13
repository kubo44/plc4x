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

// KnxNetRemoteConfigurationAndDiagnosis is the data-structure of this message
type KnxNetRemoteConfigurationAndDiagnosis struct {
	*ServiceId
	Version uint8
}

// IKnxNetRemoteConfigurationAndDiagnosis is the corresponding interface of KnxNetRemoteConfigurationAndDiagnosis
type IKnxNetRemoteConfigurationAndDiagnosis interface {
	IServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
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

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetServiceType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *KnxNetRemoteConfigurationAndDiagnosis) InitializeParent(parent *ServiceId) {}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetParent() *ServiceId {
	return m.ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetRemoteConfigurationAndDiagnosis factory function for KnxNetRemoteConfigurationAndDiagnosis
func NewKnxNetRemoteConfigurationAndDiagnosis(version uint8) *KnxNetRemoteConfigurationAndDiagnosis {
	_result := &KnxNetRemoteConfigurationAndDiagnosis{
		Version:   version,
		ServiceId: NewServiceId(),
	}
	_result.Child = _result
	return _result
}

func CastKnxNetRemoteConfigurationAndDiagnosis(structType interface{}) *KnxNetRemoteConfigurationAndDiagnosis {
	if casted, ok := structType.(KnxNetRemoteConfigurationAndDiagnosis); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxNetRemoteConfigurationAndDiagnosis); ok {
		return casted
	}
	if casted, ok := structType.(ServiceId); ok {
		return CastKnxNetRemoteConfigurationAndDiagnosis(casted.Child)
	}
	if casted, ok := structType.(*ServiceId); ok {
		return CastKnxNetRemoteConfigurationAndDiagnosis(casted.Child)
	}
	return nil
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetTypeName() string {
	return "KnxNetRemoteConfigurationAndDiagnosis"
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetRemoteConfigurationAndDiagnosisParse(readBuffer utils.ReadBuffer) (*KnxNetRemoteConfigurationAndDiagnosis, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetRemoteConfigurationAndDiagnosis"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetRemoteConfigurationAndDiagnosis")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteConfigurationAndDiagnosis"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetRemoteConfigurationAndDiagnosis")
	}

	// Create a partially initialized instance
	_child := &KnxNetRemoteConfigurationAndDiagnosis{
		Version:   version,
		ServiceId: &ServiceId{},
	}
	_child.ServiceId.Child = _child
	return _child, nil
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteConfigurationAndDiagnosis"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetRemoteConfigurationAndDiagnosis")
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteConfigurationAndDiagnosis"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetRemoteConfigurationAndDiagnosis")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
