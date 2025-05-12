import { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import styles from "./FileUploader.module.scss"; // Импорт стилей

interface FileUploaderProps {
  onFilesUploaded: (files: File[]) => void;
}

const FileUploader = ({ onFilesUploaded }: FileUploaderProps) => {
  const [preview, setPreview] = useState<string | null>(null);

  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      if (acceptedFiles.length > 0) {
        const file = acceptedFiles[0];
        onFilesUploaded([file]);

        if (file.type.startsWith("image/")) {
          const reader = new FileReader();
          reader.onload = () => setPreview(reader.result as string);
          reader.readAsDataURL(file);
        }
      }
    },
    [onFilesUploaded]
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      "image/png": [".png"],
      "image/jpeg": [".jpg", ".jpeg"],
    },
    maxFiles: 1,
  });

  return (
    <div
      {...getRootProps()}
      className={`${styles["upload-area"]} ${
        isDragActive ? styles["upload-area--active"] : ""
      }`}
    >
      <input {...getInputProps()} />

      {preview ? (
        <div className={styles["preview-container"]}>
          <img
            src={preview}
            alt="Preview"
            className={styles["preview-image"]}
          />
          <p>Нажмите или перетащите для изменения</p>
        </div>
      ) : (
        <div className={styles["upload-placeholder"]}>
          {isDragActive ? (
            <p>Отпустите файл для загрузки</p>
          ) : (
            <>
              <p>Перетащите или загрузите фото PNG/JPEG</p>
              <button type="button" className={styles["browse-button"]}>
                Выбрать файл
              </button>
            </>
          )}
        </div>
      )}
    </div>
  );
};

export default FileUploader;
