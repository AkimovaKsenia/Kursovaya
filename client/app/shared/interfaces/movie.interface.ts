export interface IMovie {
  id: number;
  name: string;
  genres: string[];
  description: string;
  directors: string[];
  operators: string[];
  photo: string;
  film_studio_name: string;
  duration_in_min: number;
  cast_list: string[];
}

export interface IMovieDto {
  name: string;
  genres: string[]; // или string, если бек принимает как одну строку
  description: string;
  photo: File | string; // теперь это File, не string
  directors: string[];
  operators: string[];
  film_studio_name: string;
  cast_list: string[];
}

export interface IListOfMovies {
  newMovies: IMovie[];
}

export interface IGenre {
  id: number;
  name: string;
}

export type IListOfGenres = IGenre[];
