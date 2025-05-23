import Layout from "@/components/ui/layout/Layout";
import { useMovieById } from "hooks/useMovieById";
import Image from "next/image";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import styles from "./MovieInfo.module.scss";

const MoviesInfo: FC = () => {
  const router = useRouter();

  const { query } = useRouter();
  const movieId = Number(query?.id);

  //АВТООБНОВЛЕНИЕ ПРОСМОТРОВ СТРАНИЦЫ
  // const { mutateAsync } = useMutation({
  //   mutationKey: ["update count opened"],
  //   mutationFn: () => ViewsService.updateViews(movieId.toString()),
  // });

  const { data, isLoading } = useMovieById(movieId);
  if (isLoading) return <div className={styles.loading}>Загрузка...</div>;
  if (!data) return <div className={styles.notFound}>Фильм не найден</div>;
  return (
    <Layout title={`${data.name}`} backgroundColor="#0f0f12">
      <div className={styles.main}>
        <h1 className={styles.neonTitle}>{data.name}</h1>

        <div className={styles.glassContainer}>
          <div className={styles.contentWrapper}>
            {data.photo && (
              <div className={styles.posterFrame}>
                <Image
                  src={data.photo}
                  alt={data.name}
                  width={300}
                  height={450}
                  className={styles.posterImage}
                />
              </div>
            )}
            {/* <div className={styles.navigation}>
              <button
                className={styles.navButton}
                onClick={() => router.push(`/movies/${movieId - 1}`)}
                disabled={movieId <= 1}
              >
                ← Назад
              </button>
              <button
                className={styles.navButton}
                onClick={() => router.push(`/movies/${movieId + 1}`)}
              >
                Вперёд →
              </button>
            </div> */}

            <div className={styles.infoSection}>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Жанр:</span>

                {data.genres && data.genres.length > 0 && (
                  <span className={styles.info}> {data.genres.join(", ")}</span>
                )}
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Режиссеры:</span>
                {data.directors && data.directors.length > 0 && (
                  <span className={styles.info}>
                    {" "}
                    {data.directors.join(", ")}
                  </span>
                )}
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Каст Актеров:</span>
                {data.cast_list && data.cast_list.length > 0 && (
                  <span className={styles.info}>
                    {" "}
                    {data.cast_list.join(", ")}
                  </span>
                )}{" "}
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Операторы:</span>
                {data.operators && data.operators.length > 0 && (
                  <span className={styles.info}>
                    {" "}
                    {data.operators.join(", ")}
                  </span>
                )}{" "}
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Киностудия:</span>
                <span className={styles.info}>{data.film_studio_name}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Длительность:</span>
                <span className={styles.info}>
                  {Math.floor(data.duration_in_min / 60)} ч{"  "}
                  {data.duration_in_min % 60} мин{" "}
                </span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Описание:</span>
                <span className={styles.infoValue}>{data.description}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default MoviesInfo;
