import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "dmolik.veriname.veriname";
export interface Ident {
    index: string;
    alias: string;
    user: string;
    kind: string;
    target: string;
    payload: string;
}
export declare const Ident: {
    encode(message: Ident, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Ident;
    fromJSON(object: any): Ident;
    toJSON(message: Ident): unknown;
    fromPartial(object: DeepPartial<Ident>): Ident;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
