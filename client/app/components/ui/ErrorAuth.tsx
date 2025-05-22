import Link from "next/link";
import { FC } from "react";
import { FaSignOutAlt } from "react-icons/fa";

const ErrorAuth: FC = ({}) => {
  return (
    <>
      <p style={{ color: "#FFFFFF", marginTop: "15px" }}>
        Пожалуйста, войдите в систему, чтобы продолжить.
      </p>
      <Link href="/">
        <button
          style={{
            marginTop: "15px",
            padding: "9px 13px",
            border: "1px solid rgba(255, 255, 255, 0.3)",
            background: "rgba(255, 255, 255, 0.1)",
            borderRadius: "4px",
            boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
            cursor: "pointer",
          }}
        >
          <FaSignOutAlt style={{ marginRight: "8px" }} />
          Вернуться на главную
        </button>
      </Link>
    </>
  );
};

export default ErrorAuth;
