import { FC, useEffect, useState } from "react";
import { useRouter } from "next/router";
import { PiPencil, PiTrash } from "react-icons/pi";
import { IUserMain } from "shared/interfaces/user.interface";
import styles from "./UserItem.module.scss";
import Modal from "../Modal";

const UserItem: FC<{ user: IUserMain }> = ({ user }) => {
  const router = useRouter();
  const [showModal, setShowModal] = useState(false);

  useEffect(() => {
    router.prefetch(`/manage/cinema/edit/${user.id}`);
  }, []);

  return (
    <div className={styles.wrapper}>
      {/* <Modal
        isOpen={showModal}
        title="Удалить кинотеатр?"
        onClose={() => setShowModal(false)}
        onConfirm={handleDelete}
      /> */}

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
