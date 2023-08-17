// @generated by protoc-gen-connect-es v0.11.0 with parameter "target=ts,import_extension=none"
// @generated from file eliza/v1/eliza.proto (package eliza.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { HelloRequest, HelloResponse, SayRequest, SayResponse } from "./eliza_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service eliza.v1.ElizaService
 */
export const ElizaService = {
  typeName: "eliza.v1.ElizaService",
  methods: {
    /**
     * @generated from rpc eliza.v1.ElizaService.Say
     */
    say: {
      name: "Say",
      I: SayRequest,
      O: SayResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc eliza.v1.ElizaService.Hello
     */
    hello: {
      name: "Hello",
      I: HelloRequest,
      O: HelloResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

