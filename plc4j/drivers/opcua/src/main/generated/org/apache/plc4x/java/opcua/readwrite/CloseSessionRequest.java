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
package org.apache.plc4x.java.opcua.readwrite;

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

public class CloseSessionRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "473";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final boolean deleteSubscriptions;

  public CloseSessionRequest(ExtensionObjectDefinition requestHeader, boolean deleteSubscriptions) {
    super();
    this.requestHeader = requestHeader;
    this.deleteSubscriptions = deleteSubscriptions;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public boolean getDeleteSubscriptions() {
    return deleteSubscriptions;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CloseSessionRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 7));

    // Simple Field (deleteSubscriptions)
    writeSimpleField("deleteSubscriptions", deleteSubscriptions, writeBoolean(writeBuffer));

    writeBuffer.popContext("CloseSessionRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CloseSessionRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (deleteSubscriptions)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("CloseSessionRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean deleteSubscriptions = readSimpleField("deleteSubscriptions", readBoolean(readBuffer));

    readBuffer.closeContext("CloseSessionRequest");
    // Create the instance
    return new CloseSessionRequestBuilderImpl(requestHeader, deleteSubscriptions);
  }

  public static class CloseSessionRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final boolean deleteSubscriptions;

    public CloseSessionRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader, boolean deleteSubscriptions) {
      this.requestHeader = requestHeader;
      this.deleteSubscriptions = deleteSubscriptions;
    }

    public CloseSessionRequest build() {
      CloseSessionRequest closeSessionRequest =
          new CloseSessionRequest(requestHeader, deleteSubscriptions);
      return closeSessionRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CloseSessionRequest)) {
      return false;
    }
    CloseSessionRequest that = (CloseSessionRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getDeleteSubscriptions() == that.getDeleteSubscriptions())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestHeader(), getDeleteSubscriptions());
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
