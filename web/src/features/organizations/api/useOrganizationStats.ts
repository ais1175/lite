import { ApiError, fetchApi } from "@/utils/http-util";
import { useQuery } from "@tanstack/react-query";

export interface OrganizationStats {
  totalFiles: number;
  totalSize: number;
  totalLogs: number;
  totalTokens: number;
  totalDataset: number;
}

export function useOrganizationStats(organizationId: string | undefined) {
  return useQuery({
    queryKey: ["organization-stats", organizationId],
    enabled: !!organizationId,
    queryFn: async () => {
      try {
        return await fetchApi<OrganizationStats>(
          `/api/dash/organization/${organizationId}/stats`,
        );
      } catch (err) {
        if (err instanceof ApiError) {
          throw new Error(err.message);
        }
      }
    },
  });
}
