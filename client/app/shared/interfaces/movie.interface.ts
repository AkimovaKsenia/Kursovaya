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

export interface IMovieDto
  extends Pick<
    IMovie,
    | "name"
    | "genres"
    | "description"
    | "photo"
    | "directors"
    | "operators"
    | "film_studio_name"
    | "cast_list"
  > {}

export interface IListOfMovies {
  newMovies: IMovie[];
}
