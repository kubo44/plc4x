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

public class MeteringDataOilConsumption extends MeteringData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final long L;

  public MeteringDataOilConsumption(
      MeteringCommandTypeContainer commandTypeContainer, byte argument, long L) {
    super(commandTypeContainer, argument);
    this.L = L;
  }

  public long getL() {
    return L;
  }

  @Override
  protected void serializeMeteringDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MeteringDataOilConsumption");

    // Simple Field (L)
    writeSimpleField("L", L, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("MeteringDataOilConsumption");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MeteringDataOilConsumption _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (L)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static MeteringDataBuilder staticParseMeteringDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("MeteringDataOilConsumption");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long L = readSimpleField("L", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("MeteringDataOilConsumption");
    // Create the instance
    return new MeteringDataOilConsumptionBuilderImpl(L);
  }

  public static class MeteringDataOilConsumptionBuilderImpl
      implements MeteringData.MeteringDataBuilder {
    private final long L;

    public MeteringDataOilConsumptionBuilderImpl(long L) {
      this.L = L;
    }

    public MeteringDataOilConsumption build(
        MeteringCommandTypeContainer commandTypeContainer, byte argument) {
      MeteringDataOilConsumption meteringDataOilConsumption =
          new MeteringDataOilConsumption(commandTypeContainer, argument, L);
      return meteringDataOilConsumption;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MeteringDataOilConsumption)) {
      return false;
    }
    MeteringDataOilConsumption that = (MeteringDataOilConsumption) o;
    return (getL() == that.getL()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getL());
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
