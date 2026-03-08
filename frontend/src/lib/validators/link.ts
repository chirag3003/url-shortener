import { z } from "zod/v4";

// ── Request Schemas ──────────────────────────────────────────────────────────

export const createLinkSchema = z.object({
  longUrl: z.url("Please enter a valid URL").max(2048, "URL is too long"),
  customAlias: z
    .string()
    .min(3, "Alias must be at least 3 characters")
    .max(64, "Alias must be at most 64 characters")
    .regex(/^[a-zA-Z0-9]+$/, "Alias must be alphanumeric")
    .optional()
    .or(z.literal("")),
  expiresAt: z.string().optional().or(z.literal("")),
  redirectType: z.union([z.literal(301), z.literal(302)]).optional(),
});

export const updateLinkSchema = z.object({
  longUrl: z
    .url("Please enter a valid URL")
    .max(2048, "URL is too long")
    .optional()
    .or(z.literal("")),
  expiresAt: z.string().optional().or(z.literal("")),
  redirectType: z.union([z.literal(301), z.literal(302)]).optional(),
  isActive: z.boolean().optional(),
});

export const listLinksQuerySchema = z.object({
  page: z.number().int().min(1).optional(),
  limit: z.number().int().min(1).max(100).optional(),
  search: z.string().max(128).optional(),
});

// ── Response Schemas ─────────────────────────────────────────────────────────

export const linkResponseSchema = z.object({
  id: z.string(),
  userId: z.string().optional(),
  longUrl: z.string(),
  shortCode: z.string(),
  shortUrl: z.string(),
  redirectType: z.number(),
  expiresAt: z.string().optional(),
  isActive: z.boolean(),
  createdAt: z.string(),
  updatedAt: z.string(),
});

export const paginatedLinksResponseSchema = z.object({
  items: z.array(linkResponseSchema),
  page: z.number(),
  limit: z.number(),
  total: z.number(),
});

export const aliasAvailabilityResponseSchema = z.object({
  alias: z.string(),
  available: z.boolean(),
});

// ── Types ────────────────────────────────────────────────────────────────────

export type CreateLinkInput = z.infer<typeof createLinkSchema>;
export type UpdateLinkInput = z.infer<typeof updateLinkSchema>;
export type ListLinksQuery = z.infer<typeof listLinksQuerySchema>;
export type LinkResponse = z.infer<typeof linkResponseSchema>;
export type PaginatedLinksResponse = z.infer<
  typeof paginatedLinksResponseSchema
>;
export type AliasAvailabilityResponse = z.infer<
  typeof aliasAvailabilityResponseSchema
>;
