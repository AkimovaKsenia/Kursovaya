import { FC } from "react";
import { useAuth } from "hooks/useAuth"; // Предполагается, что у вас есть кастомный хук useAuth
import { FaSignOutAlt } from "react-icons/fa"; // Иконка выхода

const Dashboard: FC = () => {
  const { user, setUser } = useAuth(); // Получаем данные пользователя из контекста

  const handleLogout = () => {
    // Логика выхода, например, удаление токена или очистка состояния
    setUser(null);
    alert("Вы вышли из системы");
  };

  return (
    <div
      style={{
        padding: "20px",
        fontFamily: "Arial, sans-serif",
        backgroundColor: "#181818", // Темный фон
        color: "#e0e0e0", // Светлый текст
        minHeight: "100vh", // Чтобы фон растягивался на всю высоту экрана
      }}
    >
      <h1 style={{ color: "#fff" }}>Dashboard</h1>
      {user ? (
        <div>
          <h2 style={{ color: "#e0e0e0" }}>Привет, {user.name}!</h2>
          <p>Добро пожаловать в вашу панель управления.</p>
          <button
            onClick={handleLogout}
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
            Выйти
          </button>
        </div>
      ) : (
        <p>Пожалуйста, войдите в систему, чтобы продолжить.</p>
      )}
    </div>
  );
};

export default Dashboard;
