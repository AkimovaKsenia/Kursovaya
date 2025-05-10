import Layout from "@/components/ui/layout/Layout";
import MovieItem from "@/components/ui/movie-item/MovieItem";
import { FC, useEffect, useState } from "react";
import { IListOfMovies, IMovie } from "shared/interfaces/movie.interface";
import styles from "./AllMovies.module.scss";
import { MovieService } from "services/movie.service";
import AdminLayout from "@/components/ui/layout/AdminHeader";

const Movies: FC = () => {
  const [movies, setMovies] = useState<IMovie[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const { data } = await MovieService.getAll(); // использует токен из куки
        setMovies(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setMovies([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchMovies();
  }, []);
  return (
    <AdminLayout title="Cinema" backgroundColor="#1F1F1F">
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          height: "calc(100vh - 52.8px)", // учитываем высоту header, если есть
          padding: "20px",
          boxSizing: "border-box", // 1. Добавляем для корректного расчета размеров
          overflowX: "hidden" /* запрещает горизонтальный скролл */,
        }}
      >
        <div
          className={styles.main}
          style={{
            position: "fixed",
            width: "1000px",
            height: "700px",
            marginTop: "5px",
            borderRadius: "30px",
            padding: "20px",
            fontFamily: "Arial, sans-serif",
            backgroundColor: "#A7A7B6",
            color: "#FFFFFF", // Светлый текст
            boxSizing: "border-box",
            overflow: "auto",
            boxShadow: "2px 2px 17px rgba(129, 125, 219, 0.6)",
          }}
        >
          <div
            style={{ background: "rgba(129, 125, 219, 0.8)" }}
            className={styles.catalog}
          >
            MOVIES
          </div>
          <div className={styles.items}>
            {movies.length ? (
              movies.map((movie) => <MovieItem movie={movie} key={movie.id} />)
            ) : (
              <div>Movies not found</div>
            )}
          </div>
        </div>
      </div>
    </AdminLayout>
  );
};

export default Movies;
