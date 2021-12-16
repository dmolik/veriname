/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "dmolik.veriname.veriname";
const baseMsgRegister = {
    creator: "",
    alias: "",
    user: "",
    kind: "",
    target: "",
    payload: "",
};
export const MsgRegister = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRegister };
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
    fromJSON(object) {
        const message = { ...baseMsgRegister };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = String(object.alias);
        }
        else {
            message.alias = "";
        }
        if (object.user !== undefined && object.user !== null) {
            message.user = String(object.user);
        }
        else {
            message.user = "";
        }
        if (object.kind !== undefined && object.kind !== null) {
            message.kind = String(object.kind);
        }
        else {
            message.kind = "";
        }
        if (object.target !== undefined && object.target !== null) {
            message.target = String(object.target);
        }
        else {
            message.target = "";
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = String(object.payload);
        }
        else {
            message.payload = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.alias !== undefined && (obj.alias = message.alias);
        message.user !== undefined && (obj.user = message.user);
        message.kind !== undefined && (obj.kind = message.kind);
        message.target !== undefined && (obj.target = message.target);
        message.payload !== undefined && (obj.payload = message.payload);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRegister };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = object.alias;
        }
        else {
            message.alias = "";
        }
        if (object.user !== undefined && object.user !== null) {
            message.user = object.user;
        }
        else {
            message.user = "";
        }
        if (object.kind !== undefined && object.kind !== null) {
            message.kind = object.kind;
        }
        else {
            message.kind = "";
        }
        if (object.target !== undefined && object.target !== null) {
            message.target = object.target;
        }
        else {
            message.target = "";
        }
        if (object.payload !== undefined && object.payload !== null) {
            message.payload = object.payload;
        }
        else {
            message.payload = "";
        }
        return message;
    },
};
const baseMsgRegisterResponse = {};
export const MsgRegisterResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRegisterResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgRegisterResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRegisterResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Register(request) {
        const data = MsgRegister.encode(request).finish();
        const promise = this.rpc.request("dmolik.veriname.veriname.Msg", "Register", data);
        return promise.then((data) => MsgRegisterResponse.decode(new Reader(data)));
    }
}
