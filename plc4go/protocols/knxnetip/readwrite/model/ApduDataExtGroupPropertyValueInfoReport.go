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

// ApduDataExtGroupPropertyValueInfoReport is the data-structure of this message
type ApduDataExtGroupPropertyValueInfoReport struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtGroupPropertyValueInfoReport is the corresponding interface of ApduDataExtGroupPropertyValueInfoReport
type IApduDataExtGroupPropertyValueInfoReport interface {
	IApduDataExt
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

func (m *ApduDataExtGroupPropertyValueInfoReport) GetExtApciType() uint8 {
	return 0x2B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtGroupPropertyValueInfoReport) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtGroupPropertyValueInfoReport factory function for ApduDataExtGroupPropertyValueInfoReport
func NewApduDataExtGroupPropertyValueInfoReport(length uint8) *ApduDataExtGroupPropertyValueInfoReport {
	_result := &ApduDataExtGroupPropertyValueInfoReport{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtGroupPropertyValueInfoReport(structType interface{}) *ApduDataExtGroupPropertyValueInfoReport {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueInfoReport); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtGroupPropertyValueInfoReport(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtGroupPropertyValueInfoReport(casted.Child)
	}
	return nil
}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueInfoReport"
}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtGroupPropertyValueInfoReport) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtGroupPropertyValueInfoReportParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtGroupPropertyValueInfoReport, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueInfoReport")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtGroupPropertyValueInfoReport{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtGroupPropertyValueInfoReport) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueInfoReport")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtGroupPropertyValueInfoReport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
