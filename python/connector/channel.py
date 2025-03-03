# Copyright 2024 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""An example showing how to call the MMv2 go library from Python.

This is a simple Python app that show how the MMv2 go library can
use gRPC to communicate with other languages via a minimal C shim.
"""



import grpc

from proto.connector import connector_pb2
from connector import connector_ext


class Channel(grpc.Channel):
  """Implementation of grpc.Channel for communicating with the g library.

  Mostly unimplemented since all that is necessary is unary calls.
  """

  def subscribe(self, callback, try_to_connect=False):
    """Subscribe to this Channel's connectivity state machine.

    A Channel may be in any of the states described by ChannelConnectivity.
    This method allows application to monitor the state transitions.
    The typical use case is to debug or gain better visibility into gRPC
    runtime's state.
    Args:
      callback: A callable to be invoked with ChannelConnectivity argument.
        ChannelConnectivity describes current state of the channel. The callable
        will be invoked immediately upon subscription and again for every change
        to ChannelConnectivity until it is unsubscribed or this Channel object
        goes out of scope.
      try_to_connect: A boolean indicating whether or not this Channel should
        attempt to connect immediately. If set to False, gRPC runtime decides
        when to connect.
    """
    raise NotImplementedError()

  def unsubscribe(self, callback):
    """Unsubscribes a subscribed callback from this Channel's connectivity.

    Args:
      callback: A callable previously registered with this Channel from having
        been passed to its "subscribe" method.
    """
    raise NotImplementedError()

  def unary_unary(self,
                  method,
                  request_serializer=None,
                  response_deserializer=None):
    """Creates a UnaryUnaryMultiCallable for a unary-unary method.

    Args:
      method: The name of the RPC method.
      request_serializer: Optional behaviour for serializing the request
        message. Request goes unserialized in case None is passed.
      response_deserializer: Optional behaviour for deserializing the response
        message. Response goes undeserialized in case None is passed.

    Returns:
      A UnaryUnaryMultiCallable value for the named unary-unary method.
    """
    return UnaryUnaryMultiCallable(method, request_serializer,
                                   response_deserializer)

  def unary_stream(self,
                   method,
                   request_serializer=None,
                   response_deserializer=None):
    """Creates a UnaryStreamMultiCallable for a unary-stream method.

    Args:
      method: The name of the RPC method.
      request_serializer: Optional behaviour for serializing the request
        message. Request goes unserialized in case None is passed.
      response_deserializer: Optional behaviour for deserializing the response
        message. Response goes undeserialized in case None is passed.

    Returns:
      A UnaryStreamMultiCallable value for the name unary-stream method.
    """
    raise NotImplementedError()

  def stream_unary(self,
                   method,
                   request_serializer=None,
                   response_deserializer=None):
    """Creates a StreamUnaryMultiCallable for a stream-unary method.

    Args:
      method: The name of the RPC method.
      request_serializer: Optional behaviour for serializing the request
        message. Request goes unserialized in case None is passed.
      response_deserializer: Optional behaviour for deserializing the response
        message. Response goes undeserialized in case None is passed.

    Returns:
      A StreamUnaryMultiCallable value for the named stream-unary method.
    """
    raise NotImplementedError()

  def stream_stream(self,
                    method,
                    request_serializer=None,
                    response_deserializer=None):
    """Creates a StreamStreamMultiCallable for a stream-stream method.

    Args:
      method: The name of the RPC method.
      request_serializer: Optional behaviour for serializing the request
        message. Request goes unserialized in case None is passed.
      response_deserializer: Optional behaviour for deserializing the response
        message. Response goes undeserialized in case None is passed.

    Returns:
      A StreamStreamMultiCallable value for the named stream-stream method.
    """
    raise NotImplementedError()

  def close(self):
    """Closes this Channel and releases all resources held by it."""


class UnaryUnaryMultiCallable(grpc.UnaryUnaryMultiCallable):
  """Invoke a unary-unary RPC to the go library."""

  def __init__(self,
               method,
               request_serializer=None,
               response_deserializer=None):
    self.method = method
    self.request_serializer = request_serializer
    self.response_deserializer = response_deserializer

  def __call__(self,
               request,
               timeout=None,
               metadata=None,
               credentials=None,
               wait_for_ready=None,
               compression=None):
    """Synchronously invokes the underlying RPC.

    Args:
      request: The request value for the RPC.
      timeout: An optional duration of time in seconds to allow for the RPC.
      metadata: Optional :term:`metadata` to be transmitted to the service-side
        of the RPC.
      credentials: An optional CallCredentials for the RPC. Only valid for
        secure Channel.
      wait_for_ready: This is an EXPERIMENTAL argument. An optional flag to
        enable wait for ready mechanism
      compression: An element of grpc.compression, e.g. grpc.compression.Gzip.
        This is an EXPERIMENTAL option.

    Returns:
      The response value for the RPC.
    Raises:
      RpcError: Indicating that the RPC terminated with non-OK status. The
        raised RpcError will also be a Call for the RPC affording the RPC's
        metadata, status code, and details.
    """
    call_request = connector_pb2.UnaryCallRequest()
    call_request.method = self.method
    if self.request_serializer:
      call_request.request = self.request_serializer(request)
    else:
      call_request.request = request
    call_response_str = connector_ext.Connector.Call(
        call_request.SerializeToString())
    call_response = connector_pb2.UnaryCallResponse()
    call_response.ParseFromString(call_response_str)

    if call_response.status.code:
      raise grpc.RpcError("call failed: code %d (%s)'" %
                          (call_response.status.code,
                           call_response.status.message))

    if self.response_deserializer:
      return self.response_deserializer(call_response.response)
    return call_response.response

  def with_call(self,
                request,
                timeout=None,
                metadata=None,
                credentials=None,
                wait_for_ready=None,
                compression=None):
    """Synchronously invokes the underlying RPC.

    Args:
      request: The request value for the RPC.
      timeout: An optional durating of time in seconds to allow for the RPC.
      metadata: Optional :term:`metadata` to be transmitted to the service-side
        of the RPC.
      credentials: An optional CallCredentials for the RPC. Only valid for
        secure Channel.
      wait_for_ready: This is an EXPERIMENTAL argument. An optional flag to
        enable wait for ready mechanism
      compression: An element of grpc.compression, e.g. grpc.compression.Gzip.
        This is an EXPERIMENTAL option.

    Returns:
      The response value for the RPC and a Call value for the RPC.
    Raises:
      RpcError: Indicating that the RPC terminated with non-OK status. The
        raised RpcError will also be a Call for the RPC affording the RPC's
        metadata, status code, and details.
    """
    raise NotImplementedError()

  def future(self,
             request,
             timeout=None,
             metadata=None,
             credentials=None,
             wait_for_ready=None,
             compression=None):
    """Asynchronously invokes the underlying RPC.

    Args:
      request: The request value for the RPC.
      timeout: An optional duration of time in seconds to allow for the RPC.
      metadata: Optional :term:`metadata` to be transmitted to the service-side
        of the RPC.
      credentials: An optional CallCredentials for the RPC. Only valid for
        secure Channel.
      wait_for_ready: This is an EXPERIMENTAL argument. An optional flag to
        enable wait for ready mechanism
      compression: An element of grpc.compression, e.g. grpc.compression.Gzip.
        This is an EXPERIMENTAL option.

    Returns:
        An object that is both a Call for the RPC and a Future.
        In the event of RPC completion, the return Call-Future's result
        value will be the response message of the RPC.
        Should the event terminate with non-OK status,
        the returned Call-Future's exception value will be an RpcError.
    """
    raise NotImplementedError()


def initialize():
  request = connector_pb2.InitializeRequest()
  response_str = connector_ext.Connector.Initialize(request.SerializeToString())
  response = connector_pb2.InitializeResponse()
  response.ParseFromString(response_str)
  if response.status.code:
    raise grpc.RpcError('initialization failed: code %d (%s)' %
                        (response.status.code, response.status.message))
