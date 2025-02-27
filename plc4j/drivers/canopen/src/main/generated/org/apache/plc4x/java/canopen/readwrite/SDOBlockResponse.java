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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOBlockResponse extends SDOResponse implements Message {

  // Accessors for discriminator values.
  public SDOResponseCommand getCommand() {
    return SDOResponseCommand.BLOCK;
  }

  // Properties.
  protected final SDOBlockData block;

  public SDOBlockResponse(SDOBlockData block) {
    super();
    this.block = block;
  }

  public SDOBlockData getBlock() {
    return block;
  }

  @Override
  protected void serializeSDOResponseChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SDOBlockResponse");

    // Simple Field (block)
    writeSimpleField("block", block, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SDOBlockResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SDOBlockResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (block)
    lengthInBits += block.getLengthInBits();

    return lengthInBits;
  }

  public static SDOResponseBuilder staticParseSDOResponseBuilder(
      ReadBuffer readBuffer, SDOResponseCommand command) throws ParseException {
    readBuffer.pullContext("SDOBlockResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SDOBlockData block =
        readSimpleField(
            "block",
            new DataReaderComplexDefault<>(() -> SDOBlockData.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SDOBlockResponse");
    // Create the instance
    return new SDOBlockResponseBuilderImpl(block);
  }

  public static class SDOBlockResponseBuilderImpl implements SDOResponse.SDOResponseBuilder {
    private final SDOBlockData block;

    public SDOBlockResponseBuilderImpl(SDOBlockData block) {
      this.block = block;
    }

    public SDOBlockResponse build() {
      SDOBlockResponse sDOBlockResponse = new SDOBlockResponse(block);
      return sDOBlockResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOBlockResponse)) {
      return false;
    }
    SDOBlockResponse that = (SDOBlockResponse) o;
    return (getBlock() == that.getBlock()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBlock());
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
