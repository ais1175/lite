import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/Input";
import { IS_DEV } from "@/utils/http-util";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const AuthRoute: React.FC = () => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const navigate = useNavigate();

  const handleRegister = async () => {
    const res = await fetch("/api/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    if (IS_DEV) {
      if (res.ok) {
        navigate("/app");
      }
    }
  };

  return (
    <main className="bg-background w-full h-screen dark font-medium text-3xl p-4">
      <h1 className="text-foreground">Fivemanage Lite</h1>

      <div className="space-y-4 my-6">
        <Input
          value={email}
          onChange={(e) => setEmail(e.currentTarget.value)}
          placeholder="Email"
        />
        <Input
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.currentTarget.value)}
        />
      </div>

      <Button onClick={handleRegister}>Login</Button>
    </main>
  );
};
