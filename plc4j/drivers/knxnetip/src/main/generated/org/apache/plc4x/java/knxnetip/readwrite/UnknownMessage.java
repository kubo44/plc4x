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

public class UnknownMessage extends KnxNetIpMessage implements Message {

  // Accessors for discriminator values.
  public Integer getMsgType() {
    return (int) 0x020B;
  }

  // Properties.
  protected final byte[] unknownData;

  public UnknownMessage(byte[] unknownData) {
    super();
    this.unknownData = unknownData;
  }

  public byte[] getUnknownData() {
    return unknownData;
  }

  @Override
  protected void serializeKnxNetIpMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("UnknownMessage");

    // Array Field (unknownData)
    writeByteArrayField(
        "unknownData",
        unknownData,
        writeByteArray(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("UnknownMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UnknownMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (unknownData != null) {
      lengthInBits += 8 * unknownData.length;
    }

    return lengthInBits;
  }

  public static KnxNetIpMessageBuilder staticParseKnxNetIpMessageBuilder(
      ReadBuffer readBuffer, Integer totalLength) throws ParseException {
    readBuffer.pullContext("UnknownMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte[] unknownData =
        readBuffer.readByteArray(
            "unknownData",
            Math.toIntExact((totalLength) - (6)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("UnknownMessage");
    // Create the instance
    return new UnknownMessageBuilderImpl(unknownData);
  }

  public static class UnknownMessageBuilderImpl implements KnxNetIpMessage.KnxNetIpMessageBuilder {
    private final byte[] unknownData;

    public UnknownMessageBuilderImpl(byte[] unknownData) {
      this.unknownData = unknownData;
    }

    public UnknownMessage build() {
      UnknownMessage unknownMessage = new UnknownMessage(unknownData);
      return unknownMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UnknownMessage)) {
      return false;
    }
    UnknownMessage that = (UnknownMessage) o;
    return (getUnknownData() == that.getUnknownData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnknownData());
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
