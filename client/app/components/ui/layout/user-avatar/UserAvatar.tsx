import cn from "classnames";
import Image from "next/image";
import Link from "next/link";
import { FC } from "react";
import styles from "./UserAvatar.module.scss";

const UserAvatar: FC<{ link: string; title?: string }> = ({ link, title }) => {
  return (
    <Link href={link} title={title}>
      <img
        className={styles.avatar}
        src="/User.png"
        width={40}
        height={40}
        alt=""
      />
    </Link>
  );
};
export default UserAvatar;
