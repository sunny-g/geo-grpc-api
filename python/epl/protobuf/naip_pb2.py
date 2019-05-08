# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: epl/protobuf/naip.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from epl.protobuf import query_pb2 as epl_dot_protobuf_dot_query__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='epl/protobuf/naip.proto',
  package='epl.protobuf',
  syntax='proto3',
  serialized_options=_b('\n\020com.epl.protobufB\tNaipProtoP\001Z+github.com/geo-grpc/api/golang/epl/protobuf\242\002\004NPPB'),
  serialized_pb=_b('\n\x17\x65pl/protobuf/naip.proto\x12\x0c\x65pl.protobuf\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18\x65pl/protobuf/query.proto\"R\n\x0eNaipProperties\x12\x30\n\x0csrc_img_date\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06usgsid\x18\x02 \x01(\t\"n\n\x0bNaipRequest\x12\x34\n\x0esrc_image_date\x18\x01 \x01(\x0b\x32\x1c.epl.protobuf.TimestampField\x12)\n\x06usgsid\x18\x02 \x01(\x0b\x32\x19.epl.protobuf.StringFieldBS\n\x10\x63om.epl.protobufB\tNaipProtoP\x01Z+github.com/geo-grpc/api/golang/epl/protobuf\xa2\x02\x04NPPBb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,epl_dot_protobuf_dot_query__pb2.DESCRIPTOR,])




_NAIPPROPERTIES = _descriptor.Descriptor(
  name='NaipProperties',
  full_name='epl.protobuf.NaipProperties',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='src_img_date', full_name='epl.protobuf.NaipProperties.src_img_date', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='usgsid', full_name='epl.protobuf.NaipProperties.usgsid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=100,
  serialized_end=182,
)


_NAIPREQUEST = _descriptor.Descriptor(
  name='NaipRequest',
  full_name='epl.protobuf.NaipRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='src_image_date', full_name='epl.protobuf.NaipRequest.src_image_date', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='usgsid', full_name='epl.protobuf.NaipRequest.usgsid', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=184,
  serialized_end=294,
)

_NAIPPROPERTIES.fields_by_name['src_img_date'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NAIPREQUEST.fields_by_name['src_image_date'].message_type = epl_dot_protobuf_dot_query__pb2._TIMESTAMPFIELD
_NAIPREQUEST.fields_by_name['usgsid'].message_type = epl_dot_protobuf_dot_query__pb2._STRINGFIELD
DESCRIPTOR.message_types_by_name['NaipProperties'] = _NAIPPROPERTIES
DESCRIPTOR.message_types_by_name['NaipRequest'] = _NAIPREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NaipProperties = _reflection.GeneratedProtocolMessageType('NaipProperties', (_message.Message,), dict(
  DESCRIPTOR = _NAIPPROPERTIES,
  __module__ = 'epl.protobuf.naip_pb2'
  # @@protoc_insertion_point(class_scope:epl.protobuf.NaipProperties)
  ))
_sym_db.RegisterMessage(NaipProperties)

NaipRequest = _reflection.GeneratedProtocolMessageType('NaipRequest', (_message.Message,), dict(
  DESCRIPTOR = _NAIPREQUEST,
  __module__ = 'epl.protobuf.naip_pb2'
  # @@protoc_insertion_point(class_scope:epl.protobuf.NaipRequest)
  ))
_sym_db.RegisterMessage(NaipRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)