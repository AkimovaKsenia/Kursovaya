import { FC, useEffect, useState } from "react";
import styles from "./AllCinema.module.scss";
import DashboardLayout from "../../../components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import { useRouter } from "next/router";
import { CinemaService } from "services/cinema.service";
import { ICinema, ICinemaMain } from "shared/interfaces/cinema.interface";
import CinemaItem from "@/components/ui/cinema-item/CinemaItem";
import { useAuth } from "hooks/useAuth";
import ErrorAuth from "@/components/ui/ErrorAuth";

const Cinema: FC = () => {
  const { user, setUser } = useAuth(); // Получаем данные пользователя из контекста
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

  useEffect(() => {
    router.prefetch("/manage/cinema/createcinema");
  }, []);

  return (
    <DashboardLayout>
      {user ? (
        <div>
          <div
            style={{ background: "rgba(129, 125, 219, 0.8)" }}
            className={styles.catalog}
          >
            Кинотеатры
          </div>
          <Link href="/manage/cinema/createcinema" className={styles.link}>
            <button className={styles.button}> Создать кинотеатр </button>
          </Link>

          <div className={styles.items}>
            {cinema.length ? (
              cinema.map((cinema) => (
                <CinemaItem cinema={cinema} key={cinema.id} />
              ))
            ) : (
              <div>Cinema not found</div>
            )}
          </div>
        </div>
      ) : (
        <ErrorAuth />
      )}
    </DashboardLayout>
  );
};

export default Cinema;
