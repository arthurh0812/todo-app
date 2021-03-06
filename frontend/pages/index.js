import Head from 'next/head'
import Link from 'next/link'
import Layout, {siteTitle} from '../components/layout'
import styles from '../styles/home.module.css'
import utilStyles from '../styles/utils.module.css'

export default function Home() {
  return (
      <Layout home>
        <Head>
          <title>{siteTitle}</title>
        </Head>
        <section className={utilStyles.headingMd}>
          <div className={styles.container}>
            <main className={styles.main}>
              <h1 className={styles.title}>
                Welcome to the Todo App!
              </h1>

              <p className={styles.description}>
                Get started by editing{' '}
                <code className={styles.code}>pages/index.js</code>
              </p>

              <div className={styles.grid}>
                <div className={styles.card}>
                  <h3>Hi there &rarr;</h3>
                  <p>About Marcel Davis and his team.</p>
                </div>

                <div className={styles.card}>
                  <h3>Learn &rarr;</h3>
                  <p>Learn about the todo app!</p>
                </div>

                <div className={styles.card}>
                  <h3>Examples &rarr;</h3>
                  <p>Discover amazing item examples!</p>
                </div>

                <div className={styles.card}>
                  <h3>Look at our pages &rarr;</h3>
                  <p>
                    <Link href={"/posts/first"}>
                      <a>The blog</a>
                    </Link>
                  </p>
                </div>
              </div>
            </main>
          </div>
        </section>
      </Layout>
  )
}
