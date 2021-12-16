import { Reader, Writer } from "protobufjs/minimal";
import { Ident } from "../veriname/ident";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "dmolik.veriname.veriname";
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
export declare const QueryGetIdentRequest: {
    encode(message: QueryGetIdentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetIdentRequest;
    fromJSON(object: any): QueryGetIdentRequest;
    toJSON(message: QueryGetIdentRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetIdentRequest>): QueryGetIdentRequest;
};
export declare const QueryGetIdentResponse: {
    encode(message: QueryGetIdentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetIdentResponse;
    fromJSON(object: any): QueryGetIdentResponse;
    toJSON(message: QueryGetIdentResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetIdentResponse>): QueryGetIdentResponse;
};
export declare const QueryAllIdentRequest: {
    encode(message: QueryAllIdentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllIdentRequest;
    fromJSON(object: any): QueryAllIdentRequest;
    toJSON(message: QueryAllIdentRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllIdentRequest>): QueryAllIdentRequest;
};
export declare const QueryAllIdentResponse: {
    encode(message: QueryAllIdentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllIdentResponse;
    fromJSON(object: any): QueryAllIdentResponse;
    toJSON(message: QueryAllIdentResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllIdentResponse>): QueryAllIdentResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a ident by index. */
    Ident(request: QueryGetIdentRequest): Promise<QueryGetIdentResponse>;
    /** Queries a list of ident items. */
    IdentAll(request: QueryAllIdentRequest): Promise<QueryAllIdentResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Ident(request: QueryGetIdentRequest): Promise<QueryGetIdentResponse>;
    IdentAll(request: QueryAllIdentRequest): Promise<QueryAllIdentResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
