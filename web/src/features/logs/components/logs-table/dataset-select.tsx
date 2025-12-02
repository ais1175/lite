import { useNavigate, useParams } from "react-router";
import { useListDatasets } from "../../api/dataset-api";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectValue,
  SelectItem,
} from "@/components/ui/select";

interface DatasetSelectProps {
  teamId: string;
}

export function DatasetSelect({ teamId }: DatasetSelectProps) {
  const params = useParams<{ organizationId: string; datasetId: string }>();
  const navigate = useNavigate();

  const { data: datasets } = useListDatasets(params.organizationId);

  const currentDataset = datasets?.find((d) => d.id === params.datasetId);

  function handleChange(value: string) {
    navigate(`/logs/${teamId}/${value}`);
  }

  return (
    <Select value={currentDataset?.id} onValueChange={handleChange}>
      <SelectTrigger className="min-w-[180px] text-xs">
        <SelectValue placeholder="Select a dataset" />
      </SelectTrigger>
      <SelectContent>
        {datasets?.map((dataset) => (
          <SelectItem key={dataset.id} value={dataset.id} className="text-xs">
            {dataset.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}
