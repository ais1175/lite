import { useListFields } from "../api/logs-api";
import { LogsTable } from "./logs-table/logs-table";

interface LogsStreamProps {
  organizationId: string | undefined;
  datasetId: string | undefined;
}

export function LogsStream({ organizationId, datasetId }: LogsStreamProps) {
  const { data: fields } = useListFields(organizationId, datasetId);

  if (!datasetId || !organizationId) {
    return null;
  }

  return (
    <div className="relative flex flex-col h-[calc(100vh-5rem)] overflow-hidden -m-4 -mt-0">
      <LogsTable
        fields={fields}
        teamId={organizationId}
        datasetId={datasetId}
        levels={["info", "warn", "error"]}
        defaultRowSelection={{}}
        defaultColumnSizing={{}}
        defaultColumnVisibility={{
          data: true,
        }}
      />
    </div>
  );
}
