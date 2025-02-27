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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_Block_AlarmCrReq extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.ALARM_CR_BLOCK_REQ;
  }

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final PnIoCm_AlarmCrType alarmType;
  protected final int lt;
  protected final boolean transport;
  protected final boolean priority;
  protected final int rtaTimeoutFactor;
  protected final int rtaRetries;
  protected final int localAlarmReference;
  protected final int maxAlarmDataLength;
  protected final int alarmCtrTagHeaderHigh;
  protected final int alarmCtrTagHeaderLow;

  // Reserved Fields
  private Long reservedField0;

  public PnIoCm_Block_AlarmCrReq(
      short blockVersionHigh,
      short blockVersionLow,
      PnIoCm_AlarmCrType alarmType,
      int lt,
      boolean transport,
      boolean priority,
      int rtaTimeoutFactor,
      int rtaRetries,
      int localAlarmReference,
      int maxAlarmDataLength,
      int alarmCtrTagHeaderHigh,
      int alarmCtrTagHeaderLow) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.alarmType = alarmType;
    this.lt = lt;
    this.transport = transport;
    this.priority = priority;
    this.rtaTimeoutFactor = rtaTimeoutFactor;
    this.rtaRetries = rtaRetries;
    this.localAlarmReference = localAlarmReference;
    this.maxAlarmDataLength = maxAlarmDataLength;
    this.alarmCtrTagHeaderHigh = alarmCtrTagHeaderHigh;
    this.alarmCtrTagHeaderLow = alarmCtrTagHeaderLow;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public PnIoCm_AlarmCrType getAlarmType() {
    return alarmType;
  }

  public int getLt() {
    return lt;
  }

  public boolean getTransport() {
    return transport;
  }

  public boolean getPriority() {
    return priority;
  }

  public int getRtaTimeoutFactor() {
    return rtaTimeoutFactor;
  }

  public int getRtaRetries() {
    return rtaRetries;
  }

  public int getLocalAlarmReference() {
    return localAlarmReference;
  }

  public int getMaxAlarmDataLength() {
    return maxAlarmDataLength;
  }

  public int getAlarmCtrTagHeaderHigh() {
    return alarmCtrTagHeaderHigh;
  }

  public int getAlarmCtrTagHeaderLow() {
    return alarmCtrTagHeaderLow;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Block_AlarmCrReq");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (alarmType)
    writeSimpleEnumField(
        "alarmType",
        "PnIoCm_AlarmCrType",
        alarmType,
        new DataWriterEnumDefault<>(
            PnIoCm_AlarmCrType::getValue,
            PnIoCm_AlarmCrType::name,
            writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (lt)
    writeSimpleField(
        "lt",
        lt,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (long) 0x00000000,
        writeUnsignedLong(writeBuffer, 30),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (transport)
    writeSimpleField(
        "transport",
        transport,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (priority)
    writeSimpleField(
        "priority",
        priority,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (rtaTimeoutFactor)
    writeSimpleField(
        "rtaTimeoutFactor",
        rtaTimeoutFactor,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (rtaRetries)
    writeSimpleField(
        "rtaRetries",
        rtaRetries,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (localAlarmReference)
    writeSimpleField(
        "localAlarmReference",
        localAlarmReference,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (maxAlarmDataLength)
    writeSimpleField(
        "maxAlarmDataLength",
        maxAlarmDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (alarmCtrTagHeaderHigh)
    writeSimpleField(
        "alarmCtrTagHeaderHigh",
        alarmCtrTagHeaderHigh,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (alarmCtrTagHeaderLow)
    writeSimpleField(
        "alarmCtrTagHeaderLow",
        alarmCtrTagHeaderLow,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Block_AlarmCrReq");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_AlarmCrReq _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (alarmType)
    lengthInBits += 16;

    // Simple field (lt)
    lengthInBits += 16;

    // Reserved Field (reserved)
    lengthInBits += 30;

    // Simple field (transport)
    lengthInBits += 1;

    // Simple field (priority)
    lengthInBits += 1;

    // Simple field (rtaTimeoutFactor)
    lengthInBits += 16;

    // Simple field (rtaRetries)
    lengthInBits += 16;

    // Simple field (localAlarmReference)
    lengthInBits += 16;

    // Simple field (maxAlarmDataLength)
    lengthInBits += 16;

    // Simple field (alarmCtrTagHeaderHigh)
    lengthInBits += 16;

    // Simple field (alarmCtrTagHeaderLow)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_AlarmCrReq");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    PnIoCm_AlarmCrType alarmType =
        readEnumField(
            "alarmType",
            "PnIoCm_AlarmCrType",
            new DataReaderEnumDefault<>(
                PnIoCm_AlarmCrType::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int lt =
        readSimpleField(
            "lt", readUnsignedInt(readBuffer, 16), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Long reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedLong(readBuffer, 30),
            (long) 0x00000000,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean transport =
        readSimpleField(
            "transport", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean priority =
        readSimpleField(
            "priority", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int rtaTimeoutFactor =
        readSimpleField(
            "rtaTimeoutFactor",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int rtaRetries =
        readSimpleField(
            "rtaRetries",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int localAlarmReference =
        readSimpleField(
            "localAlarmReference",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int maxAlarmDataLength =
        readSimpleField(
            "maxAlarmDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int alarmCtrTagHeaderHigh =
        readSimpleField(
            "alarmCtrTagHeaderHigh",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int alarmCtrTagHeaderLow =
        readSimpleField(
            "alarmCtrTagHeaderLow",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_AlarmCrReq");
    // Create the instance
    return new PnIoCm_Block_AlarmCrReqBuilderImpl(
        blockVersionHigh,
        blockVersionLow,
        alarmType,
        lt,
        transport,
        priority,
        rtaTimeoutFactor,
        rtaRetries,
        localAlarmReference,
        maxAlarmDataLength,
        alarmCtrTagHeaderHigh,
        alarmCtrTagHeaderLow,
        reservedField0);
  }

  public static class PnIoCm_Block_AlarmCrReqBuilderImpl
      implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final PnIoCm_AlarmCrType alarmType;
    private final int lt;
    private final boolean transport;
    private final boolean priority;
    private final int rtaTimeoutFactor;
    private final int rtaRetries;
    private final int localAlarmReference;
    private final int maxAlarmDataLength;
    private final int alarmCtrTagHeaderHigh;
    private final int alarmCtrTagHeaderLow;
    private final Long reservedField0;

    public PnIoCm_Block_AlarmCrReqBuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        PnIoCm_AlarmCrType alarmType,
        int lt,
        boolean transport,
        boolean priority,
        int rtaTimeoutFactor,
        int rtaRetries,
        int localAlarmReference,
        int maxAlarmDataLength,
        int alarmCtrTagHeaderHigh,
        int alarmCtrTagHeaderLow,
        Long reservedField0) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.alarmType = alarmType;
      this.lt = lt;
      this.transport = transport;
      this.priority = priority;
      this.rtaTimeoutFactor = rtaTimeoutFactor;
      this.rtaRetries = rtaRetries;
      this.localAlarmReference = localAlarmReference;
      this.maxAlarmDataLength = maxAlarmDataLength;
      this.alarmCtrTagHeaderHigh = alarmCtrTagHeaderHigh;
      this.alarmCtrTagHeaderLow = alarmCtrTagHeaderLow;
      this.reservedField0 = reservedField0;
    }

    public PnIoCm_Block_AlarmCrReq build() {
      PnIoCm_Block_AlarmCrReq pnIoCm_Block_AlarmCrReq =
          new PnIoCm_Block_AlarmCrReq(
              blockVersionHigh,
              blockVersionLow,
              alarmType,
              lt,
              transport,
              priority,
              rtaTimeoutFactor,
              rtaRetries,
              localAlarmReference,
              maxAlarmDataLength,
              alarmCtrTagHeaderHigh,
              alarmCtrTagHeaderLow);
      pnIoCm_Block_AlarmCrReq.reservedField0 = reservedField0;
      return pnIoCm_Block_AlarmCrReq;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_AlarmCrReq)) {
      return false;
    }
    PnIoCm_Block_AlarmCrReq that = (PnIoCm_Block_AlarmCrReq) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getAlarmType() == that.getAlarmType())
        && (getLt() == that.getLt())
        && (getTransport() == that.getTransport())
        && (getPriority() == that.getPriority())
        && (getRtaTimeoutFactor() == that.getRtaTimeoutFactor())
        && (getRtaRetries() == that.getRtaRetries())
        && (getLocalAlarmReference() == that.getLocalAlarmReference())
        && (getMaxAlarmDataLength() == that.getMaxAlarmDataLength())
        && (getAlarmCtrTagHeaderHigh() == that.getAlarmCtrTagHeaderHigh())
        && (getAlarmCtrTagHeaderLow() == that.getAlarmCtrTagHeaderLow())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getAlarmType(),
        getLt(),
        getTransport(),
        getPriority(),
        getRtaTimeoutFactor(),
        getRtaRetries(),
        getLocalAlarmReference(),
        getMaxAlarmDataLength(),
        getAlarmCtrTagHeaderHigh(),
        getAlarmCtrTagHeaderLow());
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
