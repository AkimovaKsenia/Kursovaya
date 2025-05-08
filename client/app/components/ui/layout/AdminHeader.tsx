import styles from "./AdminHeader.module.scss";
import { FC, PropsWithChildren } from "react";
import Meta from "utils/Meta";
import { IMeta } from "utils/meta.interface";
import Header from "./header/Header";
import Sidebar from "./sidebar/Sidebar";
import { useAuth } from "hooks/useAuth";

const AdminLayout: FC<PropsWithChildren<IMeta>> = ({
  children,
  backgroundColor,
  ...meta
}) => {
  const { user } = useAuth();
  return (
    <>
      <Meta {...meta} />
      <section className={user ? styles.header : ""}>
        {user && <Sidebar className={styles.sidebar} />}
        <div className={styles.content}>
          <Header className={styles["header-component"]} />
          <main className={styles.main}>{children}</main>
        </div>
      </section>
    </>
  );
};
export default AdminLayout;
