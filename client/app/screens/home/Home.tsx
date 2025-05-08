import Layout from "@/components/ui/layout/Layout";
import { FC } from "react";
import Link from "next/link";
import Button from "../../components/ui/layout/Button/Button";
import { useAuth } from "hooks/useAuth";
import { useRouter } from "next/router";
import { AuthService } from "services/auth.service";
import Head from "next/head";
import styles from "./Home.module.scss";
import cn from "classnames";

const Home: FC = () => {
  const { user, setUser } = useAuth(); // Используем контекст для получения текущего пользователя
  const router = useRouter();

  const handleLogout = () => {
    // Логика логаута
    AuthService.logout(); // Вызываем logout, который удаляет токен и данные пользователя
    setUser(null); // Очищаем состояние в контексте
  };
  console.log("Rendering Home page, user:", user); // Логируем состояние пользователя
  console.log("Rendering Home page"); // Добавьте лог

  return (
    <Layout
      title="Cinema"
      backgroundImage="/imageMain.png"
      backgroundColor="#1E1E1E"
    >
      <div
        style={{
          backgroundBlendMode: "overlay", // накладывает цвет на изображение
        }}
        className="flex flex-col items-center mr-230 justify-center min-h-[60vh] space-y-2 mt-21"
      >
        <h1 className={styles.heading}>Cinema</h1>
        <h1 className={styles.headingMastery}>MASTERY</h1>
        <div className="max-w-110">
          <h3 className={styles.text}>
            Войдите в систему и просматривайте актуальную информацию о сеансах,
            фильмах и кинотеатрах города!
          </h3>
        </div>
        <Link href="/movies" passHref>
          <Button
            className={cn("px-6 py-3 text-lg mt-5 boldText", styles.darkText)}
          >
            Перейти к фильмам
          </Button>
        </Link>
        <div>
          <button onClick={handleLogout}>Logout</button>
        </div>
      </div>
    </Layout>
  );
};
export default Home;
