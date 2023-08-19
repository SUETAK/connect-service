"use client";
import { FormEvent, useState } from "react";
import { PartialMessage } from "@bufbuild/protobuf";

import { createPromiseClient } from "@bufbuild/connect";
import { createGrpcWebTransport } from "@bufbuild/connect-web";
import { ElizaService } from "../../gen/eliza/v1/eliza_connect";
import { HelloRequest, HelloResponse, SayRequest, SayResponse } from "../../gen/eliza/v1/eliza_pb";
import useSWR from "swr";

const baseUrl = process.env.NEXT_PUBLIC_GRPC_HOST ?? "http://localhost:8080";
// gRPCクライアントの初期化
const transport = createGrpcWebTransport({
  baseUrl
});
const client = createPromiseClient(ElizaService, transport);
const headers = { "Access-Control-Allow-Origin": baseUrl };

function requestServer(): Promise<SayResponse> {
  const request: PartialMessage<SayRequest> = { sentence: 'this is a Bob ?' };
  return client.say(request, { headers });
}

async function fetcher() {
  return client.say({ sentence: 'this is a Bob ?' }, { headers });
}
let tmp: any

const Post = () => {
  console.log(baseUrl);
  const [sentence, setSentence] = useState("");
  const [text, setText] = useState("");
  const [age, setAge] = useState(0);
  const [name, setName] = useState("");

  if (tmp == null) {
    throw requestServer().then((res) => {
      tmp = res
      // console.log(tmp);
    })
  }

  const {data, error, isLoading} = useSWR([ElizaService.methods.say.name], fetcher)

  console.log('swr:', data);

  const handleSubmit = async (e: FormEvent) => {
    console.log("greetingMessage");
    e.preventDefault();
    // リクエストメッセージのオブジェクトはPartialMessageを使うと取れます
    const request: PartialMessage<SayRequest> = { sentence };
    // gRPCメソッドを呼び出す
    const greetingMessage: SayResponse = await client.say(request, { headers });
    console.log("greetingMessage: ", greetingMessage);
    setText(greetingMessage.sentence);
  };

  const handleOnClick = async (e: FormEvent) => {
    e.preventDefault();
    // gRPCメソッドを呼び出す
    const request: PartialMessage<HelloRequest> = { name };
    const helloMessage: HelloResponse = await client.hello(request, { headers });
    console.log("greetingMessage: ", helloMessage.age);
  };

  return (
    < >
      <form onSubmit={handleSubmit}>
        <input
          placeholder="name"
          value={sentence}
          onChange={(e) => setSentence(e.target.value)}
        />
        <button type="submit">submit</button>
      </form>

      <form>
        <button onClick={handleOnClick}>
          hello
        </button>
      </form>
      <div>{text}</div>
    </>
  );
};

export default Post;
