import CinemaForm from "@/components/ui/forms/CinemaForm";
import UserForm from "@/components/ui/forms/UserForm";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import { useRole } from "hooks/useRole";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { useForm } from "react-hook-form";
import { UserService } from "services/user.service";
import {
  IListOfRoles,
  IUserDto,
  IUserExportDto,
} from "shared/interfaces/user.interface";

const CreateUser: FC = () => {
  const router = useRouter();

  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    watch,
    setValue,
  } = useForm<IUserDto>({
    mode: "onChange",
  });

  const { data: rolesData, error: rolesError } = useRole();

  const { mutate, isPending } = useMutation({
    mutationKey: ["create-user"],
    mutationFn: (formData: IUserExportDto) => UserService.createUser(formData),
    onSuccess: (createUser) => {
      alert("Пользователь успешно создан!");
      router.back();
    },
    onError: (error) => {
      console.error("Ошибка при создании пользователя:", error);
      alert("Произошла ошибка при создании пользователя");
    },
  });

  const mapDtoToExportDto = (
    dto: IUserDto,
    roles: IListOfRoles
  ): IUserExportDto => {
    const inputNameSecond =
      typeof dto.role === "string" ? dto.role.trim().toLowerCase() : "";
    const role = roles.find(
      (s) => s.name.trim().toLowerCase() === inputNameSecond
    );

    if (!role) {
      console.warn("Вместимость не найдена:", dto.role);
    }
    console.log("Сравнение:");
    console.log(
      "Ищем:",
      `"${inputNameSecond}"`,
      "среди:",
      roles.map((s) => `"${s.name.trim().toLowerCase()}"`)
    );
    const exportDto: IUserExportDto = {
      name: dto.name,
      email: dto.email,
      surname: dto.surname,
      password: dto.password,
      role_id: role?.id ?? 0,
    };
    return exportDto;
  };

  const onSubmit = (formData: IUserDto) => {
    if (!rolesData) {
      alert("Справочники не загружены.");
      return;
    }

    const exportDto = mapDtoToExportDto(formData, rolesData);
    console.log("🔄 Отправка запроса на обновление фильма с телом:", exportDto);

    mutate(exportDto);
  };

  return (
    <DashboardLayout>
      <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
        <UserForm
          register={register}
          errors={errors}
          handleSubmit={handleSubmit}
          onSubmit={onSubmit}
          isPending={isPending}
          rolesData={rolesData}
        />
      </div>
    </DashboardLayout>
  );
};
export default CreateUser;
