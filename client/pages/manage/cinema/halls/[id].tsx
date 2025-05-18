import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { NextPage, GetStaticProps } from "next";
import MainHall from "screens/dashboard/cinema/halls/MainHalls";
import Movies from "screens/dashboard/movies/AllMovies";
import MovieEdit from "screens/dashboard/movies/edit/MovieEdit";

const HallsPage: NextPage = () => {
  return <MainHall />;
};

export default HallsPage;
