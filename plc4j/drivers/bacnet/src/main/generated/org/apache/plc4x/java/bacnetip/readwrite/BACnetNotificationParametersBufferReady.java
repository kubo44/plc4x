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

public class BACnetNotificationParametersBufferReady extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetDeviceObjectPropertyReferenceEnclosed bufferProperty;
  protected final BACnetContextTagUnsignedInteger previousNotification;
  protected final BACnetContextTagUnsignedInteger currentNotification;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersBufferReady(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetDeviceObjectPropertyReferenceEnclosed bufferProperty,
      BACnetContextTagUnsignedInteger previousNotification,
      BACnetContextTagUnsignedInteger currentNotification,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.bufferProperty = bufferProperty;
    this.previousNotification = previousNotification;
    this.currentNotification = currentNotification;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetDeviceObjectPropertyReferenceEnclosed getBufferProperty() {
    return bufferProperty;
  }

  public BACnetContextTagUnsignedInteger getPreviousNotification() {
    return previousNotification;
  }

  public BACnetContextTagUnsignedInteger getCurrentNotification() {
    return currentNotification;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersBufferReady");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (bufferProperty)
    writeSimpleField("bufferProperty", bufferProperty, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (previousNotification)
    writeSimpleField(
        "previousNotification", previousNotification, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (currentNotification)
    writeSimpleField(
        "currentNotification", currentNotification, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersBufferReady");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersBufferReady _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (bufferProperty)
    lengthInBits += bufferProperty.getLengthInBits();

    // Simple field (previousNotification)
    lengthInBits += previousNotification.getLengthInBits();

    // Simple field (currentNotification)
    lengthInBits += currentNotification.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersBuilder staticParseBACnetNotificationParametersBuilder(
      ReadBuffer readBuffer,
      Short peekedTagNumber,
      Short tagNumber,
      BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersBufferReady");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetDeviceObjectPropertyReferenceEnclosed bufferProperty =
        readSimpleField(
            "bufferProperty",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetDeviceObjectPropertyReferenceEnclosed.staticParse(
                        readBuffer, (short) (0)),
                readBuffer));

    BACnetContextTagUnsignedInteger previousNotification =
        readSimpleField(
            "previousNotification",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger currentNotification =
        readSimpleField(
            "currentNotification",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersBufferReady");
    // Create the instance
    return new BACnetNotificationParametersBufferReadyBuilderImpl(
        innerOpeningTag,
        bufferProperty,
        previousNotification,
        currentNotification,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersBufferReadyBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetDeviceObjectPropertyReferenceEnclosed bufferProperty;
    private final BACnetContextTagUnsignedInteger previousNotification;
    private final BACnetContextTagUnsignedInteger currentNotification;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersBufferReadyBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetDeviceObjectPropertyReferenceEnclosed bufferProperty,
        BACnetContextTagUnsignedInteger previousNotification,
        BACnetContextTagUnsignedInteger currentNotification,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.bufferProperty = bufferProperty;
      this.previousNotification = previousNotification;
      this.currentNotification = currentNotification;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersBufferReady build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersBufferReady bACnetNotificationParametersBufferReady =
          new BACnetNotificationParametersBufferReady(
              openingTag,
              peekedTagHeader,
              closingTag,
              innerOpeningTag,
              bufferProperty,
              previousNotification,
              currentNotification,
              innerClosingTag,
              tagNumber,
              objectTypeArgument);
      return bACnetNotificationParametersBufferReady;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersBufferReady)) {
      return false;
    }
    BACnetNotificationParametersBufferReady that = (BACnetNotificationParametersBufferReady) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getBufferProperty() == that.getBufferProperty())
        && (getPreviousNotification() == that.getPreviousNotification())
        && (getCurrentNotification() == that.getCurrentNotification())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getBufferProperty(),
        getPreviousNotification(),
        getCurrentNotification(),
        getInnerClosingTag());
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
