import {BeforeInsert, Column, CreateDateColumn, Entity, PrimaryGeneratedColumn} from "typeorm";

import {v4 as uuidv4} from 'uuid';


@Entity({
    name:"accounts"
})
export class Account {
    
    @PrimaryGeneratedColumn('uuid')
    id: string;
    
    @Column()
    login:string;

    @Column()
    kind:string;

    @Column()
    password:string;
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
