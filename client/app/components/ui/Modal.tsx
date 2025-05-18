import { FC, ReactNode } from "react";
import styles from "./Modal.module.scss";

interface ModalProps {
  isOpen: boolean;
  title: string;
  onClose: () => void;
  onConfirm: () => void;
  children?: ReactNode;
}

const Modal: FC<ModalProps> = ({
  isOpen,
  title,
  onClose,
  onConfirm,
  children,
}) => {
  if (!isOpen) return null;

  return (
    <div className={styles.overlay} onClick={onClose}>
      <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
        <div className={styles.header}>
          <h3>{title}</h3>
          <button onClick={onClose} className={styles.closeBtn}>
            &times;
          </button>
        </div>
        <div className={styles.body}>{children}</div>
        <div className={styles.footer}>
          <button className={styles.cancel} onClick={onClose}>
            Отмена
          </button>
          <button className={styles.confirm} onClick={onConfirm}>
            Удалить
          </button>
        </div>
      </div>
    </div>
  );
};

export default Modal;
