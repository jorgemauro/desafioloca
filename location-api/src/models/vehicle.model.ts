import {BeforeInsert, Column, CreateDateColumn, Entity, PrimaryGeneratedColumn} from "typeorm";
import {v4 as uuidv4} from 'uuid';

@Entity()
export class Vehicle {
    @PrimaryGeneratedColumn('uuid')
    id: string;
    
    @Column()
    Placa:string;

    @Column()
    Modelo:string;

    @Column()
    ValorHora:number;

    @Column()
    Combustivel:string;

    @Column()
    LimitePortaMala:string;

    @Column()
    Categoria:string;
    
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
