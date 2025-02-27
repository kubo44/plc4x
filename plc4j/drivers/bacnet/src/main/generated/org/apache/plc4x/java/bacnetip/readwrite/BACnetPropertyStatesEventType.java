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

public class BACnetPropertyStatesEventType extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetEventTypeTagged eventType;

  public BACnetPropertyStatesEventType(
      BACnetTagHeader peekedTagHeader, BACnetEventTypeTagged eventType) {
    super(peekedTagHeader);
    this.eventType = eventType;
  }

  public BACnetEventTypeTagged getEventType() {
    return eventType;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesEventType");

    // Simple Field (eventType)
    writeSimpleField("eventType", eventType, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesEventType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesEventType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (eventType)
    lengthInBits += eventType.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesEventType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetEventTypeTagged eventType =
        readSimpleField(
            "eventType",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetEventTypeTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesEventType");
    // Create the instance
    return new BACnetPropertyStatesEventTypeBuilderImpl(eventType);
  }

  public static class BACnetPropertyStatesEventTypeBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetEventTypeTagged eventType;

    public BACnetPropertyStatesEventTypeBuilderImpl(BACnetEventTypeTagged eventType) {
      this.eventType = eventType;
    }

    public BACnetPropertyStatesEventType build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesEventType bACnetPropertyStatesEventType =
          new BACnetPropertyStatesEventType(peekedTagHeader, eventType);
      return bACnetPropertyStatesEventType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesEventType)) {
      return false;
    }
    BACnetPropertyStatesEventType that = (BACnetPropertyStatesEventType) o;
    return (getEventType() == that.getEventType()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEventType());
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
