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
        {/* Неоновый заголовок */}
        <h1 className={styles.neonTitle}>{data.name}</h1>

        <div className={styles.glassContainer}>
          <div className={styles.contentWrapper}>
            {/* Постер */}
            {data.photo && (
              <div className={styles.posterFrame}>
                <Image
                  src={data.photo}
                  alt={data.name}
                  width={300} // меньше ширина
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

            {/* Информация о фильме */}
            <div className={styles.infoSection}>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Жанр:</span>
                <span>{data.genres}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Режиссер:</span>
                <span>{data.directors}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Каст Актеров:</span>
                <span>{data.cast_list}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Операторы:</span>
                <span>{data.operators}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Киностудия:</span>
                <span>{data.film_studio_name}</span>
              </div>
              <div className={styles.infoItem}>
                <span className={styles.infoLabel}>Длительность:</span>
                <span>{data.duration_in_min}</span>
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
