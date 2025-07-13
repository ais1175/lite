import { lazy } from "react";
import { List } from "@/components/ui/list";
import { DatasetList } from "./dataset-list/dataset-list";
const DatasetSheet = lazy(() => import("./dataset-sheet/dataset-sheet"));

interface DatasetActionsProps {
  organizationId: string | undefined;
}

export default function DatasetOverview({
  organizationId,
}: DatasetActionsProps) {
  return (
    <div className="w-full flex justify-center">
      <div className="w-full max-w-6xl mx-auto">
        <List>
          <List.Header title="Datasets">
            <DatasetSheet organizationId={organizationId} />
            {/*<DatasetActions
            query={searchParams.get("q")}
            datasetCount={data?.length}
            maxDatasets={maxDatasets}
            maxRetentionDays={maxRetentionDays}
          />*/}
          </List.Header>
          <DatasetList />
        </List>
      </div>
    </div>
  );
}
