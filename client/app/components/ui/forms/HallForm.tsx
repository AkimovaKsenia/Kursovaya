import {
  UseFormRegister,
  FieldErrors,
  UseFormHandleSubmit,
  Control,
} from "react-hook-form";
import { FC } from "react";
import cn from "classnames";
import styles from "../layout/MovieForm.module.scss";
import { IHallDto, IListOfTypes } from "shared/interfaces/cinema.interface";

interface HallFormProps {
  register: UseFormRegister<IHallDto>;
  errors: FieldErrors<IHallDto>;
  handleSubmit: UseFormHandleSubmit<IHallDto>;
  onSubmit: (formData: IHallDto) => void;
  isPending: boolean;
  typesData: IListOfTypes | undefined;
}

const HallForm: FC<HallFormProps> = ({
  register,
  errors,
  handleSubmit,
  onSubmit,
  isPending,
  typesData,
}) => {
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Название Зала:
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
          Тип:
        </label>
        <select
          {...register("type", { required: "Выберите тип зала" })}
          className={cn(
            "form-select w-full text-m px-3 py-2 border border-gray-300 rounded-md",
            styles.select
          )}
          size={4}
        >
          {typesData?.map((type) => (
            <option key={type.id} value={type.name}>
              {type.name}
            </option>
          ))}
        </select>
        {errors.type && (
          <p className="text-red-500 text-xs mt-1">{errors.type.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Вместимость:
        </label>
        <input
          {...register("capacity", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
        {errors.capacity && (
          <p className="text-red-500 text-xs mt-1">{errors.capacity.message}</p>
        )}
      </div>

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

export default HallForm;
