import { z } from "zod";

export const tokenSchema = z.object({
  identifier: z.string(),
  type: z.string(),
});

export type TokenParams = z.infer<typeof tokenSchema>;

export interface TokenPreview {
  token: string;
}

export interface Token {
  id: number;
  type: string;
  identifier: string;
}
