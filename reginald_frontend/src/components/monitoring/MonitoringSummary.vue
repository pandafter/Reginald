<template>
  <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6">
    <div 
      v-for="card in summaryCards" 
      :key="card.label"
      class="glass p-6 md:p-8 rounded-3xl relative group cursor-help transition-all hover:scale-[1.02] border-l-[6px]"
      :class="card.borderColor"
    >
      <MetricTooltip :title="card.title" :iso="card.iso" :description="card.description" />
      <p class="text-slate-500 dark:text-slate-400 text-[10px] font-black uppercase tracking-widest mb-1">{{ card.label }}</p>
      <p class="text-3xl md:text-4xl font-black tracking-tighter" :class="card.textColor || 'text-slate-800 dark:text-white'">
        {{ card.value }}<span v-if="card.unit" class="text-lg ml-0.5">{{ card.unit }}</span>
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import MetricTooltip from '../MetricTooltip.vue';

const props = defineProps({
  metrics: {
    type: Object,
    required: true
  }
});

const summaryCards = computed(() => [
  {
    label: 'Usuarios',
    title: 'Usuarios Totales',
    iso: 'Usabilidad',
    description: 'Alcance y adopción del aplicativo.',
    value: props.metrics.total_users || 0,
    borderColor: 'border-l-primary-500'
  },
  {
    label: 'Tareas',
    title: 'Tareas',
    iso: 'Funcionalidad',
    description: 'Volumen de trabajo gestionado.',
    value: props.metrics.total_tasks || 0,
    borderColor: 'border-l-indigo-500'
  },
  {
    label: 'Solicitudes',
    title: 'Solicitudes',
    iso: 'Funcionalidad',
    description: 'Demanda de servicios procesados.',
    value: props.metrics.total_requests || 0,
    borderColor: 'border-l-amber-500'
  },
  {
    label: 'Availability',
    title: 'Disponibilidad',
    iso: 'Fiabilidad',
    description: 'Tiempo operativo del sistema.',
    value: (props.metrics.availability || 100).toFixed(1),
    unit: '%',
    borderColor: 'border-l-emerald-500',
    textColor: 'text-emerald-600 dark:text-emerald-400'
  }
]);
</script>
