import Header from "../components/Header"
import PostCard from "../components/Card"
import Footer from "../components/Footer"
import Head from "next/head"

const Index = () => {
  return (
    <div>
    <Head>
        <title>
          5B - Bearcat events
        </title>
        </Head>
      <Header />
      <PostCard />
    <div>
      <Footer />
      </div>
    </div>
  )
}

export default Index
