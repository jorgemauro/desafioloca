import { Controller, Get } from '@nestjs/common';
import { get } from 'http';

@Controller('login')
export class LoginController {
    @Get('login')
    index(){
        return {"key":"value"}
    }
}
