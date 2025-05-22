import {
  PiArrowBendUpLeftBold,
  PiChartBarDuotone,
  PiFilmSlateDuotone,
  PiFilmStrip,
  PiUserCirclePlusLight,
  PiUserCirclePlus,
} from "react-icons/pi";
import { IconType } from "react-icons";
export interface IMenuItem {
  link: string;
  Icon: IconType;
}

export const menu: IMenuItem[] = [
  {
    link: "/",
    Icon: PiArrowBendUpLeftBold,
  },
  {
    link: "/dashboard",
    Icon: PiUserCirclePlus,
  },
  {
    link: "/manage/movies/listmovies",
    Icon: PiFilmSlateDuotone,
  },
  {
    link: "/manage/cinema/listcinema",
    Icon: PiFilmStrip,
  },
];
