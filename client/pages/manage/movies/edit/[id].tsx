import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { NextPage, GetStaticProps } from "next";
import Movies from "screens/dashboard/movies/AllMovies";
import MovieEdit from "screens/dashboard/movies/edit/MovieEdit";

const MoviesEditPage: NextPage = () => {
  return <MovieEdit />;
};

export default MoviesEditPage;
