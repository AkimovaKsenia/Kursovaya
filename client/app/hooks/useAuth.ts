import { AuthContext } from "providers/auth-provider/AuthProvider";
import { useContext } from "react";

export const useAuth = () => useContext(AuthContext);
