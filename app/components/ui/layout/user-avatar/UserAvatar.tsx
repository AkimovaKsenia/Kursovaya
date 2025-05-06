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
        src="https://i.pinimg.com/736x/3c/ae/07/3cae079ca0b9e55ec6bfc1b358c9b1e2.jpg"
        width={40}
        height={40}
        alt=""
      />
    </Link>
  );
};
export default UserAvatar;
