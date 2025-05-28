import {
  isPositiveInteger,
  addPackSize,
  removePackSize,
  formatPacks
} from 'src/helpers/ShipmentHelpers'

import { describe, it, expect } from 'vitest'

describe('isPositiveInteger', () => {
  it('returns true for valid positive integers', () => {
    expect(isPositiveInteger(1)).toBe(true)
    expect(isPositiveInteger(100)).toBe(true)
  })

  it('returns false for zero or negative or non-integer values', () => {
    expect(isPositiveInteger(0)).toBe(false)
    expect(isPositiveInteger(-1)).toBe(false)
    expect(isPositiveInteger(2.5)).toBe(false)
    expect(isPositiveInteger('10')).toBe(false)
    expect(isPositiveInteger(null)).toBe(false)
  })
})

describe('addPackSize', () => {
  it('adds a new 0 to the end of pack sizes', () => {
    expect(addPackSize([250, 500])).toEqual([250, 500, 0])
  })
})

describe('removePackSize', () => {
  it('removes a pack size by index', () => {
    expect(removePackSize([250, 500, 1000], 1)).toEqual([250, 1000])
  })
})

describe('formatPacks', () => {
  it('formats packsUsed map into array for table', () => {
    const input = { 250: 2, 500: 1 }
    const expected = [
      { pack: '250', quantity: 2 },
      { pack: '500', quantity: 1 }
    ]
    expect(formatPacks(input)).toEqual(expected)
  })

  it('returns empty array when input is empty', () => {
    expect(formatPacks({})).toEqual([])
  })
})
