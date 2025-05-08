import { FC } from "react";
import styles from "./Sidebar.module.scss";
import Link from "next/link";
import { useRouter } from "next/router";
import { menu } from "./menu.data";
import cn from "classnames";

interface SidebarProps {
  className?: string; // Явно указываем, что компонент принимает className
}

const Sidebar: FC<SidebarProps> = ({ className }) => {
  const { asPath } = useRouter();
  return (
    <aside className={styles.sidebar}>
      <Link href="/" className={styles.logo}>
        R
      </Link>
      <nav className={styles.menu}>
        <ul>
          {menu.map((item) => (
            <Link href={item.link} key={item.link}>
              <div
                className={cn(styles.item, {
                  [styles.active]: item.link == asPath,
                })}
              >
                <item.Icon />
              </div>
            </Link>
          ))}
        </ul>
      </nav>
    </aside>
  );
};
export default Sidebar;
