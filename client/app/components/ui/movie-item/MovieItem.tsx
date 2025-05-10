import { FC } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./MovieItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";

const MovieItem: FC<{ movie: IMovie }> = ({ movie }) => {
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
        <div className={styles.heading}>{movie.duration_in_min} минут</div>
      </div>

      <div className={styles.icons}>
        <PiPencil className={styles.firsticon} />
        <PiTrash className={styles.firsticon} />
      </div>
    </div>
  );
};
export default MovieItem;
