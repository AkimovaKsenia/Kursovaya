import { NextPage, GetStaticProps } from "next";
// Импортируем Movies с правильного пути
import Movies from "../app/screens/movies/Movies";
import { MovieService } from "services/movie.service";

// Создаем компонент, который рендерит Movies
const MoviesPage: NextPage = () => {
  return <Movies />;
};

export const getStaticProps: GetStaticProps = async () => {
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
