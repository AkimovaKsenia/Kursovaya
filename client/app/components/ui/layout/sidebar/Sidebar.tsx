import { FC, useEffect } from "react";
import styles from "./Sidebar.module.scss";
import Link from "next/link";
import { useRouter } from "next/router";
import { menu } from "./menu.data";
import cn from "classnames";
import { PiFilmReelLight } from "react-icons/pi";

interface SidebarProps {
  className?: string;
}

const Sidebar: FC<SidebarProps> = ({ className }) => {
  const { asPath } = useRouter();
  const router = useRouter();

  useEffect(() => {
    router.prefetch("/manage/movies/listmovies");
    router.prefetch("/manage/cinema/listcinema");
  }, []);

  const isActive = (link: string) => {
    if (asPath === link) return true;

    if (link === "/manage/movies/listmovies") {
      return asPath.startsWith("/manage/movies");
    }
    if (link === "/manage/cinema/listcinema") {
      return asPath.startsWith("/manage/cinema");
    }

    return false;
  };
  return (
    <aside className={styles.sidebar}>
      {/* <Link href="/" className={styles.logo}>
        <PiFilmReelLight />
      </Link> */}
      <div>
        <nav className={styles.menu}>
          <ul>
            {menu.map((item) => (
              <li
                key={item.link}
                className={cn(styles.item, {
                  [styles.active]: isActive(item.link),
                })}
              >
                <Link href={item.link}>
                  <item.Icon />
                </Link>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </aside>
  );
};
export default Sidebar;
