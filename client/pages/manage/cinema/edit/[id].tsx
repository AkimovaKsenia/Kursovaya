import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { NextPage, GetStaticProps } from "next";
import CinemaEdit from "screens/dashboard/cinema/edit/CinemaEdit";

const CinemaEditPage: NextPage = () => {
  return <CinemaEdit />;
};

export default CinemaEditPage;
