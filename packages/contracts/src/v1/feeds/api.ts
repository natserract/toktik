import { z } from 'zod'
import { FeedsAttributesSchema, FeedsAttributes } from '../common/feeds'

// Get
export const GetResponseSchema = z.object({
  data: FeedsAttributesSchema,
})
export type GetResponse = z.infer<typeof GetResponseSchema>

export const GetRequestSchema = z.object({
  id: z.string(),
})
export type GetRequest = z.infer<typeof GetRequestSchema>

// Search
export const SearchResponseSchema = z.object({
  data: z.array(FeedsAttributesSchema),
})
export type SearchResponse = z.infer<typeof SearchResponseSchema>

export const SearchRequestSchema = z.object({
  keywords: z.string(),
  count: z.string(),
})
export type SearchRequest = z.infer<typeof SearchRequestSchema>
