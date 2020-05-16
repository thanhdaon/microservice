import Document, { Html, Head, Main, NextScript } from "next/document";
import NextHead from "next/head";

class MyDocument extends Document {
  static async getInitialProps(ctx) {
    const initialProps = await Document.getInitialProps(ctx);
    return { ...initialProps };
  }

  render() {
    return (
      <Html>
        <Head />
        <NextHead>
          <title>App</title>
          <link rel="icon" href="/favicon.ico" />

          <script
            src="https://kit.fontawesome.com/ff796dfd7b.js"
            crossOrigin="anonymous"
          ></script>
          <link
            href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css"
            rel="stylesheet"
          />
        </NextHead>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
