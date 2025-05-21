import { FC, PropsWithChildren } from "react";
import Meta from "utils/Meta";
import { IMeta } from "utils/meta.interface";
import Header from "./header/Header";

const Layout: FC<PropsWithChildren<IMeta>> = ({
  children,
  backgroundImage,
  backgroundColor,
  ...meta
}) => {
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
};
export default Layout;
