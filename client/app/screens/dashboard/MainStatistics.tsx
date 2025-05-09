import StatisticItem from "@/components/ui/statistic-item/StatisticItem";
import { PiFilmReelLight } from "react-icons/pi";
import { PiEye } from "react-icons/pi";

import { FC } from "react";
const MainStatistic: FC = () => {
  return (
    <div className="grid grid-cols-3 gap-14">
      <StatisticItem
        item={{
          name: "Views",
          value: 200,
          Icon: PiFilmReelLight,
          image: "/Purple.png",
        }}
      />
      <StatisticItem
        item={{
          name: "Views",
          value: 200,
          Icon: PiEye,
          image: "/Blue.png",
        }}
      />
      <StatisticItem
        item={{
          name: "Views",
          value: 200,
          Icon: PiFilmReelLight,
          image: "/Green.png",
        }}
      />
    </div>
  );
};
export default MainStatistic;
