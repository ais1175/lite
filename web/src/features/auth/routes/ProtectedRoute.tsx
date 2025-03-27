import { Outlet, useLocation, Navigate, useParams } from "react-router";

type AppParams = {
  organizationId: string;
};

export const ProtectedRoute: React.FC = () => {
  const location = useLocation();
  const params = useParams<AppParams>();

  if (location.pathname === "/") {
    return <Navigate to="/workspaces" />;
  }

  console.log(params.organizationId);

  if (!params.organizationId) {
    return <Navigate to="/app/123" />;
  }

  return <Outlet />;
};
