import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import styles from './layout.module.css'
import utilStyles from '../styles/utils.module.css'

const name = "Your name"
export const siteTitle = "To-do Application"

export default function Layout({ children, home }) {
    return <div className={styles.container}>
        <Head>
            <link rel="icon" href={"/favicon.ico"} />
            <meta
                name={"description"}
                content={"Organize your day with to-do items!"}
            />
            <meta
                name={"og:title"}
                content={siteTitle}
            />
        </Head>
        <header className={styles.header}>
            {home ? (
                <>
                    <Image
                        priority
                        src={"/images/nice-pic.jpg"}
                        className={utilStyles.borderCircle}
                        width={288}
                        height={144}
                        alt={name}
                    />
                    <h1 className={utilStyles.heading2x1}>{name}</h1>
                </>
            ) : (
                <>
                    <Link href="/">
                        <a>
                            <Image
                                priority
                                src={"/images/nice-pic.jpg"}
                                className={utilStyles.borderCircle}
                                width={216}
                                height={108}
                                alt={name}
                            />
                        </a>
                    </Link>
                    <h2 className={utilStyles.headingLg}>
                        <Link href={"/"}>
                            <a className={utilStyles.colorInherit}>{name}</a>
                        </Link>
                    </h2>
                </>
            )}
        </header>
        <main>
            {children}
        </main>
        {!home && (
            <div className={styles.backToHome}>
                <Link href={"/"}>
                    <a>Back</a>
                </Link>
            </div>
        )}
    </div>
}