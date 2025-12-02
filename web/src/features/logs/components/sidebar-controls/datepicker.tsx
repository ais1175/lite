import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { Label } from "@/components/ui/label";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { format } from "date-fns";
import { useState } from "react";
import { Timepicker } from "./timepicker";

interface DatePickerProps {
  label: string;
  date: Date | undefined;
  onChange: (date: Date | undefined) => void;
}

export function DatePicker(props: DatePickerProps) {
  const { date, onChange } = props;

  const [open, setOpen] = useState(false);
  const [content, setContent] = useState<"time" | "date">("date");
  const [time, setTime] = useState<Date | undefined>(date);

  function onDateChange(date: Date | undefined) {
    if (date) {
      setTime(date);
    }
  }

  function onTimeChange(date: Date) {
    setTime(date);
  }

  function switchContent() {
    setContent(content === "time" ? "date" : "time");
  }

  function handleDateChange() {
    setOpen(false);
    onChange(time);
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger className="w-full p-2 space-y-1">
        <Label className="text-xs text-muted-foreground">{props.label}</Label>
        <Button
          variant="outline"
          className="w-full font-normal text-xs"
          size="sm"
        >
          {format(time ?? new Date(), "dd MMM yyyy, HH:mm")}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0" align="start">
        <div>
          {content === "time" ? (
            <Timepicker onChange={onTimeChange} selected={time!} />
          ) : (
            <Calendar
              mode="single"
              selected={time}
              onSelect={(date) => onDateChange(date)}
            />
          )}
          <div className="border-t flex items-center justify-between p-2">
            <Button
              variant="outline"
              size="sm"
              className="font-normal text-xs h-6 px-2"
              onClick={switchContent}
            >
              {content === "time" ? "Select Date" : "Select Time"}
            </Button>
            <Button
              size="sm"
              className="font-normal text-xs h-6 px-2"
              onClick={handleDateChange}
            >
              Ok
            </Button>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  );
}
