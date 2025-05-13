import { useQuery } from "@tanstack/react-query";
import { MovieService } from "services/movie.service";
import { IListOfDirector } from "shared/interfaces/movie.interface";

export const useDirectors = () => {
  const { data: directorsData, error: directorsError } = useQuery<
    IListOfDirector,
    Error
  >({
    queryKey: ["directors"],
    queryFn: () => MovieService.getDirectors().then((res) => res.data),
  });
  return { data: directorsData, error: directorsError };
};
