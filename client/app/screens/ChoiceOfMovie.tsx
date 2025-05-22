import Layout from "@/components/ui/layout/Layout";
import { FC, useEffect, useState } from "react";
import styles from "./dashboard/cinema/AllCinema.module.scss";
import { CinemaService } from "services/cinema.service";
import Link from "next/link";
import { useRouter } from "next/router";
import { ICinemaMain } from "shared/interfaces/cinema.interface";
import Image from "next/image";
import { MovieService } from "services/movie.service";
import { IMovie } from "shared/interfaces/movie.interface";

const Choice: FC = () => {
  const router = useRouter();
  const [movies, setMovies] = useState<IMovie[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchCinemas = async () => {
      try {
        const { data } = await MovieService.getAll();
        setMovies(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setMovies([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchCinemas();
  }, []);

  if (isLoading) return <div>Загрузка...</div>;

  return (
    <div className="bg-black">
      <div
        style={{
          borderBottom: "1px solid rgba(255, 255, 255, 0.3)",
          background: "rgba(255, 255, 255, 0.1)",
          minHeight: "50px",
        }}
        className={styles.catalog}
      >
        Фильмы
      </div>

      <div className="mx-auto max-w-[1400px] px-10 py-8 gap-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6">
          {movies.length ? (
            movies.map((movie) => (
              <Link
                href={`/movie/${movie.id}`}
                key={movie.id}
                passHref
                className="group"
              >
                <div className="relative w-[240px] aspect-[2/3] rounded-xl overflow-hidden shadow-lg transition-all duration-300 hover:shadow-2xl hover:scale-105">
                  <div className="absolute inset-0">
                    <Image
                      src={movie.photo || "/default-movie.jpg"}
                      alt={movie.name}
                      layout="fill"
                      objectFit="cover"
                      className="brightness-75"
                    />
                  </div>

                  <div className="absolute inset-0 bg-gradient-to-t from-black/90 via-black/50 to-transparent" />

                  <div className="relative z-10 h-full flex flex-col justify-end p-3">
                    <h3 className="text-md font-bold text-white mb-1 line-clamp-2">
                      {movie.name}
                    </h3>
                    <div className="text-sm text-gray-300 line-clamp-1">
                      {movie.genres?.join(", ")}
                    </div>
                  </div>

                  <div className="absolute inset-0 bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity scale-105 duration-300 flex items-center justify-center">
                    <button className="px-3 py-1.5 bg-purple-600 rounded-full text-white text-sm font-medium transform translate-y-4 group-hover:translate-y-0 transition-transform duration-300">
                      Подробнее
                    </button>
                  </div>
                </div>
              </Link>
            ))
          ) : (
            <div className="col-span-full text-center text-white py-10">
              Фильмы не найдены
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Choice;
