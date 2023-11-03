// package: board
// file: board.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class Board extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getType(): string;
  setType(value: string): void;

  getUserId(): string;
  setUserId(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getDepartureLatitude(): number;
  setDepartureLatitude(value: number): void;

  getDepartureLongitude(): number;
  setDepartureLongitude(value: number): void;

  getDestinationLatitude(): number;
  setDestinationLatitude(value: number): void;

  getDestinationLongitude(): number;
  setDestinationLongitude(value: number): void;

  getReward(): string;
  setReward(value: string): void;

  hasStartTime(): boolean;
  clearStartTime(): void;
  getStartTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Board.AsObject;
  static toObject(includeInstance: boolean, msg: Board): Board.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Board, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Board;
  static deserializeBinaryFromReader(message: Board, reader: jspb.BinaryReader): Board;
}

export namespace Board {
  export type AsObject = {
    id: string,
    type: string,
    userId: string,
    description: string,
    departureLatitude: number,
    departureLongitude: number,
    destinationLatitude: number,
    destinationLongitude: number,
    reward: string,
    startTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class BoardList extends jspb.Message {
  clearBoardsList(): void;
  getBoardsList(): Array<Board>;
  setBoardsList(value: Array<Board>): void;
  addBoards(value?: Board, index?: number): Board;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BoardList.AsObject;
  static toObject(includeInstance: boolean, msg: BoardList): BoardList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BoardList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BoardList;
  static deserializeBinaryFromReader(message: BoardList, reader: jspb.BinaryReader): BoardList;
}

export namespace BoardList {
  export type AsObject = {
    boardsList: Array<Board.AsObject>,
  }
}

export class Comment extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getPostId(): string;
  setPostId(value: string): void;

  getCommenterId(): string;
  setCommenterId(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Comment.AsObject;
  static toObject(includeInstance: boolean, msg: Comment): Comment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Comment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Comment;
  static deserializeBinaryFromReader(message: Comment, reader: jspb.BinaryReader): Comment;
}

export namespace Comment {
  export type AsObject = {
    id: string,
    postId: string,
    commenterId: string,
    content: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CommentList extends jspb.Message {
  clearCommentsList(): void;
  getCommentsList(): Array<Comment>;
  setCommentsList(value: Array<Comment>): void;
  addComments(value?: Comment, index?: number): Comment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CommentList.AsObject;
  static toObject(includeInstance: boolean, msg: CommentList): CommentList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CommentList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CommentList;
  static deserializeBinaryFromReader(message: CommentList, reader: jspb.BinaryReader): CommentList;
}

export namespace CommentList {
  export type AsObject = {
    commentsList: Array<Comment.AsObject>,
  }
}

export class ID extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ID.AsObject;
  static toObject(includeInstance: boolean, msg: ID): ID.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ID, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ID;
  static deserializeBinaryFromReader(message: ID, reader: jspb.BinaryReader): ID;
}

export namespace ID {
  export type AsObject = {
    id: string,
  }
}

