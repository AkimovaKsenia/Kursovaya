import { NextPage } from "next";
import Dashboard from "screens/dashboard/MainDasboard";
import MoviesInfo from "screens/movies/MovieInfo";

// Создаем компонент, который рендерит Movies
const DashboardPage: NextPage = () => {
  return <Dashboard />;
};

export default DashboardPage;
