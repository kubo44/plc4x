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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ApduDataExtPropertyValueRead extends ApduDataExt implements Message {

  // Accessors for discriminator values.
  public Short getExtApciType() {
    return (short) 0x15;
  }

  // Properties.
  protected final short objectIndex;
  protected final short propertyId;
  protected final byte count;
  protected final int index;

  public ApduDataExtPropertyValueRead(short objectIndex, short propertyId, byte count, int index) {
    super();
    this.objectIndex = objectIndex;
    this.propertyId = propertyId;
    this.count = count;
    this.index = index;
  }

  public short getObjectIndex() {
    return objectIndex;
  }

  public short getPropertyId() {
    return propertyId;
  }

  public byte getCount() {
    return count;
  }

  public int getIndex() {
    return index;
  }

  @Override
  protected void serializeApduDataExtChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduDataExtPropertyValueRead");

    // Simple Field (objectIndex)
    writeSimpleField("objectIndex", objectIndex, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (propertyId)
    writeSimpleField("propertyId", propertyId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (count)
    writeSimpleField("count", count, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (index)
    writeSimpleField("index", index, writeUnsignedInt(writeBuffer, 12));

    writeBuffer.popContext("ApduDataExtPropertyValueRead");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataExtPropertyValueRead _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIndex)
    lengthInBits += 8;

    // Simple field (propertyId)
    lengthInBits += 8;

    // Simple field (count)
    lengthInBits += 4;

    // Simple field (index)
    lengthInBits += 12;

    return lengthInBits;
  }

  public static ApduDataExtBuilder staticParseApduDataExtBuilder(
      ReadBuffer readBuffer, Short length) throws ParseException {
    readBuffer.pullContext("ApduDataExtPropertyValueRead");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short objectIndex = readSimpleField("objectIndex", readUnsignedShort(readBuffer, 8));

    short propertyId = readSimpleField("propertyId", readUnsignedShort(readBuffer, 8));

    byte count = readSimpleField("count", readUnsignedByte(readBuffer, 4));

    int index = readSimpleField("index", readUnsignedInt(readBuffer, 12));

    readBuffer.closeContext("ApduDataExtPropertyValueRead");
    // Create the instance
    return new ApduDataExtPropertyValueReadBuilderImpl(objectIndex, propertyId, count, index);
  }

  public static class ApduDataExtPropertyValueReadBuilderImpl
      implements ApduDataExt.ApduDataExtBuilder {
    private final short objectIndex;
    private final short propertyId;
    private final byte count;
    private final int index;

    public ApduDataExtPropertyValueReadBuilderImpl(
        short objectIndex, short propertyId, byte count, int index) {
      this.objectIndex = objectIndex;
      this.propertyId = propertyId;
      this.count = count;
      this.index = index;
    }

    public ApduDataExtPropertyValueRead build() {
      ApduDataExtPropertyValueRead apduDataExtPropertyValueRead =
          new ApduDataExtPropertyValueRead(objectIndex, propertyId, count, index);
      return apduDataExtPropertyValueRead;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataExtPropertyValueRead)) {
      return false;
    }
    ApduDataExtPropertyValueRead that = (ApduDataExtPropertyValueRead) o;
    return (getObjectIndex() == that.getObjectIndex())
        && (getPropertyId() == that.getPropertyId())
        && (getCount() == that.getCount())
        && (getIndex() == that.getIndex())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getObjectIndex(), getPropertyId(), getCount(), getIndex());
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
