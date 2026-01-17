import { Suspense } from "react";
import { CreateTokenDialog } from "../components/CreateTokenDialog";
import { TokensTable } from "../components/TokensTable";
import { Skeleton } from "@/components/ui/skeleton";

function TokensTableSkeleton() {
  return (
    <div className="rounded-xl border shadow-sm">
      <div className="h-10 bg-muted/30 border-b flex items-center px-4">
        <Skeleton className="h-4 w-24" />
        <Skeleton className="h-4 w-16 ml-auto mr-32" />
        <Skeleton className="h-4 w-16 ml-auto" />
      </div>
      {[...Array(3)].map((_, i) => (
        <div
          key={i}
          className="h-12 border-b last:border-0 flex items-center px-4"
        >
          <Skeleton className="h-4 w-48" />
          <Skeleton className="h-4 w-12 ml-auto mr-32" />
          <Skeleton className="h-8 w-8 ml-auto rounded-md" />
        </div>
      ))}
    </div>
  );
}

export default function TokensRoute() {
  return (
    <main className="container mx-auto py-10 max-w-5xl">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">API Tokens</h1>
          <p className="text-muted-foreground">
            Manage your API tokens to access the platform.
          </p>
        </div>
        <CreateTokenDialog />
      </div>
      <section>
        <Suspense fallback={<TokensTableSkeleton />}>
          <TokensTable />
        </Suspense>
      </section>
    </main>
  );
}
