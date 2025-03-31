import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/Input";

export function NewOrganizationRoute() {
  return (
    <div className="sm:mx-auto sm:max-w-2xl h-screen flex flex-col justify-center space-y-8">
      <h3 className="text-tremor-title font-semibold text-tremor-content-strong dark:text-dark-tremor-content-strong">
        Create organization
      </h3>
      <form action="#" method="POST" className="mt-8">
        <div className="space-y-6">
          <div>
            <label
              htmlFor="workspace-name"
              className="text-tremor-default font-medium text-tremor-content-strong dark:text-dark-tremor-content-strong"
            >
              Name<span className="text-red-500">*</span>
            </label>
            <Input
              type="text"
              id="workspace-name"
              name="workspace-name"
              placeholder="My organization"
              className="mt-2"
              required
            />
          </div>
        </div>
        <div className="flex items-center justify-end space-x-4 mt-4">
          <Button size="sm" type="submit">
            Create organization
          </Button>
        </div>
      </form>
    </div>
  );
}
