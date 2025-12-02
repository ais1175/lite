import { useAtom } from "jotai/react";

import {
  type CustomColumn,
  customColumnsAtom,
  rawDataColumnAtom,
  rawDataColumnWidthAtom,
} from "../store/custom-columns-store.ts";
import { useCallback } from "react";

export const useCustomColumns = () => {
  const [customColumns, setCustomColumns] = useAtom(customColumnsAtom);

  const addColumn = useCallback(
    (datasetId: string, column: CustomColumn) => {
      setCustomColumns((prev) => {
        return {
          ...prev,
          [datasetId]: [...(prev[datasetId] ?? []), column],
        };
      });
    },
    [setCustomColumns],
  );

  const removeColumn = (datasetId: string, columnName: string) => {
    setCustomColumns((prev) => {
      if (!prev[datasetId]) return prev;

      const newColumns =
        prev[datasetId]?.filter((c) => c.name !== columnName) ?? [];
      return {
        ...prev,
        [datasetId]: newColumns,
      };
    });
  };

  const updateColumnWidth = useCallback(
    (datasetId: string, column: CustomColumn) => {
      setCustomColumns((prev) => {
        if (!prev[datasetId]) return prev;

        return {
          ...prev,
          [datasetId]: prev[datasetId]?.map((c) =>
            c.name === column.name ? column : c,
          ),
        };
      });
    },
    [setCustomColumns],
  );

  return {
    columns: customColumns,
    addColumn,
    updateColumnWidth,
    removeColumn,
  };
};

export const useRawDataColumn = () => {
  const [rawDataColumn, setRawDataColumn] = useAtom(rawDataColumnAtom);

  const toogleRawDataColumn = () => {
    setRawDataColumn((prev) => !prev);
  };

  return {
    rawDataColumn,
    toogleRawDataColumn,
  };
};

export const useRawDataColumnWidth = () => {
  const [rawDataColumnWidth, setRawDataColumnWidth] = useAtom(
    rawDataColumnWidthAtom,
  );

  const updateRawDataColumnWidth = useCallback(
    (width: number) => {
      setRawDataColumnWidth(width);
    },
    [setRawDataColumnWidth],
  );

  return {
    rawDataColumnWidth,
    updateRawDataColumnWidth,
  };
};
