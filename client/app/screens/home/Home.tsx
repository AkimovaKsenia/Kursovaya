import Layout from "@/components/ui/layout/Layout";
import { FC } from "react";
import Link from "next/link";
import Button from "../../components/ui/layout/Button/Button";

const Home: FC = () => {
  console.log("Rendering Home page"); // Добавьте лог
  return (
    <Layout title="Cinema">
      <div className="flex flex-col items-center justify-center min-h-[60vh]">
        <h1 className="text-4xl font-bold mb-8">Главная страница</h1>

        <Link href="/movies" passHref>
          <Button className="px-6 py-3 text-lg">Перейти к фильмам</Button>
        </Link>
      </div>
    </Layout>
  );
};
export default Home;
