"use client";
import styles from "./page.module.css";
import { FormEvent, useState } from "react";
import { PartialMessage } from "@bufbuild/protobuf";
import { SayRequest, SayResponse } from "../../gen/eliza_pb";
import { createPromiseClient } from "@bufbuild/connect";
import { createGrpcWebTransport } from "@bufbuild/connect-web";
import { ElizaService } from "../../gen/eliza_connect";

// gRPCクライアントの初期化
const transport = createGrpcWebTransport({
  baseUrl: "http://localhost:8000",
});
const client = createPromiseClient(ElizaService, transport);

export default function Home() {
  const [sentence, setSentence] = useState("");
  const [text, setText] = useState("");

  const handleSubmit = async (e: FormEvent) => {
    console.log("greetingMessage");
    e.preventDefault();
    // リクエストメッセージのオブジェクトはPartialMessageを使うと取れます
    const response: PartialMessage<SayRequest> = { sentence };
    // gRPCメソッドを呼び出す
    const greetingMessage: SayResponse = await client.say(response);
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
