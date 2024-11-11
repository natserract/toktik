import { reactive, provide, inject } from 'vue'
import { FeedsAttributes } from '@toktik/contracts'

type FeedsState = {
  loading: boolean
  error: null
  data: FeedsAttributes[]
}
type TagsState = {
  loading: boolean
  error: null
  data: string[]
}

interface SocialWallState {
  baseUrl: string
  feeds: FeedsState
  tags: TagsState
}

const SocialWallContextSymbol = Symbol('SocialWallContext')

export function provideSocialWallContext(baseUrl: string) {
  const state = reactive<SocialWallState>({
    baseUrl,
    feeds: {
      loading: true,
      error: null,
      data: [],
    },
    tags: {
      loading: true,
      error: null,
      data: [],
    },
  })

  const getFeeds = () => state.feeds.data
  const setFeeds = (feeds: FeedsAttributes[]) => {
    state.feeds.data = feeds
  }

  const getTags = () => state.tags.data
  const setTags = (tags: string[]) => {
    state.tags.data = tags
  }

  provide(SocialWallContextSymbol, {
    state,
    getFeeds,
    setFeeds,
    getTags,
    setTags,
  })
}

interface SocialWallContext {
  state: SocialWallState
  getFeeds: () => FeedsAttributes[]
  setFeeds: (feeds: FeedsAttributes[]) => void
  getTags: () => string[]
  setTags: (tags: string[]) => void
}

export function useSocialWallContext(): SocialWallContext {
  const context = inject<SocialWallContext>(SocialWallContextSymbol)
  if (!context) {
    throw new Error('useGlobalContext must be used within a provider')
  }
  return context
}
