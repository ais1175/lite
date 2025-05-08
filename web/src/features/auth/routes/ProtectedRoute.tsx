import { useOrganizations } from "@/features/organizations/api/useOrganizations";
import { Outlet, Navigate, useParams } from "react-router";
import { useSession } from "../api/useSession";

type AppParams = {
  organizationId: string;
};

export const ProtectedRoute: React.FC = () => {
  const params = useParams<AppParams>();

  const session = useSession();
  const { data: organizations } = useOrganizations();

  if (!session.data && !session.isPending) {
    return <Navigate to="/auth" />;
  }

  if (organizations && organizations.length === 0) {
    return <Navigate to="/app/new-organization" />;
  }

  if (!params.organizationId) {
    return <Navigate to="/app" />;
  }

  return <Outlet />;
};
