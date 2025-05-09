import { FC } from "react";
import Logo from "./Logo";
import LoginForm from "./login-form/LoginForm";

interface HeaderProps {
  className?: string;
  wrapperClassName?: string;
}
import styles from "./Header.module.scss";
const Header: FC<HeaderProps> = ({ className, wrapperClassName }) => {
  return (
    <div className={wrapperClassName}>
      <header className={`${styles.header} ${className}`}>
        <Logo />
        <LoginForm />
      </header>
    </div>
  );
};
export default Header;
