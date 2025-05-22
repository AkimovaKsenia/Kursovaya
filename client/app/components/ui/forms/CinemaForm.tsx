import {
  UseFormRegister,
  FieldErrors,
  UseFormHandleSubmit,
  Control,
} from "react-hook-form";
import { FC } from "react";
import cn from "classnames";
import styles from "../layout/MovieForm.module.scss";
import FileUploader from "../layout/FileUploader";
import {
  ICinemaDto,
  IListOfCategory,
  IListOfCondition,
} from "shared/interfaces/cinema.interface";
import { validEmail } from "../layout/header/login-form/login-auth.constants";

interface CinemaFormProps {
  register: UseFormRegister<ICinemaDto>;
  errors: FieldErrors<ICinemaDto>;
  handleSubmit: UseFormHandleSubmit<ICinemaDto>;
  onSubmit: (formData: ICinemaDto) => void;
  isPending: boolean;
  categoryData: IListOfCategory | undefined;
  conditionData: IListOfCondition | undefined;
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
          <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>
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
          <p className="text-red-500 text-xs mt-1">{errors.address.message}</p>
        )}
      </div>

      <div>
        <label className="block text-xs font-medium text-gray-700 mb-1">
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
        {errors.category && (
          <p className="text-red-500 text-xs mt-1">{errors.category.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Вместимость:
        </label>
        <select
          {...register("condition", {
            required: "Введите вместимость кинотеатра",
          })}
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
        {errors.condition && (
          <p className="text-red-500 text-xs mt-1">
            {errors.condition.message}
          </p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Email:
        </label>
        <input
          {...register("email", {
            required: "Обязательное поле",
            pattern: {
              value: validEmail,
              message: "Введите корректный email",
            },
          })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
        {errors.email && (
          <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Телефон:
        </label>
        <input
          {...register("phone", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
        {errors.phone && (
          <p className="text-red-500 text-xs mt-1">{errors.phone.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Описание:
        </label>
        <textarea
          {...register("description", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
          rows={5}
        />
        {errors.description && (
          <p className="text-red-500 text-xs mt-1">
            {errors.description.message}
          </p>
        )}
      </div>

      <FileUploader onFilesUploaded={handleFileUpload} />
      <div className="pt-4">
        <button
          type="submit"
          disabled={isPending}
          style={{ cursor: "pointer" }}
          className={`px-4 py-2 ml-65 mt-6 bg-blue-600 text-white rounded-md hover:bg-blue-700 ${
            isPending ? "opacity-50 cursor-not-allowed" : ""
          }`}
        >
          {isPending ? "Сохранение..." : "Сохранить изменения"}
        </button>
      </div>
    </form>
  );
};

export default CinemaForm;
