import { useQuery } from "@tanstack/react-query";

export interface VersionStatus {
  current: string;
  latest: string;
  update_available: boolean;
  last_checked: string;
}

export interface SystemConfig {
  bucket_domain: string;
}

export const fetchVersionStatus = async (): Promise<VersionStatus> => {
  const response = await fetch("/api/dash/system/version");
  if (!response.ok) {
    throw new Error("Failed to fetch version status");
  }
  return response.json();
};

export const fetchSystemConfig = async (): Promise<SystemConfig> => {
  const response = await fetch("/api/dash/system/config");
  if (!response.ok) {
    throw new Error("Failed to fetch system config");
  }
  return response.json();
};

export const useVersion = () => {
  return useQuery({
    queryKey: ["version-status"],
    queryFn: fetchVersionStatus,
    staleTime: 1000 * 60 * 60, // 1 hour
  });
};

export const useConfig = () => {
  return useQuery({
    queryKey: ["system-config"],
    queryFn: fetchSystemConfig,
    staleTime: Infinity, // Config shouldn't change during session
  });
};
