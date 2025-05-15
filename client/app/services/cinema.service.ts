import instance from "api/interceptor";
import axios, { axiosClassic } from "api/interceptor";
import {
  ICinemaDto,
  ICinemaExportDto,
  ICinemaMain,
  IHall,
  IListOfCategory,
  IListofCinema,
  IListOfCondition,
} from "shared/interfaces/cinema.interface";
import Cookies from "js-cookie";

export const CinemaService = {
  async getAllCinema() {
    const fullUrl = instance.defaults.baseURL + "/auth/film";
    console.log("Полный URL (вручную):", fullUrl);
    try {
      const response = await instance.get<ICinemaMain[]>(
        "/auth/cinema/address_name",
        {
          headers: {
            Accept: "application/json",
          },
        }
      );
      console.log("Ответ получен!"); // Доходит ли сюда?
      return response;
    } catch (error) {
      console.error("Ошибка:", error); // Ловим ошибки
      throw error;
    }
  },
  async getHallsById(id: number) {
    return instance.get<IHall[]>(`/auth/cinema/halls/${id}`);
  },

  async getCinemaById(id: number) {
    return instance.get<ICinemaDto>(`/auth/cinema/id/${id}`);
  },
  async getCategory() {
    try {
      const response = await instance.get<IListOfCategory>(
        "/auth/cinema/categories",
        {
          headers: {
            Authorization: `Bearer ${Cookies.get("accessToken")}`,
          },
        }
      );
      return response;
    } catch (error) {
      console.error("Categories fetch error:", error);
      throw error;
    }
  },
  async getCondition() {
    try {
      const response = await instance.get<IListOfCondition>(
        "/auth/cinema/conditions",
        {
          headers: {
            Authorization: `Bearer ${Cookies.get("accessToken")}`,
          },
        }
      );
      return response;
    } catch (error) {
      console.error("Conditions fetch error:", error);
      throw error;
    }
  },
  async updateCinema(id: number, body: ICinemaExportDto) {
    try {
      const formData = new FormData();

      formData.append("id", id.toString());
      formData.append("name", body.name);
      formData.append("address", body.address);
      formData.append("email", body.email);
      formData.append("phone", body.phone);
      formData.append("description", body.description);
      formData.append("category_id", body.category_id.toString());
      formData.append("condition_id", body.condition_id.toString());

      // Файл
      if (body.photo) {
        formData.append("film_photo", body.photo);
      }

      console.log("📦 Формируем formData для отправки:");
      for (let pair of formData.entries()) {
        console.log(`${pair[0]}:`, pair[1]);
      }
      const response = await instance.put<string>("/auth/cinema", formData, {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async createCinema(body: ICinemaExportDto) {
    try {
      const formData = new FormData();

      formData.append("name", body.name);
      formData.append("address", body.address);
      formData.append("email", body.email);
      formData.append("phone", body.phone);
      formData.append("description", body.description);
      formData.append("category_id", body.category_id.toString());
      formData.append("condition_id", body.condition_id.toString());

      // Файл
      if (body.photo) {
        formData.append("photo", body.photo);
      }

      console.log("📦 Формируем formData для отправки:");
      for (let pair of formData.entries()) {
        console.log(`${pair[0]}:`, pair[1]);
      }
      const response = await instance.post<string>("/auth/cinema", formData, {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      console.error("FilmStudio fetch error:", error);
      throw error;
    }
  },
};
