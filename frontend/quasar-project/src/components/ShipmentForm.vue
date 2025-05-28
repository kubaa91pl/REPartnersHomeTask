<template>
  <div class="q-pa-md">
    <div class="q-gutter-md" style="max-width: 400px">
      <h2>Order Packs</h2>

      <q-input
        outlined
        v-model.number="items"
        label="Number of Items"
        type="number"
        placeholder="e.g. 250"
        :rules="[positiveIntegerRule]"
      />

      <q-toggle
        v-model="useDefaultPacks"
        label="Use default pack sizes"
      />

      <div v-if="!useDefaultPacks">
        <label>Custom Pack Sizes</label>
        <div v-for="(size, index) in packSizes" :key="index" class="q-mb-sm">
          <q-input
            outlined
            v-model.number="packSizes[index]"
            type="number"
            placeholder="e.g. 250"
            :rules="[positiveIntegerRule]"
          />
        </div>
        <q-btn flat color="primary" @click="addPackSize">Add Pack Size</q-btn>
      </div>

      <q-btn color="primary" @click="submit" :loading="loading" :disable="items <= 0">
        Submit
      </q-btn>


      <div v-if="result">
        <h3 class="q-mt-lg">Shipment Result</h3>
        <div><strong>ID:</strong> {{ result.id }}</div>

        <q-table
          :rows="formattedPacks"
          :columns="columns"
          row-key="pack"
          dense
          flat
        />
      </div>

      <q-banner v-if="error" class="q-mt-md" color="negative">
        {{ error }}
      </q-banner>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'

const positiveIntegerRule = val =>
  Number.isInteger(val) && val > 0 || 'Enter a positive whole number'

const items = ref(0)
const useDefaultPacks = ref(true)
const packSizes = ref([250, 500, 1000])
const result = ref(null)
const error = ref(null)
const loading = ref(false)

const columns = [
  { name: 'pack', label: 'Pack', field: 'pack', align: 'left' },
  { name: 'quantity', label: 'Quantity', field: 'quantity', align: 'right' }
]

const formattedPacks = computed(() => {
  if (!result.value) return []
  return Object.entries(result.value.packs_used).map(([pack, quantity]) => ({
    pack,
    quantity
  }))
})

function addPackSize() {
  packSizes.value.push(0)
}

async function submit() {
  error.value = null
  result.value = null
  loading.value = true

  try {
    const response = await axios.post('http://localhost:8080/shipment', {
      items: items.value,
      packs: useDefaultPacks.value ? [] : packSizes.value
    })
    result.value = response.data
  } catch (err) {
    error.value = err.response?.data || 'Failed to calculate shipment.'
  } finally {
    loading.value = false
  }
}
</script>
