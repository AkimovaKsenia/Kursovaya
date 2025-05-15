import CinemaForm from "@/components/ui/forms/CinemaForm";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useCategory } from "hooks/useCategory";
import { useCinemaById } from "hooks/useCinemaById";
import { useCondition } from "hooks/useCondition";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { useForm } from "react-hook-form";
import { CinemaService } from "services/cinema.service";
import {
  ICinemaDto,
  ICinemaExportDto,
  IListOfCategory,
  IListOfCondition,
} from "shared/interfaces/cinema.interface";

const CinemaEdit: FC = () => {
  const router = useRouter();
  const { id } = router.query;
  const queryClient = useQueryClient();
  useEffect(() => {
    console.log("router.query.id =", id);
    console.log("router.isReady =", router.isReady);
  }, [id, router.isReady]);

  const cinemaId = router.isReady ? Number(id) : undefined;
  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    watch,
    setValue,
  } = useForm<ICinemaDto>({
    mode: "onChange",
  });

  const { data: categoryData, error: categoryError } = useCategory();
  const { data: conditionData, error: conditionError } = useCondition();

  const { data, isLoading } = useCinemaById(cinemaId);

  // Обработка успешного запроса через useEffect
  useEffect(() => {
    console.log("Сработал useEffect, data =", data);

    if (data) {
      console.log(data.photo);
      //   const castListString = Array.isArray(data.cast_list)
      //     ? data.cast_list.join(", ")
      //     : "";

      //   setValue("cast_list", castListString);
      setValue("name", data.name);
      setValue("address", data.address);
      setValue("description", data.description);
      setValue("photo", data.photo);
      setValue("category", data.category);
      setValue("condition", data.condition);
      setValue("email", data.email);
      // setValue("cast_list", data.cast_list);
      setValue("phone", data.phone);
    } else {
      console.log("Данные кинотеатра не получены");
    }
  }, [data, setValue]);

  const { mutate, isPending } = useMutation({
    mutationKey: ["update-cinema", cinemaId],
    mutationFn: (formData: ICinemaExportDto) =>
      CinemaService.updateCinema(Number(cinemaId), formData),
    onSuccess: () => {
      alert("Кионотеатр успешно обновлен!");
      router.push("/manage/cinema/listcinema");
    },
    onError: (error) => {
      console.error("Ошибка при обновлении кинотеатра:", error);
      alert("Произошла ошибка при обновлении кинотеатра");
    },
  });

  const mapDtoToExportDto = (
    dto: ICinemaDto,
    categories: IListOfCategory,
    conditions: IListOfCondition
  ): ICinemaExportDto => {
    const inputName =
      typeof dto.category === "string" ? dto.category.trim().toLowerCase() : "";
    const category = categories.find(
      (s) => s.name.trim().toLowerCase() === inputName
    );

    if (!category) {
      console.warn("Квтегория не найдена:", dto.category);
    }
    console.log("Сравнение категорий:");
    console.log(
      "Ищем:",
      `"${inputName}"`,
      "среди:",
      categories.map((s) => `"${s.name.trim().toLowerCase()}"`)
    );

    const inputNameSecond =
      typeof dto.condition === "string"
        ? dto.condition.trim().toLowerCase()
        : "";
    const condition = conditions.find(
      (s) => s.name.trim().toLowerCase() === inputNameSecond
    );

    if (!condition) {
      console.warn("Вместимость не найдена:", dto.condition);
    }
    console.log("Сравнение:");
    console.log(
      "Ищем:",
      `"${inputNameSecond}"`,
      "среди:",
      categories.map((s) => `"${s.name.trim().toLowerCase()}"`)
    );
    const exportDto: ICinemaExportDto = {
      name: dto.name,
      description: dto.description,
      address: dto.address,
      email: dto.email,
      phone: dto.phone,
      category_id: category?.id ?? 0,
      condition_id: condition?.id ?? 0,
    };
    if (dto.photo) {
      exportDto.photo = dto.photo;
    }
    return exportDto;
  };

  const onSubmit = (formData: ICinemaDto) => {
    if (!categoryData || !conditionData) {
      alert("Справочники не загружены.");

      // if (!(formData.photo instanceof File)) {
      //   console.error("Ошибка: photo не является файлом", formData.photo);
      //   alert("Загрузите изображение заново");
      //   return;
      // }
      return;
    }

    //   const inputName =
    //     Array.isArray(formData.film_studio_name) &&
    //     formData.film_studio_name.length > 0
    //       ? formData.film_studio_name[0].trim().toLowerCase()
    //       : "";

    //   const studio = filmStudioData.find(
    //     (s) => s.name.trim().toLowerCase() === inputName
    //   );

    const exportDto = mapDtoToExportDto(formData, categoryData, conditionData);
    console.log("🔄 Отправка запроса на обновление фильма с телом:", exportDto);

    mutate(exportDto);
  };

  const handleFileUpload = (files: File[]) => {
    if (files.length > 0) {
      setValue("photo", files[0], { shouldDirty: true });
    } else {
      setValue("photo", "", { shouldDirty: true }); // Очистка поля, если файл не выбран
    }
  };
  // const genreOptions = genresData || [];
  if (isLoading) return <div>Загрузка...</div>;

  return (
    <DashboardLayout>
      <div>Редактирование кинотеатра с ID: {cinemaId}</div>
      <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
        <CinemaForm
          register={register}
          errors={errors}
          handleSubmit={handleSubmit}
          onSubmit={onSubmit}
          isPending={isPending}
          categoryData={categoryData}
          conditionData={conditionData}
          handleFileUpload={handleFileUpload}
        />
      </div>
    </DashboardLayout>
  );
};
export default CinemaEdit;
