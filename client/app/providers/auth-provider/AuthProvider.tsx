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
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

export const AuthContext = createContext({} as IContext);

const AuthProvider: FC<PropsWithChildren<unknown>> = ({ children }) => {
  const [user, setUser] = useState<TypeUserState>(null);
  const { pathname } = useRouter();

  useEffect(() => {
    const accessToken = Cookies.get("accessToken");
    if (accessToken) {
      const user = JSON.parse(localStorage.getItem("user") || "");

      setUser(user);
    }
  }, []);

  useEffect(() => {
    const accessToken = Cookies.get("accessToken");
    if (!accessToken && !user) {
      AuthService.logout();
      setUser(null);
    }
  }, [pathname]);

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
};
export default AuthProvider;
