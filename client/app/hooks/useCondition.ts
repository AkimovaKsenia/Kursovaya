import { useQuery } from "@tanstack/react-query";
import { CinemaService } from "services/cinema.service";
import { IListOfCondition } from "shared/interfaces/cinema.interface";

export const useCondition = () => {
  const { data: conditionData, error: conditionError } = useQuery<
    IListOfCondition,
    Error
  >({
    queryKey: ["condition"],
    queryFn: () => CinemaService.getCondition().then((res) => res.data),
  });
  return { data: conditionData, error: conditionError };
};
