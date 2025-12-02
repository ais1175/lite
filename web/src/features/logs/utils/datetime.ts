import { subDays, subHours, subMinutes } from "date-fns";
import { TIME_OPTIONS, TimeDuration } from "./constants";

export function getDatetimeQuery(
  quickRange: string | null,
  fromDate: string | null,
  toDate: string | null,
): { from: string; to: string } | undefined {
  if (quickRange) {
    const timeOption = TIME_OPTIONS.find((option) => option.id === quickRange);
    if (timeOption) {
      const fromDate = new Date();
      const toDate = new Date();

      switch (timeOption.type) {
        case TimeDuration.Minutes: {
          const rangeMinutes = subMinutes(toDate, timeOption.duration);

          fromDate.setTime(rangeMinutes.getTime());
          toDate.setTime(toDate.getTime());

          break;
        }
        case TimeDuration.Hours: {
          const rangeHours = subHours(toDate, timeOption.duration);

          fromDate.setTime(rangeHours.getTime());
          toDate.setTime(toDate.getTime());

          break;
        }
        case TimeDuration.Days: {
          const rangeDays = subDays(toDate, timeOption.duration);

          fromDate.setTime(rangeDays.getTime());
          toDate.setTime(toDate.getTime());

          break;
        }
      }

      return { from: fromDate.toUTCString(), to: toDate.toUTCString() };
    }
  } else if (fromDate && toDate) {
    return { from: fromDate, to: toDate };
  }
}
