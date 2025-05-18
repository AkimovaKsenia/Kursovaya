import Layout from "@/components/ui/layout/Layout";
import { FC, useEffect, useState } from "react";
import styles from "./dashboard/cinema/AllCinema.module.scss";
import { CinemaService } from "services/cinema.service";
import Link from "next/link";
import { useRouter } from "next/router";
import { ICinemaMain } from "shared/interfaces/cinema.interface";
import Image from "next/image";

const Choice: FC = () => {
  const router = useRouter();
  const [cinemas, setCinemas] = useState<ICinemaMain[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchCinemas = async () => {
      try {
        const { data } = await CinemaService.getAllCinema();
        setCinemas(data);
      } catch (error) {
        console.error("Ошибка при загрузке кинотеатров", error);
        setCinemas([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchCinemas();
  }, []);

  if (isLoading) return <div>Загрузка...</div>;

  return (
    <div className=" bg-black">
      <div
        style={{ background: "rgba(129, 125, 219, 0.2)" }}
        className={styles.catalog}
      >
        Кинотеатры
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
        {cinemas.length ? (
          cinemas.map((cinema) => (
            <Link href={`/manage/cinema/${cinema.id}`} key={cinema.id} passHref>
              <div className="relative bg-gray-400 opacity-100 rounded-lg overflow-hidden shadow-lg cursor-pointer transition-all duration-300 hover:shadow-xl hover:transform hover:-translate-y-1 group h-full min-h-[300px]">
                {/* Основной контент карточки */}
                <div className="p-6 z-10 relative h-full flex flex-col">
                  <h3 className="text-2xl font-bold mb-2 text-white">
                    {cinema.name}
                  </h3>
                  <p className="text-gray-900 mb-4">{cinema.address}</p>
                </div>

                {/* Фоновое изображение, появляющееся при наведении */}
                <div className="absolute inset-0 overflow-hidden opacity-0 group-hover:opacity-20 transition-opacity duration-300">
                  <Image
                    src="/choicePhoto.jpg"
                    alt={cinema.name}
                    layout="fill"
                    objectFit="cover"
                    className=""
                  />
                </div>
              </div>
            </Link>
          ))
        ) : (
          <div className="text-white col-span-full text-center py-10">
            Кинотеатры не найдены
          </div>
        )}
      </div>
    </div>
  );
};

export default Choice;
