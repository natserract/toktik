import './style.css'
import { defineCustomElement } from 'vue'

// Components
import SocialWall from '@/components/social-wall/SocialWall.vue'
import SocialWallSearch from '@/components/social-wall/SocialWallSearch.vue'
import SocialWallResults from '@/components/social-wall/SocialWallResults.vue'
import SocialWallTags from '@/components/social-wall/SocialWallTags.vue'

// Define custom elements
const SocialWallElement = defineCustomElement(SocialWall)
const SocialWallSearchElement = defineCustomElement(SocialWallSearch)
const SocialWallResultsElement = defineCustomElement(SocialWallResults)
const SocialWallTagsElement = defineCustomElement(SocialWallTags)

// Register web component
customElements.define('social-wall', SocialWallElement)
customElements.define('social-wall-search', SocialWallSearchElement)
customElements.define('social-wall-results', SocialWallResultsElement)
customElements.define('social-wall-tags', SocialWallTagsElement)

export { SocialWall, SocialWallSearch, SocialWallResults, SocialWallTags }
