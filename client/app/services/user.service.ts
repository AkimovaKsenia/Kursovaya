import { axiosClassic } from "api/interceptor";
import {
  IAuthResponse,
  IListOfRoles,
  IUserExportDto,
} from "shared/interfaces/user.interface";
import { removeTokenFromStorage, saveToStorage } from "./auth.helper";
import instance from "api/interceptor";
import Cookies from "js-cookie";

export const UserService = {
  async createUser(body: IUserExportDto) {
    try {
      const requestData = {
        name: body.name,
        role_id: body.role_id,
        surname: body.surname,
        password: body.password,
        email: body.email,
      };

      console.log("📦 Отправляемые данные:", requestData);

      const response = await instance.post<string>(
        "/auth/register",
        requestData, // Отправляем как JSON
        {
          headers: {
            Authorization: `Bearer ${Cookies.get("accessToken")}`,
            "Content-Type": "application/json", // Явно указываем тип контента
          },
        }
      );
      return response;
    } catch (error) {
      throw error;
    }
  },
  async getRoles() {
    try {
      const response = await instance.get<IListOfRoles>("/auth/user/role", {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      console.error("roles fetch error:", error);
      throw error;
    }
  },
};
