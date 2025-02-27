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

public class BACnetHostNPort implements Message {

  // Properties.
  protected final BACnetHostAddressEnclosed host;
  protected final BACnetContextTagUnsignedInteger port;

  public BACnetHostNPort(BACnetHostAddressEnclosed host, BACnetContextTagUnsignedInteger port) {
    super();
    this.host = host;
    this.port = port;
  }

  public BACnetHostAddressEnclosed getHost() {
    return host;
  }

  public BACnetContextTagUnsignedInteger getPort() {
    return port;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetHostNPort");

    // Simple Field (host)
    writeSimpleField("host", host, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (port)
    writeSimpleField("port", port, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetHostNPort");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetHostNPort _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (host)
    lengthInBits += host.getLengthInBits();

    // Simple field (port)
    lengthInBits += port.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetHostNPort staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetHostNPort staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetHostNPort");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetHostAddressEnclosed host =
        readSimpleField(
            "host",
            new DataReaderComplexDefault<>(
                () -> BACnetHostAddressEnclosed.staticParse(readBuffer, (short) (0)), readBuffer));

    BACnetContextTagUnsignedInteger port =
        readSimpleField(
            "port",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetHostNPort");
    // Create the instance
    BACnetHostNPort _bACnetHostNPort;
    _bACnetHostNPort = new BACnetHostNPort(host, port);
    return _bACnetHostNPort;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetHostNPort)) {
      return false;
    }
    BACnetHostNPort that = (BACnetHostNPort) o;
    return (getHost() == that.getHost()) && (getPort() == that.getPort()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHost(), getPort());
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
