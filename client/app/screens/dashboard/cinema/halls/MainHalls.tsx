import Button from "@/components/ui/layout/Button/Button";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { useRouter } from "next/router";
import { FC, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { CinemaService } from "services/cinema.service";
import { MovieService } from "services/movie.service";
import { IHall } from "shared/interfaces/cinema.interface";
import styles from "./MainHalls.module.scss";
import HallItem from "@/components/ui/hall-item/HallItem";
import Link from "next/link";
import { useAuth } from "hooks/useAuth";
import ErrorAuth from "@/components/ui/ErrorAuth";

const MainHall: FC = () => {
  const router = useRouter();
  const { user, setUser } = useAuth();
  const [halls, setHalls] = useState<IHall[]>([]); // Всегда массив
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const cinemaId = Number(router.query.id);
  useEffect(() => {
    if (!router.isReady) return;

    const fetchHalls = async () => {
      setIsLoading(true);
      try {
        const cinemaId = Number(router.query.id);

        // Если ID отсутствует или невалиден, просто возвращаем пустой список
        if (!router.query.id || isNaN(cinemaId)) {
          setHalls([]);
          return;
        }

        const response = await CinemaService.getHallsById(cinemaId);
        const data = Array.isArray(response?.data) ? response.data : [];
        setHalls(data);
        setError(null);
      } catch (error) {
        console.error("Ошибка при загрузке залов:", error);
        setHalls([]);
        setError("Ошибка загрузки данных");
      } finally {
        setIsLoading(false);
      }
    };

    fetchHalls();
  }, [router.isReady, router.query.id]);

  if (!router.isReady || isLoading) {
    return (
      <DashboardLayout>
        <div className={styles.loading}>Загрузка данных...</div>
      </DashboardLayout>
    );
  }

  return (
    <DashboardLayout>
      {user ? (
        <div>
          <div
            style={{ background: "rgba(129, 125, 219, 0.8)" }}
            className={styles.catalog}
          >
            Залы
          </div>
          <Link href={`/manage/cinema/halls/create/${cinemaId}`}>
            <button className={styles.button}> Создать зал </button>
          </Link>
          {error ? (
            <div className={styles.error}>{error}</div>
          ) : (
            <>
              <div className={styles.items}>
                {halls.length === 0 ? (
                  <div className={styles.empty}>Залы не найдены</div>
                ) : (
                  halls.map((hall) => <HallItem key={hall.id} halls={hall} />)
                )}
              </div>
            </>
          )}{" "}
        </div>
      ) : (
        <ErrorAuth />
      )}
    </DashboardLayout>
  );
};
export default MainHall;
