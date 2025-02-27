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

public class DeviceConfigurationAckDataBlock implements Message {

  // Properties.
  protected final short communicationChannelId;
  protected final short sequenceCounter;
  protected final Status status;

  public DeviceConfigurationAckDataBlock(
      short communicationChannelId, short sequenceCounter, Status status) {
    super();
    this.communicationChannelId = communicationChannelId;
    this.sequenceCounter = sequenceCounter;
    this.status = status;
  }

  public short getCommunicationChannelId() {
    return communicationChannelId;
  }

  public short getSequenceCounter() {
    return sequenceCounter;
  }

  public Status getStatus() {
    return status;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DeviceConfigurationAckDataBlock");

    // Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short structureLength = (short) (getLengthInBytes());
    writeImplicitField("structureLength", structureLength, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (communicationChannelId)
    writeSimpleField(
        "communicationChannelId", communicationChannelId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (sequenceCounter)
    writeSimpleField("sequenceCounter", sequenceCounter, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleEnumField(
        "status",
        "Status",
        status,
        new DataWriterEnumDefault<>(
            Status::getValue, Status::name, writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("DeviceConfigurationAckDataBlock");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DeviceConfigurationAckDataBlock _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (structureLength)
    lengthInBits += 8;

    // Simple field (communicationChannelId)
    lengthInBits += 8;

    // Simple field (sequenceCounter)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static DeviceConfigurationAckDataBlock staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DeviceConfigurationAckDataBlock staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("DeviceConfigurationAckDataBlock");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short structureLength = readImplicitField("structureLength", readUnsignedShort(readBuffer, 8));

    short communicationChannelId =
        readSimpleField("communicationChannelId", readUnsignedShort(readBuffer, 8));

    short sequenceCounter = readSimpleField("sequenceCounter", readUnsignedShort(readBuffer, 8));

    Status status =
        readEnumField(
            "status",
            "Status",
            new DataReaderEnumDefault<>(Status::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("DeviceConfigurationAckDataBlock");
    // Create the instance
    DeviceConfigurationAckDataBlock _deviceConfigurationAckDataBlock;
    _deviceConfigurationAckDataBlock =
        new DeviceConfigurationAckDataBlock(communicationChannelId, sequenceCounter, status);
    return _deviceConfigurationAckDataBlock;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeviceConfigurationAckDataBlock)) {
      return false;
    }
    DeviceConfigurationAckDataBlock that = (DeviceConfigurationAckDataBlock) o;
    return (getCommunicationChannelId() == that.getCommunicationChannelId())
        && (getSequenceCounter() == that.getSequenceCounter())
        && (getStatus() == that.getStatus())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommunicationChannelId(), getSequenceCounter(), getStatus());
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
