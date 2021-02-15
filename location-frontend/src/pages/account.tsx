import * as React from 'react'
import { GetServerSideProps } from "next";
import Layout from "../components/Layout";
import Title from "../components/Title";
import { bankHttp } from "../util/http";
const UserAccount =(props)=>{
  const { userAccounts } = props;
    return(<Layout>
        <Title>Sua conta</Title>
        <div className="row">
        {userAccounts&&<div className="card">
            <h5 className="card-title">{userAccounts[0].name}</h5>
            <div className="card-body">
              <ul className="list-group">
                <li className="list-group-item"><b>CPF</b>:{userAccounts[0].cpf}</li>
                <li className="list-group-item"><b>matricula</b>:{userAccounts[0].matricula}</li>
                <li className="list-group-item"><b>aniversario</b>:{userAccounts[0].birth}</li>
              </ul>
            </div>
          </div>}
        </div>
      </Layout>)
    
    }

    export default UserAccount

    export const getServerSideProps: GetServerSideProps = async () => {
      const { data: userAccounts } = await bankHttp.get("api/user-accounts/userAccount");
      return {
        props: {
          userAccounts,
        },
      };
    };