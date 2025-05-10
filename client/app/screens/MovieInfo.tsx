import Layout from "@/components/ui/layout/Layout";
import { useMutation, useQuery } from "@tanstack/react-query";
import Image from "next/image";
import { useRouter } from "next/router";
import { FC, useEffect } from "react";
import { MovieService } from "services/movie.service";
import { ViewsService } from "services/views.service";

const MoviesInfo: FC = () => {
  const { query } = useRouter();
  const movieId = Number(query?.id);
  const { refetch, data: movie } = useQuery({
    queryKey: ["get movie", movieId],
    queryFn: () => MovieService.getMovieById(movieId),
    select: ({ data }) => data,
  });

  //АВТООБНОВЛЕНИЕ ПРОСМОТРОВ СТРАНИЦЫ
  const { mutateAsync } = useMutation({
    mutationKey: ["update count opened"],
    mutationFn: () => ViewsService.updateViews(movieId.toString()),
  });

  useEffect(() => {
    if (movieId) mutateAsync();
  }, [movieId]);

  return (
    <Layout title={`${movie?.name}`}>
      <div style={{ background: "blue" }}>MOVIES</div>
      <div>
        <h1>{movie?.name}</h1>
        {movie?.photo && movie?.name && (
          <Image
            src={movie.photo}
            alt={movie.name}
            width={220}
            height={330}
            layout="responsive"
          />
        )}
        <div>
          <ul>
            <li>
              <span>Жанр</span>
              <span>{movie?.genres}</span>
              <span>Режиссер</span>
              <span>{movie?.directors}</span>
            </li>
          </ul>
        </div>
      </div>
    </Layout>
  );
};

export default MoviesInfo;
