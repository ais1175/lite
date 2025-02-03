import { QueryKeys } from "@/typings/query";
import { fetchApi } from "@/utils/http-util";
import { useQuery } from "@tanstack/react-query";

export function useSession() {
  console.log("getting session");
  const { data } = useQuery({
    queryKey: [QueryKeys.Session],
    queryFn: () => fetchApi("/api/auth/session"),
  });

  return data;
}
