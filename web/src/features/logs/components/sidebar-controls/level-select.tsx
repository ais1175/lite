import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { CheckIcon } from "lucide-react";
import { useState } from "react";
import { useSearchParams } from "react-router";

interface LevelSelectProps {
  levels: string[] | null | undefined;
}

export function LevelSelect({ levels }: LevelSelectProps) {
  const [open, setOpen] = useState(false);

  const [searchParams, setSearchParams] = useSearchParams();

  const hasLevel = searchParams?.has("level");
  const urlLevel = searchParams?.get("level");

  const [selectedValues, setSelectedValues] = useState<string[]>(
    hasLevel && urlLevel ? urlLevel.split(",") : [],
  );

  const handleItemSelect = (value: string) => {
    const params = new URLSearchParams(searchParams);
    const isSelect = selectedValues.includes(value);

    if (isSelect) {
      setSelectedValues((prev) => [...prev.filter((v) => v !== value)]);

      const cpLevels = [...selectedValues];
      const newLevels = cpLevels.filter((v) => v !== value);

      if (newLevels.length === 0) {
        params.delete("level");
      } else {
        params.set("level", newLevels.join(","));
      }
    } else {
      setSelectedValues((prev) => [...prev, value]);

      const cpLevels = [...selectedValues];
      const newLevels = [...cpLevels, value];

      params.set("level", newLevels.join(","));
    }

    setSearchParams(params);
  };

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant="outline" className="justify-start">
          {selectedValues ? (
            <div className="flex space-x-2">
              <>+ Set level</>
              <div className="ml-2 flex space-x-2">
                {selectedValues.map((value) => (
                  <Badge variant="secondary" key={value}>
                    {value}
                  </Badge>
                ))}
              </div>
            </div>
          ) : (
            <>+ Set level</>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="p-0" side="bottom" align="start">
        <Command>
          <CommandInput placeholder="Change level..." />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup>
              {levels?.map((level: string) => {
                const isSelected = selectedValues.includes(level);

                return (
                  <CommandItem
                    key={level}
                    onSelect={(value) => {
                      handleItemSelect(value);
                    }}
                  >
                    <div
                      className={cn(
                        "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
                        isSelected
                          ? "bg-primary text-primary-foreground"
                          : "opacity-50 [&_svg]:invisible",
                      )}
                    >
                      <CheckIcon className={cn("h-4 w-4")} />
                    </div>
                    <span>{level}</span>
                  </CommandItem>
                );
              })}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
