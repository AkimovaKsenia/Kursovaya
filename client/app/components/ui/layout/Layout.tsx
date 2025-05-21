import { FC, PropsWithChildren, useEffect, useState } from "react";
import Meta from "utils/Meta";
import { IMeta } from "utils/meta.interface";
import Header from "./header/Header";

const Layout: FC<PropsWithChildren<IMeta>> = ({
  children,
  backgroundImage,
  backgroundColor,
  ...meta
}) => {
  const [isImageLoaded, setIsImageLoaded] = useState(false);
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    setIsMounted(true);

    if (!backgroundImage) {
      setIsImageLoaded(true);
      return;
    }

    const img = new Image();
    img.src = backgroundImage;
    img.onload = () => setIsImageLoaded(true);
    img.onerror = () => setIsImageLoaded(true);

    return () => {
      img.onload = null;
      img.onerror = null;
    };
  }, [backgroundImage]);

  // На сервере всегда рендерим без прелоадера
  if (!isMounted) {
    return (
      <>
        <Meta {...meta} />
        <div
          className="min-h-screen bg-gray-300 text-white"
          style={{
            backgroundImage: backgroundImage
              ? `url(${backgroundImage})`
              : undefined,
            backgroundColor: backgroundColor || "rgb(163, 163, 181)",
            backgroundSize: backgroundImage ? "auto 100%" : undefined,
            backgroundRepeat: "no-repeat",
            backgroundPosition: "120% center",
          }}
        >
          <Header />
          <main className="p-4">{children}</main>
        </div>
      </>
    );
  }
  return (
    <>
      <Meta {...meta} />
      <div
        className="min-h-screen bg-gray-300 text-white"
        // style={{
        //   backgroundImage: backgroundImage
        //     ? `url(${backgroundImage})`
        //     : undefined,
        //   backgroundColor: backgroundColor || "rgb(163, 163, 181)",
        //   backgroundSize: backgroundImage ? "auto 100%" : undefined,
        //   backgroundRepeat: "no-repeat",
        //   backgroundPosition: "120% center",
        // }}
      >
        {/* Прелоадер только на клиенте */}
        {!isImageLoaded && (
          <div className="fixed inset-0 z-50 flex items-center justify-center bg-black">
            <div className="text-white">Загрузка...</div>
          </div>
        )}

        {/* Фоновое изображение */}
        {backgroundImage && (
          <div
            className={`absolute inset-0 transition-opacity duration-500 ${
              isImageLoaded ? "opacity-100" : "opacity-0"
            }`}
            style={{
              backgroundImage: `url(${backgroundImage})`,
              backgroundSize: "auto 100%",
              backgroundRepeat: "no-repeat",
              backgroundPosition: "120% center",
            }}
          />
        )}
        <div
          className={`relative z-10 transition-opacity duration-500 ${
            isImageLoaded ? "opacity-100" : "opacity-0"
          }`}
        >
          <Header />
          <main className="p-4">{children}</main>
        </div>
      </div>
    </>
  );
};
export default Layout;
