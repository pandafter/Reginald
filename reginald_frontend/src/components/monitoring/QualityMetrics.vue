<template>
  <div class="grid grid-cols-1 xl:grid-cols-2 gap-4 md:gap-8">
    <!-- Reliability -->
    <div class="glass p-6 md:p-10 rounded-[2.5rem] space-y-6">
        <div class="flex items-center justify-between">
            <h3 class="text-lg md:text-xl font-black dark:text-white flex items-center gap-2">
                <ShieldCheck class="text-emerald-500" />
                Confiabilidad
            </h3>
            <span class="text-emerald-600 dark:text-emerald-400 font-black bg-emerald-50 dark:bg-emerald-900/30 px-3 py-1 rounded-lg text-[10px] md:text-xs uppercase tracking-widest">{{ labels.health }}</span>
        </div>
        <div class="space-y-5">
            <div v-for="(val, key) in reliabilityMetrics" :key="key" class="relative group cursor-help">
                <MetricTooltip :title="val.title" :iso="val.iso" :description="val.desc" />
                <div class="flex justify-between text-[11px] md:text-xs font-bold uppercase tracking-widest mb-2">
                    <span class="text-slate-500 dark:text-slate-400">{{ val.label }}</span>
                    <span class="dark:text-white">{{ val.value.toFixed(1) }}%</span>
                </div>
                <div class="w-full bg-slate-100 dark:bg-slate-800 h-2 rounded-full overflow-hidden shadow-inner">
                    <div class="bg-emerald-500 h-full rounded-full transition-all duration-1000 ease-out" :style="{ width: val.value + '%' }"></div>
                </div>
            </div>
        </div>
    </div>

    <!-- Performance -->
    <div class="glass p-6 md:p-10 rounded-[2.5rem] space-y-6">
        <div class="flex items-center justify-between">
            <h3 class="text-lg md:text-xl font-black dark:text-white flex items-center gap-2">
                <Zap class="text-amber-500" />
                Desempeño
            </h3>
            <span class="text-amber-600 dark:text-amber-400 font-black bg-amber-50 dark:bg-amber-900/30 px-3 py-1 rounded-lg text-[10px] md:text-xs uppercase tracking-widest">{{ labels.performance }}</span>
        </div>
        <div class="space-y-5">
            <div v-for="(val, key) in performanceMetrics" :key="key" class="relative group cursor-help">
                <MetricTooltip :title="val.title" :iso="val.iso" :description="val.desc" />
                <div class="flex justify-between text-[11px] md:text-xs font-bold uppercase tracking-widest mb-2">
                    <span class="text-slate-500 dark:text-slate-400">{{ val.label }}</span>
                    <span class="dark:text-white">{{ val.display || val.value.toFixed(1) + '%' }}</span>
                </div>
                <div class="w-full bg-slate-100 dark:bg-slate-800 h-2 rounded-full overflow-hidden shadow-inner">
                    <div class="bg-amber-500 h-full rounded-full transition-all duration-1000 ease-out" :style="{ width: val.percent + '%' }"></div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { ShieldCheck, Zap } from 'lucide-vue-next';
import MetricTooltip from '../MetricTooltip.vue';

const props = defineProps({
  metrics: { type: Object, required: true }
});

const reliabilityMetrics = computed(() => ({
    avail: { label: 'Disponibilidad', title: 'Disponibilidad', value: props.metrics.availability || 100, iso: 'ISO/IEC 25010 - Fiabilidad', desc: 'Fiabilidad del servicio.' },
    matu: { label: 'Madurez', title: 'Madurez', value: props.metrics.reliability?.maturity || 0, iso: 'ISO/IEC 25010 - Fiabilidad', desc: 'Estabilidad del software.' },
    fault: { label: 'Tolerancia', title: 'Tolerancia a Fallos', value: props.metrics.reliability?.fault_tolerance || 0, iso: 'ISO/IEC 25010 - Fiabilidad', desc: 'Capacidad ante errores.' }
}));

const performanceMetrics = computed(() => ({
    lat: { label: 'Latencia', title: 'Latencia Media', value: props.metrics.avg_latency || 0, percent: Math.min((props.metrics.avg_latency || 0) / 10, 100), display: (props.metrics.avg_latency || 0).toFixed(1) + 'ms', iso: 'ISO/IEC 25010 - Eficiencia', desc: 'Comportamiento temporal.' },
    res: { label: 'Recursos', title: 'Uso de Memoria', value: props.metrics.performance?.resource_utilization || 0, percent: props.metrics.performance?.resource_utilization || 0, iso: 'ISO/IEC 25010 - Eficiencia', desc: 'Eficiencia de hardware.' },
    cap: { label: 'Capacidad', title: 'Capacidad', value: props.metrics.performance?.capacity || 0, percent: props.metrics.performance?.capacity || 0, iso: 'ISO/IEC 25010 - Eficiencia', desc: 'Saturación del sistema.' }
}));

const labels = computed(() => ({
  health: props.metrics.availability >= 99 ? 'Optimized' : props.metrics.availability >= 95 ? 'Stable' : 'Critical',
  performance: props.metrics.avg_latency <= 50 ? 'Peak' : props.metrics.avg_latency <= 150 ? 'Steady' : 'High Load'
}));
</script>
