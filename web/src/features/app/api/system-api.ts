import { $api } from "@/lib/api/client";

export const useVersion = () => {
  return $api.useQuery("get", "/dash/system/version", {
    staleTime: 1000 * 60 * 60, // 1 hour
  });
};

export const useConfig = () => {
  return $api.useQuery("get", "/dash/system/config", {
    staleTime: Infinity, // Config shouldn't change during session
  });
};
