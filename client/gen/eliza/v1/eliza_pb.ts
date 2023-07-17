// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file eliza/v1/eliza.proto (package eliza.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message eliza.v1.SayRequest
 */
export class SayRequest extends Message<SayRequest> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<SayRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.SayRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayRequest {
    return new SayRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayRequest {
    return new SayRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayRequest {
    return new SayRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SayRequest | PlainMessage<SayRequest> | undefined, b: SayRequest | PlainMessage<SayRequest> | undefined): boolean {
    return proto3.util.equals(SayRequest, a, b);
  }
}

/**
 * @generated from message eliza.v1.SayResponse
 */
export class SayResponse extends Message<SayResponse> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<SayResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.SayResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayResponse {
    return new SayResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayResponse {
    return new SayResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayResponse {
    return new SayResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SayResponse | PlainMessage<SayResponse> | undefined, b: SayResponse | PlainMessage<SayResponse> | undefined): boolean {
    return proto3.util.equals(SayResponse, a, b);
  }
}

