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
      {/* Заголовок (оставлен как в оригинале) */}
      <div
        style={{ background: "rgba(129, 125, 219, 0.2)" }}
        className={styles.catalog}
      >
        Фильмы
      </div>

      {/* Сетка фильмов */}
      <div className=" mt-15 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-8 px-6 pt-6">
        {movies.length ? (
          movies.map((movie) => (
            <Link
              href={`/movie/${movie.id}`}
              key={movie.id}
              passHref
              className="group"
            >
              <div className="relative h-full min-h-[400px] rounded-xl overflow-hidden shadow-lg transition-all duration-300 hover:shadow-2xl hover:scale-105">
                {/* Фоновое изображение фильма */}
                <div className="absolute inset-0">
                  <Image
                    src={movie.photo || "/default-movie.jpg"} // Заглушка, если нет фото
                    alt={movie.name}
                    layout="fill"
                    objectFit="cover"
                    className="brightness-75"
                  />
                </div>

                {/* Градиент для текста */}
                <div className="absolute inset-0 bg-gradient-to-t from-black/90 via-black/50 to-transparent" />

                {/* Информация о фильме */}
                <div className="relative z-10 h-full flex flex-col justify-end p-4">
                  <h3 className="text-xl font-bold text-white mb-1 line-clamp-2">
                    {movie.name}
                  </h3>
                  <div className="flex flex-wrap gap-2 mb-2">
                    {movie.genres && movie.genres.length > 0 && (
                      <span> {movie.genres.join(", ")}</span>
                    )}
                  </div>
                  <div className="flex items-center justify-between text-sm text-gray-300"></div>
                </div>

                {/* Эффект при наведении */}
                <div className="absolute inset-0 bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center">
                  <button className="px-4 py-2 bg-purple-600 rounded-full text-white font-medium transform translate-y-4 group-hover:translate-y-0 transition-transform duration-300">
                    Подробнее
                  </button>
                </div>
              </div>
            </Link>
          ))
        ) : (
          <div className="col-span-full text-center py-16">
            <div className="text-white text-xl">Фильмы не найдены</div>
            <div className="text-gray-400 mt-2">
              Попробуйте обновить страницу
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Choice;
