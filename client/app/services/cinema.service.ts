import instance from "api/interceptor";
import axios, { axiosClassic } from "api/interceptor";
import {
  ICinemaMain,
  IHall,
  IListofCinema,
} from "shared/interfaces/cinema.interface";

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
};
