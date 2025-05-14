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
  photo?: File | string; // теперь это File, не string
  directors: string[];
  operators: string[];
  film_studio_name: string;
  cast_list: string;
  duration_in_min: number;
}

export interface IMovieExportDto {
  name: string;
  description: string;
  film_photo?: File | string;
  cast_list: string[];
  film_studio_id: number;
  duration_in_min: number;
  director_ids: number[];
  operator_ids: number[];
  genre_ids: number[];
}

export interface IListOfMovies {
  newMovies: IMovie[];
}

export interface IGenre {
  id: number;
  name: string;
}

export type IListOfGenres = IGenre[];

export interface IOperator {
  fio: string;
  id: number;
}

export type IListOfOperators = IOperator[];

export interface IFilmStudio {
  id: number;
  name: string;
}

export type IListOfFilmStudio = IFilmStudio[];

export interface IDirector {
  fio: string;
  id: number;
}

export type IListOfDirector = IDirector[];
