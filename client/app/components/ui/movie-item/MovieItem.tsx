import { FC } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./MovieItem.module.scss";
import Image from "next/image";

const MovieItem: FC<{ movie: IMovie }> = ({ movie }) => {
  return (
    <Link href={`/movie/${movie.id}`} className={styles.item}>
      <div>
        <div className={styles.poster}>
          <Image
            src={movie.poster}
            alt={movie.name}
            width={220}
            height={330}
            layout="responsive"
          />
        </div>
        <div className={styles.heading}>{movie.name}</div>
      </div>
    </Link>
  );
};
export default MovieItem;
