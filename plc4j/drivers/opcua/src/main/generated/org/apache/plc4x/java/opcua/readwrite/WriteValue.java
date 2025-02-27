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
package org.apache.plc4x.java.opcua.readwrite;

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

public class WriteValue extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "670";
  }

  // Properties.
  protected final NodeId nodeId;
  protected final long attributeId;
  protected final PascalString indexRange;
  protected final DataValue value;

  public WriteValue(NodeId nodeId, long attributeId, PascalString indexRange, DataValue value) {
    super();
    this.nodeId = nodeId;
    this.attributeId = attributeId;
    this.indexRange = indexRange;
    this.value = value;
  }

  public NodeId getNodeId() {
    return nodeId;
  }

  public long getAttributeId() {
    return attributeId;
  }

  public PascalString getIndexRange() {
    return indexRange;
  }

  public DataValue getValue() {
    return value;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("WriteValue");

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (attributeId)
    writeSimpleField("attributeId", attributeId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (indexRange)
    writeSimpleField("indexRange", indexRange, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (value)
    writeSimpleField("value", value, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("WriteValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    WriteValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Simple field (attributeId)
    lengthInBits += 32;

    // Simple field (indexRange)
    lengthInBits += indexRange.getLengthInBits();

    // Simple field (value)
    lengthInBits += value.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("WriteValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId nodeId =
        readSimpleField(
            "nodeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    long attributeId = readSimpleField("attributeId", readUnsignedLong(readBuffer, 32));

    PascalString indexRange =
        readSimpleField(
            "indexRange",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    DataValue value =
        readSimpleField(
            "value",
            new DataReaderComplexDefault<>(() -> DataValue.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("WriteValue");
    // Create the instance
    return new WriteValueBuilderImpl(nodeId, attributeId, indexRange, value);
  }

  public static class WriteValueBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId nodeId;
    private final long attributeId;
    private final PascalString indexRange;
    private final DataValue value;

    public WriteValueBuilderImpl(
        NodeId nodeId, long attributeId, PascalString indexRange, DataValue value) {
      this.nodeId = nodeId;
      this.attributeId = attributeId;
      this.indexRange = indexRange;
      this.value = value;
    }

    public WriteValue build() {
      WriteValue writeValue = new WriteValue(nodeId, attributeId, indexRange, value);
      return writeValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof WriteValue)) {
      return false;
    }
    WriteValue that = (WriteValue) o;
    return (getNodeId() == that.getNodeId())
        && (getAttributeId() == that.getAttributeId())
        && (getIndexRange() == that.getIndexRange())
        && (getValue() == that.getValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getNodeId(), getAttributeId(), getIndexRange(), getValue());
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
