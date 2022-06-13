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

// BACnetActionCommand is the data-structure of this message
type BACnetActionCommand struct {
	DeviceIdentifier   *BACnetContextTagObjectIdentifier
	ObjectIdentifier   *BACnetContextTagObjectIdentifier
	PropertyIdentifier *BACnetPropertyIdentifierTagged
	ArrayIndex         *BACnetContextTagUnsignedInteger
	PropertyValue      *BACnetConstructedData
	Priority           *BACnetContextTagUnsignedInteger
	PostDelay          *BACnetContextTagBoolean
	QuitOnFailure      *BACnetContextTagBoolean
	WriteSuccessful    *BACnetContextTagBoolean
}

// IBACnetActionCommand is the corresponding interface of BACnetActionCommand
type IBACnetActionCommand interface {
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() *BACnetContextTagUnsignedInteger
	// GetPostDelay returns PostDelay (property field)
	GetPostDelay() *BACnetContextTagBoolean
	// GetQuitOnFailure returns QuitOnFailure (property field)
	GetQuitOnFailure() *BACnetContextTagBoolean
	// GetWriteSuccessful returns WriteSuccessful (property field)
	GetWriteSuccessful() *BACnetContextTagBoolean
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

func (m *BACnetActionCommand) GetDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *BACnetActionCommand) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetActionCommand) GetPropertyIdentifier() *BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *BACnetActionCommand) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetActionCommand) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

func (m *BACnetActionCommand) GetPriority() *BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *BACnetActionCommand) GetPostDelay() *BACnetContextTagBoolean {
	return m.PostDelay
}

func (m *BACnetActionCommand) GetQuitOnFailure() *BACnetContextTagBoolean {
	return m.QuitOnFailure
}

func (m *BACnetActionCommand) GetWriteSuccessful() *BACnetContextTagBoolean {
	return m.WriteSuccessful
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetActionCommand factory function for BACnetActionCommand
func NewBACnetActionCommand(deviceIdentifier *BACnetContextTagObjectIdentifier, objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetPropertyIdentifierTagged, arrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedData, priority *BACnetContextTagUnsignedInteger, postDelay *BACnetContextTagBoolean, quitOnFailure *BACnetContextTagBoolean, writeSuccessful *BACnetContextTagBoolean) *BACnetActionCommand {
	return &BACnetActionCommand{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, PostDelay: postDelay, QuitOnFailure: quitOnFailure, WriteSuccessful: writeSuccessful}
}

func CastBACnetActionCommand(structType interface{}) *BACnetActionCommand {
	if casted, ok := structType.(BACnetActionCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetActionCommand); ok {
		return casted
	}
	return nil
}

func (m *BACnetActionCommand) GetTypeName() string {
	return "BACnetActionCommand"
}

func (m *BACnetActionCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetActionCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += (*m.DeviceIdentifier).GetLengthInBits()
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).GetLengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += (*m.Priority).GetLengthInBits()
	}

	// Optional Field (postDelay)
	if m.PostDelay != nil {
		lengthInBits += (*m.PostDelay).GetLengthInBits()
	}

	// Simple field (quitOnFailure)
	lengthInBits += m.QuitOnFailure.GetLengthInBits()

	// Simple field (writeSuccessful)
	lengthInBits += m.WriteSuccessful.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetActionCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetActionCommandParse(readBuffer utils.ReadBuffer) (*BACnetActionCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field")
		default:
			deviceIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
			}
		}
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetPropertyIdentifierTagged(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(4), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue(), CastBACnetTagPayloadUnsignedInteger(CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((*arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(5), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field")
		default:
			priority = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	// Optional Field (postDelay) (Can be skipped, if a given expression evaluates to false)
	var postDelay *BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("postDelay"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for postDelay")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(6), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'postDelay' field")
		default:
			postDelay = CastBACnetContextTagBoolean(_val)
			if closeErr := readBuffer.CloseContext("postDelay"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for postDelay")
			}
		}
	}

	// Simple Field (quitOnFailure)
	if pullErr := readBuffer.PullContext("quitOnFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for quitOnFailure")
	}
	_quitOnFailure, _quitOnFailureErr := BACnetContextTagParse(readBuffer, uint8(uint8(7)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _quitOnFailureErr != nil {
		return nil, errors.Wrap(_quitOnFailureErr, "Error parsing 'quitOnFailure' field")
	}
	quitOnFailure := CastBACnetContextTagBoolean(_quitOnFailure)
	if closeErr := readBuffer.CloseContext("quitOnFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for quitOnFailure")
	}

	// Simple Field (writeSuccessful)
	if pullErr := readBuffer.PullContext("writeSuccessful"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writeSuccessful")
	}
	_writeSuccessful, _writeSuccessfulErr := BACnetContextTagParse(readBuffer, uint8(uint8(8)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _writeSuccessfulErr != nil {
		return nil, errors.Wrap(_writeSuccessfulErr, "Error parsing 'writeSuccessful' field")
	}
	writeSuccessful := CastBACnetContextTagBoolean(_writeSuccessful)
	if closeErr := readBuffer.CloseContext("writeSuccessful"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writeSuccessful")
	}

	if closeErr := readBuffer.CloseContext("BACnetActionCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionCommand")
	}

	// Create the instance
	return NewBACnetActionCommand(deviceIdentifier, objectIdentifier, propertyIdentifier, arrayIndex, propertyValue, priority, postDelay, quitOnFailure, writeSuccessful), nil
}

func (m *BACnetActionCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetActionCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionCommand")
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	if m.DeviceIdentifier != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		deviceIdentifier = m.DeviceIdentifier
		_deviceIdentifierErr := writeBuffer.WriteSerializable(deviceIdentifier)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(m.ObjectIdentifier)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(m.PropertyIdentifier)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedData = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := writeBuffer.WriteSerializable(propertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority *BACnetContextTagUnsignedInteger = nil
	if m.Priority != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.Priority
		_priorityErr := writeBuffer.WriteSerializable(priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	// Optional Field (postDelay) (Can be skipped, if the value is null)
	var postDelay *BACnetContextTagBoolean = nil
	if m.PostDelay != nil {
		if pushErr := writeBuffer.PushContext("postDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for postDelay")
		}
		postDelay = m.PostDelay
		_postDelayErr := writeBuffer.WriteSerializable(postDelay)
		if popErr := writeBuffer.PopContext("postDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for postDelay")
		}
		if _postDelayErr != nil {
			return errors.Wrap(_postDelayErr, "Error serializing 'postDelay' field")
		}
	}

	// Simple Field (quitOnFailure)
	if pushErr := writeBuffer.PushContext("quitOnFailure"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for quitOnFailure")
	}
	_quitOnFailureErr := writeBuffer.WriteSerializable(m.QuitOnFailure)
	if popErr := writeBuffer.PopContext("quitOnFailure"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for quitOnFailure")
	}
	if _quitOnFailureErr != nil {
		return errors.Wrap(_quitOnFailureErr, "Error serializing 'quitOnFailure' field")
	}

	// Simple Field (writeSuccessful)
	if pushErr := writeBuffer.PushContext("writeSuccessful"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for writeSuccessful")
	}
	_writeSuccessfulErr := writeBuffer.WriteSerializable(m.WriteSuccessful)
	if popErr := writeBuffer.PopContext("writeSuccessful"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for writeSuccessful")
	}
	if _writeSuccessfulErr != nil {
		return errors.Wrap(_writeSuccessfulErr, "Error serializing 'writeSuccessful' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionCommand")
	}
	return nil
}

func (m *BACnetActionCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
