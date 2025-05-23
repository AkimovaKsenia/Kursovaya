import Layout from "@/components/ui/layout/Layout";
import MovieItem from "@/components/ui/movie-item/MovieItem";
import { FC, useEffect, useState } from "react";
import { IListOfMovies, IMovie } from "shared/interfaces/movie.interface";
import styles from "./AllMovies.module.scss";
import { MovieService } from "services/movie.service";
import AdminLayout from "@/components/ui/layout/AdminHeader";
import DashboardLayout from "../../../components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import Button from "@/components/ui/layout/Button/Button";
import Link from "next/link";
import { useRouter } from "next/router";
import { useAuth } from "hooks/useAuth";
import ErrorAuth from "@/components/ui/ErrorAuth";

const Movies: FC = () => {
  const router = useRouter();
  const { user, setUser } = useAuth();
  const [movies, setMovies] = useState<IMovie[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const { data } = await MovieService.getAll();
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
    <DashboardLayout>
      {user ? (
        <div>
          <div
            style={{ background: "rgba(129, 125, 219, 0.8)" }}
            className={styles.catalog}
          >
            Фильмы
          </div>
          <Link href="/manage/movies/createmovie" className={styles.link}>
            <button className={styles.button}> Создать фильм </button>
          </Link>

          <div className={styles.items}>
            {movies.length ? (
              movies.map((movie) => <MovieItem movie={movie} key={movie.id} />)
            ) : (
              <div>Movies not found</div>
            )}
          </div>
        </div>
      ) : (
        <ErrorAuth />
      )}
    </DashboardLayout>
  );
};

export default Movies;
