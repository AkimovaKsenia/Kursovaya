import styles from "./AdminHeader.module.scss";
import { FC, PropsWithChildren } from "react";
import Meta from "utils/Meta";
import { IMeta } from "utils/meta.interface";
import Header from "./header/Header";
import Sidebar from "./sidebar/Sidebar";
import { useAuth } from "hooks/useAuth";
import cn from "classnames";

const AdminLayout: FC<PropsWithChildren<IMeta>> = ({
  children,
  backgroundColor,
  ...meta
}) => {
  const { user } = useAuth();
  return (
    <>
      <Meta {...meta} />
      <section
        className={cn("flex min-h-screen", user && styles.wrapper)}
        style={{ backgroundColor: backgroundColor }}
      >
        {user && user.role === "admin" && (
          <Sidebar className={styles.sidebar} />
        )}

        <div className={cn("flex-1 flex flex-col", styles.content)}>
          <Header className="header-wrapper" />

          <main className={cn("flex-1 overflow-auto", styles.main)}>
            {children}
          </main>
        </div>
      </section>
    </>
  );
};
export default AdminLayout;
