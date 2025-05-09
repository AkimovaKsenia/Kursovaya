import Layout from "@/components/ui/layout/Layout";
import MovieItem from "@/components/ui/movie-item/MovieItem";
import { FC } from "react";
import { IListOfMovies } from "shared/interfaces/movie.interface";
import styles from "./AllMovies.module.scss";

const Movies: FC<IListOfMovies> = ({ newMovies }) => {
  return (
    <Layout title="Cinema">
      <div style={{ background: "blue" }} className={styles.catalog}>
        MOVIES
      </div>
      {newMovies.length ? (
        newMovies.map((movie) => <MovieItem movie={movie} key={movie.id} />)
      ) : (
        <div>Movies not found</div>
      )}
    </Layout>
  );
};

export default Movies;
