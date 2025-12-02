import { Button } from "@/components/ui/button";
import { addDays } from "date-fns";
import { useState } from "react";
import { type DateRange } from "react-day-picker";
import { DatePicker } from "./datepicker";
import { useSearchParams } from "react-router";

interface LogsDatetimePickerProps {
  onApply: () => void;
}

export function LogsDatetimePicker(props: LogsDatetimePickerProps) {
  const { onApply } = props;
  const [searchParams, setSearchParams] = useSearchParams();

  const hasRange = searchParams?.has("from") && searchParams?.has("to");

  const [selectedDays, setSelectedDays] = useState<DateRange | undefined>(
    hasRange
      ? {
          from: new Date(searchParams?.get("from") ?? Date.now()),
          to: new Date(searchParams?.get("to") ?? addDays(Date.now(), 1)),
        }
      : {
          from: new Date(Date.now()),
          to: new Date(addDays(Date.now(), 1)),
        },
  );

  const handleRangeChange = () => {
    onApply();
    const params = new URLSearchParams(searchParams);

    params.delete("qr");

    if (selectedDays?.from && selectedDays?.to) {
      params.set("from", selectedDays.from.toUTCString());
      params.set("to", selectedDays.to.toUTCString());
    } else {
      params.delete("from");
      params.delete("to");
    }

    setSearchParams(params);
  };

  function handleStartTimeChange(date: Date | undefined) {
    setSelectedDays((prev) => {
      if (!prev) return;

      return {
        ...prev,
        from: date,
      };
    });
  }

  function handleEndTimeChange(date: Date | undefined) {
    setSelectedDays((prev) => {
      if (!prev) return;

      return {
        ...prev,
        to: date,
      };
    });
  }

  return (
    <div>
      <DatePicker
        date={selectedDays?.from}
        label="START TIME"
        onChange={handleStartTimeChange}
      />
      <DatePicker
        date={selectedDays?.to}
        label="END TIME"
        onChange={handleEndTimeChange}
      />

      <div className="flex justify-end p-2">
        <Button size="sm" onClick={handleRangeChange}>
          Apply
        </Button>
      </div>
    </div>
  );
}
