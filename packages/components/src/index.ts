import { defineCustomElement } from 'vue'
import SearchInput from '@/components/search/SearchInput.vue'

// Register web component
const SearchInputElement = defineCustomElement(SearchInput)
customElements.define('search-input', SearchInputElement)

export { default as SearchInput } from '@/components/search/SearchInput.vue'
