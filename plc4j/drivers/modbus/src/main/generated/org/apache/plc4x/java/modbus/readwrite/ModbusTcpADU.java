/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.modbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ModbusTcpADU extends ModbusADU implements Message {

  // Accessors for discriminator values.
  public DriverType getDriverType() {
    return DriverType.MODBUS_TCP;
  }

  // Constant values.
  public static final Integer PROTOCOLIDENTIFIER = 0x0000;

  // Properties.
  protected final int transactionIdentifier;
  protected final short unitIdentifier;
  protected final ModbusPDU pdu;

  public ModbusTcpADU(int transactionIdentifier, short unitIdentifier, ModbusPDU pdu) {
    super();
    this.transactionIdentifier = transactionIdentifier;
    this.unitIdentifier = unitIdentifier;
    this.pdu = pdu;
  }

  public int getTransactionIdentifier() {
    return transactionIdentifier;
  }

  public short getUnitIdentifier() {
    return unitIdentifier;
  }

  public ModbusPDU getPdu() {
    return pdu;
  }

  public int getProtocolIdentifier() {
    return PROTOCOLIDENTIFIER;
  }

  @Override
  protected void serializeModbusADUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModbusTcpADU");

    // Simple Field (transactionIdentifier)
    writeSimpleField(
        "transactionIdentifier",
        transactionIdentifier,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (protocolIdentifier)
    writeConstField(
        "protocolIdentifier",
        PROTOCOLIDENTIFIER,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int length = (int) ((getPdu().getLengthInBytes()) + (1));
    writeImplicitField(
        "length",
        length,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (unitIdentifier)
    writeSimpleField(
        "unitIdentifier",
        unitIdentifier,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (pdu)
    writeSimpleField(
        "pdu",
        pdu,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("ModbusTcpADU");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusTcpADU _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (transactionIdentifier)
    lengthInBits += 16;

    // Const Field (protocolIdentifier)
    lengthInBits += 16;

    // Implicit Field (length)
    lengthInBits += 16;

    // Simple field (unitIdentifier)
    lengthInBits += 8;

    // Simple field (pdu)
    lengthInBits += pdu.getLengthInBits();

    return lengthInBits;
  }

  public static ModbusADUBuilder staticParseModbusADUBuilder(
      ReadBuffer readBuffer, DriverType driverType, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusTcpADU");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int transactionIdentifier =
        readSimpleField(
            "transactionIdentifier",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int protocolIdentifier =
        readConstField(
            "protocolIdentifier",
            readUnsignedInt(readBuffer, 16),
            ModbusTcpADU.PROTOCOLIDENTIFIER,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int length =
        readImplicitField(
            "length",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short unitIdentifier =
        readSimpleField(
            "unitIdentifier",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    ModbusPDU pdu =
        readSimpleField(
            "pdu",
            new DataReaderComplexDefault<>(
                () -> ModbusPDU.staticParse(readBuffer, (boolean) (response)), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("ModbusTcpADU");
    // Create the instance
    return new ModbusTcpADUBuilderImpl(transactionIdentifier, unitIdentifier, pdu);
  }

  public static class ModbusTcpADUBuilderImpl implements ModbusADU.ModbusADUBuilder {
    private final int transactionIdentifier;
    private final short unitIdentifier;
    private final ModbusPDU pdu;

    public ModbusTcpADUBuilderImpl(int transactionIdentifier, short unitIdentifier, ModbusPDU pdu) {
      this.transactionIdentifier = transactionIdentifier;
      this.unitIdentifier = unitIdentifier;
      this.pdu = pdu;
    }

    public ModbusTcpADU build() {
      ModbusTcpADU modbusTcpADU = new ModbusTcpADU(transactionIdentifier, unitIdentifier, pdu);
      return modbusTcpADU;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusTcpADU)) {
      return false;
    }
    ModbusTcpADU that = (ModbusTcpADU) o;
    return (getTransactionIdentifier() == that.getTransactionIdentifier())
        && (getUnitIdentifier() == that.getUnitIdentifier())
        && (getPdu() == that.getPdu())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getTransactionIdentifier(), getUnitIdentifier(), getPdu());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
