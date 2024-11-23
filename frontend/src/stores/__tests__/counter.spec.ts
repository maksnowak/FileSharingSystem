import { expect, describe, test } from 'vitest'

import { useCounterStore } from '../counter'
import { setActivePinia, createPinia } from 'pinia'
import { beforeEach } from 'vitest'

describe('Counter store tests', () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    })
    
    test('increments count', () => {
        const store = useCounterStore()
        store.increment()
        expect(store.count).toBe(1)
    })

    test('doubles count', () => {
        const store = useCounterStore()
        store.increment()
        expect(store.doubleCount).toBe(2)
    })

    test('initial count is zero', () => {
        const store = useCounterStore()
        expect(store.count).toBe(0)
    })

    test('double count is zero initially', () => {
        const store = useCounterStore()
        expect(store.doubleCount).toBe(0)
    })

    test('increments count multiple times', () => {
        const store = useCounterStore()
        store.increment()
        store.increment()
        store.increment()
        expect(store.count).toBe(3)
    })

    test('double count after multiple increments', () => {
        const store = useCounterStore()
        store.increment()
        store.increment()
        expect(store.doubleCount).toBe(4)
    })
})