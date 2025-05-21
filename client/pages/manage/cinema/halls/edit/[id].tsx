import { NextPage, GetStaticProps } from "next";
import HallEdit from "screens/dashboard/cinema/halls/edit/HallEdit";
import MainHall from "screens/dashboard/cinema/halls/MainHalls";
import MovieEdit from "screens/dashboard/movies/edit/MovieEdit";

const HallEditPage: NextPage = () => {
  return <HallEdit />;
};

export default HallEditPage;
