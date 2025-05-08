import {
  createContext,
  FC,
  PropsWithChildren,
  useEffect,
  useState,
} from "react";
import { IContext, TypeUserState } from "./auth.interface";
import { useRouter } from "next/router";
import Cookies from "js-cookie";
import { AuthService } from "services/auth.service";

export const AuthContext = createContext({} as IContext);

const AuthProvider: FC<PropsWithChildren<unknown>> = ({ children }) => {
  const [user, setUser] = useState<TypeUserState>(null);
  const { pathname } = useRouter(); //Содержит текущий путь

  useEffect(() => {
    const token = Cookies.get("token");
    if (token) {
      const userData = localStorage.getItem("user");
      if (userData) {
        try {
          const user = JSON.parse(userData);
          setUser(user);
        } catch (error) {
          console.error("Ошибка при парсинге данных пользователя:", error);
        }
      }
    }
  }, []);

  useEffect(() => {
    const token = Cookies.get("token");
    if (!token) {
      AuthService.logout();
      setUser(null); // Сбрасываем состояние пользователя
    }
  }, [pathname]);

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
};
export default AuthProvider;
