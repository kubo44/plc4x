#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from ctypes import c_uint8
from enum import Enum


class SimulatedDataTypeSizes(Enum):
    BOOL: c_uint8 = (1, c_uint8(1))
    BYTE: c_uint8 = (2, c_uint8(1))
    WORD: c_uint8 = (3, c_uint8(2))
    DWORD: c_uint8 = (4, c_uint8(4))
    LWORD: c_uint8 = (5, c_uint8(8))
    SINT: c_uint8 = (6, c_uint8(1))
    INT: c_uint8 = (7, c_uint8(2))
    DINT: c_uint8 = (8, c_uint8(4))
    LINT: c_uint8 = (9, c_uint8(8))
    USINT: c_uint8 = (10, c_uint8(1))
    UINT: c_uint8 = (11, c_uint8(2))
    UDINT: c_uint8 = (12, c_uint8(4))
    ULINT: c_uint8 = (13, c_uint8(8))
    REAL: c_uint8 = (14, c_uint8(4))
    LREAL: c_uint8 = (15, c_uint8(8))
    TIME: c_uint8 = (16, c_uint8(8))
    LTIME: c_uint8 = (17, c_uint8(8))
    DATE: c_uint8 = (18, c_uint8(8))
    LDATE: c_uint8 = (19, c_uint8(8))
    TIME_OF_DAY: c_uint8 = (20, c_uint8(8))
    LTIME_OF_DAY: c_uint8 = (21, c_uint8(8))
    DATE_AND_TIME: c_uint8 = (22, c_uint8(8))
    LDATE_AND_TIME: c_uint8 = (23, c_uint8(8))
    CHAR: c_uint8 = (24, c_uint8(1))
    WCHAR: c_uint8 = (25, c_uint8(2))
    STRING: c_uint8 = (26, c_uint8(255))
    WSTRING: c_uint8 = (27, c_uint8(127))

    def __new__(cls, value, dataTypeSize):
        obj = object.__new__(cls)
        obj._value_ = value
        obj.dataTypeSize = dataTypeSize
        return obj
