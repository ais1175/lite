import { Button } from "@/components/ui/Button";
import { KeyRound } from "lucide-react";

export function TokensRoute() {
  return (
    <main>
      <Button size="sm">
        <KeyRound size={16} className="mr-2" />
        Create token
      </Button>
    </main>
  );
}
