import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { NextPage, GetStaticProps } from "next";
import CreateCinema from "screens/dashboard/cinema/create/CreateCinema";
import Movies from "screens/dashboard/movies/AllMovies";
import CreateMovie from "screens/dashboard/movies/create/CreateMovie";
import MovieEdit from "screens/dashboard/movies/edit/MovieEdit";

const UserCreatePage: NextPage = () => {
  return <CreateCinema />;
};

export default UserCreatePage;
