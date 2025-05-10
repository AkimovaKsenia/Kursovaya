import instance from "api/interceptor";
import axios, { axiosClassic } from "api/interceptor";
import { IMovie, IMovieDto } from "shared/interfaces/movie.interface";
import Cookies from "js-cookie";

export const MovieService = {
  async getMovieById(id: number) {
    return axiosClassic.get<IMovie>(`/movie/${id}`);
  },
  async getAll() {
    const fullUrl = instance.defaults.baseURL + "/auth/film";
    console.log("Полный URL (вручную):", fullUrl); // Что получилось?

    try {
      const response = await instance.get<IMovie[]>("/auth/film", {
        headers: {
          Accept: "application/json",
        },
      });
      console.log("Ответ получен!"); // Доходит ли сюда?
      return response;
    } catch (error) {
      console.error("Ошибка:", error); // Ловим ошибки
      throw error;
    }
  },

  async createMovie() {
    return axios.post<string>("/movie");
  },
  async updateMovie(id: number, body: IMovieDto) {
    return axios.patch<string>("/movie");
  },
  async deleteMovie(id: number) {
    return instance.delete<string>(`/auth/film/${id}`);
  },
};
