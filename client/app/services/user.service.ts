import { axiosClassic } from "api/interceptor";
import {
  IAuthResponse,
  IListOfRoles,
  IListOfUsers,
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
        requestData,
        {
          headers: {
            Authorization: `Bearer ${Cookies.get("accessToken")}`,
            "Content-Type": "application/json",
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
  async getAllUsers() {
    const fullUrl = instance.defaults.baseURL + "/auth/user";
    console.log("Полный URL (вручную):", fullUrl);

    try {
      const response = await instance.get<IListOfUsers>("/auth/user", {
        headers: {
          Accept: "application/json",
        },
      });
      console.log("Ответ получен!");
      return response;
    } catch (error) {
      console.error("Ошибка:", error);
      throw error;
    }
  },
  async deleteUser(id: number) {
    return instance.delete<string>(`/auth/user/${id}`);
  },
};
