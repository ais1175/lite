import { addDays, endOfDay, startOfDay } from "date-fns";
import { PanelLeftCloseIcon, PanelLeftOpenIcon } from "lucide-react";
import { useEffect } from "react";

import { Button } from "@/components/ui/button";
//import { QueryModeSwitcher } from "../[datasetId]/_components/QueryModeSwitcher";
//import { LogsColumnsSettingsDropdown } from "./LogsColumnsSettingsDropdown";
import { useSidebarControls } from "../../hooks/use-sidebar-controls";
import { useSearchParams } from "react-router";
//import { AutoRefresh } from "./auto-refresh";

export function LogsFilterWrapper() {
  const [searchParams, setSearchParams] = useSearchParams();

  const { sidebarControls, toggleSidebarControls } = useSidebarControls();

  function handleShowControls() {
    toggleSidebarControls();
  }

  useEffect(() => {
    const params = new URLSearchParams(searchParams);
    const qr = params.get("qr");
    const from = params.get("from");
    const to = params.get("to");

    if (qr) {
      return;
    }

    if (!from || !to) {
      params.set("from", startOfDay(new Date(Date.now())).toUTCString());
      params.set(
        "to",
        endOfDay(new Date(addDays(Date.now(), 1))).toUTCString(),
      );
    }

    setSearchParams(params);
  }, [searchParams, setSearchParams]);

  return (
    <div className="mb-2 flex items-center justify-between">
      <div className="flex items-center justify-between space-x-4">
        <div className="flex items-center gap-2">
          <Button variant="ghost" onClick={handleShowControls}>
            {sidebarControls.showSidebarControls ? (
              <PanelLeftCloseIcon />
            ) : (
              <PanelLeftOpenIcon />
            )}
            {sidebarControls.showSidebarControls
              ? "Hide controls"
              : "Show controls"}
          </Button>

          {/*<QueryModeSwitcher />*/}
        </div>
      </div>
      <div className="flex flex-row gap-1">
        {/*<AutoRefresh />*/}
        {/*<LogsColumnsSettingsDropdown />*/}
      </div>
    </div>
  );
}
