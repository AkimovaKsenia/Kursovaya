import instance from "api/interceptor";
import axios, { axiosClassic } from "api/interceptor";
import {
  IListOfDirector,
  IListOfFilmStudio,
  IListOfGenres,
  IListOfOperators,
  IMovie,
  IMovieDto,
  IMovieExportDto,
} from "shared/interfaces/movie.interface";
import Cookies from "js-cookie";

export const MovieService = {
  async getMovieById(id: number) {
    return instance.get<IMovie>(`/auth/film/id/${id}`);
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

  async createMovie(body: IMovieExportDto) {
    try {
      const formData = new FormData();

      formData.append("name", body.name);
      formData.append("description", body.description);
      formData.append("duration_in_min", body.duration_in_min.toString());
      formData.append("film_studio_id", body.film_studio_id.toString());

      // Файл
      if (body.film_photo && body.film_photo instanceof File) {
        formData.append("film_photo", body.film_photo);
      }

      // Массивы
      console.log("cast_list:", body.cast_list);
      console.log("Array.isArray(cast_list):", Array.isArray(body.cast_list));

      if (Array.isArray(body.cast_list)) {
        body.cast_list.forEach((actor) => formData.append("cast_list", actor));
      } else {
        console.warn("cast_list is not array:", body.cast_list);
      }

      body.genre_ids.forEach((id) =>
        formData.append("genre_ids", id.toString())
      );
      body.operator_ids.forEach((id) =>
        formData.append("operator_ids", id.toString())
      );
      body.director_ids.forEach((id) =>
        formData.append("director_ids", id.toString())
      );

      console.log("📦 Формируем formData для отправки:");
      for (let pair of formData.entries()) {
        console.log(`${pair[0]}:`, pair[1]);
      }
      const response = await instance.post<string>("/auth/film", formData, {
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

  async updateMovie(id: number, body: IMovieExportDto) {
    try {
      const formData = new FormData();

      formData.append("id", id.toString());
      formData.append("name", body.name);
      formData.append("description", body.description);
      formData.append("duration_in_min", body.duration_in_min.toString());
      formData.append("film_studio_id", body.film_studio_id.toString());

      // Файл
      if (body.film_photo && body.film_photo instanceof File) {
        formData.append("film_photo", body.film_photo);
      }

      // Массивы
      console.log("cast_list:", body.cast_list);
      console.log("Array.isArray(cast_list):", Array.isArray(body.cast_list));

      if (Array.isArray(body.cast_list)) {
        body.cast_list.forEach((actor: string) => {
          formData.append("cast_list[]", actor);
        });
      }
      body.genre_ids.forEach((id) =>
        formData.append("genre_ids", id.toString())
      );
      body.operator_ids.forEach((id) =>
        formData.append("operator_ids", id.toString())
      );
      body.director_ids.forEach((id) =>
        formData.append("director_ids", id.toString())
      );

      console.log("📦 Формируем formData для отправки:");
      for (let pair of formData.entries()) {
        console.log(`${pair[0]}:`, pair[1]);
      }
      const response = await instance.put<string>("/auth/film", formData, {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async deleteMovie(id: number) {
    return instance.delete<string>(`/auth/film/${id}`);
  },
  async getGenres() {
    try {
      const response = await instance.get<IListOfGenres>("/auth/genres", {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      console.error("Genres fetch error:", error);
      throw error;
    }
  },
  async getOperators() {
    try {
      const response = await instance.get<IListOfOperators>("/auth/operators", {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      console.error("Operators fetch error:", error);
      throw error;
    }
  },
  async getFilmStudios() {
    try {
      const response = await instance.get<IListOfFilmStudio>(
        "/auth/film-studios",
        {
          headers: {
            Authorization: `Bearer ${Cookies.get("accessToken")}`,
          },
        }
      );
      return response;
    } catch (error) {
      console.error("FilmStudio fetch error:", error);
      throw error;
    }
  },
  async getDirectors() {
    try {
      const response = await instance.get<IListOfDirector>("/auth/directors", {
        headers: {
          Authorization: `Bearer ${Cookies.get("accessToken")}`,
        },
      });
      return response;
    } catch (error) {
      console.error("Directors fetch error:", error);
      throw error;
    }
  },
};
