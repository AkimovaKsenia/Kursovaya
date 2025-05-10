import { FC, PropsWithChildren } from "react";
import { FaSignOutAlt } from "react-icons/fa"; // Иконка выхода
import AdminLayout from "@/components/ui/layout/AdminHeader";
import { useAuth } from "hooks/useAuth";
import Link from "next/link";
import MainStatistic from "../../../screens/dashboard/MainStatistics";
import styles from "./DashboardLayout.module.scss";

const DashboardLayout: FC<PropsWithChildren> = ({ children }) => {
  return (
    <AdminLayout title="Dashboard" backgroundColor="#1F1F1F">
      <div className={styles.container}>
        <div className={styles.main}>
          <main>{children}</main>
        </div>
      </div>
    </AdminLayout>
  );
};

export default DashboardLayout;
