import {
  PiArrowBendUpLeftBold,
  PiChartBarDuotone,
  PiFilmSlateDuotone,
  PiFilmStrip,
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
    Icon: PiChartBarDuotone,
  },
  {
    link: "/manage/movies",
    Icon: PiFilmSlateDuotone,
  },
  {
    link: "/manage/cinema",
    Icon: PiFilmStrip,
  },
];
