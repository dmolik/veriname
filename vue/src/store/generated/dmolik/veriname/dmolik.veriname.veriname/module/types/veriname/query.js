/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Ident } from "../veriname/ident";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
export const protobufPackage = "dmolik.veriname.veriname";
const baseQueryGetIdentRequest = { index: "" };
export const QueryGetIdentRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== "") {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetIdentRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetIdentRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetIdentRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        return message;
    },
};
const baseQueryGetIdentResponse = {};
export const QueryGetIdentResponse = {
    encode(message, writer = Writer.create()) {
        if (message.ident !== undefined) {
            Ident.encode(message.ident, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetIdentResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetIdentResponse };
        if (object.ident !== undefined && object.ident !== null) {
            message.ident = Ident.fromJSON(object.ident);
        }
        else {
            message.ident = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.ident !== undefined &&
            (obj.ident = message.ident ? Ident.toJSON(message.ident) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetIdentResponse };
        if (object.ident !== undefined && object.ident !== null) {
            message.ident = Ident.fromPartial(object.ident);
        }
        else {
            message.ident = undefined;
        }
        return message;
    },
};
const baseQueryAllIdentRequest = {};
export const QueryAllIdentRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllIdentRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllIdentRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllIdentRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllIdentResponse = {};
export const QueryAllIdentResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.ident) {
            Ident.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllIdentResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryAllIdentResponse };
        message.ident = [];
        if (object.ident !== undefined && object.ident !== null) {
            for (const e of object.ident) {
                message.ident.push(Ident.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.ident) {
            obj.ident = message.ident.map((e) => (e ? Ident.toJSON(e) : undefined));
        }
        else {
            obj.ident = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllIdentResponse };
        message.ident = [];
        if (object.ident !== undefined && object.ident !== null) {
            for (const e of object.ident) {
                message.ident.push(Ident.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Ident(request) {
        const data = QueryGetIdentRequest.encode(request).finish();
        const promise = this.rpc.request("dmolik.veriname.veriname.Query", "Ident", data);
        return promise.then((data) => QueryGetIdentResponse.decode(new Reader(data)));
    }
    IdentAll(request) {
        const data = QueryAllIdentRequest.encode(request).finish();
        const promise = this.rpc.request("dmolik.veriname.veriname.Query", "IdentAll", data);
        return promise.then((data) => QueryAllIdentResponse.decode(new Reader(data)));
    }
}
