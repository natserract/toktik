// Get
import { z } from 'zod'

export const GetAllResponseSchema = z.object({
  data: z.object({
    tags: z.array(z.string()),
  }),
})
export type GetAllResponse = z.infer<typeof GetAllResponseSchema>
