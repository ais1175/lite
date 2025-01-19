import { useSession } from "@/features/auth/api/useSession";

export const AppDashboard: React.FC = () => {
  const data = useSession();

  return (
    <div>
      <h1>dashboard</h1>
      <pre>{JSON.stringify(data)}</pre>
    </div>
  );
};
