import { Card } from "@/components/ui/card";
import { Asset } from "@/typings/asset";
import { FileIcon } from "lucide-react";

interface FileViewerProps {
  fileUrl: string;
  fileName: string;
  file: Asset;
}

export function FileViewer({ file, ...rest }: FileViewerProps) {
  return (
    <Card className="overflow-hidden">
      <div className="relative bg-muted/30 min-h-[400px] flex items-center justify-center p-4">
        {file.type === "image" ? (
          <img
            src={rest.fileUrl}
            alt={rest.fileName}
            className="max-h-[600px] w-auto h-auto object-contain rounded-md shadow-sm"
          />
        ) : file.type === "video" ? (
          <video controls className="max-h-[600px] w-full rounded-md shadow-sm">
            <source src={rest.fileUrl} type={file.type} />
            Your browser does not support the video tag.
          </video>
        ) : (
          <div className="flex flex-col items-center gap-4 text-muted-foreground py-20">
            <div className="p-6 rounded-full bg-background shadow-sm">
              <FileIcon className="h-16 w-16 text-primary/40" />
            </div>
            <p className="text-lg font-medium">
              No preview available for this file type
            </p>
          </div>
        )}
      </div>
    </Card>
  );
}
