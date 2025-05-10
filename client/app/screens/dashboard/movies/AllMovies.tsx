import Layout from "@/components/ui/layout/Layout";
import MovieItem from "@/components/ui/movie-item/MovieItem";
import { FC, useEffect, useState } from "react";
import { IListOfMovies, IMovie } from "shared/interfaces/movie.interface";
import styles from "./AllMovies.module.scss";
import { MovieService } from "services/movie.service";
import AdminLayout from "@/components/ui/layout/AdminHeader";
import DashboardLayout from "../../../components/ui/layout/DashboardLayout";

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
    <AdminLayout title="Dashboard" backgroundColor="#1F1F1F">
      <div className={styles.containerfirst}>
        <div className={styles.main}>
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
