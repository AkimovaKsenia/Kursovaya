import { useQuery } from "@tanstack/react-query";
import { NextRouter, useRouter } from "next/router";
import { useEffect } from "react";
import { CinemaService } from "services/cinema.service";

export const useCinemaById = (cinemaId?: number) => {
  const { data, isLoading } = useQuery({
    queryKey: ["cinema", cinemaId],
    queryFn: () =>
      CinemaService.getCinemaById(Number(cinemaId)).then((res) => res.data),
    enabled: !!cinemaId && !isNaN(Number(cinemaId)),
  });
  return { data, isLoading };
};
