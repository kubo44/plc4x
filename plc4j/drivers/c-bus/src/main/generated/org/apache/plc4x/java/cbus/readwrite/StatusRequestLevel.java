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
package org.apache.plc4x.java.cbus.readwrite;

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

public class StatusRequestLevel extends StatusRequest implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final ApplicationIdContainer application;
  protected final byte startingGroupAddressLabel;

  // Reserved Fields
  private Byte reservedField0;
  private Byte reservedField1;

  public StatusRequestLevel(
      byte statusType, ApplicationIdContainer application, byte startingGroupAddressLabel) {
    super(statusType);
    this.application = application;
    this.startingGroupAddressLabel = startingGroupAddressLabel;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public byte getStartingGroupAddressLabel() {
    return startingGroupAddressLabel;
  }

  @Override
  protected void serializeStatusRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("StatusRequestLevel");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x73,
        writeByte(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (byte) 0x07,
        writeByte(writeBuffer, 8));

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (startingGroupAddressLabel)
    writeSimpleField(
        "startingGroupAddressLabel", startingGroupAddressLabel, writeByte(writeBuffer, 8));

    writeBuffer.popContext("StatusRequestLevel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    StatusRequestLevel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (application)
    lengthInBits += 8;

    // Simple field (startingGroupAddressLabel)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static StatusRequestBuilder staticParseStatusRequestBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("StatusRequestLevel");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0x73);

    Byte reservedField1 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0x07);

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    byte startingGroupAddressLabel =
        readSimpleField("startingGroupAddressLabel", readByte(readBuffer, 8));
    // Validation
    if (!(((((((((startingGroupAddressLabel) == (0x00)) || ((startingGroupAddressLabel) == (0x20)))
                            || ((startingGroupAddressLabel) == (0x40)))
                        || ((startingGroupAddressLabel) == (0x60)))
                    || ((startingGroupAddressLabel) == (0x80)))
                || ((startingGroupAddressLabel) == (0xA0)))
            || ((startingGroupAddressLabel) == (0xC0)))
        || ((startingGroupAddressLabel) == (0xE0)))) {
      throw new ParseValidationException("invalid label");
    }

    readBuffer.closeContext("StatusRequestLevel");
    // Create the instance
    return new StatusRequestLevelBuilderImpl(
        application, startingGroupAddressLabel, reservedField0, reservedField1);
  }

  public static class StatusRequestLevelBuilderImpl implements StatusRequest.StatusRequestBuilder {
    private final ApplicationIdContainer application;
    private final byte startingGroupAddressLabel;
    private final Byte reservedField0;
    private final Byte reservedField1;

    public StatusRequestLevelBuilderImpl(
        ApplicationIdContainer application,
        byte startingGroupAddressLabel,
        Byte reservedField0,
        Byte reservedField1) {
      this.application = application;
      this.startingGroupAddressLabel = startingGroupAddressLabel;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
    }

    public StatusRequestLevel build(byte statusType) {
      StatusRequestLevel statusRequestLevel =
          new StatusRequestLevel(statusType, application, startingGroupAddressLabel);
      statusRequestLevel.reservedField0 = reservedField0;
      statusRequestLevel.reservedField1 = reservedField1;
      return statusRequestLevel;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof StatusRequestLevel)) {
      return false;
    }
    StatusRequestLevel that = (StatusRequestLevel) o;
    return (getApplication() == that.getApplication())
        && (getStartingGroupAddressLabel() == that.getStartingGroupAddressLabel())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getApplication(), getStartingGroupAddressLabel());
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
