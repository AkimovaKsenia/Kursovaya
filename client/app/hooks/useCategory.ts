import { useQuery } from "@tanstack/react-query";
import { CinemaService } from "services/cinema.service";
import { MovieService } from "services/movie.service";
import { IListOfCategory } from "shared/interfaces/cinema.interface";
import { IListOfGenres } from "shared/interfaces/movie.interface";

export const useCategory = () => {
  const { data: categoryData, error: categoryError } = useQuery<
    IListOfCategory,
    Error
  >({
    queryKey: ["category"],
    queryFn: () => CinemaService.getCategory().then((res) => res.data),
  });
  return { data: categoryData, error: categoryError };
};
