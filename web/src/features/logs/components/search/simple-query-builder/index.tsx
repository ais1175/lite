import { Input } from "@/components/ui/input";
import { useDebounce } from "@/hooks/use-debounce";
import { useCallback, useEffect, useState } from "react";
import { useSearchParams } from "react-router";

export function SimpleQueryBuilder() {
  const [query, setQuery] = useState("");
  const debouncedQuery = useDebounce(query, 300);

  const [searchParams, setSearchParams] = useSearchParams();
  function handleQueryChange(event: React.ChangeEvent<HTMLInputElement>) {
    const newQuery = event.currentTarget.value;
    setQuery(newQuery);
  }

  const handleSetFilter = useCallback(
    (query: string) => {
      const params = new URLSearchParams(searchParams.toString());
      if (query.trim()) {
        params.set("filter", query);
      } else {
        params.delete("filter");
      }

      setSearchParams(params);
    },
    [searchParams, setSearchParams],
  );

  useEffect(() => {
    handleSetFilter(debouncedQuery);
  }, [debouncedQuery, handleSetFilter]);

  return (
    <div>
      <Input
        placeholder="Search for logs..."
        value={query}
        onChange={handleQueryChange}
      />
    </div>
  );
}
