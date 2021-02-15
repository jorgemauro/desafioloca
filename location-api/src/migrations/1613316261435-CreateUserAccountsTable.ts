import {MigrationInterface, QueryRunner, Table} from "typeorm";

export class CreateUserAccountsTable1613316261435 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        queryRunner.createTable(
            new Table({
                name: 'user_accounts',
                columns: [
                    {
                        name: 'id',
                        type: 'uuid',
                        isPrimary: true
                    },
                    {
                        name: 'name',
                        type: 'varchar'
                    },
                    {
                        name: 'cpf',
                        type: 'varchar'
                    },
                    {
                        name: 'matricula',
                        type: 'varchar'
                    },
                    {
                        name: 'adress_id',
                        type: 'varchar',
                    },
                    {
                        name: 'account_id',
                        type: 'varchar'
                    },
                    {
                        name: 'birth',
                        type: 'timestamp'
                    },
                    {
                        name: 'created_at',
                        type: 'timestamp',
                        default: 'CURRENT_TIMESTAMP'
                    }
                ]
            })
        )
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropTable('user_accounts')
    }

}
