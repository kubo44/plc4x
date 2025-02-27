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
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
import math


@dataclass
class ModbusPDUWriteSingleCoilRequest(PlcMessage, ModbusPDU):
    address: c_uint16
    value: c_uint16
    # Accessors for discriminator values.
    error_flag: c_bool = False
    function_flag: c_uint8 = 0x05
    response: c_bool = False

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusPDUWriteSingleCoilRequest")

        # Simple Field (address)
        write_simple_field(
            "address", self.address, write_unsigned_int(write_buffer, 16)
        )

        # Simple Field (value)
        write_simple_field("value", self.value, write_unsigned_int(write_buffer, 16))

        write_buffer.pop_context("ModbusPDUWriteSingleCoilRequest")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUWriteSingleCoilRequest = self

        # Simple field (address)
        length_in_bits += 16

        # Simple field (value)
        length_in_bits += 16

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: c_bool):
        read_buffer.pull_context("ModbusPDUWriteSingleCoilRequest")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        address: c_uint16 = read_simple_field(
            "address", read_unsigned_int(read_buffer, 16)
        )

        value: c_uint16 = read_simple_field("value", read_unsigned_int(read_buffer, 16))

        read_buffer.close_context("ModbusPDUWriteSingleCoilRequest")
        # Create the instance
        return ModbusPDUWriteSingleCoilRequestBuilder(address, value)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUWriteSingleCoilRequest):
            return False

        that: ModbusPDUWriteSingleCoilRequest = ModbusPDUWriteSingleCoilRequest(o)
        return (
            (self.address == that.address)
            and (self.value == that.value)
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
class ModbusPDUWriteSingleCoilRequestBuilder(ModbusPDUBuilder):
    address: c_uint16
    value: c_uint16

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUWriteSingleCoilRequest:
        modbus_pdu_write_single_coil_request: ModbusPDUWriteSingleCoilRequest = (
            ModbusPDUWriteSingleCoilRequest(self.address, self.value)
        )
        return modbus_pdu_write_single_coil_request
