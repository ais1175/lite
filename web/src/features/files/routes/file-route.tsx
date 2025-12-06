import { useParams } from "react-router";
import { useFile } from "../api/use-file";

export default function FileRoute() {
  const { organizationId, fileId } = useParams<{
    organizationId: string;
    fileId: string;
  }>();

  const { data, isPending } = useFile(organizationId, fileId);

  if (isPending) {
    return <div>Loading...</div>;
  }

  if (!data) {
    return <div>File not found</div>;
  }

  return (
    <div>
      <div>{data.key}</div>
    </div>
  );
}
