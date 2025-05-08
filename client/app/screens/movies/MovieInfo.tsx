import Layout from "@/components/ui/layout/Layout";
import { useQuery } from "@tanstack/react-query";
import Image from "next/image";
import { useRouter } from "next/router";
import { FC } from "react";
import { MovieService } from "services/movie.service";

const MoviesInfo: FC = () => {
  const { query } = useRouter();
  const movieId = Number(query?.id);
  const { refetch, data: movie } = useQuery({
    queryKey: ["get movie", movieId],
    queryFn: () => MovieService.getMovieById(movieId),
    select: ({ data }) => data,
  });
  return (
    <Layout title={`${movie?.name}`}>
      <div style={{ background: "blue" }}>MOVIES</div>
      <div>
        <h1>{movie?.name}</h1>
        {movie?.poster && movie?.name && (
          <Image
            src={movie.poster}
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
              <span>{movie?.genre}</span>
              <span>Режиссер</span>
              <span>{movie?.director}</span>
            </li>
          </ul>
        </div>
      </div>
    </Layout>
  );
};

export default MoviesInfo;
