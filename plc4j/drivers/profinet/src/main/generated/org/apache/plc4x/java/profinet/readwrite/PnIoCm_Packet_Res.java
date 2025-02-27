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

public class PnIoCm_Packet_Res extends PnIoCm_Packet implements Message {

  // Accessors for discriminator values.
  public DceRpc_PacketType getPacketType() {
    return DceRpc_PacketType.RESPONSE;
  }

  // Properties.
  protected final short errorCode2;
  protected final short errorCode1;
  protected final short errorDecode;
  protected final short errorCode;
  protected final long arrayMaximumCount;
  protected final long arrayOffset;
  protected final List<PnIoCm_Block> blocks;

  public PnIoCm_Packet_Res(
      short errorCode2,
      short errorCode1,
      short errorDecode,
      short errorCode,
      long arrayMaximumCount,
      long arrayOffset,
      List<PnIoCm_Block> blocks) {
    super();
    this.errorCode2 = errorCode2;
    this.errorCode1 = errorCode1;
    this.errorDecode = errorDecode;
    this.errorCode = errorCode;
    this.arrayMaximumCount = arrayMaximumCount;
    this.arrayOffset = arrayOffset;
    this.blocks = blocks;
  }

  public short getErrorCode2() {
    return errorCode2;
  }

  public short getErrorCode1() {
    return errorCode1;
  }

  public short getErrorDecode() {
    return errorDecode;
  }

  public short getErrorCode() {
    return errorCode;
  }

  public long getArrayMaximumCount() {
    return arrayMaximumCount;
  }

  public long getArrayOffset() {
    return arrayOffset;
  }

  public List<PnIoCm_Block> getBlocks() {
    return blocks;
  }

  @Override
  protected void serializePnIoCm_PacketChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Packet_Res");

    // Simple Field (errorCode2)
    writeSimpleField("errorCode2", errorCode2, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (errorCode1)
    writeSimpleField("errorCode1", errorCode1, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (errorDecode)
    writeSimpleField("errorDecode", errorDecode, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (errorCode)
    writeSimpleField("errorCode", errorCode, writeUnsignedShort(writeBuffer, 8));

    // Implicit Field (argsLength) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    long argsLength =
        (long)
            (((((((((getLengthInBytes()) - (1L)) - (1L)) - (1L)) - (1L)) - (4L)) - (4L)) - (4L))
                - (4L));
    writeImplicitField("argsLength", argsLength, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (arrayMaximumCount)
    writeSimpleField("arrayMaximumCount", arrayMaximumCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (arrayOffset)
    writeSimpleField("arrayOffset", arrayOffset, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (arrayActualCount) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    long arrayActualCount =
        (long)
            (((((((((getLengthInBytes()) - (1L)) - (1L)) - (1L)) - (1L)) - (4L)) - (4L)) - (4L))
                - (4L));
    writeImplicitField("arrayActualCount", arrayActualCount, writeUnsignedLong(writeBuffer, 32));

    // Array Field (blocks)
    writeComplexTypeArrayField("blocks", blocks, writeBuffer);

    writeBuffer.popContext("PnIoCm_Packet_Res");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Packet_Res _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (errorCode2)
    lengthInBits += 8;

    // Simple field (errorCode1)
    lengthInBits += 8;

    // Simple field (errorDecode)
    lengthInBits += 8;

    // Simple field (errorCode)
    lengthInBits += 8;

    // Implicit Field (argsLength)
    lengthInBits += 32;

    // Simple field (arrayMaximumCount)
    lengthInBits += 32;

    // Simple field (arrayOffset)
    lengthInBits += 32;

    // Implicit Field (arrayActualCount)
    lengthInBits += 32;

    // Array field
    if (blocks != null) {
      for (Message element : blocks) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static PnIoCm_PacketBuilder staticParsePnIoCm_PacketBuilder(
      ReadBuffer readBuffer, DceRpc_PacketType packetType) throws ParseException {
    readBuffer.pullContext("PnIoCm_Packet_Res");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short errorCode2 = readSimpleField("errorCode2", readUnsignedShort(readBuffer, 8));

    short errorCode1 = readSimpleField("errorCode1", readUnsignedShort(readBuffer, 8));

    short errorDecode = readSimpleField("errorDecode", readUnsignedShort(readBuffer, 8));

    short errorCode = readSimpleField("errorCode", readUnsignedShort(readBuffer, 8));

    long argsLength = readImplicitField("argsLength", readUnsignedLong(readBuffer, 32));

    long arrayMaximumCount = readSimpleField("arrayMaximumCount", readUnsignedLong(readBuffer, 32));

    long arrayOffset = readSimpleField("arrayOffset", readUnsignedLong(readBuffer, 32));

    long arrayActualCount = readImplicitField("arrayActualCount", readUnsignedLong(readBuffer, 32));

    List<PnIoCm_Block> blocks =
        readLengthArrayField(
            "blocks",
            new DataReaderComplexDefault<>(() -> PnIoCm_Block.staticParse(readBuffer), readBuffer),
            argsLength);

    readBuffer.closeContext("PnIoCm_Packet_Res");
    // Create the instance
    return new PnIoCm_Packet_ResBuilderImpl(
        errorCode2, errorCode1, errorDecode, errorCode, arrayMaximumCount, arrayOffset, blocks);
  }

  public static class PnIoCm_Packet_ResBuilderImpl implements PnIoCm_Packet.PnIoCm_PacketBuilder {
    private final short errorCode2;
    private final short errorCode1;
    private final short errorDecode;
    private final short errorCode;
    private final long arrayMaximumCount;
    private final long arrayOffset;
    private final List<PnIoCm_Block> blocks;

    public PnIoCm_Packet_ResBuilderImpl(
        short errorCode2,
        short errorCode1,
        short errorDecode,
        short errorCode,
        long arrayMaximumCount,
        long arrayOffset,
        List<PnIoCm_Block> blocks) {
      this.errorCode2 = errorCode2;
      this.errorCode1 = errorCode1;
      this.errorDecode = errorDecode;
      this.errorCode = errorCode;
      this.arrayMaximumCount = arrayMaximumCount;
      this.arrayOffset = arrayOffset;
      this.blocks = blocks;
    }

    public PnIoCm_Packet_Res build() {
      PnIoCm_Packet_Res pnIoCm_Packet_Res =
          new PnIoCm_Packet_Res(
              errorCode2,
              errorCode1,
              errorDecode,
              errorCode,
              arrayMaximumCount,
              arrayOffset,
              blocks);
      return pnIoCm_Packet_Res;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Packet_Res)) {
      return false;
    }
    PnIoCm_Packet_Res that = (PnIoCm_Packet_Res) o;
    return (getErrorCode2() == that.getErrorCode2())
        && (getErrorCode1() == that.getErrorCode1())
        && (getErrorDecode() == that.getErrorDecode())
        && (getErrorCode() == that.getErrorCode())
        && (getArrayMaximumCount() == that.getArrayMaximumCount())
        && (getArrayOffset() == that.getArrayOffset())
        && (getBlocks() == that.getBlocks())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getErrorCode2(),
        getErrorCode1(),
        getErrorDecode(),
        getErrorCode(),
        getArrayMaximumCount(),
        getArrayOffset(),
        getBlocks());
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
