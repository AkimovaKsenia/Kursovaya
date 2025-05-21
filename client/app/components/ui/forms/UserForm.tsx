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
import {} from "shared/interfaces/cinema.interface";
import { IListOfRoles, IUserDto } from "shared/interfaces/user.interface";

interface UserFormProps {
  register: UseFormRegister<IUserDto>;
  errors: FieldErrors<IUserDto>;
  handleSubmit: UseFormHandleSubmit<IUserDto>;
  onSubmit: (formData: IUserDto) => void;
  isPending: boolean;
  rolesData: IListOfRoles | undefined;
}

const UserForm: FC<UserFormProps> = ({
  register,
  errors,
  handleSubmit,
  onSubmit,
  isPending,
  rolesData,
}) => {
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Имя:
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
          Фамилия:
        </label>
        <input
          {...register("surname", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
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
          Пароль:
        </label>
        <input
          {...register("password", { required: "Обязательное поле" })}
          className="w-115 px-3 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Роль:
        </label>
        <select
          {...register("role", { required: "Выберите категорию" })}
          className={cn(
            "form-select w-full text-m px-3 py-2 border border-gray-300 rounded-md",
            styles.select
          )}
          size={4}
        >
          {rolesData?.map((role) => (
            <option key={role.id} value={role.name}>
              {role.name}
            </option>
          ))}
        </select>
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

export default UserForm;
