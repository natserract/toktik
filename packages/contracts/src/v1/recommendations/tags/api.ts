// Get
import { z } from 'zod'

export const GetResponseSchema = z.object({
  data: z.object({
    tags: z.array(z.string()),
  }),
})
export type GetResponse = z.infer<typeof GetResponseSchema>
