export default function debounce(fn: (...args: any[]) => void, delay: number) {
  let timeoutId: number | undefined
  return function (...args: any[]) {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }
    timeoutId = setTimeout(() => {
      fn(...args)
    }, delay)
  }
}
