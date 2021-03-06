# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from epl.protobuf import stac_pb2 as epl_dot_protobuf_dot_stac__pb2


class StacServiceStub(object):
  """
  gRPC Interfaces for working with stac metadata
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Search = channel.unary_stream(
        '/epl.protobuf.StacService/Search',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
        )
    self.Insert = channel.stream_stream(
        '/epl.protobuf.StacService/Insert',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )
    self.Update = channel.stream_stream(
        '/epl.protobuf.StacService/Update',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )
    self.Count = channel.unary_unary(
        '/epl.protobuf.StacService/Count',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )
    self.DeleteOne = channel.unary_unary(
        '/epl.protobuf.StacService/DeleteOne',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )
    self.SearchOne = channel.unary_unary(
        '/epl.protobuf.StacService/SearchOne',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
        )
    self.InsertOne = channel.unary_unary(
        '/epl.protobuf.StacService/InsertOne',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )
    self.UpdateOne = channel.unary_unary(
        '/epl.protobuf.StacService/UpdateOne',
        request_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
        response_deserializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.FromString,
        )


class StacServiceServicer(object):
  """
  gRPC Interfaces for working with stac metadata
  """

  def Search(self, request, context):
    """
    using a search request, stream all the results that match the search filter
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Insert(self, request_iterator, context):
    """
    insert a stream of items into the STAC service
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Update(self, request_iterator, context):
    """
    update a stream of items in the STAC service
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Count(self, request, context):
    """
    count all the items in the Stac service according to the StacRequest filter
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteOne(self, request, context):
    """
    delete an item from the STAC service
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SearchOne(self, request, context):
    """
    using a search request get the first item that matches the request
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def InsertOne(self, request, context):
    """
    Insert one item into the STAC service
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateOne(self, request, context):
    """
    Update one item in the STAC service
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_StacServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Search': grpc.unary_stream_rpc_method_handler(
          servicer.Search,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
      ),
      'Insert': grpc.stream_stream_rpc_method_handler(
          servicer.Insert,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
      'Update': grpc.stream_stream_rpc_method_handler(
          servicer.Update,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
      'Count': grpc.unary_unary_rpc_method_handler(
          servicer.Count,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
      'DeleteOne': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteOne,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
      'SearchOne': grpc.unary_unary_rpc_method_handler(
          servicer.SearchOne,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacRequest.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacItem.SerializeToString,
      ),
      'InsertOne': grpc.unary_unary_rpc_method_handler(
          servicer.InsertOne,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
      'UpdateOne': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateOne,
          request_deserializer=epl_dot_protobuf_dot_stac__pb2.StacItem.FromString,
          response_serializer=epl_dot_protobuf_dot_stac__pb2.StacDbResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'epl.protobuf.StacService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
