// Get
import { z } from 'zod'

export const GetAllResponseSchema = z.object({
  data: z.object({
    keywords: z.array(z.string()),
  }),
})
export type GetAllResponse = z.infer<typeof GetAllResponseSchema>
