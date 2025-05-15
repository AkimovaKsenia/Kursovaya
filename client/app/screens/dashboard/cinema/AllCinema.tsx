import Layout from "@/components/ui/layout/Layout";
import MovieItem from "@/components/ui/movie-item/MovieItem";
import { FC, useEffect, useState } from "react";
import { IListOfMovies, IMovie } from "shared/interfaces/movie.interface";
import styles from "./AllCinema.module.scss";
import { MovieService } from "services/movie.service";
import AdminLayout from "@/components/ui/layout/AdminHeader";
import DashboardLayout from "../../../components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import Button from "@/components/ui/layout/Button/Button";
import Link from "next/link";
import { useRouter } from "next/router";
import { CinemaService } from "services/cinema.service";
import { ICinema, ICinemaMain } from "shared/interfaces/cinema.interface";
import CinemaItem from "@/components/ui/cinema-item/CinemaItem";

const Cinema: FC = () => {
  const router = useRouter();
  const [cinema, setCinema] = useState<ICinemaMain[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchCinema = async () => {
      try {
        const { data } = await CinemaService.getAllCinema(); // использует токен из куки
        setCinema(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setCinema([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchCinema();
  }, []);

  return (
    <DashboardLayout>
      <div
        style={{ background: "rgba(129, 125, 219, 0.8)" }}
        className={styles.catalog}
      >
        Cinema
      </div>
      <Link href="/manage/cinema/createcinema">
        <button className={styles.button}> Создать кинотеатр </button>
      </Link>
      <div className={styles.items}>
        {cinema.length ? (
          cinema.map((cinema) => <CinemaItem cinema={cinema} key={cinema.id} />)
        ) : (
          <div>Movies not found</div>
        )}
      </div>
    </DashboardLayout>
  );
};

export default Cinema;
