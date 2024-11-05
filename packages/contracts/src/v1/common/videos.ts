import { z } from 'zod'

const MusicInfoSchema = z.object({
  id: z.string(),
  title: z.string(),
  play: z.string(),
  cover: z.string(),
  author: z.string(),
  original: z.boolean(),
  duration: z.number().optional(),
  album: z.string(),
})

const CommerceInfoSchema = z.object({
  adv_promotable: z.boolean(),
  auction_ad_invited: z.boolean(),
  branded_content_type: z.number(),
  with_comment_filter_words: z.boolean(),
})

const AuthorSchema = z.object({
  id: z.string(),
  unique_id: z.string(),
  nickname: z.string(),
  avatar: z.string(),
})

export const VideoAttributesSchema = z.object({
  id: z.string().optional(),
  aweme_id: z.string(),
  video_id: z.string(),
  region: z.string(),
  title: z.string(),
  cover: z.string(),
  ai_dynamic_cover: z.string(),
  origin_cover: z.string(),
  duration: z.number(),
  play: z.string(),
  wmplay: z.string(),
  size: z.bigint(),
  wm_size: z.bigint(),
  music: z.string(),
  music_info: MusicInfoSchema,
  play_count: z.number(),
  digg_count: z.number(),
  comment_count: z.number(),
  share_count: z.number(),
  download_count: z.number(),
  create_time: z.bigint(),
  anchors: z.unknown(),
  anchors_extras: z.string(),
  is_ad: z.boolean(),
  commerce_info: CommerceInfoSchema,
  commercial_video_info: z.string(),
  item_comment_settings: z.number(),
  mentioned_users: z.string(),
  author: AuthorSchema,
  is_top: z.number(),
})

export type VideoAttributes = z.infer<typeof VideoAttributesSchema>
