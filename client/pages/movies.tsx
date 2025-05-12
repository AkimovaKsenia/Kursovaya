import { NextPage, GetStaticProps } from "next";
// Импортируем Movies с правильного пути
import Movies from "../app/screens/dashboard/movies/AllMovies";
import { MovieService } from "services/movie.service";
import { IListOfMovies } from "shared/interfaces/movie.interface";
import { axiosClassic } from "api/interceptor";

const MoviesPage: NextPage = () => {
  return <div className="bg-black text-green">Сеансы</div>;
};

// export const getStaticProps: GetStaticProps<IListOfMovies> = async () => {
//   try {
//     const { data: newMovies } = await MovieService.getAll();
//     return {
//       props: {
//         newMovies,
//       },
//       revalidate: 60,
//     };
//   } catch (e) {
//     return {
//       props: {
//         newMovies: [],
//       },
//     };
//   }
// };

export default MoviesPage;
