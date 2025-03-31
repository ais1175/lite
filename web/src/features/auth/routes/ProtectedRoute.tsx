import { useOrganizations } from "@/features/organizations/api/useOrganizations";
import { Outlet, Navigate, useParams } from "react-router";

type AppParams = {
  organizationId: string;
};

export const ProtectedRoute: React.FC = () => {
  const params = useParams<AppParams>();

  const { data: organizations } = useOrganizations();
  if (organizations && organizations.length === 0) {
    return <Navigate to="/app/new-organization" />;
  }

  if (!params.organizationId) {
    return <Navigate to="/app" />;
  }

  return <Outlet />;
};
