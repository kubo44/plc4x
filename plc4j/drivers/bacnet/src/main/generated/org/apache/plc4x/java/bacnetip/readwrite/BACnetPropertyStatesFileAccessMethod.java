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

public class BACnetPropertyStatesFileAccessMethod extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetFileAccessMethodTagged fileAccessMethod;

  public BACnetPropertyStatesFileAccessMethod(
      BACnetTagHeader peekedTagHeader, BACnetFileAccessMethodTagged fileAccessMethod) {
    super(peekedTagHeader);
    this.fileAccessMethod = fileAccessMethod;
  }

  public BACnetFileAccessMethodTagged getFileAccessMethod() {
    return fileAccessMethod;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesFileAccessMethod");

    // Simple Field (fileAccessMethod)
    writeSimpleField(
        "fileAccessMethod", fileAccessMethod, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesFileAccessMethod");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesFileAccessMethod _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (fileAccessMethod)
    lengthInBits += fileAccessMethod.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesFileAccessMethod");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetFileAccessMethodTagged fileAccessMethod =
        readSimpleField(
            "fileAccessMethod",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetFileAccessMethodTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesFileAccessMethod");
    // Create the instance
    return new BACnetPropertyStatesFileAccessMethodBuilderImpl(fileAccessMethod);
  }

  public static class BACnetPropertyStatesFileAccessMethodBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetFileAccessMethodTagged fileAccessMethod;

    public BACnetPropertyStatesFileAccessMethodBuilderImpl(
        BACnetFileAccessMethodTagged fileAccessMethod) {
      this.fileAccessMethod = fileAccessMethod;
    }

    public BACnetPropertyStatesFileAccessMethod build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesFileAccessMethod bACnetPropertyStatesFileAccessMethod =
          new BACnetPropertyStatesFileAccessMethod(peekedTagHeader, fileAccessMethod);
      return bACnetPropertyStatesFileAccessMethod;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesFileAccessMethod)) {
      return false;
    }
    BACnetPropertyStatesFileAccessMethod that = (BACnetPropertyStatesFileAccessMethod) o;
    return (getFileAccessMethod() == that.getFileAccessMethod()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFileAccessMethod());
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
