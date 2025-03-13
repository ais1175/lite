import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTrigger,
} from "@/components/ui/Dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/Form";
import { type TokenParams, tokenSchema } from "@/typings/token";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useCreateToken } from "../api/useCreateToken";
import { useState, useTransition } from "react";
import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/Input";
import { KeyRound } from "lucide-react";
import { DialogTitle } from "@radix-ui/react-dialog";

export function CreateTokenDialog() {
  const [open, setOpen] = useState(false);
  const [isPending, startTransition] = useTransition();

  const { mutateAsync, reset, data, isSuccess } = useCreateToken();
  const form = useForm<TokenParams>({
    defaultValues: {
      type: "media",
    },
    resolver: zodResolver(tokenSchema),
  });

  function handleOnSubmit(data: TokenParams) {
    startTransition(async () => {
      await mutateAsync(data);
    });
  }

  function resetForm() {
    form.reset();
    reset();
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button size="sm">
          <KeyRound size={16} className="mr-2" />
          Create token
        </Button>
      </DialogTrigger>
      <DialogContent
        id="create-token-dialog"
        onAnimationEnd={resetForm}
        aria-description="Create a new token"
      >
        <DialogHeader>
          <DialogTitle>Create token</DialogTitle>
        </DialogHeader>
        {isSuccess && data ? (
          <div className="border p-2">{data.token}</div>
        ) : (
          <Form {...form}>
            <form onSubmit={form.handleSubmit(handleOnSubmit)}>
              <div>
                <FormField
                  control={form.control}
                  name="identifier"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Identifier</FormLabel>
                      <FormControl>
                        <Input {...field} type="text" />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </div>

              <DialogFooter className="mt-4">
                <Button type="submit" disabled={isPending}>
                  Create
                </Button>
              </DialogFooter>
            </form>
          </Form>
        )}
      </DialogContent>
    </Dialog>
  );
}
