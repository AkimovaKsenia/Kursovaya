import { useQuery } from "@tanstack/react-query";
import { CinemaService } from "services/cinema.service";
import {
  IListOfCondition,
  IListOfTypes,
} from "shared/interfaces/cinema.interface";

export const useTypes = () => {
  const { data: typesData, error: typesError } = useQuery<IListOfTypes, Error>({
    queryKey: ["types"],
    queryFn: () => CinemaService.getTypes().then((res) => res.data),
  });
  return { data: typesData, error: typesError };
};
