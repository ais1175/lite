import { atomWithStorage } from "jotai/utils";

export const LS_LOGS_SIDEBAR_CONTROLS_SETTINGS_KEY =
  "logsSidebarControlsSettings";

interface LogsSidebarControlsSettings {
  showSidebarControls: boolean;
}

export const logsSidebarControlsSettingsStore =
  atomWithStorage<LogsSidebarControlsSettings>(
    LS_LOGS_SIDEBAR_CONTROLS_SETTINGS_KEY,
    {
      showSidebarControls: true,
    },
    {
      getItem(key) {
        const item = window.localStorage.getItem(key);
        return item ? JSON.parse(item) : {};
      },
      setItem(key, newValue) {
        window.localStorage.setItem(key, JSON.stringify(newValue));
      },
      removeItem(key) {
        window.localStorage.removeItem(key);
      },
    },
  );
