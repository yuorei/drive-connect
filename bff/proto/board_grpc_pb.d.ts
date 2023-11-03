// package: board
// file: board.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as board_pb from "./board_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface IBoardServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createBoard: IBoardServiceService_ICreateBoard;
    readBoard: IBoardServiceService_IReadBoard;
    readAllBoard: IBoardServiceService_IReadAllBoard;
    updateBoard: IBoardServiceService_IUpdateBoard;
    deleteBoard: IBoardServiceService_IDeleteBoard;
    createComment: IBoardServiceService_ICreateComment;
    readComment: IBoardServiceService_IReadComment;
    readAllComment: IBoardServiceService_IReadAllComment;
    updateComment: IBoardServiceService_IUpdateComment;
    deleteComment: IBoardServiceService_IDeleteComment;
}

interface IBoardServiceService_ICreateBoard extends grpc.MethodDefinition<board_pb.Board, board_pb.Board> {
    path: "/board.BoardService/CreateBoard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.Board>;
    requestDeserialize: grpc.deserialize<board_pb.Board>;
    responseSerialize: grpc.serialize<board_pb.Board>;
    responseDeserialize: grpc.deserialize<board_pb.Board>;
}
interface IBoardServiceService_IReadBoard extends grpc.MethodDefinition<board_pb.ID, board_pb.Board> {
    path: "/board.BoardService/ReadBoard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.ID>;
    requestDeserialize: grpc.deserialize<board_pb.ID>;
    responseSerialize: grpc.serialize<board_pb.Board>;
    responseDeserialize: grpc.deserialize<board_pb.Board>;
}
interface IBoardServiceService_IReadAllBoard extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, board_pb.BoardList> {
    path: "/board.BoardService/ReadAllBoard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<board_pb.BoardList>;
    responseDeserialize: grpc.deserialize<board_pb.BoardList>;
}
interface IBoardServiceService_IUpdateBoard extends grpc.MethodDefinition<board_pb.Board, board_pb.Board> {
    path: "/board.BoardService/UpdateBoard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.Board>;
    requestDeserialize: grpc.deserialize<board_pb.Board>;
    responseSerialize: grpc.serialize<board_pb.Board>;
    responseDeserialize: grpc.deserialize<board_pb.Board>;
}
interface IBoardServiceService_IDeleteBoard extends grpc.MethodDefinition<board_pb.ID, google_protobuf_empty_pb.Empty> {
    path: "/board.BoardService/DeleteBoard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.ID>;
    requestDeserialize: grpc.deserialize<board_pb.ID>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IBoardServiceService_ICreateComment extends grpc.MethodDefinition<board_pb.Comment, board_pb.Comment> {
    path: "/board.BoardService/CreateComment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.Comment>;
    requestDeserialize: grpc.deserialize<board_pb.Comment>;
    responseSerialize: grpc.serialize<board_pb.Comment>;
    responseDeserialize: grpc.deserialize<board_pb.Comment>;
}
interface IBoardServiceService_IReadComment extends grpc.MethodDefinition<board_pb.ID, board_pb.Comment> {
    path: "/board.BoardService/ReadComment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.ID>;
    requestDeserialize: grpc.deserialize<board_pb.ID>;
    responseSerialize: grpc.serialize<board_pb.Comment>;
    responseDeserialize: grpc.deserialize<board_pb.Comment>;
}
interface IBoardServiceService_IReadAllComment extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, board_pb.CommentList> {
    path: "/board.BoardService/ReadAllComment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<board_pb.CommentList>;
    responseDeserialize: grpc.deserialize<board_pb.CommentList>;
}
interface IBoardServiceService_IUpdateComment extends grpc.MethodDefinition<board_pb.Comment, board_pb.Comment> {
    path: "/board.BoardService/UpdateComment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.Comment>;
    requestDeserialize: grpc.deserialize<board_pb.Comment>;
    responseSerialize: grpc.serialize<board_pb.Comment>;
    responseDeserialize: grpc.deserialize<board_pb.Comment>;
}
interface IBoardServiceService_IDeleteComment extends grpc.MethodDefinition<board_pb.ID, google_protobuf_empty_pb.Empty> {
    path: "/board.BoardService/DeleteComment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<board_pb.ID>;
    requestDeserialize: grpc.deserialize<board_pb.ID>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const BoardServiceService: IBoardServiceService;

export interface IBoardServiceServer {
    createBoard: grpc.handleUnaryCall<board_pb.Board, board_pb.Board>;
    readBoard: grpc.handleUnaryCall<board_pb.ID, board_pb.Board>;
    readAllBoard: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, board_pb.BoardList>;
    updateBoard: grpc.handleUnaryCall<board_pb.Board, board_pb.Board>;
    deleteBoard: grpc.handleUnaryCall<board_pb.ID, google_protobuf_empty_pb.Empty>;
    createComment: grpc.handleUnaryCall<board_pb.Comment, board_pb.Comment>;
    readComment: grpc.handleUnaryCall<board_pb.ID, board_pb.Comment>;
    readAllComment: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, board_pb.CommentList>;
    updateComment: grpc.handleUnaryCall<board_pb.Comment, board_pb.Comment>;
    deleteComment: grpc.handleUnaryCall<board_pb.ID, google_protobuf_empty_pb.Empty>;
}

export interface IBoardServiceClient {
    createBoard(request: board_pb.Board, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    createBoard(request: board_pb.Board, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    createBoard(request: board_pb.Board, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    readBoard(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    readBoard(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    readBoard(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    readAllBoard(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    readAllBoard(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    readAllBoard(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    updateBoard(request: board_pb.Board, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    updateBoard(request: board_pb.Board, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    updateBoard(request: board_pb.Board, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    deleteBoard(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteBoard(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteBoard(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    createComment(request: board_pb.Comment, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    createComment(request: board_pb.Comment, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    createComment(request: board_pb.Comment, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    readComment(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    readComment(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    readComment(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    readAllComment(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    readAllComment(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    readAllComment(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    updateComment(request: board_pb.Comment, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    updateComment(request: board_pb.Comment, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    updateComment(request: board_pb.Comment, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    deleteComment(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteComment(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteComment(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class BoardServiceClient extends grpc.Client implements IBoardServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createBoard(request: board_pb.Board, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public createBoard(request: board_pb.Board, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public createBoard(request: board_pb.Board, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public readBoard(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public readBoard(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public readBoard(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public readAllBoard(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    public readAllBoard(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    public readAllBoard(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.BoardList) => void): grpc.ClientUnaryCall;
    public updateBoard(request: board_pb.Board, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public updateBoard(request: board_pb.Board, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public updateBoard(request: board_pb.Board, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Board) => void): grpc.ClientUnaryCall;
    public deleteBoard(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteBoard(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteBoard(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public createComment(request: board_pb.Comment, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public createComment(request: board_pb.Comment, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public createComment(request: board_pb.Comment, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public readComment(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public readComment(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public readComment(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public readAllComment(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    public readAllComment(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    public readAllComment(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.CommentList) => void): grpc.ClientUnaryCall;
    public updateComment(request: board_pb.Comment, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public updateComment(request: board_pb.Comment, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public updateComment(request: board_pb.Comment, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: board_pb.Comment) => void): grpc.ClientUnaryCall;
    public deleteComment(request: board_pb.ID, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteComment(request: board_pb.ID, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteComment(request: board_pb.ID, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}
