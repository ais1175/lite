import { Navigate, NavLink } from "react-router";
import { useOrganizations } from "../api/useOrganizations";
import { useSession } from "@/features/auth/api/useSession";

export function OrganizationSelectRoute() {
  const { data, isLoading } = useOrganizations();

  // todo: move this to a layout route
  const session = useSession();

  if (!session.data && !session.isPending) {
    return <Navigate to="/auth" />;
  }

  // switch out to use suspense
  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!data || data.length === 0) {
    return <Navigate to="/app/new-organization" />;
  }

  return (
    <div className="h-screen flex flex-col items-center justify-center">
      <div className="border max-w-2xl w-full p-2 rounded-md">
        {data &&
          data.map((org) => (
            <NavLink
              to={`/app/${org.id}`}
              state={{ organization: org }}
              key={org.id}
            >
              <div
                key={org.id}
                className="p-4 hover:bg-accent bg-accent/10 rounded-md"
              >
                <h2 className="text-lg font-medium">{org.name}</h2>
              </div>
            </NavLink>
          ))}
      </div>
    </div>
  );
}
