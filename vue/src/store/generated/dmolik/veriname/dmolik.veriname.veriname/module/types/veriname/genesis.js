/* eslint-disable */
import { Ident } from "../veriname/ident";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "dmolik.veriname.veriname";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.identList) {
            Ident.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.identList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.identList.push(Ident.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.identList = [];
        if (object.identList !== undefined && object.identList !== null) {
            for (const e of object.identList) {
                message.identList.push(Ident.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.identList) {
            obj.identList = message.identList.map((e) => e ? Ident.toJSON(e) : undefined);
        }
        else {
            obj.identList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.identList = [];
        if (object.identList !== undefined && object.identList !== null) {
            for (const e of object.identList) {
                message.identList.push(Ident.fromPartial(e));
            }
        }
        return message;
    },
};
