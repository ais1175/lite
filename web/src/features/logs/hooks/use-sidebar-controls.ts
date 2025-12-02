import { useAtom } from "jotai/react";
import { logsSidebarControlsSettingsStore } from "../store/logs-sidebar-controls-settings";

export function useSidebarControls() {
  const [sidebarControls, setSidebarControls] = useAtom(
    logsSidebarControlsSettingsStore,
  );

  const toggleSidebarControls = () => {
    setSidebarControls((prev) => ({
      ...prev,
      showSidebarControls: !prev.showSidebarControls,
    }));
  };

  return {
    sidebarControls,
    toggleSidebarControls,
  };
}
