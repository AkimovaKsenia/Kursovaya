import { Dispatch, SetStateAction } from "react";
import { IAuthResponse, IUser } from "shared/interfaces/user.interface";

export type TypeUserState = IAuthResponse | null;
export interface IContext {
  user: TypeUserState;
  setUser: Dispatch<SetStateAction<TypeUserState>>;
}
