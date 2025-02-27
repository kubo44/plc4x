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
package org.apache.plc4x.java.s7.readwrite;

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

public class AlarmMessageObjectPushType implements Message {

  // Constant values.
  public static final Short VARIABLESPEC = 0x12;

  // Properties.
  protected final short lengthSpec;
  protected final SyntaxIdType syntaxId;
  protected final short numberOfValues;
  protected final long eventId;
  protected final State eventState;
  protected final State localState;
  protected final State ackStateGoing;
  protected final State ackStateComing;
  protected final List<AssociatedValueType> AssociatedValues;

  public AlarmMessageObjectPushType(
      short lengthSpec,
      SyntaxIdType syntaxId,
      short numberOfValues,
      long eventId,
      State eventState,
      State localState,
      State ackStateGoing,
      State ackStateComing,
      List<AssociatedValueType> AssociatedValues) {
    super();
    this.lengthSpec = lengthSpec;
    this.syntaxId = syntaxId;
    this.numberOfValues = numberOfValues;
    this.eventId = eventId;
    this.eventState = eventState;
    this.localState = localState;
    this.ackStateGoing = ackStateGoing;
    this.ackStateComing = ackStateComing;
    this.AssociatedValues = AssociatedValues;
  }

  public short getLengthSpec() {
    return lengthSpec;
  }

  public SyntaxIdType getSyntaxId() {
    return syntaxId;
  }

  public short getNumberOfValues() {
    return numberOfValues;
  }

  public long getEventId() {
    return eventId;
  }

  public State getEventState() {
    return eventState;
  }

  public State getLocalState() {
    return localState;
  }

  public State getAckStateGoing() {
    return ackStateGoing;
  }

  public State getAckStateComing() {
    return ackStateComing;
  }

  public List<AssociatedValueType> getAssociatedValues() {
    return AssociatedValues;
  }

  public short getVariableSpec() {
    return VARIABLESPEC;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AlarmMessageObjectPushType");

    // Const Field (variableSpec)
    writeConstField("variableSpec", VARIABLESPEC, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (lengthSpec)
    writeSimpleField("lengthSpec", lengthSpec, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (syntaxId)
    writeSimpleEnumField(
        "syntaxId",
        "SyntaxIdType",
        syntaxId,
        new DataWriterEnumDefault<>(
            SyntaxIdType::getValue, SyntaxIdType::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (numberOfValues)
    writeSimpleField("numberOfValues", numberOfValues, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (eventId)
    writeSimpleField("eventId", eventId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (eventState)
    writeSimpleField("eventState", eventState, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (localState)
    writeSimpleField("localState", localState, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (ackStateGoing)
    writeSimpleField("ackStateGoing", ackStateGoing, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (ackStateComing)
    writeSimpleField("ackStateComing", ackStateComing, new DataWriterComplexDefault<>(writeBuffer));

    // Array Field (AssociatedValues)
    writeComplexTypeArrayField("AssociatedValues", AssociatedValues, writeBuffer);

    writeBuffer.popContext("AlarmMessageObjectPushType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AlarmMessageObjectPushType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (variableSpec)
    lengthInBits += 8;

    // Simple field (lengthSpec)
    lengthInBits += 8;

    // Simple field (syntaxId)
    lengthInBits += 8;

    // Simple field (numberOfValues)
    lengthInBits += 8;

    // Simple field (eventId)
    lengthInBits += 32;

    // Simple field (eventState)
    lengthInBits += eventState.getLengthInBits();

    // Simple field (localState)
    lengthInBits += localState.getLengthInBits();

    // Simple field (ackStateGoing)
    lengthInBits += ackStateGoing.getLengthInBits();

    // Simple field (ackStateComing)
    lengthInBits += ackStateComing.getLengthInBits();

    // Array field
    if (AssociatedValues != null) {
      int i = 0;
      for (AssociatedValueType element : AssociatedValues) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= AssociatedValues.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static AlarmMessageObjectPushType staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AlarmMessageObjectPushType staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AlarmMessageObjectPushType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short variableSpec =
        readConstField(
            "variableSpec",
            readUnsignedShort(readBuffer, 8),
            AlarmMessageObjectPushType.VARIABLESPEC);

    short lengthSpec = readSimpleField("lengthSpec", readUnsignedShort(readBuffer, 8));

    SyntaxIdType syntaxId =
        readEnumField(
            "syntaxId",
            "SyntaxIdType",
            new DataReaderEnumDefault<>(
                SyntaxIdType::enumForValue, readUnsignedShort(readBuffer, 8)));

    short numberOfValues = readSimpleField("numberOfValues", readUnsignedShort(readBuffer, 8));

    long eventId = readSimpleField("eventId", readUnsignedLong(readBuffer, 32));

    State eventState =
        readSimpleField(
            "eventState",
            new DataReaderComplexDefault<>(() -> State.staticParse(readBuffer), readBuffer));

    State localState =
        readSimpleField(
            "localState",
            new DataReaderComplexDefault<>(() -> State.staticParse(readBuffer), readBuffer));

    State ackStateGoing =
        readSimpleField(
            "ackStateGoing",
            new DataReaderComplexDefault<>(() -> State.staticParse(readBuffer), readBuffer));

    State ackStateComing =
        readSimpleField(
            "ackStateComing",
            new DataReaderComplexDefault<>(() -> State.staticParse(readBuffer), readBuffer));

    List<AssociatedValueType> AssociatedValues =
        readCountArrayField(
            "AssociatedValues",
            new DataReaderComplexDefault<>(
                () -> AssociatedValueType.staticParse(readBuffer), readBuffer),
            numberOfValues);

    readBuffer.closeContext("AlarmMessageObjectPushType");
    // Create the instance
    AlarmMessageObjectPushType _alarmMessageObjectPushType;
    _alarmMessageObjectPushType =
        new AlarmMessageObjectPushType(
            lengthSpec,
            syntaxId,
            numberOfValues,
            eventId,
            eventState,
            localState,
            ackStateGoing,
            ackStateComing,
            AssociatedValues);
    return _alarmMessageObjectPushType;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AlarmMessageObjectPushType)) {
      return false;
    }
    AlarmMessageObjectPushType that = (AlarmMessageObjectPushType) o;
    return (getLengthSpec() == that.getLengthSpec())
        && (getSyntaxId() == that.getSyntaxId())
        && (getNumberOfValues() == that.getNumberOfValues())
        && (getEventId() == that.getEventId())
        && (getEventState() == that.getEventState())
        && (getLocalState() == that.getLocalState())
        && (getAckStateGoing() == that.getAckStateGoing())
        && (getAckStateComing() == that.getAckStateComing())
        && (getAssociatedValues() == that.getAssociatedValues())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getLengthSpec(),
        getSyntaxId(),
        getNumberOfValues(),
        getEventId(),
        getEventState(),
        getLocalState(),
        getAckStateGoing(),
        getAckStateComing(),
        getAssociatedValues());
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
