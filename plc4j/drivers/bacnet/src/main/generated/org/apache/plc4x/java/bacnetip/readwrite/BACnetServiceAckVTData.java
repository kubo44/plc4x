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

public class BACnetServiceAckVTData extends BACnetServiceAck implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.VT_DATA;
  }

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger vtSessionIdentifier;
  protected final BACnetApplicationTagOctetString vtNewData;
  protected final BACnetApplicationTagUnsignedInteger vtDataFlag;

  // Arguments.
  protected final Long serviceAckLength;

  public BACnetServiceAckVTData(
      BACnetApplicationTagUnsignedInteger vtSessionIdentifier,
      BACnetApplicationTagOctetString vtNewData,
      BACnetApplicationTagUnsignedInteger vtDataFlag,
      Long serviceAckLength) {
    super(serviceAckLength);
    this.vtSessionIdentifier = vtSessionIdentifier;
    this.vtNewData = vtNewData;
    this.vtDataFlag = vtDataFlag;
    this.serviceAckLength = serviceAckLength;
  }

  public BACnetApplicationTagUnsignedInteger getVtSessionIdentifier() {
    return vtSessionIdentifier;
  }

  public BACnetApplicationTagOctetString getVtNewData() {
    return vtNewData;
  }

  public BACnetApplicationTagUnsignedInteger getVtDataFlag() {
    return vtDataFlag;
  }

  @Override
  protected void serializeBACnetServiceAckChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetServiceAckVTData");

    // Simple Field (vtSessionIdentifier)
    writeSimpleField(
        "vtSessionIdentifier", vtSessionIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (vtNewData)
    writeSimpleField("vtNewData", vtNewData, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (vtDataFlag)
    writeSimpleField("vtDataFlag", vtDataFlag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetServiceAckVTData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetServiceAckVTData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (vtSessionIdentifier)
    lengthInBits += vtSessionIdentifier.getLengthInBits();

    // Simple field (vtNewData)
    lengthInBits += vtNewData.getLengthInBits();

    // Simple field (vtDataFlag)
    lengthInBits += vtDataFlag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetServiceAckBuilder staticParseBACnetServiceAckBuilder(
      ReadBuffer readBuffer, Long serviceAckLength) throws ParseException {
    readBuffer.pullContext("BACnetServiceAckVTData");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger vtSessionIdentifier =
        readSimpleField(
            "vtSessionIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagOctetString vtNewData =
        readSimpleField(
            "vtNewData",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagOctetString) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagUnsignedInteger vtDataFlag =
        readSimpleField(
            "vtDataFlag",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetServiceAckVTData");
    // Create the instance
    return new BACnetServiceAckVTDataBuilderImpl(
        vtSessionIdentifier, vtNewData, vtDataFlag, serviceAckLength);
  }

  public static class BACnetServiceAckVTDataBuilderImpl
      implements BACnetServiceAck.BACnetServiceAckBuilder {
    private final BACnetApplicationTagUnsignedInteger vtSessionIdentifier;
    private final BACnetApplicationTagOctetString vtNewData;
    private final BACnetApplicationTagUnsignedInteger vtDataFlag;
    private final Long serviceAckLength;

    public BACnetServiceAckVTDataBuilderImpl(
        BACnetApplicationTagUnsignedInteger vtSessionIdentifier,
        BACnetApplicationTagOctetString vtNewData,
        BACnetApplicationTagUnsignedInteger vtDataFlag,
        Long serviceAckLength) {
      this.vtSessionIdentifier = vtSessionIdentifier;
      this.vtNewData = vtNewData;
      this.vtDataFlag = vtDataFlag;
      this.serviceAckLength = serviceAckLength;
    }

    public BACnetServiceAckVTData build(Long serviceAckLength) {

      BACnetServiceAckVTData bACnetServiceAckVTData =
          new BACnetServiceAckVTData(vtSessionIdentifier, vtNewData, vtDataFlag, serviceAckLength);
      return bACnetServiceAckVTData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetServiceAckVTData)) {
      return false;
    }
    BACnetServiceAckVTData that = (BACnetServiceAckVTData) o;
    return (getVtSessionIdentifier() == that.getVtSessionIdentifier())
        && (getVtNewData() == that.getVtNewData())
        && (getVtDataFlag() == that.getVtDataFlag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getVtSessionIdentifier(), getVtNewData(), getVtDataFlag());
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
