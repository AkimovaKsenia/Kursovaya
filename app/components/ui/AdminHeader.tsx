import { FC } from "react";
import Logo from "./layout/header/Logo";
import LoginForm from "./layout/header/login-form/LoginForm";

import styles from "./AdminHeader.module.scss";
const AdminHeader: FC = () => {
  return (
    <header className={styles.header}>
      <Logo />
      <LoginForm />
    </header>
  );
};
export default AdminHeader;
