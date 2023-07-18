"use client";
import styles from "./page.module.css";
import { FormEvent, useState } from "react";
import { PartialMessage } from "@bufbuild/protobuf";

import { createPromiseClient } from "@bufbuild/connect";
import { createGrpcWebTransport } from "@bufbuild/connect-web";
import { ElizaService } from "../../gen/eliza/v1/eliza_connect";
import { SayRequest, SayResponse } from "../../gen/eliza/v1/eliza_pb";

const baseUrl = process.env.NEXT_PUBLIC_GRPC_HOST ?? 'http://localhost:8080';
// gRPCクライアントの初期化
const transport = createGrpcWebTransport({
  baseUrl
});
const client = createPromiseClient(ElizaService, transport);

export default function Home() {
  console.log(baseUrl);
  const [sentence, setSentence] = useState("");
  const [text, setText] = useState("");

  const handleSubmit = async (e: FormEvent) => {
    console.log("greetingMessage");
    e.preventDefault();
    // リクエストメッセージのオブジェクトはPartialMessageを使うと取れます
    const response: PartialMessage<SayRequest> = { sentence };
    // gRPCメソッドを呼び出す
    const greetingMessage: SayResponse = await client.say(response, {headers: { "Access-Control-Allow-Origin": baseUrl }});
    console.log("greetingMessage: ", greetingMessage);
    setText(greetingMessage.sentence);
  };
  return (
    <main className={styles.main}>
      <form onSubmit={handleSubmit}>
        <input
          placeholder="name"
          value={sentence}
          onChange={(e) => setSentence(e.target.value)}
        />
        <button type="submit">submit</button>
      </form>
      <div>{text}</div>
    </main>
  );
}
