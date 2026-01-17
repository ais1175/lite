import { Card, CardFooter } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

export function AssetGridSkeleton({ rowCount = 10 }: { rowCount?: number }) {
  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
      {Array.from({ length: rowCount }).map((_, i) => (
        <Card key={i} className="overflow-hidden py-0 gap-0">
          <Skeleton className="aspect-video w-full" />
          <CardFooter className="p-3 flex flex-col items-start gap-1">
            <Skeleton className="h-4 w-3/4" />
            <Skeleton className="h-3 w-1/4" />
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
