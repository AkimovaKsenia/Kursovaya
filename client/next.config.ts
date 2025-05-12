import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactStrictMode: true,
  poweredByHeader: false,
  env: {
    APP_URL: process.env.REACT_APP_URL,
  },
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "s1ru1.kinoplan24.ru",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "s2ru1.kinoplan24.ru",
        pathname: "/**",
      },
      {
        protocol: "http",
        hostname: "localhost",
        port: "10000",
        pathname: "/**",
      },
      {
        protocol: "http",
        hostname: "minio",
        port: "9000",
        pathname: "/**",
      },
    ],
  },
  async rewrites() {
    return [
      {
        source: "/auth/:path*", // только API-запросы
        destination: "http://localhost:10000/auth/:path*",
      },
      {
        source: "/login", // авторизация
        destination: "http://localhost:10000/login",
      },
    ];
  },
};

export default nextConfig;
