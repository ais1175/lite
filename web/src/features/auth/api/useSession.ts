import { fetchApi } from "@/utils/http-util";
import { useQuery } from "@tanstack/react-query";

export const useSession = () => {
  const { data } = useQuery({
    queryKey: ["session"],
    queryFn: async ({ signal }) => {
      return await fetchApi("/api/auth/session", {
        signal,
      });
    },
  });

  return data;
};
