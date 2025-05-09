import { FC } from "react";
import styles from "./Sidebar.module.scss";
import Link from "next/link";
import { useRouter } from "next/router";
import { menu } from "./menu.data";
import cn from "classnames";
import { PiFilmReelLight } from "react-icons/pi";

interface SidebarProps {
  className?: string; // Явно указываем, что компонент принимает className
}

const Sidebar: FC<SidebarProps> = ({ className }) => {
  const { asPath } = useRouter();
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
                  [styles.active]: item.link == asPath,
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
