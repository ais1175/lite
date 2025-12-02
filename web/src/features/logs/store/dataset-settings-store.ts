import { atomWithStorage } from "jotai/utils";

const LS_DATASET_SETTINGS_KEY = "datasetSettings";

export interface DatasetSettings {
  hideRawData: boolean;
}

export const datasetSettingsStore = atomWithStorage<DatasetSettings>(
  LS_DATASET_SETTINGS_KEY,
  {
    hideRawData: true,
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
