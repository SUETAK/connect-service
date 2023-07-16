import { Controller } from '@nestjs/common';
import { AppService } from '@/app.service';

import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';
import { GrpcMethod } from '@nestjs/microservices';
import { SayRequest, SayResponse } from '../gen/eliza_pb';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @GrpcMethod('ElizaService', 'Say')
  findOne(
    data: SayRequest,
    metadata: Metadata,
    call: ServerUnaryCall<any, any>,
  ): SayResponse {
    return SayResponse.fromJson({
      sentence: `おっけー。うけとりましたよ。${data.sentence}`,
    });
  }
}
