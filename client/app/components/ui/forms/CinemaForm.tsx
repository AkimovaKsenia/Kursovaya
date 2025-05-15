import {
  UseFormRegister,
  FieldErrors,
  UseFormHandleSubmit,
  Control,
} from "react-hook-form";
import {
  IListOfDirector,
  IListOfFilmStudio,
  IListOfGenres,
  IListOfOperators,
  IMovieDto,
} from "shared/interfaces/movie.interface";
import { FC } from "react";
import cn from "classnames";
import styles from "../layout/MovieForm.module.scss";
import FileUploader from "../layout/FileUploader";
import {
  ICinemaDto,
  IListOfCategory,
  IListOfCondition,
} from "shared/interfaces/cinema.interface";

interface CinemaFormProps {
  register: UseFormRegister<ICinemaDto>;
  errors: FieldErrors<ICinemaDto>;
  handleSubmit: UseFormHandleSubmit<ICinemaDto>;
  onSubmit: (formData: ICinemaDto) => void;
  isPending: boolean;
  categoryData: IListOfCategory | undefined;
  conditionData: IListOfCondition | undefined;
  //   filmStudioData: IListOfFilmStudio | undefined;
  //   directorsData: IListOfDirector | undefined;
  handleFileUpload: (files: File[]) => void;
}

const CinemaForm: FC<CinemaFormProps> = ({
  register,
  errors,
  handleSubmit,
  onSubmit,
  isPending,
  categoryData,
  conditionData,
  //   filmStudioData,
  //   directorsData,
  handleFileUpload,
}) => {
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Название кинотеатра:
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
          Адрес:
        </label>
        <input
          {...register("address", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
        {errors.address && (
          <p className="text-red-500 text-sm mt-1">{errors.address.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Категория:
        </label>
        <select
          {...register("category", { required: "Выберите категорию" })}
          className={cn(
            "form-select w-full text-m px-3 py-2 border border-gray-300 rounded-md",
            styles.select
          )}
          size={4}
        >
          {categoryData?.map((category) => (
            <option key={category.id} value={category.name}>
              {category.name}
            </option>
          ))}
        </select>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Вместимость:
        </label>
        <select
          {...register("condition", { required: "Выберите категорию" })}
          className={cn(
            "form-select w-full text-m px-3 py-2 border border-gray-300 rounded-md",
            styles.select
          )}
          size={4}
        >
          {conditionData?.map((condition) => (
            <option key={condition.id} value={condition.name}>
              {condition.name}
            </option>
          ))}
        </select>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Email:
        </label>
        <input
          {...register("email", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Телефон:
        </label>
        <input
          {...register("phone", { required: "Обязательное поле" })}
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

export default CinemaForm;
