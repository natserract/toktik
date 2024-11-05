// Get
import { z } from 'zod'

export const GetResponseSchema = z.object({
  data: z.object({
    keywords: z.array(z.string()),
  }),
})
export type GetResponse = z.infer<typeof GetResponseSchema>
