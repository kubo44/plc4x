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

from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
import math


@dataclass
class ModbusPDUReadFileRecordRequestItem(PlcMessage):
    reference_type: c_uint8
    file_number: c_uint16
    record_number: c_uint16
    record_length: c_uint16

    def __post_init__(self):
        super().__init__()

    def serialize(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusPDUReadFileRecordRequestItem")

        # Simple Field (referenceType)
        write_simple_field(
            "referenceType", self.reference_type, write_unsigned_short(write_buffer, 8)
        )

        # Simple Field (fileNumber)
        write_simple_field(
            "fileNumber", self.file_number, write_unsigned_int(write_buffer, 16)
        )

        # Simple Field (recordNumber)
        write_simple_field(
            "recordNumber", self.record_number, write_unsigned_int(write_buffer, 16)
        )

        # Simple Field (recordLength)
        write_simple_field(
            "recordLength", self.record_length, write_unsigned_int(write_buffer, 16)
        )

        write_buffer.pop_context("ModbusPDUReadFileRecordRequestItem")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = 0
        _value: ModbusPDUReadFileRecordRequestItem = self

        # Simple field (referenceType)
        length_in_bits += 8

        # Simple field (fileNumber)
        length_in_bits += 16

        # Simple field (recordNumber)
        length_in_bits += 16

        # Simple field (recordLength)
        length_in_bits += 16

        return length_in_bits

    def static_parse(read_buffer: ReadBuffer, args):
        position_aware: PositionAware = read_buffer
        return staticParse(read_buffer)

    @staticmethod
    def static_parse_context(read_buffer: ReadBuffer):
        read_buffer.pull_context("ModbusPDUReadFileRecordRequestItem")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        reference_type: c_uint8 = read_simple_field(
            "referenceType", read_unsigned_short(read_buffer, 8)
        )

        file_number: c_uint16 = read_simple_field(
            "fileNumber", read_unsigned_int(read_buffer, 16)
        )

        record_number: c_uint16 = read_simple_field(
            "recordNumber", read_unsigned_int(read_buffer, 16)
        )

        record_length: c_uint16 = read_simple_field(
            "recordLength", read_unsigned_int(read_buffer, 16)
        )

        read_buffer.close_context("ModbusPDUReadFileRecordRequestItem")
        # Create the instance
        _modbus_pdu_read_file_record_request_item: ModbusPDUReadFileRecordRequestItem = ModbusPDUReadFileRecordRequestItem(
            reference_type, file_number, record_number, record_length
        )
        return _modbus_pdu_read_file_record_request_item

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadFileRecordRequestItem):
            return False

        that: ModbusPDUReadFileRecordRequestItem = ModbusPDUReadFileRecordRequestItem(o)
        return (
            (self.reference_type == that.reference_type)
            and (self.file_number == that.file_number)
            and (self.record_number == that.record_number)
            and (self.record_length == that.record_length)
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
