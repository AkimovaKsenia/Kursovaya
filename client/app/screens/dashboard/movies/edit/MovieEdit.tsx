import DashboardLayout from "@/components/ui/layout/DashboardLayout";
import FileUploader from "@/components/ui/layout/FileUploader";
import MovieForm from "@/components/ui/layout/MovieForm";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useGenres } from "hooks/UseGenres";
import { useMovieById } from "hooks/useMovieById";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { useForm } from "react-hook-form";
import { MovieService } from "services/movie.service";
import {
  IListOfGenres,
  IMovie,
  IMovieDto,
} from "shared/interfaces/movie.interface";

const MovieEdit: FC = () => {
  const router = useRouter();
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
  const { data, isLoading } = useMovieById(movieId);

  useEffect(() => {
    if (genresData) {
      console.log("Жанры успешно загружены:", genresData);
    }
  }, [genresData]);

  // Обработка успешного запроса через useEffect
  useEffect(() => {
    console.log("Сработал useEffect, data =", data);
    if (data) {
      setValue("name", data.name);
      setValue("genres", data.genres);
      setValue("description", data.description);
      setValue("photo", data.photo);
      setValue("directors", data.directors);
      setValue("operators", data.operators);
      setValue("film_studio_name", data.film_studio_name);
      setValue("cast_list", data.cast_list);
    } else {
      console.log("Данные фильма не получены");
    }
  }, [data, setValue]);

  const { mutate, isPending } = useMutation({
    mutationKey: ["update-movie", movieId],
    mutationFn: (formData: IMovieDto) =>
      MovieService.updateMovie(Number(movieId), formData),
    onSuccess: (updatedMovie) => {
      queryClient.setQueryData(["movie", movieId], updatedMovie);
      alert("Фильм успешно обновлен!");
    },
    onError: (error) => {
      console.error("Ошибка при обновлении фильма:", error);
      alert("Произошла ошибка при обновлении фильма");
    },
  });
  const onSubmit = (formData: IMovieDto) => {
    mutate(formData);
  };
  const handleFileUpload = (files: File[]) => {
    if (files.length > 0) {
      setValue("photo", files[0], { shouldDirty: true });

      const reader = new FileReader();
    }
  };
  // const genreOptions = genresData || [];
  if (isLoading) return <div>Загрузка...</div>;

  return (
    <DashboardLayout>
      <div>Редактирование фильма с ID: {movieId}</div>
      <div className=" flex flex-col items-center justify-start max-w-xl mx-auto p-6 bg-none rounded-lg shadow mt-6 ">
        <MovieForm
          register={register}
          errors={errors}
          handleSubmit={handleSubmit}
          onSubmit={onSubmit}
          isPending={isPending}
          genresData={genresData}
          handleFileUpload={handleFileUpload}
        />
      </div>
    </DashboardLayout>
  );
};
export default MovieEdit;
