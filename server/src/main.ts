import { NestFactory } from '@nestjs/core';
import { AppModule } from '@/app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      transport: Transport.GRPC,
      options: {
        url: 'localhost:8000',
        package: 'eliza.v1',
        protoPath: join(__dirname, '../proto/eliza.proto'),
      },
    },
  );
  await app.listen();
}

bootstrap();
