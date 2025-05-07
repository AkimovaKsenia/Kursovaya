import { axiosClassic } from "api/interceptor";

export const ViewsService = {
  async updateViews(Movieid: string) {
    return axiosClassic.patch(`/views/update/${Movieid}`);
  },
};
