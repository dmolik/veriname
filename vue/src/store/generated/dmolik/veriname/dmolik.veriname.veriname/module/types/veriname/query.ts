/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Ident } from "../veriname/ident";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "dmolik.veriname.veriname";

export interface QueryGetIdentRequest {
  index: string;
}

export interface QueryGetIdentResponse {
  ident: Ident | undefined;
}

export interface QueryAllIdentRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllIdentResponse {
  ident: Ident[];
  pagination: PageResponse | undefined;
}

const baseQueryGetIdentRequest: object = { index: "" };

export const QueryGetIdentRequest = {
  encode(
    message: QueryGetIdentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetIdentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetIdentRequest } as QueryGetIdentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIdentRequest {
    const message = { ...baseQueryGetIdentRequest } as QueryGetIdentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetIdentRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetIdentRequest>): QueryGetIdentRequest {
    const message = { ...baseQueryGetIdentRequest } as QueryGetIdentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetIdentResponse: object = {};

export const QueryGetIdentResponse = {
  encode(
    message: QueryGetIdentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ident !== undefined) {
      Ident.encode(message.ident, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetIdentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetIdentResponse } as QueryGetIdentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ident = Ident.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIdentResponse {
    const message = { ...baseQueryGetIdentResponse } as QueryGetIdentResponse;
    if (object.ident !== undefined && object.ident !== null) {
      message.ident = Ident.fromJSON(object.ident);
    } else {
      message.ident = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetIdentResponse): unknown {
    const obj: any = {};
    message.ident !== undefined &&
      (obj.ident = message.ident ? Ident.toJSON(message.ident) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetIdentResponse>
  ): QueryGetIdentResponse {
    const message = { ...baseQueryGetIdentResponse } as QueryGetIdentResponse;
    if (object.ident !== undefined && object.ident !== null) {
      message.ident = Ident.fromPartial(object.ident);
    } else {
      message.ident = undefined;
    }
    return message;
  },
};

const baseQueryAllIdentRequest: object = {};

export const QueryAllIdentRequest = {
  encode(
    message: QueryAllIdentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllIdentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllIdentRequest } as QueryAllIdentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllIdentRequest {
    const message = { ...baseQueryAllIdentRequest } as QueryAllIdentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIdentRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllIdentRequest>): QueryAllIdentRequest {
    const message = { ...baseQueryAllIdentRequest } as QueryAllIdentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllIdentResponse: object = {};

export const QueryAllIdentResponse = {
  encode(
    message: QueryAllIdentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ident) {
      Ident.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllIdentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllIdentResponse } as QueryAllIdentResponse;
    message.ident = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ident.push(Ident.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllIdentResponse {
    const message = { ...baseQueryAllIdentResponse } as QueryAllIdentResponse;
    message.ident = [];
    if (object.ident !== undefined && object.ident !== null) {
      for (const e of object.ident) {
        message.ident.push(Ident.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIdentResponse): unknown {
    const obj: any = {};
    if (message.ident) {
      obj.ident = message.ident.map((e) => (e ? Ident.toJSON(e) : undefined));
    } else {
      obj.ident = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllIdentResponse>
  ): QueryAllIdentResponse {
    const message = { ...baseQueryAllIdentResponse } as QueryAllIdentResponse;
    message.ident = [];
    if (object.ident !== undefined && object.ident !== null) {
      for (const e of object.ident) {
        message.ident.push(Ident.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a ident by index. */
  Ident(request: QueryGetIdentRequest): Promise<QueryGetIdentResponse>;
  /** Queries a list of ident items. */
  IdentAll(request: QueryAllIdentRequest): Promise<QueryAllIdentResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Ident(request: QueryGetIdentRequest): Promise<QueryGetIdentResponse> {
    const data = QueryGetIdentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dmolik.veriname.veriname.Query",
      "Ident",
      data
    );
    return promise.then((data) =>
      QueryGetIdentResponse.decode(new Reader(data))
    );
  }

  IdentAll(request: QueryAllIdentRequest): Promise<QueryAllIdentResponse> {
    const data = QueryAllIdentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dmolik.veriname.veriname.Query",
      "IdentAll",
      data
    );
    return promise.then((data) =>
      QueryAllIdentResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
