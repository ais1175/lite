import { useSession } from "@/features/auth/api/useSession";

export const AppDashboard: React.FC = () => {
  const data = useSession();

  return (
    <div>
      <pre>{JSON.stringify(data)}</pre>
    </div>
  );
};
