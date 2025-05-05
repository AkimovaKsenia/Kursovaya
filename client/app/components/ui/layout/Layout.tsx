import { FC, PropsWithChildren } from "react";
import Meta from "utils/Meta";
import { IMeta } from "utils/meta.interface";
import Header from "./header/Header";
import AdminHeader from "../AdminHeader";

const Layout: FC<PropsWithChildren<IMeta>> = ({ children, ...meta }) => {
  return (
    <>
      <Meta {...meta} />
      <div className="min-h-screen bg-gray-300 text-white">
        <Header />
        <main className="p-4">{children}</main>
      </div>
    </>
  );
};
export default Layout;
