import { lazy } from "react";
import { List } from "@/components/ui/list";
import { useNavigate } from "react-router";
import { $api } from "@/lib/api/client";
const DatasetSheet = lazy(() => import("./dataset-sheet/dataset-sheet"));

interface DatasetActionsProps {
  organizationId: string;
}

export default function DatasetOverview({
  organizationId,
}: DatasetActionsProps) {
  const { data: response } = $api.useQuery(
    "get",
    "/dash/{organizationId}/dataset",
    {
      params: {
        path: { organizationId },
      },
    },
  );
  const datasets = response?.data;
  const navigate = useNavigate();

  function navigateToDataset(id: string) {
    navigate(`${id}`);
  }

  return (
    <div className="w-full flex justify-center">
      <div className="w-full max-w-6xl mx-auto">
        <List>
          <List.Header title="Datasets">
            <DatasetSheet organizationId={organizationId} />
          </List.Header>
          <div>
            {datasets &&
              datasets.map((dataset) => (
                <List.Item
                  key={dataset.id}
                  title={dataset.name}
                  subtitle={dataset.description}
                  onClick={() => navigateToDataset(dataset.id)}
                />
              ))}
          </div>
        </List>
      </div>
    </div>
  );
}
