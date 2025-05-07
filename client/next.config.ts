import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactStrictMode: true,
  poweredByHeader: false,
  env: {
    APP_URL: process.env.REACT_APP_URL,
  },
  images: {
    domains: ["localhost"],
  },
  async rewrites() {
    return [
      {
        source: "/:path*",
        destination: "http://localhost:10000/uploads/:path*",
      },
    ];
  },
};

export default nextConfig;
