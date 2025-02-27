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

public class UnregisterNodesRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "566";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfNodesToUnregister;
  protected final List<NodeId> nodesToUnregister;

  public UnregisterNodesRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfNodesToUnregister,
      List<NodeId> nodesToUnregister) {
    super();
    this.requestHeader = requestHeader;
    this.noOfNodesToUnregister = noOfNodesToUnregister;
    this.nodesToUnregister = nodesToUnregister;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfNodesToUnregister() {
    return noOfNodesToUnregister;
  }

  public List<NodeId> getNodesToUnregister() {
    return nodesToUnregister;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("UnregisterNodesRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfNodesToUnregister)
    writeSimpleField(
        "noOfNodesToUnregister", noOfNodesToUnregister, writeSignedInt(writeBuffer, 32));

    // Array Field (nodesToUnregister)
    writeComplexTypeArrayField("nodesToUnregister", nodesToUnregister, writeBuffer);

    writeBuffer.popContext("UnregisterNodesRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UnregisterNodesRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfNodesToUnregister)
    lengthInBits += 32;

    // Array field
    if (nodesToUnregister != null) {
      int i = 0;
      for (NodeId element : nodesToUnregister) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= nodesToUnregister.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("UnregisterNodesRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfNodesToUnregister =
        readSimpleField("noOfNodesToUnregister", readSignedInt(readBuffer, 32));

    List<NodeId> nodesToUnregister =
        readCountArrayField(
            "nodesToUnregister",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer),
            noOfNodesToUnregister);

    readBuffer.closeContext("UnregisterNodesRequest");
    // Create the instance
    return new UnregisterNodesRequestBuilderImpl(
        requestHeader, noOfNodesToUnregister, nodesToUnregister);
  }

  public static class UnregisterNodesRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfNodesToUnregister;
    private final List<NodeId> nodesToUnregister;

    public UnregisterNodesRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        int noOfNodesToUnregister,
        List<NodeId> nodesToUnregister) {
      this.requestHeader = requestHeader;
      this.noOfNodesToUnregister = noOfNodesToUnregister;
      this.nodesToUnregister = nodesToUnregister;
    }

    public UnregisterNodesRequest build() {
      UnregisterNodesRequest unregisterNodesRequest =
          new UnregisterNodesRequest(requestHeader, noOfNodesToUnregister, nodesToUnregister);
      return unregisterNodesRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UnregisterNodesRequest)) {
      return false;
    }
    UnregisterNodesRequest that = (UnregisterNodesRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfNodesToUnregister() == that.getNoOfNodesToUnregister())
        && (getNodesToUnregister() == that.getNodesToUnregister())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestHeader(), getNoOfNodesToUnregister(), getNodesToUnregister());
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
