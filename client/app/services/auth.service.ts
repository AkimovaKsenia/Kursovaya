import { axiosClassic } from "api/interceptor";
import { IAuthResponse } from "shared/interfaces/user.interface";
import { removeTokenFromStorage, saveToStorage } from "./auth.helper";

export const AuthService = {
  async login(email: string, password: string) {
    const response = await axiosClassic.post<IAuthResponse>("/login", {
      email,
      password,
    });

    if (response.data.token) {
      saveToStorage(response.data);
    }
    return response.data;
  },
  logout() {
    removeTokenFromStorage();
    localStorage.removeItem("user");
  },
};
