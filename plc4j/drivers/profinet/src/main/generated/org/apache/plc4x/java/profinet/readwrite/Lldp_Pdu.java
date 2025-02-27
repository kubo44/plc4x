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
package org.apache.plc4x.java.profinet.readwrite;

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

public class Lldp_Pdu implements Message {

  // Properties.
  protected final List<LldpUnit> lldpParameters;

  public Lldp_Pdu(List<LldpUnit> lldpParameters) {
    super();
    this.lldpParameters = lldpParameters;
  }

  public List<LldpUnit> getLldpParameters() {
    return lldpParameters;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Lldp_Pdu");

    // Manual Array Field (lldpParameters)
    writeManualArrayField(
        "lldpParameters",
        lldpParameters,
        (LldpUnit _value) ->
            org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.serializeSysexString(
                writeBuffer, _value),
        writeBuffer);

    writeBuffer.popContext("Lldp_Pdu");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Lldp_Pdu _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Manual Array Field (lldpParameters)
    lengthInBits +=
        org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.lengthSysexString(
                lldpParameters)
            * 8;

    return lengthInBits;
  }

  public static Lldp_Pdu staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static Lldp_Pdu staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Lldp_Pdu");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<LldpUnit> lldpParameters =
        readManualArrayField(
            "lldpParameters",
            readBuffer,
            (List<LldpUnit> _values) ->
                (boolean)
                    (org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.isSysexEnd(
                        readBuffer)),
            () ->
                (LldpUnit)
                    (org.apache.plc4x.java.profinet.readwrite.utils.StaticHelper.parseSysexString(
                        readBuffer)));

    readBuffer.closeContext("Lldp_Pdu");
    // Create the instance
    Lldp_Pdu _lldp_Pdu;
    _lldp_Pdu = new Lldp_Pdu(lldpParameters);
    return _lldp_Pdu;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Lldp_Pdu)) {
      return false;
    }
    Lldp_Pdu that = (Lldp_Pdu) o;
    return (getLldpParameters() == that.getLldpParameters()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getLldpParameters());
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
