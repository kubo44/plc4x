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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetVTSession implements Message {

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger localVtSessionId;
  protected final BACnetApplicationTagUnsignedInteger removeVtSessionId;
  protected final BACnetAddress remoteVtAddress;

  public BACnetVTSession(
      BACnetApplicationTagUnsignedInteger localVtSessionId,
      BACnetApplicationTagUnsignedInteger removeVtSessionId,
      BACnetAddress remoteVtAddress) {
    super();
    this.localVtSessionId = localVtSessionId;
    this.removeVtSessionId = removeVtSessionId;
    this.remoteVtAddress = remoteVtAddress;
  }

  public BACnetApplicationTagUnsignedInteger getLocalVtSessionId() {
    return localVtSessionId;
  }

  public BACnetApplicationTagUnsignedInteger getRemoveVtSessionId() {
    return removeVtSessionId;
  }

  public BACnetAddress getRemoteVtAddress() {
    return remoteVtAddress;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetVTSession");

    // Simple Field (localVtSessionId)
    writeSimpleField(
        "localVtSessionId", localVtSessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (removeVtSessionId)
    writeSimpleField(
        "removeVtSessionId", removeVtSessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (remoteVtAddress)
    writeSimpleField(
        "remoteVtAddress", remoteVtAddress, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetVTSession");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetVTSession _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (localVtSessionId)
    lengthInBits += localVtSessionId.getLengthInBits();

    // Simple field (removeVtSessionId)
    lengthInBits += removeVtSessionId.getLengthInBits();

    // Simple field (remoteVtAddress)
    lengthInBits += remoteVtAddress.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetVTSession staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetVTSession staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetVTSession");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger localVtSessionId =
        readSimpleField(
            "localVtSessionId",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagUnsignedInteger removeVtSessionId =
        readSimpleField(
            "removeVtSessionId",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetAddress remoteVtAddress =
        readSimpleField(
            "remoteVtAddress",
            new DataReaderComplexDefault<>(
                () -> BACnetAddress.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("BACnetVTSession");
    // Create the instance
    BACnetVTSession _bACnetVTSession;
    _bACnetVTSession = new BACnetVTSession(localVtSessionId, removeVtSessionId, remoteVtAddress);
    return _bACnetVTSession;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetVTSession)) {
      return false;
    }
    BACnetVTSession that = (BACnetVTSession) o;
    return (getLocalVtSessionId() == that.getLocalVtSessionId())
        && (getRemoveVtSessionId() == that.getRemoveVtSessionId())
        && (getRemoteVtAddress() == that.getRemoteVtAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getLocalVtSessionId(), getRemoveVtSessionId(), getRemoteVtAddress());
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
