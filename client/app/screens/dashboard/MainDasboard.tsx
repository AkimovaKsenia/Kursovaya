import { FC } from "react";
import { FaSignOutAlt } from "react-icons/fa"; // Иконка выхода
import AdminLayout from "@/components/ui/layout/AdminHeader";
import { useAuth } from "hooks/useAuth";
import Link from "next/link";
import MainStatistic from "./MainStatistics";
import styles from "./MainDashboard.module.scss";

const Dashboard: FC = () => {
  const { user, setUser } = useAuth(); // Получаем данные пользователя из контекста
  return (
    <AdminLayout title="Dashboard" backgroundColor="#1F1F1F">
      <div className={styles.container}>
        <div className={styles.main}>
          {user ? (
            <div>
              <h1 style={{ color: "#FFFFFF" }}>Dashboard</h1>
              <h2 style={{ color: "#FFFFFF" }}>Привет!</h2>
              <p>Добро пожаловать в вашу панель управления.</p>
              <div className="w-full flex justify-center">
                <MainStatistic />
              </div>
            </div>
          ) : (
            <>
              <p style={{ color: "#FFFFFF", marginTop: "15px" }}>
                Пожалуйста, войдите в систему, чтобы продолжить.
              </p>
              <Link href="/">
                <button className={styles.button}>
                  <FaSignOutAlt style={{ marginRight: "8px" }} />
                  Вернуться на главную
                </button>
              </Link>
            </>
          )}
        </div>
      </div>
    </AdminLayout>
  );
};

export default Dashboard;
