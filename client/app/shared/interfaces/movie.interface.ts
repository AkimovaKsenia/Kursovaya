export interface IMovie {
  id: number;
  name: string;
  genre: string;
  description: string;
  director: string;
  producer: string;
  poster: string;
  film_studio: string;
  duration: number;
  views?: number;
  actors?: [];
}
