import './style.css'
import 'vue-loading-skeleton/dist/style.css'
import { defineCustomElement } from 'vue'

// Components
import SocialWall from '@/components/social-wall/SocialWall.vue'
import SocialWallSearch from '@/components/social-wall/SocialWallSearch.vue'
import SocialWallVideos from '@/components/social-wall/SocialWallVideos.vue'
import SocialWallTags from '@/components/social-wall/SocialWallTags.vue'

// Low level api
import { useSocialWallContext } from './components/social-wall/context'

// Define custom elements
const SocialWallElement = defineCustomElement(SocialWall)
const SocialWallSearchElement = defineCustomElement(SocialWallSearch)
const SocialWallVideosElement = defineCustomElement(SocialWallVideos)
const SocialWallTagsElement = defineCustomElement(SocialWallTags)

// Register web component
customElements.define('social-wall', SocialWallElement)
customElements.define('social-wall-search', SocialWallSearchElement)
customElements.define('social-wall-videos', SocialWallVideosElement)
customElements.define('social-wall-tags', SocialWallTagsElement)

export { SocialWall, SocialWallSearch, SocialWallVideos, SocialWallTags, useSocialWallContext }
