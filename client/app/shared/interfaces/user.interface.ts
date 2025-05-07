export interface IUser {
  id: number;
  name: string;
  email: string;
}

export interface IAuthResponse {
  role: string;
  token: string;
}
