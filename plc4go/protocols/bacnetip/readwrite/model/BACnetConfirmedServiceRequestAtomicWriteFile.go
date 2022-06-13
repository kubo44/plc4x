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

// BACnetConfirmedServiceRequestAtomicWriteFile is the data-structure of this message
type BACnetConfirmedServiceRequestAtomicWriteFile struct {
	*BACnetConfirmedServiceRequest
	DeviceIdentifier  *BACnetApplicationTagObjectIdentifier
	OpeningTag        *BACnetOpeningTag
	FileStartPosition *BACnetApplicationTagSignedInteger
	FileData          *BACnetApplicationTagOctetString
	ClosingTag        *BACnetClosingTag

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestAtomicWriteFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicWriteFile
type IBACnetConfirmedServiceRequestAtomicWriteFile interface {
	IBACnetConfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() *BACnetApplicationTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() *BACnetApplicationTagSignedInteger
	// GetFileData returns FileData (property field)
	GetFileData() *BACnetApplicationTagOctetString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
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

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetDeviceIdentifier() *BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetFileStartPosition() *BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetFileData() *BACnetApplicationTagOctetString {
	return m.FileData
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicWriteFile factory function for BACnetConfirmedServiceRequestAtomicWriteFile
func NewBACnetConfirmedServiceRequestAtomicWriteFile(deviceIdentifier *BACnetApplicationTagObjectIdentifier, openingTag *BACnetOpeningTag, fileStartPosition *BACnetApplicationTagSignedInteger, fileData *BACnetApplicationTagOctetString, closingTag *BACnetClosingTag, serviceRequestLength uint16) *BACnetConfirmedServiceRequestAtomicWriteFile {
	_result := &BACnetConfirmedServiceRequestAtomicWriteFile{
		DeviceIdentifier:              deviceIdentifier,
		OpeningTag:                    openingTag,
		FileStartPosition:             fileStartPosition,
		FileData:                      fileData,
		ClosingTag:                    closingTag,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestAtomicWriteFile(structType interface{}) *BACnetConfirmedServiceRequestAtomicWriteFile {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicWriteFile); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestAtomicWriteFile(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestAtomicWriteFile(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicWriteFile"
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits()

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += (*m.OpeningTag).GetLengthInBits()
	}

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits()

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits()

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += (*m.ClosingTag).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicWriteFileParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestAtomicWriteFile, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field")
	}
	deviceIdentifier := CastBACnetApplicationTagObjectIdentifier(_deviceIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
	}

	// Optional Field (openingTag) (Can be skipped, if a given expression evaluates to false)
	var openingTag *BACnetOpeningTag = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
		}
		_val, _err := BACnetOpeningTagParse(readBuffer, uint8(0))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'openingTag' field")
		default:
			openingTag = CastBACnetOpeningTag(_val)
			if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for openingTag")
			}
		}
	}

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartPosition")
	}
	_fileStartPosition, _fileStartPositionErr := BACnetApplicationTagParse(readBuffer)
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field")
	}
	fileStartPosition := CastBACnetApplicationTagSignedInteger(_fileStartPosition)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartPosition")
	}

	// Simple Field (fileData)
	if pullErr := readBuffer.PullContext("fileData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileData")
	}
	_fileData, _fileDataErr := BACnetApplicationTagParse(readBuffer)
	if _fileDataErr != nil {
		return nil, errors.Wrap(_fileDataErr, "Error parsing 'fileData' field")
	}
	fileData := CastBACnetApplicationTagOctetString(_fileData)
	if closeErr := readBuffer.CloseContext("fileData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileData")
	}

	// Optional Field (closingTag) (Can be skipped, if a given expression evaluates to false)
	var closingTag *BACnetClosingTag = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
		}
		_val, _err := BACnetClosingTagParse(readBuffer, uint8(0))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'closingTag' field")
		default:
			closingTag = CastBACnetClosingTag(_val)
			if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for closingTag")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicWriteFile")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestAtomicWriteFile{
		DeviceIdentifier:              CastBACnetApplicationTagObjectIdentifier(deviceIdentifier),
		OpeningTag:                    CastBACnetOpeningTag(openingTag),
		FileStartPosition:             CastBACnetApplicationTagSignedInteger(fileStartPosition),
		FileData:                      CastBACnetApplicationTagOctetString(fileData),
		ClosingTag:                    CastBACnetClosingTag(closingTag),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicWriteFile")
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		_deviceIdentifierErr := writeBuffer.WriteSerializable(m.DeviceIdentifier)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Optional Field (openingTag) (Can be skipped, if the value is null)
		var openingTag *BACnetOpeningTag = nil
		if m.OpeningTag != nil {
			if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for openingTag")
			}
			openingTag = m.OpeningTag
			_openingTagErr := writeBuffer.WriteSerializable(openingTag)
			if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for openingTag")
			}
			if _openingTagErr != nil {
				return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
			}
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileStartPosition")
		}
		_fileStartPositionErr := writeBuffer.WriteSerializable(m.FileStartPosition)
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileStartPosition")
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		// Simple Field (fileData)
		if pushErr := writeBuffer.PushContext("fileData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileData")
		}
		_fileDataErr := writeBuffer.WriteSerializable(m.FileData)
		if popErr := writeBuffer.PopContext("fileData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileData")
		}
		if _fileDataErr != nil {
			return errors.Wrap(_fileDataErr, "Error serializing 'fileData' field")
		}

		// Optional Field (closingTag) (Can be skipped, if the value is null)
		var closingTag *BACnetClosingTag = nil
		if m.ClosingTag != nil {
			if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for closingTag")
			}
			closingTag = m.ClosingTag
			_closingTagErr := writeBuffer.WriteSerializable(closingTag)
			if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for closingTag")
			}
			if _closingTagErr != nil {
				return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicWriteFile")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
