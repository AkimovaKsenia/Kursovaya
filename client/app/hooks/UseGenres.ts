import { useQuery } from "@tanstack/react-query";
import { MovieService } from "services/movie.service";
import { IListOfGenres } from "shared/interfaces/movie.interface";

export const useGenres = () => {
  const { data: genresData, error: genresError } = useQuery<
    IListOfGenres,
    Error
  >({
    queryKey: ["genres"],
    queryFn: () => MovieService.getGenres().then((res) => res.data),
  });
  return { data: genresData, error: genresError };
};
