import Layout from "@/components/ui/layout/Layout";
import { FC, useEffect, useState } from "react";
import styles from "./ChoiceOfCinema.module.scss";
import { CinemaService } from "services/cinema.service";
import Link from "next/link";
import { useRouter } from "next/router";
import { ICinemaMain, ListOfCinema } from "shared/interfaces/cinema.interface";
import Image from "next/image";
import { ListSeparator } from "sass";
import cn from "classnames";

const Choice: FC = () => {
  const router = useRouter();
  const [cinemas, setCinemas] = useState<ListOfCinema>([]);
  const [isLoading, setIsLoading] = useState(true);

  // useEffect(() => {
  //   const fetchCinemas = async () => {
  //     try {
  //       const { data } = await CinemaService.getAllCinema();
  //       setCinemas(data);
  //     } catch (error) {
  //       console.error("Ошибка при загрузке кинотеатров", error);
  //       setCinemas([]);
  //     } finally {
  //       setIsLoading(false);
  //     }
  //   };

  //   fetchCinemas();
  // }, []);

  useEffect(() => {
    const fetchCinemas = async () => {
      try {
        const { data } = await CinemaService.getListOfCinema();
        setCinemas(data);
      } catch (error) {
        console.error("Ошибка при загрузке фильмов", error);
        setCinemas([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchCinemas();
  }, []);

  if (isLoading) return <div>Загрузка...</div>;

  if (isLoading) return <div>Загрузка...</div>;

  return (
    <div
      className=" bg-black"
      style={
        {
          // backgroundImage:
          //   "url('https://i.pinimg.com/736x/33/53/97/335397688d991327387f2b7b0640289d.jpg')",
          // backgroundSize: "cover",
          // backgroundPosition: "center",
          // backgroundRepeat: "no-repeat",
        }
      }
    >
      <div
        style={{
          borderBottom: "1px solid rgba(255, 255, 255, 0.3)",
          background: "rgba(255, 255, 255, 0.1)",
          minHeight: "50px",
        }}
        className={cn(styles.catalog)}
      >
        Кинотеатры
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
        {cinemas.length ? (
          cinemas.map((cinema) => (
            <Link href={`/manage/cinema/${cinema.id}`} key={cinema.id} passHref>
              <div className="group relative z-10 h-full flex flex-col justify-between p-6 bg-gray-600 rounded-lg bg-opacity-70 text-white min-h-63 ">
                {/* Основной контент карточки */}
                <div>
                  <h3 className={cn("text-2xl font-bold mb-2", styles.heading)}>
                    {cinema.name}
                  </h3>
                  <p className="text-gray-300 mb-2">{cinema.address}</p>
                  <p className="text-gray-300">{cinema.phone}</p>
                </div>
                {/* Дополнительная информация (как в вашем примере) */}
                <div className="mt-4 flex justify-between items-center">
                  <span className="bg-purple-600 text-white px-3 py-1 rounded-full text-sm">
                    {cinema.category}
                  </span>
                  <span className="text-gray-300 text-sm">
                    {cinema.condition}
                  </span>
                </div>

                {/* Фоновое изображение, появляющееся при наведении */}
                <div className="absolute inset-0 overflow-hidden opacity-0 group-hover:opacity-20 transition-opacity duration-300">
                  <Image
                    src={cinema.photo || "/choicePhoto.jpg"}
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
