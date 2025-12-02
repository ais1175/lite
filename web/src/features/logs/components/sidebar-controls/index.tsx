import { DatetimePopover } from "./datetime-popover";
import { LevelSelect } from "./level-select";

interface LogsSidebarControlsProps {
  levels: string[] | null | undefined;
}

export function LogsSidebarControls({ levels }: LogsSidebarControlsProps) {
  return (
    <div className="h-full w-96 bg-background border-r flex flex-col">
      <div className="p-4 border-b">
        <h1 className="font-medium text-sm">Controls</h1>
      </div>
      <div className="p-2 space-y-4 flex-shrink-0">
        <div>
          <DatetimePopover />
        </div>
        <div>
          <LevelSelect levels={levels} />
        </div>
      </div>
    </div>
  );
}
