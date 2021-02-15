import Head from 'next/head'
import Layout from "../components/Layout";
import Title from "../components/Title";

export default function Home() {
  return (
    <Layout>
    <Title>Bem vindo ao desafio</Title>
    <a href='/account'>tela dos dados do usuario</a>
    </Layout>
  )
}
