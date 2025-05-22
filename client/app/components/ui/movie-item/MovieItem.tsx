import { FC, useState } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./MovieItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";
import { MovieService } from "services/movie.service";
import { useRouter } from "next/router";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import Modal from "../Modal";
import cn from "classnames";

const MovieItem: FC<{ movie: IMovie }> = ({ movie }) => {
  const router = useRouter();
  const [showModal, setShowModal] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      await MovieService.deleteMovie(movie.id);
      router.reload();
    } catch (error) {
      alert("Ошибка при удалении");
    } finally {
      setIsDeleting(false);
      setShowModal(false);
    }
  };
  return (
    <div className={styles.main}>
      <Modal
        isOpen={showModal}
        title="Удалить фильм"
        onClose={() => setShowModal(false)}
        onConfirm={handleDelete}
      >
        <p>Вы уверены, что хотите удалить фильм "{movie.name}"?</p>
      </Modal>
      <Link href={`/movie/${movie.id}`} className={styles.item}>
        <div className={styles.imageWrapper}>
          {movie.photo && (
            <div className={styles.posterFrame}>
              <Image
                src={movie.photo}
                alt={movie.name}
                width={300}
                height={450}
                className={cn(styles.posterImage, "rounded-lg")}
              />
            </div>
          )}
        </div>
      </Link>

      <div className={styles.content}>
        <div className={styles.heading}>{movie.name}</div>
        <div className={styles.headingsecond}>
          Длительность: {Math.floor(movie.duration_in_min / 60)} ч{"  "}
          {movie.duration_in_min % 60} мин
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
          onClick={() => setShowModal(true)}
        />
      </div>
    </div>
  );
};
export default MovieItem;
