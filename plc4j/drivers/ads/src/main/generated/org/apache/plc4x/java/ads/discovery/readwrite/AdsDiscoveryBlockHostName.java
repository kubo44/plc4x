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
package org.apache.plc4x.java.ads.discovery.readwrite;

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

public class AdsDiscoveryBlockHostName extends AdsDiscoveryBlock implements Message {

  // Accessors for discriminator values.
  public AdsDiscoveryBlockType getBlockType() {
    return AdsDiscoveryBlockType.HOST_NAME;
  }

  // Properties.
  protected final AmsString hostName;

  public AdsDiscoveryBlockHostName(AmsString hostName) {
    super();
    this.hostName = hostName;
  }

  public AmsString getHostName() {
    return hostName;
  }

  @Override
  protected void serializeAdsDiscoveryBlockChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AdsDiscoveryBlockHostName");

    // Simple Field (hostName)
    writeSimpleField("hostName", hostName, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("AdsDiscoveryBlockHostName");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsDiscoveryBlockHostName _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (hostName)
    lengthInBits += hostName.getLengthInBits();

    return lengthInBits;
  }

  public static AdsDiscoveryBlockBuilder staticParseAdsDiscoveryBlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsDiscoveryBlockHostName");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    AmsString hostName =
        readSimpleField(
            "hostName",
            new DataReaderComplexDefault<>(() -> AmsString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("AdsDiscoveryBlockHostName");
    // Create the instance
    return new AdsDiscoveryBlockHostNameBuilderImpl(hostName);
  }

  public static class AdsDiscoveryBlockHostNameBuilderImpl
      implements AdsDiscoveryBlock.AdsDiscoveryBlockBuilder {
    private final AmsString hostName;

    public AdsDiscoveryBlockHostNameBuilderImpl(AmsString hostName) {
      this.hostName = hostName;
    }

    public AdsDiscoveryBlockHostName build() {
      AdsDiscoveryBlockHostName adsDiscoveryBlockHostName = new AdsDiscoveryBlockHostName(hostName);
      return adsDiscoveryBlockHostName;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsDiscoveryBlockHostName)) {
      return false;
    }
    AdsDiscoveryBlockHostName that = (AdsDiscoveryBlockHostName) o;
    return (getHostName() == that.getHostName()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getHostName());
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
