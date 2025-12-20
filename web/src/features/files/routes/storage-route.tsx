import { Suspense } from "react";
import { DataTableSkeleton } from "@/components/data/DataTableSkeleton";
import { UploadDialog } from "../components/upload-dialog";
import { AssetList } from "../components/asset-list";

export default function StorageRoute() {
  return (
    <div>
      <UploadDialog />
      <div className="mt-4">
        <Suspense fallback={<DataTableSkeleton columnCount={10} />}>
          <AssetList />
        </Suspense>
      </div>
    </div>
  );
}
