import { z } from "zod/v4";

export const userResponseSchema = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string(),
  phoneNo: z.string().optional(),
});

export const loginResponseSchema = z.object({
  token: z.string(),
  user: userResponseSchema,
});

export type UserResponse = z.infer<typeof userResponseSchema>;
export type LoginResponse = z.infer<typeof loginResponseSchema>;
