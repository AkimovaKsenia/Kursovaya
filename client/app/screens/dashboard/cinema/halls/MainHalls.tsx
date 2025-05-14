import Button from "@/components/ui/layout/Button/Button";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import FileUploader from "@/components/ui/layout/FileUploader";
import MovieForm from "@/components/ui/layout/MovieForm";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import { useRouter } from "next/router";
import { FC, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { CinemaService } from "services/cinema.service";
import { MovieService } from "services/movie.service";
import { IHall } from "shared/interfaces/cinema.interface";
import {
  IListOfDirector,
  IListOfFilmStudio,
  IListOfGenres,
  IListOfOperators,
  IMovie,
  IMovieDto,
  IMovieExportDto,
} from "shared/interfaces/movie.interface";
import styles from "./MainHalls.module.scss";
import HallItem from "@/components/ui/hall-item/HallItem";

const MainHall: FC = () => {
  const router = useRouter();
  const [halls, setHalls] = useState<IHall[]>([]);
  const { id } = router.query;
  const queryClient = useQueryClient();
  useEffect(() => {
    console.log("router.query.id =", id);
    console.log("router.isReady =", router.isReady);
  }, [id, router.isReady]);

  const cinemaId = router.isReady ? Number(id) : undefined;

  //  const HallsById = (cinemaId?: number) => {
  //     const { data, isLoading } = useQuery({
  //       queryKey: ["movie", cinemaId],
  //       queryFn: () =>
  //         CinemaService.getHallsById(Number(cinemaId)).then((res) => res.data),
  //       enabled: !!cinemaId && !isNaN(Number(cinemaId)),
  //     });
  //     return { data, isLoading };
  //   };

  //   const { mutate } = useMutation({
  //     mutationKey: ["halls", cinemaId],
  //     mutationFn: () =>
  //       CinemaService.getHallsById(Number(cinemaId)),
  //     onSuccess: () => {
  //       alert("Фильм успешно обновлен!");
  //     },
  //     onError: (error) => {
  //       console.error("Ошибка ", error);
  //       alert("Произошла ошибка ");
  //     },
  //   });

  useEffect(() => {
    const fetchCinema = async () => {
      try {
        const { data } = await CinemaService.getHallsById(Number(cinemaId)); // использует токен из куки
        setHalls(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setHalls([]);
      } finally {
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
        Залы
      </div>
      <div>Залы кинотеатра с ID: {cinemaId}</div>

      <div className={styles.items}>
        {halls.length ? (
          halls.map((hall) => <HallItem halls={hall} key={hall.id} />)
        ) : (
          <div>Movies not found</div>
        )}
      </div>
    </DashboardLayout>
  );
};
export default MainHall;
