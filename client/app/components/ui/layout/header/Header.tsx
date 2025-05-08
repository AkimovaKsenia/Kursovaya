import { FC } from "react";
import Logo from "./Logo";
import LoginForm from "./login-form/LoginForm";

interface HeaderProps {
  className?: string; // Явно указываем, что компонент принимает className
}
import styles from "./Header.module.scss";
const Header: FC<HeaderProps> = ({ className }) => {
  return (
    <header className={styles.header}>
      <Logo />
      <LoginForm />
    </header>
  );
};
export default Header;
