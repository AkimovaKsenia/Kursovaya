import { NextPage, GetStaticProps } from "next";
// Импортируем Movies с правильного пути
import Movies from "../app/screens/movies/AllMovies";
import { MovieService } from "services/movie.service";
import { IListOfMovies } from "shared/interfaces/movie.interface";

// Создаем компонент, который рендерит Movies
const MoviesPage: NextPage<IListOfMovies> = (props) => {
  return <Movies {...props} />;
};

export const getStaticProps: GetStaticProps<IListOfMovies> = async () => {
  try {
    const { data: newMovies } = await MovieService.getAll();
    return {
      props: {
        newMovies,
      },
      revalidate: 60,
    };
  } catch (e) {
    return {
      props: {
        newMovies: [],
      },
    };
  }
};

export default MoviesPage;
