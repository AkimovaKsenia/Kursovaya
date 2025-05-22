import ErrorAuth from "@/components/ui/ErrorAuth";
import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import FileUploader from "@/components/ui/layout/FileUploader";
import MovieForm from "@/components/ui/layout/MovieForm";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "hooks/useAuth";
import { useDirectors } from "hooks/useDirectors";
import { useFilmStudio } from "hooks/UseFilmStudio";
import { useGenres } from "hooks/UseGenres";
import { useMovieById } from "hooks/useMovieById";
import { useOperators } from "hooks/useOperators";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { useForm } from "react-hook-form";
import { MovieService } from "services/movie.service";
import {
  IListOfDirector,
  IListOfFilmStudio,
  IListOfGenres,
  IListOfOperators,
  IMovie,
  IMovieDto,
  IMovieExportDto,
} from "shared/interfaces/movie.interface";

const MovieEdit: FC = () => {
  const router = useRouter();
  const { user, setUser } = useAuth();
  const { id } = router.query;
  const queryClient = useQueryClient();
  useEffect(() => {
    console.log("router.query.id =", id);
    console.log("router.isReady =", router.isReady);
  }, [id, router.isReady]);

  const movieId = router.isReady ? Number(id) : undefined;
  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    watch,
    setValue,
  } = useForm<IMovieDto>({
    mode: "onChange",
  });

  const { data: genresData, error: genresError } = useGenres();
  const { data: operatorsData, error: operatorsError } = useOperators();
  const { data: filmStudioData, error: filmStudioError } = useFilmStudio();
  const { data: directorsData, error: directorsError } = useDirectors();

  const { data, isLoading } = useMovieById(movieId);

  // Обработка успешного запроса через useEffect
  useEffect(() => {
    console.log("Сработал useEffect, data =", data);

    if (data) {
      console.log(data.photo);
      const castListString = Array.isArray(data.cast_list)
        ? data.cast_list.join(", ")
        : "";

      setValue("cast_list", castListString);
      setValue("name", data.name);
      setValue("genres", data.genres);
      setValue("description", data.description);
      // setValue("photo", data.photo);
      setValue("directors", data.directors);
      setValue("operators", data.operators);
      setValue("film_studio_name", data.film_studio_name);
      // setValue("cast_list", data.cast_list);
      setValue("duration_in_min", data.duration_in_min);
    } else {
      console.log("Данные фильма не получены");
    }
  }, [data, setValue]);

  const { mutate, isPending } = useMutation({
    mutationKey: ["update-movie", movieId],
    mutationFn: (formData: IMovieExportDto) =>
      MovieService.updateMovie(Number(movieId), formData),
    onSuccess: (updatedMovie) => {
      queryClient.setQueryData(["movie", movieId], updatedMovie);
      alert("Фильм успешно обновлен!");
      router.push("/manage/movies/listmovies");
    },
    onError: (error) => {
      console.error("Ошибка при обновлении фильма:", error);
      alert("Произошла ошибка при обновлении фильма");
    },
  });

  const mapDtoToExportDto = (
    dto: IMovieDto,
    genres: IListOfGenres,
    operators: IListOfOperators,
    filmStudios: IListOfFilmStudio,
    directors: IListOfDirector
  ): IMovieExportDto => {
    const castListArray = dto.cast_list
      .split(",")
      .map((actor) => actor.trim())
      .filter((actor) => actor.length > 0);
    const inputName =
      typeof dto.film_studio_name === "string"
        ? dto.film_studio_name.trim().toLowerCase()
        : "";

    const studio = filmStudios.find(
      (s) => s.name.trim().toLowerCase() === inputName
    );

    if (!studio) {
      console.warn("Киностудия не найдена:", dto.film_studio_name);
    }
    console.log("Сравнение киностудий:");
    console.log(
      "Ищем:",
      `"${inputName}"`,
      "среди:",
      filmStudios.map((s) => `"${s.name.trim().toLowerCase()}"`)
    );
    const exportDto: IMovieExportDto = {
      name: dto.name,
      description: dto.description,
      cast_list: castListArray,
      film_studio_id: studio?.id ?? 0,
      duration_in_min: dto.duration_in_min,
      genre_ids: genres
        .filter((g) => dto.genres.includes(g.name))
        .map((g) => g.id),
      operator_ids: operators
        .filter((op) => dto.operators.includes(op.fio))
        .map((op) => op.id),
      director_ids: directors
        .filter((d) => dto.directors.includes(d.fio))
        .map((d) => d.id),
    };
    if (dto.photo) {
      exportDto.film_photo = dto.photo;
    }
    return exportDto;
  };

  const onSubmit = (formData: IMovieDto) => {
    console.log("film_studio_name из формы:", formData.film_studio_name);

    if (!genresData || !operatorsData || !filmStudioData || !directorsData) {
      alert("Справочники не загружены.");

      if (!(formData.photo instanceof File)) {
        console.error("Ошибка: photo не является файлом", formData.photo);
        alert("Загрузите изображение заново");
        return;
      }
      return;
    }

    const inputName =
      Array.isArray(formData.film_studio_name) &&
      formData.film_studio_name.length > 0
        ? formData.film_studio_name[0].trim().toLowerCase()
        : "";

    const studio = filmStudioData.find(
      (s) => s.name.trim().toLowerCase() === inputName
    );

    const exportDto = mapDtoToExportDto(
      formData,
      genresData,
      operatorsData,
      filmStudioData,
      directorsData
    );
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
      {user ? (
        <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
          <MovieForm
            register={register}
            errors={errors}
            handleSubmit={handleSubmit}
            onSubmit={onSubmit}
            isPending={isPending}
            genresData={genresData}
            operatorsData={operatorsData}
            filmStudioData={filmStudioData}
            directorsData={directorsData}
            handleFileUpload={handleFileUpload}
          />
        </div>
      ) : (
        <ErrorAuth />
      )}
    </DashboardLayout>
  );
};
export default MovieEdit;
