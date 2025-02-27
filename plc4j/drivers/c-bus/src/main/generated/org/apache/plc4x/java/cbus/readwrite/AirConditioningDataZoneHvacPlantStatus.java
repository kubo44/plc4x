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

public class AirConditioningDataZoneHvacPlantStatus extends AirConditioningData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte zoneGroup;
  protected final HVACZoneList zoneList;
  protected final HVACType hvacType;
  protected final HVACStatusFlags hvacStatus;
  protected final HVACError hvacErrorCode;

  public AirConditioningDataZoneHvacPlantStatus(
      AirConditioningCommandTypeContainer commandTypeContainer,
      byte zoneGroup,
      HVACZoneList zoneList,
      HVACType hvacType,
      HVACStatusFlags hvacStatus,
      HVACError hvacErrorCode) {
    super(commandTypeContainer);
    this.zoneGroup = zoneGroup;
    this.zoneList = zoneList;
    this.hvacType = hvacType;
    this.hvacStatus = hvacStatus;
    this.hvacErrorCode = hvacErrorCode;
  }

  public byte getZoneGroup() {
    return zoneGroup;
  }

  public HVACZoneList getZoneList() {
    return zoneList;
  }

  public HVACType getHvacType() {
    return hvacType;
  }

  public HVACStatusFlags getHvacStatus() {
    return hvacStatus;
  }

  public HVACError getHvacErrorCode() {
    return hvacErrorCode;
  }

  @Override
  protected void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AirConditioningDataZoneHvacPlantStatus");

    // Simple Field (zoneGroup)
    writeSimpleField("zoneGroup", zoneGroup, writeByte(writeBuffer, 8));

    // Simple Field (zoneList)
    writeSimpleField("zoneList", zoneList, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (hvacType)
    writeSimpleEnumField(
        "hvacType",
        "HVACType",
        hvacType,
        new DataWriterEnumDefault<>(
            HVACType::getValue, HVACType::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (hvacStatus)
    writeSimpleField("hvacStatus", hvacStatus, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (hvacErrorCode)
    writeSimpleEnumField(
        "hvacErrorCode",
        "HVACError",
        hvacErrorCode,
        new DataWriterEnumDefault<>(
            HVACError::getValue, HVACError::name, writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("AirConditioningDataZoneHvacPlantStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AirConditioningDataZoneHvacPlantStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (zoneGroup)
    lengthInBits += 8;

    // Simple field (zoneList)
    lengthInBits += zoneList.getLengthInBits();

    // Simple field (hvacType)
    lengthInBits += 8;

    // Simple field (hvacStatus)
    lengthInBits += hvacStatus.getLengthInBits();

    // Simple field (hvacErrorCode)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static AirConditioningDataBuilder staticParseAirConditioningDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AirConditioningDataZoneHvacPlantStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte zoneGroup = readSimpleField("zoneGroup", readByte(readBuffer, 8));

    HVACZoneList zoneList =
        readSimpleField(
            "zoneList",
            new DataReaderComplexDefault<>(() -> HVACZoneList.staticParse(readBuffer), readBuffer));

    HVACType hvacType =
        readEnumField(
            "hvacType",
            "HVACType",
            new DataReaderEnumDefault<>(HVACType::enumForValue, readUnsignedShort(readBuffer, 8)));

    HVACStatusFlags hvacStatus =
        readSimpleField(
            "hvacStatus",
            new DataReaderComplexDefault<>(
                () -> HVACStatusFlags.staticParse(readBuffer), readBuffer));

    HVACError hvacErrorCode =
        readEnumField(
            "hvacErrorCode",
            "HVACError",
            new DataReaderEnumDefault<>(HVACError::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("AirConditioningDataZoneHvacPlantStatus");
    // Create the instance
    return new AirConditioningDataZoneHvacPlantStatusBuilderImpl(
        zoneGroup, zoneList, hvacType, hvacStatus, hvacErrorCode);
  }

  public static class AirConditioningDataZoneHvacPlantStatusBuilderImpl
      implements AirConditioningData.AirConditioningDataBuilder {
    private final byte zoneGroup;
    private final HVACZoneList zoneList;
    private final HVACType hvacType;
    private final HVACStatusFlags hvacStatus;
    private final HVACError hvacErrorCode;

    public AirConditioningDataZoneHvacPlantStatusBuilderImpl(
        byte zoneGroup,
        HVACZoneList zoneList,
        HVACType hvacType,
        HVACStatusFlags hvacStatus,
        HVACError hvacErrorCode) {
      this.zoneGroup = zoneGroup;
      this.zoneList = zoneList;
      this.hvacType = hvacType;
      this.hvacStatus = hvacStatus;
      this.hvacErrorCode = hvacErrorCode;
    }

    public AirConditioningDataZoneHvacPlantStatus build(
        AirConditioningCommandTypeContainer commandTypeContainer) {
      AirConditioningDataZoneHvacPlantStatus airConditioningDataZoneHvacPlantStatus =
          new AirConditioningDataZoneHvacPlantStatus(
              commandTypeContainer, zoneGroup, zoneList, hvacType, hvacStatus, hvacErrorCode);
      return airConditioningDataZoneHvacPlantStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningDataZoneHvacPlantStatus)) {
      return false;
    }
    AirConditioningDataZoneHvacPlantStatus that = (AirConditioningDataZoneHvacPlantStatus) o;
    return (getZoneGroup() == that.getZoneGroup())
        && (getZoneList() == that.getZoneList())
        && (getHvacType() == that.getHvacType())
        && (getHvacStatus() == that.getHvacStatus())
        && (getHvacErrorCode() == that.getHvacErrorCode())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getZoneGroup(),
        getZoneList(),
        getHvacType(),
        getHvacStatus(),
        getHvacErrorCode());
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
