import { FC } from "react";
import { FaSignOutAlt } from "react-icons/fa"; // Иконка выхода
import AdminLayout from "@/components/ui/layout/AdminHeader";
import { useAuth } from "hooks/useAuth";
import Link from "next/link";

const Dashboard: FC = () => {
  const { user, setUser } = useAuth(); // Получаем данные пользователя из контекста
  return (
    <AdminLayout title="Dashboard" backgroundColor="#1F1F1F">
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          height: "calc(100vh - 52.8px)", // учитываем высоту header, если есть
          padding: "20px",
          boxSizing: "border-box", // 1. Добавляем для корректного расчета размеров
        }}
      >
        <div
          style={{
            position: "fixed",
            width: "1000px",
            height: "700px",
            marginTop: "5px",
            borderRadius: "30px",
            padding: "20px",
            fontFamily: "Arial, sans-serif",
            backgroundColor: "#A7A7B6",
            color: "#1F1F1F", // Светлый текст
            boxSizing: "border-box",
            overflow: "auto",
            boxShadow: "2px 2px 17px rgba(129, 125, 219, 0.6)",
          }}
        >
          <h1 style={{ color: "#1F1F1F" }}>Dashboard</h1>
          {user ? (
            <div>
              <h2 style={{ color: "#1F1F1F" }}>Привет!</h2>
              <p>Добро пожаловать в вашу панель управления.</p>
            </div>
          ) : (
            <>
              <p>Пожалуйста, войдите в систему, чтобы продолжить.</p>
              <Link href="/">
                <button
                  style={{
                    padding: "10px 15px",
                    backgroundColor: "#bb86fc", // Цвет кнопки для темной темы
                    color: "#121212", // Темный текст на кнопке
                    border: "none",
                    cursor: "pointer",
                    display: "flex",
                    alignItems: "center",
                    borderRadius: "5px", // Скругление углов
                  }}
                >
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
