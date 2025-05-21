import { useQuery } from "@tanstack/react-query";
import { CinemaService } from "services/cinema.service";

export const useHallById = (hallId?: number) => {
  const { data, isLoading } = useQuery({
    queryKey: ["hall", hallId],
    queryFn: () =>
      CinemaService.getHallById(Number(hallId)).then((res) => res.data),
    enabled: !!hallId && !isNaN(Number(hallId)),
  });
  return { data, isLoading };
};
