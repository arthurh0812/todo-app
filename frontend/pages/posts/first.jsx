import Link from 'next/link'
import Image from 'next/image'
import Head from 'next/head'
import Layout from '../../components/layout'

export default function FirstPost() {
    return (
        <Layout>
            <Head>
                <title>My first blog post</title>
            </Head>
            <h1>My first blog!</h1>
            <Image
                src="/images/nice-pic.jpg"
                height={144}
                width={288}
                alt={"nice picture"}
            />
            <h2>
                <Link href="/">
                    <a>Back</a>
                </Link>
            </h2>
        </Layout>
    )
}
