import StatisticItem from "@/components/ui/statistic-item/StatisticItem";
import { PiFilmReelLight } from "react-icons/pi";

import { FC } from "react";
const MainStatistic: FC = () => {
  return (
    <div className="grid gris-cols-3 gap-8">
      <StatisticItem
        item={{
          name: "Views",
          value: 200,
          Icon: PiFilmReelLight,
        }}
      />
    </div>
  );
};
export default MainStatistic;
