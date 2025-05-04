export interface IMovie {
  id: number;
  name: string;
  genre: string;
  director: string;
  operator: string;
  poster: string;
  film_studio: string;
  duration: number;
  views?: number;
  fees?: number;
  actors?: [];
}
