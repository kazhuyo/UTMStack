// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: agent.proto

package com.park.utmstack.service.grpc;

public interface PingResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:agent.PingResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>bool is_alive = 1;</code>
   * @return The isAlive.
   */
  boolean getIsAlive();

  /**
   * <code>string agent_key = 2;</code>
   * @return The agentKey.
   */
  java.lang.String getAgentKey();
  /**
   * <code>string agent_key = 2;</code>
   * @return The bytes for agentKey.
   */
  com.google.protobuf.ByteString
      getAgentKeyBytes();
}