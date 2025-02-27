#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusDeviceInformationLevel import (
    ModbusDeviceInformationLevel,
)
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
import math


@dataclass
class ModbusPDUReadDeviceIdentificationRequest(PlcMessage, ModbusPDU):
    level: ModbusDeviceInformationLevel
    object_id: c_uint8
    MEITYPE: c_uint8 = 0x0E
    # Accessors for discriminator values.
    error_flag: c_bool = False
    function_flag: c_uint8 = 0x2B
    response: c_bool = False

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusPDUReadDeviceIdentificationRequest")

        # Const Field (meiType)
        write_const_field(
            "meiType", self.mei_type, write_unsigned_short(write_buffer, 8)
        )

        # Simple Field (level)
        write_simple_enum_field(
            "level",
            "ModbusDeviceInformationLevel",
            self.level,
            DataWriterEnumDefault(
                ModbusDeviceInformationLevel.value,
                ModbusDeviceInformationLevel.name,
                write_unsigned_short(write_buffer, 8),
            ),
        )

        # Simple Field (objectId)
        write_simple_field(
            "objectId", self.object_id, write_unsigned_short(write_buffer, 8)
        )

        write_buffer.pop_context("ModbusPDUReadDeviceIdentificationRequest")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUReadDeviceIdentificationRequest = self

        # Const Field (meiType)
        length_in_bits += 8

        # Simple field (level)
        length_in_bits += 8

        # Simple field (objectId)
        length_in_bits += 8

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: c_bool):
        read_buffer.pull_context("ModbusPDUReadDeviceIdentificationRequest")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        mei_type: c_uint8 = read_const_field(
            "meiType",
            read_unsigned_short(read_buffer, 8),
            ModbusPDUReadDeviceIdentificationRequest.MEITYPE,
        )

        level: ModbusDeviceInformationLevel = read_enum_field(
            "level",
            "ModbusDeviceInformationLevel",
            DataReaderEnumDefault(
                ModbusDeviceInformationLevel.enumForValue,
                read_unsigned_short(read_buffer, 8),
            ),
        )

        object_id: c_uint8 = read_simple_field(
            "objectId", read_unsigned_short(read_buffer, 8)
        )

        read_buffer.close_context("ModbusPDUReadDeviceIdentificationRequest")
        # Create the instance
        return ModbusPDUReadDeviceIdentificationRequestBuilder(level, object_id)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadDeviceIdentificationRequest):
            return False

        that: ModbusPDUReadDeviceIdentificationRequest = (
            ModbusPDUReadDeviceIdentificationRequest(o)
        )
        return (
            (self.level == that.level)
            and (self.object_id == that.object_id)
            and super().equals(that)
            and True
        )

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUReadDeviceIdentificationRequestBuilder(ModbusPDUBuilder):
    level: ModbusDeviceInformationLevel
    objectId: c_uint8

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUReadDeviceIdentificationRequest:
        modbus_pdu_read_device_identification_request: ModbusPDUReadDeviceIdentificationRequest = ModbusPDUReadDeviceIdentificationRequest(
            self.level, self.object_id
        )
        return modbus_pdu_read_device_identification_request
