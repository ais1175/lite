import { QueryKeys } from "@/typings/query";
import { Session } from "@/typings/user";
import { fetchApi } from "@/utils/http-util";
import { useQuery } from "@tanstack/react-query";

export function useSession() {
  const { data } = useQuery({
    queryKey: [QueryKeys.Session],
    queryFn: () => fetchApi<Session>("/api/auth/session"),
  });

  return data;
}
