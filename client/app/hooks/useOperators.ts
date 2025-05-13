import { useQuery } from "@tanstack/react-query";
import { MovieService } from "services/movie.service";
import { IListOfOperators } from "shared/interfaces/movie.interface";

export const useOperators = () => {
  const { data: operatorsData, error: operatorsError } = useQuery<
    IListOfOperators,
    Error
  >({
    queryKey: ["operators"],
    queryFn: () => MovieService.getOperators().then((res) => res.data),
  });
  return { data: operatorsData, error: operatorsError };
};
