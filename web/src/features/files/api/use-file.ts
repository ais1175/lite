import { Asset } from "@/typings/asset";
import { fetchApi } from "@/utils/http-util";
import { useQuery } from "@tanstack/react-query";

export function useFile(
  organizationId: string | undefined,
  fileId: string | undefined,
) {
  const { data, isPending } = useQuery({
    queryKey: ["file"],
    queryFn: async () => {
      return fetchApi<Asset>(
        `/api/dash/storage/${organizationId}/file/${fileId}`,
      );
    },
  });

  return { data, isPending };
}
