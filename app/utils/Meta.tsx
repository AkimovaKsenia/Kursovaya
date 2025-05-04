import { FC } from "react";
import { IMeta } from "./meta.interface";
import Head from "next/head";
import Link from "next/link";

const Meta: FC<IMeta> = ({ description, title }) => {
  return (
    <>
      <Head>
        <title>{title}</title>
        <Link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
        {description ? (
          <meta
            itemProp="description"
            name="description"
            content={description}
          />
        ) : (
          <meta name="robots" content="noindex,nofollow" />
        )}
      </Head>
    </>
  );
};
export default Meta;
