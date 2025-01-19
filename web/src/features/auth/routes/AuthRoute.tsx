import { AuthForm } from "../components/AuthForm";

export const AuthRoute: React.FC = () => {
  return (
    <main className="h-full bg-background dark">
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-sm">
          <AuthForm />
        </div>
      </div>
    </main>
  );
};
