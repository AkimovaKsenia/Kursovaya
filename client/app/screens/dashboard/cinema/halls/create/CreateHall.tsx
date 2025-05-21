import CinemaForm from "@/components/ui/forms/CinemaForm";
import HallForm from "@/components/ui/forms/HallForm";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useTypes } from "hooks/useTypes";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { useForm } from "react-hook-form";
import { CinemaService } from "services/cinema.service";
import {
  IHallDto,
  IHallExportDto,
  IListOfTypes,
} from "shared/interfaces/cinema.interface";

const CreateHall: FC = () => {
  const router = useRouter();

  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    watch,
    setValue,
  } = useForm<IHallDto>({
    mode: "onChange",
  });

  const { data: typesData, error: typesError } = useTypes();

  const { mutate, isPending } = useMutation({
    mutationKey: ["create-hall"],
    mutationFn: (formData: IHallExportDto) =>
      CinemaService.createHall(formData),
    onSuccess: (createHall) => {
      alert("Зал успешно создан!");
      router.push("/manage/cinema/listcinema");
    },
    onError: (error) => {
      console.error("Ошибка при создании зала:", error);
      alert("Произошла ошибка при создании зала");
    },
  });

  const mapDtoToExportDto = (
    dto: IHallDto,
    types: IListOfTypes
  ): IHallExportDto => {
    const inputNameSecond =
      typeof dto.type === "string" ? dto.type.trim().toLowerCase() : "";
    const type = types.find(
      (s) => s.name.trim().toLowerCase() === inputNameSecond
    );

    if (!type) {
      console.warn("тип не найден:", dto.type);
    }
    console.log("Сравнение:");
    console.log(
      "Ищем:",
      `"${inputNameSecond}"`,
      "среди:",
      types.map((s) => `"${s.name.trim().toLowerCase()}"`)
    );
    const exportDto: IHallExportDto = {
      name: dto.name,
      capacity: dto.capacity,
      type_id: type?.id ?? 0,
    };

    return exportDto;
  };

  const onSubmit = (formData: IHallDto) => {
    if (!typesData) {
      alert("Справочники не загружены.");
      return;
    }

    const exportDto = mapDtoToExportDto(formData, typesData);
    console.log("🔄 Отправка запроса на обновление фильма с телом:", exportDto);

    mutate(exportDto);
  };

  return (
    <DashboardLayout>
      <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
        <HallForm
          register={register}
          errors={errors}
          handleSubmit={handleSubmit}
          onSubmit={onSubmit}
          isPending={isPending}
          typesData={typesData}
        />
      </div>
    </DashboardLayout>
  );
};
export default CreateHall;
