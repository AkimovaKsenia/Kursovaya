import { FC, useState } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./HallItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";
import { MovieService } from "services/movie.service";
import { useRouter } from "next/router";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ICinemaMain, IHall } from "shared/interfaces/cinema.interface";
import Modal from "../Modal";
import { CinemaService } from "services/cinema.service";

const HallItem: FC<{ halls: IHall }> = ({ halls }) => {
  const router = useRouter();
  const [showModal, setShowModal] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  // const handleDelete = async (id: number) => {
  //   try {
  //     console.log("Удаляем фильм с ID:", id);
  //     await MovieService.deleteMovie(id);
  //     router.reload();

  //     alert("Фильм удалён");
  //   } catch (error) {
  //     console.error("Ошибка при удалении фильма:", error);
  //     alert("Ошибка при удалении");
  //   }
  // };
  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      await CinemaService.deleteHall(halls.id);
      router.reload();
    } catch (error) {
      alert("Ошибка при удалении");
    } finally {
      setIsDeleting(false);
      setShowModal(false);
    }
  };

  return (
    <div className={styles.main}>
      <Modal
        isOpen={showModal}
        title="Удалить зал?"
        onClose={() => setShowModal(false)}
        onConfirm={handleDelete}
      ></Modal>
      <div className={styles.content}>
        <div className={styles.heading}>{halls.name}</div>
        <div className={styles.hallInfo}>
          <span className={styles.hallType}>{halls.type}</span>
          <span className={styles.hallCapacity}>{halls.capacity} мест</span>
        </div>
      </div>

      {/* <div className={styles.buttons}>
        <button className={styles.button}>
          <Link href="/" style={{ cursor: "pointer" }}>
            Сеансы
          </Link>
        </button>
      </div> */}

      <div className={styles.icons}>
        <Link
          href={`/manage/cinema/halls/edit/${halls.id}`}
          passHref
          legacyBehavior
        >
          <PiPencil
            className={styles.firsticon}
            onClick={() => router.push(`/manage/cinema/halls/edit/${halls.id}`)}
          />
        </Link>
        <PiTrash
          className={styles.firsticon}
          onClick={() => setShowModal(true)}
        />
      </div>
    </div>
  );
};
export default HallItem;
