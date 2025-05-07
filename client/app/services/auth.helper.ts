import Cookies from "js-cookie";
import { IAuthResponse } from "shared/interfaces/user.interface";

export const saveTokenToStorage = (token: string) => {
  Cookies.set("token", token);
};
export const removeTokenFromStorage = () => {
  Cookies.remove("token");
};

export const saveToStorage = (data: IAuthResponse) => {
  saveTokenToStorage(data.token);
  localStorage.setItem(
    "user",
    JSON.stringify({ role: data.role, token: data.token })
  );
};
