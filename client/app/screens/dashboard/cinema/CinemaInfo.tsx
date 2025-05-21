import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/router";
import { CinemaService } from "services/cinema.service";
import Image from "next/image";
import Head from "next/head";
import { FC, useEffect, useState } from "react";

const CinemaInfo: FC = () => {
  const router = useRouter();
  const { id } = router.query;
  const [imageUrl, setImageUrl] = useState<string>("/default-cinema.jpg");

  const { data: cinema, isLoading } = useQuery({
    queryKey: ["cinema", id],
    queryFn: () =>
      CinemaService.getCinemaById(Number(id)).then((res) => res.data),
    enabled: !!id && !isNaN(Number(id)),
  });
  useEffect(() => {
    if (cinema?.photo instanceof File) {
      const url = URL.createObjectURL(cinema.photo);
      setImageUrl(url);

      // Очистка при размонтировании
      return () => URL.revokeObjectURL(url);
    } else if (typeof cinema?.photo === "string") {
      setImageUrl(cinema.photo);
    }
  }, [cinema?.photo]);
  if (isLoading) return <div>Загрузка...</div>;
  if (!cinema) return <div>Кинотеатр не найден</div>;

  return (
    <div className="bg-black text-white min-h-screen">
      <Head>
        <title>{cinema.name} | Кинотеатр</title>
        <meta
          name="description"
          content={`Информация о кинотеатре ${cinema.name}`}
        />
      </Head>

      {/* Hero Section */}
      <div className="relative h-90 w-full">
        <Image
          src={imageUrl || "/default-cinema.jpg"}
          alt={cinema.name}
          layout="fill"
          objectFit="cover"
          objectPosition="left 1%"
          className="opacity-40"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-black to-transparent" />

        {/* Эллипс (фоновой декоративный элемент) */}
        <div className="absolute mt-105 ml-30 z-0 pointer-events-none w-[100] h-[500px]">
          {/* <Image
            src="/Ellipse3.png"
            alt="Cinema ellipse"
            width={800}
            height={700}
            className="rounded-full"
          /> */}
        </div>
        <div className="absolute bottom-[1.5rem] left-32 right-8">
          <h1 className="text-4xl md:text-6xl font-bold mb-2">{cinema.name}</h1>
          <p className="text-xl">{cinema.address}</p>
        </div>
      </div>

      {/* Main Content */}
      <div className="container mx-auto px-4 py-3 mt-10 relative z-20">
        <div className="flex flex-col items-center ml-5 gap-8 pt-0">
          {/* Картинка и блок "О кинотеатре" */}
          <div className="flex flex-col md:flex-row gap-6">
            <div className="flex-shrink-0 max-h-[600px] overflow-hidden  rounded-xl  ">
              <Image
                src={imageUrl || "/default-cinema.jpg"}
                alt="Cinema"
                width={700}
                height={700}
                objectPosition="bottom"
                className="object-center rounded-xl  max-h-md "
              />
            </div>
            <div className="flex-grow">
              <div className=" p-6 rounded-lg h-full max-w-lg ml-15">
                <div className="absolute mt-30 ml-90 z-0 pointer-events-none w-[900] h-[700px]">
                  {/* <Image
                    src="/Ellipse3.png"
                    alt="Cinema ellipse"
                    width={900}
                    height={700}
                    className="rounded-full"
                  /> */}
                </div>
                <h2 className="text-3xl font-bold mb-6 pb-2 relative inline-block">
                  <span className="relative z-10">О кинотеатре</span>
                  <span className="absolute bottom-0 left-0 w-full h-1 bg-gradient-to-r from-purple-500 to-transparent rounded-full" />
                </h2>
                <p className="text-gray-300 leading-relaxed">
                  {cinema.description}
                </p>
                <div className="rounded-lg mt-8">
                  <h3 className="text-xl font-semibold mb-2">Условия</h3>
                  <p className="text-gray-400">
                    <span className="font-medium">Тип:</span> {cinema.category}
                  </p>
                  <p className="text-gray-400">
                    <span className="font-medium">Состояние:</span>{" "}
                    {cinema.condition}
                  </p>
                </div>
                <div className="rounded-lg mt-8">
                  <h3 className="text-xl font-semibold mb-2">
                    Контактная информация
                  </h3>
                  <p className="text-gray-400">
                    <span className="font-medium">Телефон:</span> {cinema.phone}
                  </p>
                  <p className="text-gray-400">
                    <span className="font-medium">Email:</span> {cinema.email}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-gray-900 rounded-lg overflow-hidden max-w-2xl min-h-xl"></div>
        </div>
      </div>
    </div>
  );
};

export default CinemaInfo;
