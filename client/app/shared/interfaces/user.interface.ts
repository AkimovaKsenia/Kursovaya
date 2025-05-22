export interface IUser {
  id: number;
  name: string;
  email: string;
}

export interface IAuthResponse {
  role: string;
  token: string;
}

export interface IUserDto {
  email: string;
  password: string;
  name: string;
  role: string;
  surname: string;
}

export interface IUserMain {
  id: number;
  name: string;
  surname: string; // Добавляем фамилию
  email: string;
  role: string; // Добавляем роль
}

export type IListOfUsers = IUserMain[];

export interface IUserExportDto {
  email: string;
  password: string;
  name: string;
  role_id: number;
  surname: string;
}

export interface IRole {
  id: number;
  name: string;
}

export type IListOfRoles = IRole[];
