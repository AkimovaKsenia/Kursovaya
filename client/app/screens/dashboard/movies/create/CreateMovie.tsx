// import DashboardLayout from "@/components/ui/layout/DashboardLayout";
// import MovieForm from "@/components/ui/layout/MovieForm";
// import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
// import { useDirectors } from "hooks/useDirectors";
// import { useFilmStudio } from "hooks/UseFilmStudio";
// import { useGenres } from "hooks/UseGenres";
// import { useMovieById } from "hooks/useMovieById";
// import { useOperators } from "hooks/useOperators";
// import { useRouter } from "next/router";
// import { FC, useEffect } from "react";
// import { useForm } from "react-hook-form";
// import { MovieService } from "services/movie.service";
// import {
//   IListOfDirector,
//   IListOfFilmStudio,
//   IListOfGenres,
//   IListOfOperators,
//   IMovie,
//   IMovieDto,
//   IMovieExportDto,
// } from "shared/interfaces/movie.interface";

// const CreateMovie: FC = () => {
//   const router = useRouter();
//   const queryClient = useQueryClient();

//   const {
//     register,
//     formState: { errors },
//     control,
//     handleSubmit,
//     setValue,
//     reset,
//   } = useForm<IMovieDto>({
//     mode: "onChange",
//     defaultValues: {
//       genres: [],
//       operators: [],
//       directors: [],
//       cast_list: [],
//     },
//   });

//   const { data: genresData, error: genresError } = useGenres();
//   const { data: operatorsData, error: operatorsError } = useOperators();
//   const { data: filmStudioData, error: filmStudioError } = useFilmStudio();
//   const { data: directorsData, error: directorsError } = useDirectors();

//   const { mutate, isPending } = useMutation({
//     mutationKey: ["create-movie"],
//     mutationFn: (formData: IMovieExportDto) =>
//       MovieService.createMovie(formData),
//     onSuccess: () => {
//       queryClient.invalidateQueries({ queryKey: ["movies"] });
//       alert("Фильм успешно создан!");
//       reset();
//       router.push("/movies");
//     },
//     onError: (error: any) => {
//       console.error(
//         "Ошибка создания:",
//         error.response?.data?.message || error.message
//       );
//       alert(error.response?.data?.message || "Ошибка при создании фильма");
//     },
//   });
//   const handleFileUpload = (files: File[]) => {
//     setValue("photo", files[0] || null, { shouldDirty: true });
//   };

//   const onSubmit = (formData: IMovieDto) => {
//     if (!genresData || !operatorsData || !filmStudioData || !directorsData) {
//       alert("Данные справочников не загружены");
//       return;
//     }

//     // ✅ Безопасное преобразование cast_list в массив строк
//     const normalizedCastList: string[] = Array.isArray(formData.cast_list)
//       ? formData.cast_list
//       : typeof formData.cast_list === "string"
//       ? formData.cast_list
//       : [];

//     const exportDto = mapDtoToExportDto(
//       { ...formData, cast_list: normalizedCastList },
//       genresData,
//       operatorsData,
//       filmStudioData,
//       directorsData
//     );

//     console.log("🔄 Отправка запроса на создание фильма с телом:", exportDto);
//     mutate(exportDto);
//   };

//   const mapDtoToExportDto = (
//     dto: IMovieDto,
//     genres: IListOfGenres,
//     operators: IListOfOperators,
//     filmStudios: IListOfFilmStudio,
//     directors: IListOfDirector
//   ): IMovieExportDto => {
//     const exportDto: IMovieExportDto = {
//       name: dto.name,
//       description: dto.description,
//       cast_list: dto.cast_list,
//       film_studio_id:
//         filmStudios.find((studio) => studio.name === dto.film_studio_name)
//           ?.id ?? 0,
//       duration_in_min: dto.duration_in_min,
//       genre_ids: genres
//         .filter((g) => dto.genres.includes(g.name))
//         .map((g) => g.id),
//       operator_ids: operators
//         .filter((op) => dto.operators.includes(op.fio))
//         .map((op) => op.id),
//       director_ids: directors
//         .filter((d) => dto.directors.includes(d.fio))
//         .map((d) => d.id),
//     };
//     if (dto.photo) {
//       exportDto.film_photo = dto.photo;
//     }
//     return exportDto;
//   };

//   return (
//     <DashboardLayout>
//       <div>Создание фильма</div>
//       <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
//         <MovieForm
//           register={register}
//           errors={errors}
//           handleSubmit={handleSubmit}
//           onSubmit={onSubmit}
//           isPending={isPending}
//           genresData={genresData}
//           operatorsData={operatorsData}
//           filmStudioData={filmStudioData}
//           directorsData={directorsData}
//           handleFileUpload={handleFileUpload}
//         />
//       </div>
//     </DashboardLayout>
//   );
// };
// export default CreateMovie;
