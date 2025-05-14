import { FC } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./MovieItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";
import { MovieService } from "services/movie.service";
import { useRouter } from "next/router";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const MovieItem: FC<{ movie: IMovie }> = ({ movie }) => {
  const router = useRouter();

  const handleDelete = async (id: number) => {
    try {
      console.log("Удаляем фильм с ID:", id);
      await MovieService.deleteMovie(id);
      router.reload();

      alert("Фильм удалён");
    } catch (error) {
      console.error("Ошибка при удалении фильма:", error);
      alert("Ошибка при удалении");
    }
  };

  return (
    <div className={styles.main}>
      <Link href={`/movie/${movie.id}`} className={styles.item}>
        <div className={styles.imageWrapper}>
          <Image
            className={styles.poster}
            src={movie.photo}
            alt={movie.name}
            width={600}
            height={733}
            style={{ objectFit: "contain" }}
            // layout="responsive"
          />
        </div>
      </Link>

      <div className={styles.content}>
        <div className={styles.heading}>{movie.name}</div>
        <div className={styles.headingsecond}>
          Длительность: {movie.duration_in_min} минут
        </div>
        <div className={styles.headingsecond}>
          Киностудия: {movie.film_studio_name}
        </div>
        {movie.genres && movie.genres.length > 0 && (
          <div className={styles.headingsecond}>
            Жанры: {movie.genres.join(", ")}
          </div>
        )}
        {movie.directors && movie.directors.length > 0 && (
          <div className={styles.headingsecond}>
            Продюссеры: {movie.directors.join(", ")}
          </div>
        )}
        {movie.operators && movie.operators.length > 0 && (
          <div className={styles.headingsecond}>
            Операторы: {movie.operators.join(", ")}
          </div>
        )}
      </div>

      <div className={styles.icons}>
        <Link href={`/manage/movies/edit/${movie.id}`} passHref legacyBehavior>
          <PiPencil className={styles.firsticon} />
        </Link>
        <PiTrash
          className={styles.firsticon}
          onClick={() => handleDelete(movie.id)}
        />
      </div>
    </div>
  );
};
export default MovieItem;
