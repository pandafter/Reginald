<template>
  <div class="glass p-6 md:p-10 rounded-[2.5rem] space-y-6">
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div class="flex items-center gap-3">
        <h3 class="text-lg md:text-xl font-black dark:text-white flex items-center gap-2">
          <Zap class="text-amber-500" />
          Rendimiento (24h)
        </h3>
      </div>
      <div class="flex items-center gap-4">
        <div class="flex items-center gap-2">
          <span class="w-2.5 h-2.5 bg-emerald-500 rounded-full"></span>
          <span class="text-slate-500 dark:text-slate-400 text-[9px] font-black uppercase tracking-widest">Normal</span>
        </div>
        <div class="flex items-center gap-2">
          <span class="w-2.5 h-2.5 bg-amber-400 rounded-full animate-pulse"></span>
          <span class="text-slate-500 dark:text-slate-400 text-[9px] font-black uppercase tracking-widest">Real Time</span>
        </div>
      </div>
    </div>
    <div class="h-[240px] md:h-[300px] w-full relative">
       <canvas ref="chartRef"></canvas>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { Zap } from 'lucide-vue-next';
import Chart from 'chart.js/auto';

const props = defineProps({
  history: {
    type: Array,
    default: () => []
  }
});

const chartRef = ref(null);
let chartInstance = null;

const initChart = () => {
    if (!chartRef.value) return;
    const ctx = chartRef.value.getContext('2d');
    const isDarkMode = document.documentElement.classList.contains('dark');
    
    const sortedHistory = [...props.history].sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp));
    const labels = sortedHistory.map(p => {
        const d = new Date(p.timestamp);
        return `${d.getHours()}:00`;
    });
    const data = sortedHistory.map(p => p.latency);

    if (chartInstance) chartInstance.destroy();

    chartInstance = new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels.length ? labels : ['24h'],
            datasets: [{
                label: 'Responsiveness (ms)',
                data: data.length ? data : [0],
                borderColor: isDarkMode ? '#38bdf8' : '#3b82f6',
                backgroundColor: isDarkMode ? 'rgba(56, 189, 248, 0.05)' : 'rgba(59, 130, 246, 0.05)',
                borderWidth: 3,
                fill: true,
                tension: 0.4,
                pointRadius: 4,
                pointBackgroundColor: isDarkMode ? '#0f172a' : '#fff',
                pointBorderWidth: 2
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                legend: { display: false },
                tooltip: {
                    backgroundColor: isDarkMode ? '#1e293b' : '#111827',
                    titleFont: { size: 10, weight: '900' },
                    bodyFont: { size: 12, weight: '700' },
                    padding: 12,
                    cornerRadius: 12
                }
            },
            scales: {
                y: { 
                    beginAtZero: true, 
                    grid: { color: isDarkMode ? 'rgba(255,255,255,0.03)' : 'rgba(0,0,0,0.03)' },
                    ticks: { color: '#64748b', font: { size: 10, weight: 'bold' } }
                },
                x: { 
                    grid: { display: false },
                    ticks: { color: '#64748b', font: { size: 10, weight: 'bold' } }
                }
            }
        }
    });
};

// Re-init chart when theme changes or history updates
watch(() => props.history, () => initChart(), { deep: true });

onMounted(() => {
    initChart();
    // Watch for system theme changes via observation in parent or simple interval/mutation
});

onUnmounted(() => {
    if (chartInstance) {
        chartInstance.destroy();
        chartInstance = null;
    }
});
</script>
