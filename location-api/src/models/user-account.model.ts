import {BeforeInsert, Column, CreateDateColumn, Entity, PrimaryGeneratedColumn} from "typeorm";

import {v4 as uuidv4} from 'uuid';


@Entity({
    name:"user_accounts"
})
export class UserAccount {
    
    @PrimaryGeneratedColumn('uuid')
    id: string;
    
    @Column()
    name:string;

    @Column()
    cpf:string;

    @Column()
    matricula:string;

    @Column()
    adress_id:string;

    @Column()
    account_id:string;


@CreateDateColumn({type: 'timestamp'})
birth:Date;
@CreateDateColumn({type: 'timestamp'})
created_at:Date;

@BeforeInsert()
    generateId(){
        if(this.id){
            return;
        }

        this.id = uuidv4();
    }
}
