import axios from "axios";
import Cookies from "js-cookie";

export const getContentType = () => ({
  "Content-Type": "application/json",
});

export const API_URL = `${process.env.APP_URL}`;

export const axiosClassic = axios.create({
  baseURL: API_URL,
  headers: getContentType(),
});

const instance = axios.create({
  baseURL: API_URL,
  headers: getContentType(),
});

instance.interceptors.request.use((config) => {
  const token = Cookies.get("token");

  if (config.headers && token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default instance;
