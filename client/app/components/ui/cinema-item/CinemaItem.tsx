import { FC, useEffect, useState } from "react";
import { IMovie } from "../../../shared/interfaces/movie.interface";
import Link from "next/link";
import styles from "./CinemaItem.module.scss";
import Image from "next/image";
import { PiPencil, PiTrash } from "react-icons/pi";
import { MovieService } from "services/movie.service";
import { useRouter } from "next/router";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ICinemaMain } from "shared/interfaces/cinema.interface";
import { CinemaService } from "services/cinema.service";
import Modal from "../Modal";

const CinemaItem: FC<{ cinema: ICinemaMain }> = ({ cinema }) => {
  const router = useRouter();
  const [showModal, setShowModal] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  // const handleDelete = async (id: number) => {
  //   try {
  //     console.log("Удаляем кинотеартр с ID:", id);
  //     await CinemaService.deleteCinema(id);
  //     router.reload();

  //     alert("Кинотеатр удалён");
  //   } catch (error) {
  //     console.error("Ошибка при удалении фильма:", error);
  //     alert("Ошибка при удалении");
  //   }
  // };
  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      await CinemaService.deleteCinema(cinema.id);
      router.reload();
    } catch (error) {
      alert("Ошибка при удалении");
    } finally {
      setIsDeleting(false);
      setShowModal(false);
    }
  };
  useEffect(() => {
    router.prefetch(`/manage/cinema/halls/${cinema.id}`);
    router.prefetch(`/manage/cinema/edit/${cinema.id}`);
  }, []);

  return (
    <div className={styles.main}>
      <Modal
        isOpen={showModal}
        title="Удалить кинотеатр?"
        onClose={() => setShowModal(false)}
        onConfirm={handleDelete}
      ></Modal>
      <div className={styles.content}>
        <Link href={`/manage/cinema/${cinema.id}`}>
          <div className={styles.heading}>{cinema.name}</div>
        </Link>
        <div className={styles.address}>{cinema.address} </div>
      </div>

      <div className={styles.buttons}>
        <button className={styles.button}>
          <Link href="/" style={{ cursor: "pointer" }}>
            Сеансы
          </Link>
        </button>

        <button className={styles.button}>
          <Link
            href={`/manage/cinema/halls/${cinema.id}`}
            style={{
              cursor: "pointer",
            }}
          >
            Залы
          </Link>
        </button>
      </div>

      <div className={styles.icons}>
        <Link href={`/manage/cinema/edit/${cinema.id}`} passHref legacyBehavior>
          <PiPencil className={styles.firsticon} />
        </Link>
        <PiTrash
          className={styles.firsticon}
          onClick={() => setShowModal(true)}
        />
      </div>
    </div>
  );
};
export default CinemaItem;
