import { useQuery } from "@tanstack/react-query";
import { CinemaService } from "services/cinema.service";
import { UserService } from "services/user.service";
import { IListOfRoles } from "shared/interfaces/user.interface";

export const useRole = () => {
  const { data: rolesData, error: rolesError } = useQuery<IListOfRoles, Error>({
    queryKey: ["roles"],
    queryFn: () => UserService.getRoles().then((res) => res.data),
  });
  return { data: rolesData, error: rolesError };
};
