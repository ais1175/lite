import { datasetSettingsStore } from "../store/dataset-settings-store";
import { useAtom } from "jotai/react";

export const useDatasetSettings = () => {
  const [datasetSettings, setDatasetSettings] = useAtom(datasetSettingsStore);

  const toggleRawData = () => {
    setDatasetSettings((prev) => ({
      ...prev,
      hideRawData: !prev.hideRawData,
    }));
  };

  return {
    datasetSettings,
    toggleRawData,
  };
};
