import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConsoleModule } from 'nestjs-console';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { LoginController } from './controllers/login/login.controller';
import { UserAccountController } from './controllers/user-account/user-account.controller';
import { UserAccount } from './models/user-account.model';
import { FixturesCommand } from './fixtures/fixtures.command';
import { Account } from './models/account.model';
import { Vehicle } from './models/vehicle.model';
@Module({
  imports: [
    ConfigModule.forRoot(),
    ConsoleModule,
    TypeOrmModule.forRoot({
      type: process.env.TYPEORM_CONNECTION as any,
      host: process.env.TYPEORM_HOST,
      port: parseInt(process.env.TYPEORM_PORT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.TYPEORM_DATABASE,
      entities: [UserAccount,Account,Vehicle]
    }),
    TypeOrmModule.forFeature([UserAccount,Account,Vehicle]),],
  controllers: [AppController, LoginController,UserAccountController],
  providers: [AppService,FixturesCommand],
})
export class AppModule {}
