/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "dmolik.veriname.veriname";

export interface MsgRegister {
  creator: string;
  alias: string;
  user: string;
  kind: string;
  target: string;
  payload: string;
}

export interface MsgRegisterResponse {}

const baseMsgRegister: object = {
  creator: "",
  alias: "",
  user: "",
  kind: "",
  target: "",
  payload: "",
};

export const MsgRegister = {
  encode(message: MsgRegister, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.alias !== "") {
      writer.uint32(18).string(message.alias);
    }
    if (message.user !== "") {
      writer.uint32(26).string(message.user);
    }
    if (message.kind !== "") {
      writer.uint32(34).string(message.kind);
    }
    if (message.target !== "") {
      writer.uint32(42).string(message.target);
    }
    if (message.payload !== "") {
      writer.uint32(50).string(message.payload);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegister {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegister } as MsgRegister;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.alias = reader.string();
          break;
        case 3:
          message.user = reader.string();
          break;
        case 4:
          message.kind = reader.string();
          break;
        case 5:
          message.target = reader.string();
          break;
        case 6:
          message.payload = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegister {
    const message = { ...baseMsgRegister } as MsgRegister;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = String(object.user);
    } else {
      message.user = "";
    }
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = String(object.kind);
    } else {
      message.kind = "";
    }
    if (object.target !== undefined && object.target !== null) {
      message.target = String(object.target);
    } else {
      message.target = "";
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = String(object.payload);
    } else {
      message.payload = "";
    }
    return message;
  },

  toJSON(message: MsgRegister): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.alias !== undefined && (obj.alias = message.alias);
    message.user !== undefined && (obj.user = message.user);
    message.kind !== undefined && (obj.kind = message.kind);
    message.target !== undefined && (obj.target = message.target);
    message.payload !== undefined && (obj.payload = message.payload);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegister>): MsgRegister {
    const message = { ...baseMsgRegister } as MsgRegister;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = object.user;
    } else {
      message.user = "";
    }
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = object.kind;
    } else {
      message.kind = "";
    }
    if (object.target !== undefined && object.target !== null) {
      message.target = object.target;
    } else {
      message.target = "";
    }
    if (object.payload !== undefined && object.payload !== null) {
      message.payload = object.payload;
    } else {
      message.payload = "";
    }
    return message;
  },
};

const baseMsgRegisterResponse: object = {};

export const MsgRegisterResponse = {
  encode(_: MsgRegisterResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterResponse } as MsgRegisterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRegisterResponse {
    const message = { ...baseMsgRegisterResponse } as MsgRegisterResponse;
    return message;
  },

  toJSON(_: MsgRegisterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRegisterResponse>): MsgRegisterResponse {
    const message = { ...baseMsgRegisterResponse } as MsgRegisterResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Register(request: MsgRegister): Promise<MsgRegisterResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Register(request: MsgRegister): Promise<MsgRegisterResponse> {
    const data = MsgRegister.encode(request).finish();
    const promise = this.rpc.request(
      "dmolik.veriname.veriname.Msg",
      "Register",
      data
    );
    return promise.then((data) => MsgRegisterResponse.decode(new Reader(data)));
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
