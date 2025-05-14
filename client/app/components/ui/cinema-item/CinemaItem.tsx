import { FC } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "../movie-item/MovieItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";
import { MovieService } from "services/movie.service";
import { useRouter } from "next/router";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ICinemaMain } from "shared/interfaces/cinema.interface";

const CinemaItem: FC<{ cinema: ICinemaMain }> = ({ cinema }) => {
  const router = useRouter();

  // const handleDelete = async (id: number) => {
  //   try {
  //     console.log("Удаляем фильм с ID:", id);
  //     await MovieService.deleteMovie(id);
  //     router.reload();

  //     alert("Фильм удалён");
  //   } catch (error) {
  //     console.error("Ошибка при удалении фильма:", error);
  //     alert("Ошибка при удалении");
  //   }
  // };

  return (
    <div className={styles.main}>
      <Link href={`/movie/${cinema.id}`} className={styles.item}></Link>

      <div className={styles.content}>
        <div className={styles.heading}>{cinema.name}</div>
        <div className={styles.heading}>{cinema.address} </div>
      </div>

      <div className={styles.icons}>
        {/* <Link href={`/manage/movies/edit/${cinema.id}`} passHref legacyBehavior> */}
        <PiPencil className={styles.firsticon} />
        {/* </Link> */}
        <PiTrash
          className={styles.firsticon}
          // onClick={() => handleDelete(cinema.id)}
        />
      </div>
    </div>
  );
};
export default CinemaItem;
