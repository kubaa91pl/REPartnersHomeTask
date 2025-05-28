
export function isPositiveInteger(val) {
  return Number.isInteger(val) && val > 0
}

export function addPackSize(sizes) {
  return [...sizes, 0]
}

export function removePackSize(sizes, index) {
  return sizes.filter((_, i) => i !== index)
}

export function formatPacks(packsUsed) {
  return Object.entries(packsUsed).map(([pack, quantity]) => ({
    pack,
    quantity
  }))
}