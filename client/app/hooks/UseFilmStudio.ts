import { useQuery } from "@tanstack/react-query";
import { MovieService } from "services/movie.service";
import { IListOfFilmStudio } from "shared/interfaces/movie.interface";

export const useFilmStudio = () => {
  const { data: filmStudioData, error: filmStudioError } = useQuery<
    IListOfFilmStudio,
    Error
  >({
    queryKey: ["filmStudio"],
    queryFn: () => MovieService.getFilmStudios().then((res) => res.data),
  });
  return { data: filmStudioData, error: filmStudioError };
};
