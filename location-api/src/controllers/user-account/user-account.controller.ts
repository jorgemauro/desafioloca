import { Controller, Get } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { UserAccount } from 'src/models/user-account.model';
import { Repository } from 'typeorm';

@Controller('user-accounts')
export class UserAccountController {
    
    constructor(
        @InjectRepository(UserAccount)
        private userAccountRepo: Repository<UserAccount>
    ){

    }
    @Get('userAccount')
    index(){
        return this.userAccountRepo.find()
    }
}
