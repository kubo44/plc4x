/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.base.connection;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandler;
import io.netty.util.HashedWheelTimer;
import io.netty.util.Timer;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcException;
import org.apache.plc4x.java.api.exceptions.PlcIoException;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

public abstract class NettyPlcConnection extends AbstractPlcConnection {

    /**
     * a {@link HashedWheelTimer} shall be only instantiated once.
     */
    // TODO: maybe find a way to make this configurable per jvm
    protected final static Timer timer = new HashedWheelTimer();

    protected final ChannelFactory channelFactory;

    protected final boolean awaitSessionSetupComplete;

    protected Channel channel;

    protected boolean connected;

    protected NettyPlcConnection(ChannelFactory channelFactory) {
        this(channelFactory, false);
    }

    protected NettyPlcConnection(ChannelFactory channelFactory, boolean awaitSessionSetupComplete) {
        this.channelFactory = channelFactory;
        this.awaitSessionSetupComplete = awaitSessionSetupComplete;
        this.connected = false;
    }

    @Override
    public void connect() throws PlcConnectionException {
        try {
            // As we don't just want to wait till the connection is established,
            // define a future we can use to signal back that the s7 session is
            // finished initializing.
            CompletableFuture<Void> sessionSetupCompleteFuture = new CompletableFuture<>();

            // Have the channel factory create a new channel instance.
            channel = channelFactory.createChannel(getChannelHandler(sessionSetupCompleteFuture));
            channel.closeFuture().addListener(future -> {
                if (!sessionSetupCompleteFuture.isDone()) {
                    sessionSetupCompleteFuture.completeExceptionally(
                        new PlcIoException("Connection terminated by remote"));
                }
            });
            // Send an event to the pipeline telling the Protocol filters what's going on.
            sendChannelCreatedEvent();

            // Wait till the connection is established.
            if (awaitSessionSetupComplete) {
                sessionSetupCompleteFuture.get();
            }

            // Set the connection to "connected"
            connected = true;
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new PlcConnectionException(e);
        } catch (ExecutionException e) {
            throw new PlcConnectionException(e);
        }
    }

    @Override
    public CompletableFuture<Void> ping() {
        CompletableFuture<Void> future = new CompletableFuture<>();
        try {
            // Relay the actual pinging to the channel factory ...
            channelFactory.ping();
            // If we got here, the ping was successful.
            future.complete(null);
        } catch(PlcException e) {
            // If we got here, something went wrong.
            future.completeExceptionally(e);
        }
        return future;
    }

    @Override
    public void close() throws PlcConnectionException {
        channel = null;
        connected = false;
    }

    /**
     * Check if the communication channel is active (channel.isActive()) and the driver for a given protocol
     * has finished establishing the connection.
     */
    @Override
    public boolean isConnected() {
        return connected && channel.isActive();
    }

    public Channel getChannel() {
        return channel;
    }

    protected abstract ChannelHandler getChannelHandler(CompletableFuture<Void> sessionSetupCompleteFuture);

    protected void sendChannelCreatedEvent() {
        // Implemented in sub-classes, if needed.
    }

}
