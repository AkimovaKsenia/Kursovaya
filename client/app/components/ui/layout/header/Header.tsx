import { FC } from "react";
import Logo from "./Logo";
import LoginForm from "./login-form/LoginForm";

import styles from "./Header.module.scss";
const Header: FC = () => {
  return (
    <header className={styles.header}>
      <Logo />
      <LoginForm />
    </header>
  );
};
export default Header;
