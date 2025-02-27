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

public abstract class S7DataAlarmMessage implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getCpuFunctionType();

  // Constant values.
  public static final Short FUNCTIONID = 0x00;
  public static final Short NUMBERMESSAGEOBJ = 0x01;

  public S7DataAlarmMessage() {
    super();
  }

  public short getFunctionId() {
    return FUNCTIONID;
  }

  public short getNumberMessageObj() {
    return NUMBERMESSAGEOBJ;
  }

  protected abstract void serializeS7DataAlarmMessageChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7DataAlarmMessage");

    // Const Field (functionId)
    writeConstField("functionId", FUNCTIONID, writeUnsignedShort(writeBuffer, 8));

    // Const Field (numberMessageObj)
    writeConstField("numberMessageObj", NUMBERMESSAGEOBJ, writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeS7DataAlarmMessageChild(writeBuffer);

    writeBuffer.popContext("S7DataAlarmMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    S7DataAlarmMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (functionId)
    lengthInBits += 8;

    // Const Field (numberMessageObj)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static S7DataAlarmMessage staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Byte cpuFunctionType;
    if (args[0] instanceof Byte) {
      cpuFunctionType = (Byte) args[0];
    } else if (args[0] instanceof String) {
      cpuFunctionType = Byte.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Byte or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, cpuFunctionType);
  }

  public static S7DataAlarmMessage staticParse(ReadBuffer readBuffer, Byte cpuFunctionType)
      throws ParseException {
    readBuffer.pullContext("S7DataAlarmMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short functionId =
        readConstField(
            "functionId", readUnsignedShort(readBuffer, 8), S7DataAlarmMessage.FUNCTIONID);

    short numberMessageObj =
        readConstField(
            "numberMessageObj",
            readUnsignedShort(readBuffer, 8),
            S7DataAlarmMessage.NUMBERMESSAGEOBJ);

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    S7DataAlarmMessageBuilder builder = null;
    if (EvaluationHelper.equals(cpuFunctionType, (byte) 0x04)) {
      builder =
          S7MessageObjectRequest.staticParseS7DataAlarmMessageBuilder(readBuffer, cpuFunctionType);
    } else if (EvaluationHelper.equals(cpuFunctionType, (byte) 0x08)) {
      builder =
          S7MessageObjectResponse.staticParseS7DataAlarmMessageBuilder(readBuffer, cpuFunctionType);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "cpuFunctionType="
              + cpuFunctionType
              + "]");
    }

    readBuffer.closeContext("S7DataAlarmMessage");
    // Create the instance
    S7DataAlarmMessage _s7DataAlarmMessage = builder.build();
    return _s7DataAlarmMessage;
  }

  public interface S7DataAlarmMessageBuilder {
    S7DataAlarmMessage build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7DataAlarmMessage)) {
      return false;
    }
    S7DataAlarmMessage that = (S7DataAlarmMessage) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
