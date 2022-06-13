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

// CBusOptions is the data-structure of this message
type CBusOptions struct {
	Connect bool
	Smart   bool
	Idmon   bool
	Exstat  bool
	Monitor bool
	Monall  bool
	Pun     bool
	Pcn     bool
}

// ICBusOptions is the corresponding interface of CBusOptions
type ICBusOptions interface {
	// GetConnect returns Connect (property field)
	GetConnect() bool
	// GetSmart returns Smart (property field)
	GetSmart() bool
	// GetIdmon returns Idmon (property field)
	GetIdmon() bool
	// GetExstat returns Exstat (property field)
	GetExstat() bool
	// GetMonitor returns Monitor (property field)
	GetMonitor() bool
	// GetMonall returns Monall (property field)
	GetMonall() bool
	// GetPun returns Pun (property field)
	GetPun() bool
	// GetPcn returns Pcn (property field)
	GetPcn() bool
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

func (m *CBusOptions) GetConnect() bool {
	return m.Connect
}

func (m *CBusOptions) GetSmart() bool {
	return m.Smart
}

func (m *CBusOptions) GetIdmon() bool {
	return m.Idmon
}

func (m *CBusOptions) GetExstat() bool {
	return m.Exstat
}

func (m *CBusOptions) GetMonitor() bool {
	return m.Monitor
}

func (m *CBusOptions) GetMonall() bool {
	return m.Monall
}

func (m *CBusOptions) GetPun() bool {
	return m.Pun
}

func (m *CBusOptions) GetPcn() bool {
	return m.Pcn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusOptions factory function for CBusOptions
func NewCBusOptions(connect bool, smart bool, idmon bool, exstat bool, monitor bool, monall bool, pun bool, pcn bool) *CBusOptions {
	return &CBusOptions{Connect: connect, Smart: smart, Idmon: idmon, Exstat: exstat, Monitor: monitor, Monall: monall, Pun: pun, Pcn: pcn}
}

func CastCBusOptions(structType interface{}) *CBusOptions {
	if casted, ok := structType.(CBusOptions); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusOptions); ok {
		return casted
	}
	return nil
}

func (m *CBusOptions) GetTypeName() string {
	return "CBusOptions"
}

func (m *CBusOptions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusOptions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (connect)
	lengthInBits += 1

	// Simple field (smart)
	lengthInBits += 1

	// Simple field (idmon)
	lengthInBits += 1

	// Simple field (exstat)
	lengthInBits += 1

	// Simple field (monitor)
	lengthInBits += 1

	// Simple field (monall)
	lengthInBits += 1

	// Simple field (pun)
	lengthInBits += 1

	// Simple field (pcn)
	lengthInBits += 1

	return lengthInBits
}

func (m *CBusOptions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusOptionsParse(readBuffer utils.ReadBuffer) (*CBusOptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (connect)
	_connect, _connectErr := readBuffer.ReadBit("connect")
	if _connectErr != nil {
		return nil, errors.Wrap(_connectErr, "Error parsing 'connect' field")
	}
	connect := _connect

	// Simple Field (smart)
	_smart, _smartErr := readBuffer.ReadBit("smart")
	if _smartErr != nil {
		return nil, errors.Wrap(_smartErr, "Error parsing 'smart' field")
	}
	smart := _smart

	// Simple Field (idmon)
	_idmon, _idmonErr := readBuffer.ReadBit("idmon")
	if _idmonErr != nil {
		return nil, errors.Wrap(_idmonErr, "Error parsing 'idmon' field")
	}
	idmon := _idmon

	// Simple Field (exstat)
	_exstat, _exstatErr := readBuffer.ReadBit("exstat")
	if _exstatErr != nil {
		return nil, errors.Wrap(_exstatErr, "Error parsing 'exstat' field")
	}
	exstat := _exstat

	// Simple Field (monitor)
	_monitor, _monitorErr := readBuffer.ReadBit("monitor")
	if _monitorErr != nil {
		return nil, errors.Wrap(_monitorErr, "Error parsing 'monitor' field")
	}
	monitor := _monitor

	// Simple Field (monall)
	_monall, _monallErr := readBuffer.ReadBit("monall")
	if _monallErr != nil {
		return nil, errors.Wrap(_monallErr, "Error parsing 'monall' field")
	}
	monall := _monall

	// Simple Field (pun)
	_pun, _punErr := readBuffer.ReadBit("pun")
	if _punErr != nil {
		return nil, errors.Wrap(_punErr, "Error parsing 'pun' field")
	}
	pun := _pun

	// Simple Field (pcn)
	_pcn, _pcnErr := readBuffer.ReadBit("pcn")
	if _pcnErr != nil {
		return nil, errors.Wrap(_pcnErr, "Error parsing 'pcn' field")
	}
	pcn := _pcn

	if closeErr := readBuffer.CloseContext("CBusOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusOptions")
	}

	// Create the instance
	return NewCBusOptions(connect, smart, idmon, exstat, monitor, monall, pun, pcn), nil
}

func (m *CBusOptions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusOptions")
	}

	// Simple Field (connect)
	connect := bool(m.Connect)
	_connectErr := writeBuffer.WriteBit("connect", (connect))
	if _connectErr != nil {
		return errors.Wrap(_connectErr, "Error serializing 'connect' field")
	}

	// Simple Field (smart)
	smart := bool(m.Smart)
	_smartErr := writeBuffer.WriteBit("smart", (smart))
	if _smartErr != nil {
		return errors.Wrap(_smartErr, "Error serializing 'smart' field")
	}

	// Simple Field (idmon)
	idmon := bool(m.Idmon)
	_idmonErr := writeBuffer.WriteBit("idmon", (idmon))
	if _idmonErr != nil {
		return errors.Wrap(_idmonErr, "Error serializing 'idmon' field")
	}

	// Simple Field (exstat)
	exstat := bool(m.Exstat)
	_exstatErr := writeBuffer.WriteBit("exstat", (exstat))
	if _exstatErr != nil {
		return errors.Wrap(_exstatErr, "Error serializing 'exstat' field")
	}

	// Simple Field (monitor)
	monitor := bool(m.Monitor)
	_monitorErr := writeBuffer.WriteBit("monitor", (monitor))
	if _monitorErr != nil {
		return errors.Wrap(_monitorErr, "Error serializing 'monitor' field")
	}

	// Simple Field (monall)
	monall := bool(m.Monall)
	_monallErr := writeBuffer.WriteBit("monall", (monall))
	if _monallErr != nil {
		return errors.Wrap(_monallErr, "Error serializing 'monall' field")
	}

	// Simple Field (pun)
	pun := bool(m.Pun)
	_punErr := writeBuffer.WriteBit("pun", (pun))
	if _punErr != nil {
		return errors.Wrap(_punErr, "Error serializing 'pun' field")
	}

	// Simple Field (pcn)
	pcn := bool(m.Pcn)
	_pcnErr := writeBuffer.WriteBit("pcn", (pcn))
	if _pcnErr != nil {
		return errors.Wrap(_pcnErr, "Error serializing 'pcn' field")
	}

	if popErr := writeBuffer.PopContext("CBusOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusOptions")
	}
	return nil
}

func (m *CBusOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
