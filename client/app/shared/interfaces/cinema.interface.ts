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

export type IListofCinema = ICinemaMain[];

export interface IHall {
  capacity: number;
  id: number;
  name: string;
  type: string;
}
