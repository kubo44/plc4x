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

public class ContentFilterResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "609";
  }

  // Properties.
  protected final int noOfElementResults;
  protected final List<ExtensionObjectDefinition> elementResults;
  protected final int noOfElementDiagnosticInfos;
  protected final List<DiagnosticInfo> elementDiagnosticInfos;

  public ContentFilterResult(
      int noOfElementResults,
      List<ExtensionObjectDefinition> elementResults,
      int noOfElementDiagnosticInfos,
      List<DiagnosticInfo> elementDiagnosticInfos) {
    super();
    this.noOfElementResults = noOfElementResults;
    this.elementResults = elementResults;
    this.noOfElementDiagnosticInfos = noOfElementDiagnosticInfos;
    this.elementDiagnosticInfos = elementDiagnosticInfos;
  }

  public int getNoOfElementResults() {
    return noOfElementResults;
  }

  public List<ExtensionObjectDefinition> getElementResults() {
    return elementResults;
  }

  public int getNoOfElementDiagnosticInfos() {
    return noOfElementDiagnosticInfos;
  }

  public List<DiagnosticInfo> getElementDiagnosticInfos() {
    return elementDiagnosticInfos;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ContentFilterResult");

    // Simple Field (noOfElementResults)
    writeSimpleField("noOfElementResults", noOfElementResults, writeSignedInt(writeBuffer, 32));

    // Array Field (elementResults)
    writeComplexTypeArrayField("elementResults", elementResults, writeBuffer);

    // Simple Field (noOfElementDiagnosticInfos)
    writeSimpleField(
        "noOfElementDiagnosticInfos", noOfElementDiagnosticInfos, writeSignedInt(writeBuffer, 32));

    // Array Field (elementDiagnosticInfos)
    writeComplexTypeArrayField("elementDiagnosticInfos", elementDiagnosticInfos, writeBuffer);

    writeBuffer.popContext("ContentFilterResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ContentFilterResult _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (noOfElementResults)
    lengthInBits += 32;

    // Array field
    if (elementResults != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : elementResults) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= elementResults.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfElementDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (elementDiagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : elementDiagnosticInfos) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= elementDiagnosticInfos.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ContentFilterResult");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfElementResults = readSimpleField("noOfElementResults", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> elementResults =
        readCountArrayField(
            "elementResults",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("606")),
                readBuffer),
            noOfElementResults);

    int noOfElementDiagnosticInfos =
        readSimpleField("noOfElementDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> elementDiagnosticInfos =
        readCountArrayField(
            "elementDiagnosticInfos",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfElementDiagnosticInfos);

    readBuffer.closeContext("ContentFilterResult");
    // Create the instance
    return new ContentFilterResultBuilderImpl(
        noOfElementResults, elementResults, noOfElementDiagnosticInfos, elementDiagnosticInfos);
  }

  public static class ContentFilterResultBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfElementResults;
    private final List<ExtensionObjectDefinition> elementResults;
    private final int noOfElementDiagnosticInfos;
    private final List<DiagnosticInfo> elementDiagnosticInfos;

    public ContentFilterResultBuilderImpl(
        int noOfElementResults,
        List<ExtensionObjectDefinition> elementResults,
        int noOfElementDiagnosticInfos,
        List<DiagnosticInfo> elementDiagnosticInfos) {
      this.noOfElementResults = noOfElementResults;
      this.elementResults = elementResults;
      this.noOfElementDiagnosticInfos = noOfElementDiagnosticInfos;
      this.elementDiagnosticInfos = elementDiagnosticInfos;
    }

    public ContentFilterResult build() {
      ContentFilterResult contentFilterResult =
          new ContentFilterResult(
              noOfElementResults,
              elementResults,
              noOfElementDiagnosticInfos,
              elementDiagnosticInfos);
      return contentFilterResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ContentFilterResult)) {
      return false;
    }
    ContentFilterResult that = (ContentFilterResult) o;
    return (getNoOfElementResults() == that.getNoOfElementResults())
        && (getElementResults() == that.getElementResults())
        && (getNoOfElementDiagnosticInfos() == that.getNoOfElementDiagnosticInfos())
        && (getElementDiagnosticInfos() == that.getElementDiagnosticInfos())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNoOfElementResults(),
        getElementResults(),
        getNoOfElementDiagnosticInfos(),
        getElementDiagnosticInfos());
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
