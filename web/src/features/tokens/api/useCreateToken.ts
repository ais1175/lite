import { TokenParams, TokenPreview } from "@/typings/token";
import { ApiError, fetchApi } from "@/utils/http-util";
import { useMutation } from "@tanstack/react-query";

export function useCreateToken() {
  const { mutateAsync, data, isSuccess } = useMutation({
    mutationFn: async (params: TokenParams) => {
      try {
        return await fetchApi<TokenPreview>("/api/token", {
          method: "POST",
          body: JSON.stringify(params),
        });
      } catch (err) {
        if (err instanceof ApiError) {
          throw new Error(err.message);
        }
      }
    },
  });

  return { mutateAsync, data, isSuccess };
}
