import { createConnectTransport } from "@bufbuild/connect-web";
import { createPromiseClient } from "@bufbuild/connect";
import { ElizaService } from "../../gen/eliza/v1/eliza_connect";
import { SayResponse } from "../../gen/eliza/v1/eliza_pb";


export const elizaService = () => {
  const transport = createConnectTransport({
    baseUrl: "http://localhost:8000",
  });

  const client = createPromiseClient(ElizaService, transport);

  const say = async (sentence: string) => {
    const res: SayResponse = await client.say({ sentence });
    return res.sentence;
  };

  return {
    say,
  };
};
