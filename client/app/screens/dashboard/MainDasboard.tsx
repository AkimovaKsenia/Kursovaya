import { FC, useEffect, useState } from "react";
import { FaSignOutAlt } from "react-icons/fa"; // Иконка выхода
import AdminLayout from "@/components/ui/layout/AdminHeader";
import { useAuth } from "hooks/useAuth";
import Link from "next/link";
import MainStatistic from "./MainStatistics";
import styles from "./MainDashboard.module.scss";
import DashboardLayout from "../../components/ui/layout/DashboardLayout";
import { IListOfUsers } from "shared/interfaces/user.interface";
import { UserService } from "services/user.service";
import UserItem from "@/components/ui/user-item/UserItem";
import ErrorAuth from "@/components/ui/ErrorAuth";

const Dashboard: FC = () => {
  const { user, setUser } = useAuth();
  const [users, setUsers] = useState<IListOfUsers>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchCinema = async () => {
      try {
        const { data } = await UserService.getAllUsers(); // использует токен из куки
        setUsers(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setUsers([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchCinema();
  }, []);

  return (
    <DashboardLayout>
      {user ? (
        <div>
          <h1 className={styles.heading}>
            Добро пожаловать в вашу панель управления
          </h1>
          <Link href={`/manage/user/create/createuser`}>
            <button className={styles.firstButton}>
              {" "}
              Создать нового пользователя{" "}
            </button>
          </Link>
          {/* <div className="w-full flex justify-center">
                <MainStatistic />
              </div> */}
          <div className={styles.items}>
            {users.length ? (
              users.map((user) => <UserItem user={user} key={user.id} />)
            ) : (
              <div>not found</div>
            )}
          </div>
        </div>
      ) : (
        <ErrorAuth />
      )}
    </DashboardLayout>
  );
};

export default Dashboard;
