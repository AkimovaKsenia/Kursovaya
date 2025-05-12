import {
  UseFormRegister,
  FieldErrors,
  UseFormHandleSubmit,
  Control,
} from "react-hook-form";
import { IListOfGenres, IMovieDto } from "shared/interfaces/movie.interface";
import { FC } from "react";
import FileUploader from "./FileUploader";
import cn from "classnames";
import styles from "./MovieForm.module.scss";
interface MovieFormProps {
  register: UseFormRegister<IMovieDto>;
  errors: FieldErrors<IMovieDto>;
  handleSubmit: UseFormHandleSubmit<IMovieDto>;
  onSubmit: (formData: IMovieDto) => void;
  isPending: boolean;
  genresData: IListOfGenres | undefined;
  handleFileUpload: (files: File[]) => void;
}

const MovieForm: FC<MovieFormProps> = ({
  register,
  errors,
  handleSubmit,
  onSubmit,
  isPending,
  genresData,
  handleFileUpload,
}) => {
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Название фильма:
        </label>
        <input
          {...register("name", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
        {errors.name && (
          <p className="text-red-500 text-sm mt-1">{errors.name.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Жанр:
        </label>
        <select
          {...register("genres", { required: "Выберите жанры" })}
          className={cn(
            "form-select w-full text-m px-3 py-2 border border-gray-300 rounded-md",
            styles.select
          )}
          multiple
          size={4}
        >
          {genresData?.map((genre) => (
            <option key={genre.id} value={genre.name}>
              {genre.name}
            </option>
          ))}
        </select>
        {errors.genres && (
          <p className="text-red-500 text-sm mt-1">{errors.genres.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Режиссер:
        </label>
        <input
          {...register("directors")}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Оператор:
        </label>
        <input
          {...register("operators")}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Список актеров:
        </label>
        <textarea
          {...register("cast_list")}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
          rows={3}
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Киностудия:
        </label>
        <input
          {...register("film_studio_name")}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Описание:
        </label>
        <textarea
          {...register("description")}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
          rows={5}
        />
      </div>

      <div className="pt-4">
        <button
          type="submit"
          disabled={isPending}
          className={`px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 ${
            isPending ? "opacity-50 cursor-not-allowed" : ""
          }`}
        >
          {isPending ? "Сохранение..." : "Сохранить изменения"}
        </button>
      </div>
      <FileUploader onFilesUploaded={handleFileUpload} />
    </form>
  );
};

export default MovieForm;
