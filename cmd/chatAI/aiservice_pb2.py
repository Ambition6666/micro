# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: aiservice.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x61iservice.proto\x12\x04\x63hat\"M\n\x08\x42\x61seResp\x12\x13\n\x0bstatus_code\x18\x01 \x01(\x03\x12\x16\n\x0estatus_message\x18\x02 \x01(\t\x12\x14\n\x0cservice_time\x18\x03 \x01(\x03\"*\n\x0b\x43hatRequest\x12\x0b\n\x03msg\x18\x01 \x01(\t\x12\x0e\n\x06userID\x18\x02 \x01(\x03\"9\n\x0c\x43hatResponse\x12\x0b\n\x03msg\x18\x01 \x01(\t\x12\x1c\n\x04resp\x18\x02 \x01(\x0b\x32\x0e.chat.BaseResp2>\n\x0b\x43hatService\x12/\n\x04\x43hat\x12\x11.chat.ChatRequest\x1a\x12.chat.ChatResponse\"\x00\x42\x13Z\x11./chatwithbotdemob\x06proto3')



_BASERESP = DESCRIPTOR.message_types_by_name['BaseResp']
_CHATREQUEST = DESCRIPTOR.message_types_by_name['ChatRequest']
_CHATRESPONSE = DESCRIPTOR.message_types_by_name['ChatResponse']
BaseResp = _reflection.GeneratedProtocolMessageType('BaseResp', (_message.Message,), {
  'DESCRIPTOR' : _BASERESP,
  '__module__' : 'aiservice_pb2'
  # @@protoc_insertion_point(class_scope:chat.BaseResp)
  })
_sym_db.RegisterMessage(BaseResp)

ChatRequest = _reflection.GeneratedProtocolMessageType('ChatRequest', (_message.Message,), {
  'DESCRIPTOR' : _CHATREQUEST,
  '__module__' : 'aiservice_pb2'
  # @@protoc_insertion_point(class_scope:chat.ChatRequest)
  })
_sym_db.RegisterMessage(ChatRequest)

ChatResponse = _reflection.GeneratedProtocolMessageType('ChatResponse', (_message.Message,), {
  'DESCRIPTOR' : _CHATRESPONSE,
  '__module__' : 'aiservice_pb2'
  # @@protoc_insertion_point(class_scope:chat.ChatResponse)
  })
_sym_db.RegisterMessage(ChatResponse)

_CHATSERVICE = DESCRIPTOR.services_by_name['ChatService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\021./chatwithbotdemo'
  _BASERESP._serialized_start=25
  _BASERESP._serialized_end=102
  _CHATREQUEST._serialized_start=104
  _CHATREQUEST._serialized_end=146
  _CHATRESPONSE._serialized_start=148
  _CHATRESPONSE._serialized_end=205
  _CHATSERVICE._serialized_start=207
  _CHATSERVICE._serialized_end=269
# @@protoc_insertion_point(module_scope)