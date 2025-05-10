import { Suspense } from "react";
import { AssetList } from "./components/AssetList";
import { DataTableSkeleton } from "@/components/data/DataTableSkeleton";

export default function StorageRoute() {
  return (
    <div>
      <Suspense fallback={<DataTableSkeleton columnCount={10} />}>
        <AssetList />
      </Suspense>
    </div>
  );
}
