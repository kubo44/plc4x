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

public class BACnetEventParameterChangeOfCharacterString extends BACnetEventParameter
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger timeDelay;
  protected final BACnetEventParameterChangeOfCharacterStringListOfAlarmValues listOfAlarmValues;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterChangeOfCharacterString(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger timeDelay,
      BACnetEventParameterChangeOfCharacterStringListOfAlarmValues listOfAlarmValues,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.timeDelay = timeDelay;
    this.listOfAlarmValues = listOfAlarmValues;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getTimeDelay() {
    return timeDelay;
  }

  public BACnetEventParameterChangeOfCharacterStringListOfAlarmValues getListOfAlarmValues() {
    return listOfAlarmValues;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventParameterChangeOfCharacterString");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (listOfAlarmValues)
    writeSimpleField(
        "listOfAlarmValues", listOfAlarmValues, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterChangeOfCharacterString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterChangeOfCharacterString _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (listOfAlarmValues)
    lengthInBits += listOfAlarmValues.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterChangeOfCharacterString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (17)), readBuffer));

    BACnetContextTagUnsignedInteger timeDelay =
        readSimpleField(
            "timeDelay",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetEventParameterChangeOfCharacterStringListOfAlarmValues listOfAlarmValues =
        readSimpleField(
            "listOfAlarmValues",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetEventParameterChangeOfCharacterStringListOfAlarmValues.staticParse(
                        readBuffer, (short) (1)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (17)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterChangeOfCharacterString");
    // Create the instance
    return new BACnetEventParameterChangeOfCharacterStringBuilderImpl(
        openingTag, timeDelay, listOfAlarmValues, closingTag);
  }

  public static class BACnetEventParameterChangeOfCharacterStringBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetContextTagUnsignedInteger timeDelay;
    private final BACnetEventParameterChangeOfCharacterStringListOfAlarmValues listOfAlarmValues;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterChangeOfCharacterStringBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetContextTagUnsignedInteger timeDelay,
        BACnetEventParameterChangeOfCharacterStringListOfAlarmValues listOfAlarmValues,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.timeDelay = timeDelay;
      this.listOfAlarmValues = listOfAlarmValues;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterChangeOfCharacterString build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterChangeOfCharacterString bACnetEventParameterChangeOfCharacterString =
          new BACnetEventParameterChangeOfCharacterString(
              peekedTagHeader, openingTag, timeDelay, listOfAlarmValues, closingTag);
      return bACnetEventParameterChangeOfCharacterString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterChangeOfCharacterString)) {
      return false;
    }
    BACnetEventParameterChangeOfCharacterString that =
        (BACnetEventParameterChangeOfCharacterString) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getTimeDelay() == that.getTimeDelay())
        && (getListOfAlarmValues() == that.getListOfAlarmValues())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getOpeningTag(), getTimeDelay(), getListOfAlarmValues(), getClosingTag());
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
