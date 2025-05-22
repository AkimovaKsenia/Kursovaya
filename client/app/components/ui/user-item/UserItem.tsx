import { FC, useEffect, useState } from "react";
import { useRouter } from "next/router";
import { PiPencil, PiTrash } from "react-icons/pi";
import { IUserMain } from "shared/interfaces/user.interface";
import styles from "./UserItem.module.scss";
import Modal from "../Modal";
import { UserService } from "services/user.service";

const UserItem: FC<{ user: IUserMain }> = ({ user }) => {
  const router = useRouter();
  const [isDeleting, setIsDeleting] = useState(false);
  const [showModal, setShowModal] = useState(false);

  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      await UserService.deleteUser(user.id);
      router.reload();
    } catch (error) {
      alert("Ошибка при удалении");
    } finally {
      setIsDeleting(false);
      setShowModal(false);
    }
  };
  return (
    <div className={styles.wrapper}>
      <Modal
        isOpen={showModal}
        title="Удалить пользователя?"
        onClose={() => setShowModal(false)}
        onConfirm={handleDelete}
      />

      <div className={styles.card}>
        <div className={styles.content}>
          <div className={styles.heading}>
            {user.name} {user.surname}
          </div>
          <div className={styles.address}>{user.email}</div>
          <div className={styles.address}>{user.role}</div>
        </div>

        <div className={styles.icons}>
          <PiPencil className={styles.firsticon} />
          <PiTrash
            className={styles.firsticon}
            onClick={() => setShowModal(true)}
          />
        </div>
      </div>
    </div>
  );
};

export default UserItem;
