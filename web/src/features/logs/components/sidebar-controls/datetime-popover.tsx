import { Button } from "@/components/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { format } from "date-fns";
import { CalendarIcon } from "lucide-react";
import { useMemo, useState } from "react";
import { TIME_OPTIONS } from "../../utils/constants";
import { useSearchParams } from "react-router";
import { QuickRangePicker } from "./quick-range-picker";
import { LogsDatetimePicker } from "./logs-datetime-picker";

export function DatetimePopover() {
  const [datePickerOpen, setDatePickerOpen] = useState(false);
  const [searchParams] = useSearchParams();

  const durationLabel = useMemo(() => {
    const params = new URLSearchParams(searchParams);
    const qr = params.get("qr");
    const from = params.get("from");
    const to = params.get("to");

    if (qr) {
      return TIME_OPTIONS.find((option) => option.id === qr)?.label;
    } else if (from && to) {
      const fromDate = new Date(from);
      const toDate = new Date(to);

      return `${format(fromDate, "PP")} - ${format(toDate, "PP")}`;
    }
  }, [searchParams]);

  function toggleDatePicker() {
    setDatePickerOpen((prev) => !prev);
  }

  return (
    <Popover open={datePickerOpen} onOpenChange={setDatePickerOpen}>
      <PopoverTrigger asChild className="w-full">
        <Button variant="outline">
          <span>
            <CalendarIcon className="mr-2 h-4 w-4" />
          </span>
          {durationLabel ? durationLabel : "Select date range"}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0" align="start">
        <QuickRangePicker onSelect={toggleDatePicker} />
        <LogsDatetimePicker onApply={toggleDatePicker} />
      </PopoverContent>
    </Popover>
  );
}
