// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var board_pb = require('./board_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_board_Board(arg) {
  if (!(arg instanceof board_pb.Board)) {
    throw new Error('Expected argument of type board.Board');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_board_Board(buffer_arg) {
  return board_pb.Board.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_board_BoardList(arg) {
  if (!(arg instanceof board_pb.BoardList)) {
    throw new Error('Expected argument of type board.BoardList');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_board_BoardList(buffer_arg) {
  return board_pb.BoardList.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_board_Comment(arg) {
  if (!(arg instanceof board_pb.Comment)) {
    throw new Error('Expected argument of type board.Comment');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_board_Comment(buffer_arg) {
  return board_pb.Comment.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_board_CommentList(arg) {
  if (!(arg instanceof board_pb.CommentList)) {
    throw new Error('Expected argument of type board.CommentList');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_board_CommentList(buffer_arg) {
  return board_pb.CommentList.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_board_ID(arg) {
  if (!(arg instanceof board_pb.ID)) {
    throw new Error('Expected argument of type board.ID');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_board_ID(buffer_arg) {
  return board_pb.ID.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


var BoardServiceService = exports.BoardServiceService = {
  createBoard: {
    path: '/board.BoardService/CreateBoard',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.Board,
    responseType: board_pb.Board,
    requestSerialize: serialize_board_Board,
    requestDeserialize: deserialize_board_Board,
    responseSerialize: serialize_board_Board,
    responseDeserialize: deserialize_board_Board,
  },
  readBoard: {
    path: '/board.BoardService/ReadBoard',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.ID,
    responseType: board_pb.Board,
    requestSerialize: serialize_board_ID,
    requestDeserialize: deserialize_board_ID,
    responseSerialize: serialize_board_Board,
    responseDeserialize: deserialize_board_Board,
  },
  readAllBoard: {
    path: '/board.BoardService/ReadAllBoard',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: board_pb.BoardList,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_board_BoardList,
    responseDeserialize: deserialize_board_BoardList,
  },
  updateBoard: {
    path: '/board.BoardService/UpdateBoard',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.Board,
    responseType: board_pb.Board,
    requestSerialize: serialize_board_Board,
    requestDeserialize: deserialize_board_Board,
    responseSerialize: serialize_board_Board,
    responseDeserialize: deserialize_board_Board,
  },
  deleteBoard: {
    path: '/board.BoardService/DeleteBoard',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.ID,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_board_ID,
    requestDeserialize: deserialize_board_ID,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  createComment: {
    path: '/board.BoardService/CreateComment',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.Comment,
    responseType: board_pb.Comment,
    requestSerialize: serialize_board_Comment,
    requestDeserialize: deserialize_board_Comment,
    responseSerialize: serialize_board_Comment,
    responseDeserialize: deserialize_board_Comment,
  },
  readComment: {
    path: '/board.BoardService/ReadComment',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.ID,
    responseType: board_pb.Comment,
    requestSerialize: serialize_board_ID,
    requestDeserialize: deserialize_board_ID,
    responseSerialize: serialize_board_Comment,
    responseDeserialize: deserialize_board_Comment,
  },
  readAllComment: {
    path: '/board.BoardService/ReadAllComment',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: board_pb.CommentList,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_board_CommentList,
    responseDeserialize: deserialize_board_CommentList,
  },
  updateComment: {
    path: '/board.BoardService/UpdateComment',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.Comment,
    responseType: board_pb.Comment,
    requestSerialize: serialize_board_Comment,
    requestDeserialize: deserialize_board_Comment,
    responseSerialize: serialize_board_Comment,
    responseDeserialize: deserialize_board_Comment,
  },
  deleteComment: {
    path: '/board.BoardService/DeleteComment',
    requestStream: false,
    responseStream: false,
    requestType: board_pb.ID,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_board_ID,
    requestDeserialize: deserialize_board_ID,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.BoardServiceClient = grpc.makeGenericClientConstructor(BoardServiceService);
