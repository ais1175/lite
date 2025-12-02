import { useState } from "react";

import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import dayjs from "dayjs";
import timezone from "dayjs/plugin/timezone";
import utc from "dayjs/plugin/utc";
import { TIME_OPTIONS, TimeDuration, TimeOption } from "../../utils/constants";
import { useSearchParams } from "react-router";

dayjs.extend(utc);
dayjs.extend(timezone);

interface QuickRangePickerProps {
  onSelect: () => void;
}

export function QuickRangePicker(props: QuickRangePickerProps) {
  const { onSelect } = props;
  const [duration, setDuration] = useState<TimeOption | null>(null);

  const [searchParams, setSearchParams] = useSearchParams();

  const handleChangeDuration = ({ type, duration, label, id }: TimeOption) => {
    onSelect();

    const params = new URLSearchParams(searchParams);
    params.delete("from");
    params.delete("to");

    switch (type) {
      case TimeDuration.Minutes: {
        params.set("qr", id);
        setDuration({ type, duration, label, id });

        setSearchParams(params);
        break;
      }
      case TimeDuration.Hours: {
        params.set("qr", id);
        setDuration({ type, duration, label, id });

        setSearchParams(params);
        break;
      }
      case TimeDuration.Days: {
        params.set("qr", id);
        setDuration({ type, duration, label, id });

        setSearchParams(params);
        break;
      }
    }
  };

  return (
    <div className="grid grid-cols-3 gap-2 p-2">
      {TIME_OPTIONS.map((time) => {
        const isActive = duration?.id === time.id;

        return (
          <Button
            size="sm"
            variant="outline"
            key={time.label}
            className={cn(isActive && "bg-accent", "text-xs")}
            onClick={() => handleChangeDuration(time)}
          >
            {time.label}
          </Button>
        );
      })}
    </div>
  );
}
