import Layout from "@/components/ui/layout/Layout";
import { FC } from "react";

// Экспортируем компонент как default
const Movies: FC = () => {
  return (
    <Layout title="Cinema">
      <div style={{ background: "blue" }}>MOVIES</div>
    </Layout>
  );
};

export default Movies;
