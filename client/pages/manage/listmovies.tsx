import { NextPage, GetStaticProps } from "next";

import { MovieService } from "services/movie.service";
import { IListOfMovies } from "shared/interfaces/movie.interface";
import { axiosClassic } from "api/interceptor";
import Movies from "screens/movies/AllMovies";

const MoviesPage: NextPage = () => {
  return <Movies />;
};

export default MoviesPage;
