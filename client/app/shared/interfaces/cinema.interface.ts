export interface ICinema {
  id: number;
  name: string;
  photo: string;
  address: string;
  category: string;
  condition: string;
  description: string;
  email: string;
  phone: string;
}

export type ICinemaMain = Omit<
  ICinema,
  "photo" | "category" | "condition" | "description" | "email" | "phone"
>;

export interface ICinemaDto {
  address: string;
  category: string;
  condition: string;
  description: string;
  email: string;
  name: string;
  phone: string;
  photo?: File | string;
}

export interface ICinemaExportDto {
  address: string;
  category_id: number;
  condition_id: number;
  description: string;
  email: string;
  name: string;
  phone: string;
  photo?: File | string;
}

export type IListofCinema = ICinemaMain[];

export interface IHall {
  capacity: number;
  id: number;
  name: string;
  type: string;
}

export interface ICategory {
  id: number;
  name: string;
}

export type IListOfCategory = ICategory[];

export interface ICondition {
  id: number;
  name: string;
}

export type IListOfCondition = ICondition[];
