import Image from "next/image";
import { Geist, Geist_Mono } from "next/font/google";
import Layout from "@/components/ui/layout/Layout";
import { GetStaticProps, NextPage } from "next";
import Home from "../app/screens/home/Home";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const HomePage: NextPage = () => {
  return <Home />;
};

export const getStaticProps: GetStaticProps = async () => {
  return {
    props: {}, // никаких данных не требуется
    revalidate: 86400, // опционально: регенерация раз в день (если потом добавите контент)
  };
};

export default HomePage;
