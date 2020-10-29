import Layout from "layout";

function Home() {
  return <div className="text-xl">Hello World</div>;
}

export default function HomaePage() {
  return (
    <Layout>
      <Home />
    </Layout>
  );
}
