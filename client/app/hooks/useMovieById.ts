import { useQuery } from "@tanstack/react-query";
import { NextRouter, useRouter } from "next/router";
import { useEffect } from "react";
import { MovieService } from "services/movie.service";

export const useMovieById = (movieId?: number) => {
  const { data, isLoading } = useQuery({
    queryKey: ["movie", movieId],
    queryFn: () =>
      MovieService.getMovieById(Number(movieId)).then((res) => res.data),
    enabled: !!movieId && !isNaN(Number(movieId)),
  });
  return { data, isLoading };
};
