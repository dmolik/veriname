import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "dmolik.veriname.veriname";
export interface MsgRegister {
    creator: string;
    alias: string;
    user: string;
    kind: string;
    target: string;
    payload: string;
}
export interface MsgRegisterResponse {
}
export declare const MsgRegister: {
    encode(message: MsgRegister, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegister;
    fromJSON(object: any): MsgRegister;
    toJSON(message: MsgRegister): unknown;
    fromPartial(object: DeepPartial<MsgRegister>): MsgRegister;
};
export declare const MsgRegisterResponse: {
    encode(_: MsgRegisterResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterResponse;
    fromJSON(_: any): MsgRegisterResponse;
    toJSON(_: MsgRegisterResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterResponse>): MsgRegisterResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
