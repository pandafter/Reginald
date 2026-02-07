<template>
  <div class="flex-1 flex flex-col h-screen overflow-hidden">
    <Navbar />
    
    <main class="flex-1 overflow-y-auto p-4 md:p-8 lg:p-12 transition-colors duration-500">
      <div class="max-w-[1600px] mx-auto space-y-6 md:space-y-10">
        <header class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <h2 class="text-3xl md:text-4xl font-black text-slate-900 dark:text-white tracking-tight">Métricas del Sistema</h2>
            <p class="text-slate-500 dark:text-slate-400 mt-1 text-sm md:text-base font-bold">Control de calidad y seguridad (ISO/IEC 25010)</p>
          </div>
          
          <div class="flex items-center gap-4">
            <!-- Archivados Sunken Button -->
            <button 
              @click="showArchived = !showArchived"
              class="relative group transition-all duration-300 active:scale-95"
            >
              <div 
                class="flex items-center gap-3 px-6 py-3 rounded-2xl shadow-sunken bg-white dark:bg-slate-950 border border-slate-100 dark:border-slate-900/50 overflow-hidden"
              >
                <div class="relative">
                  <Archive class="w-5 h-5 transition-colors duration-500" :class="showArchived ? 'text-primary-500' : 'text-slate-500 dark:text-white'" />
                  <div v-if="archivedAlerts.length > 0" class="absolute -top-2 -right-2 bg-red-600 text-white text-[8px] font-black w-4 h-4 rounded-full flex items-center justify-center border-2 border-white dark:border-slate-950">
                    {{ archivedAlerts.length }}
                  </div>
                </div>
                <span class=" text-[10px] font-black uppercase tracking-widest dark:text-white transition-colors duration-500" :class="showArchived ? 'text-primary-500' : 'text-slate-500 dark:text-white'">
                  {{ showArchived ? 'Ver Activos' : 'Archivados' }}
                </span>
              </div>
            </button>

            <!-- Alerts Badge -->
            <div v-if="filteredAlerts.length > 0 && !showArchived" class="bg-rose-50 dark:bg-rose-900/20 border border-rose-100 dark:border-rose-900/30 px-4 py-2.5 rounded-2xl flex items-center gap-3 animate-pulse">
              <ShieldAlert class="text-rose-600 dark:text-rose-400 w-5 h-5" />
              <span class="text-rose-700 dark:text-rose-400 text-[10px] md:text-xs font-black uppercase tracking-widest">¡{{ filteredAlerts.length }} Alertas Activas!</span>
            </div>
          </div>
        </header>

        <div v-if="loading && !metrics.total_users" class="flex flex-col items-center justify-center py-20 gap-4">
          <Loader2 class="w-10 h-10 animate-spin text-primary-600" />
          <p class="text-slate-400 font-bold uppercase tracking-widest text-xs">Calculando indicadores...</p>
        </div>

        <div v-else class="space-y-6 md:space-y-10 pb-20">
          <template v-if="!showArchived">
            <!-- Summary Cards - 4 columns -->
            <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6">
              <div class="glass p-6 md:p-8 rounded-3xl border-l-[6px] border-l-primary-500 relative group cursor-help transition-all hover:scale-[1.02]">
                  <MetricTooltip title="Usuarios Totales" iso="Usabilidad" description="Alcance y adopción del aplicativo." />
                  <p class="text-slate-500 dark:text-slate-400 text-[10px] font-black uppercase tracking-widest mb-1">Usuarios</p>
                  <p class="text-3xl md:text-4xl font-black text-slate-800 dark:text-white tracking-tighter">{{ metrics.total_users || 0 }}</p>
              </div>
              <div class="glass p-6 md:p-8 rounded-3xl border-l-[6px] border-l-indigo-500 relative group cursor-help transition-all hover:scale-[1.02]">
                  <MetricTooltip title="Tareas" iso="Funcionalidad" description="Volumen de trabajo gestionado." />
                  <p class="text-slate-500 dark:text-slate-400 text-[10px] font-black uppercase tracking-widest mb-1">Tareas</p>
                  <p class="text-3xl md:text-4xl font-black text-slate-800 dark:text-white tracking-tighter">{{ metrics.total_tasks || 0 }}</p>
              </div>
              <div class="glass p-6 md:p-8 rounded-3xl border-l-[6px] border-l-amber-500 relative group cursor-help transition-all hover:scale-[1.02]">
                  <MetricTooltip title="Solicitudes" iso="Funcionalidad" description="Demanda de servicios procesados." />
                  <p class="text-slate-500 dark:text-slate-400 text-[10px] font-black uppercase tracking-widest mb-1">Solicitudes</p>
                  <p class="text-3xl md:text-4xl font-black text-slate-800 dark:text-white tracking-tighter">{{ metrics.total_requests || 0 }}</p>
              </div>
              <div class="glass p-6 md:p-8 rounded-3xl border-l-[6px] border-l-emerald-500 relative group cursor-help transition-all hover:scale-[1.02]">
                  <MetricTooltip title="Disponibilidad" iso="Fiabilidad" description="Tiempo operativo del sistema." />
                  <p class="text-slate-500 dark:text-slate-400 text-[10px] font-black uppercase tracking-widest mb-1">Availability</p>
                  <p class="text-3xl md:text-4xl font-black text-emerald-600 dark:text-emerald-400 tracking-tighter">{{ (metrics.availability || 100).toFixed(1) }}%</p>
              </div>
            </div>

            <!-- History Chart Section -->
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
                 <canvas ref="latencyChartCanvas"></canvas>
              </div>
            </div>

            <!-- Quality Metrics Grid (RESTORING THIS) -->
            <div class="grid grid-cols-1 xl:grid-cols-2 gap-4 md:gap-8">
              <!-- Reliability -->
              <div class="glass p-6 md:p-10 rounded-[2.5rem] space-y-6">
                  <div class="flex items-center justify-between">
                      <h3 class="text-lg md:text-xl font-black dark:text-white flex items-center gap-2">
                          <ShieldCheck class="text-emerald-500" />
                          Confiabilidad
                      </h3>
                      <span class="text-emerald-600 dark:text-emerald-400 font-black bg-emerald-50 dark:bg-emerald-900/30 px-3 py-1 rounded-lg text-[10px] md:text-xs uppercase tracking-widest">{{ getHealthLabel(metrics.availability) }}</span>
                  </div>
                  <div class="space-y-5">
                      <div v-for="(val, key) in reliabilityMetrics" :key="key" class="relative group cursor-help">
                          <MetricTooltip :title="val.title" :iso="'ISO/IEC 25010 - Fiabilidad'" :description="val.desc" />
                          <div class="flex justify-between text-[11px] md:text-xs font-bold uppercase tracking-widest mb-2">
                              <span class="text-slate-500 dark:text-slate-400">{{ val.label }}</span>
                              <span class="dark:text-white">{{ val.value.toFixed(1) }}%</span>
                          </div>
                          <div class="w-full bg-slate-100 dark:bg-slate-800 h-2 rounded-full overflow-hidden">
                              <div class="bg-emerald-500 h-full rounded-full transition-all duration-1000" :style="{ width: val.value + '%' }"></div>
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
                      <span class="text-amber-600 dark:text-amber-400 font-black bg-amber-50 dark:bg-amber-900/30 px-3 py-1 rounded-lg text-[10px] md:text-xs uppercase tracking-widest">{{ getPerformanceLabel(metrics.avg_latency) }}</span>
                  </div>
                  <div class="space-y-5">
                      <div v-for="(val, key) in performanceMetrics" :key="key" class="relative group cursor-help">
                          <MetricTooltip :title="val.title" :iso="'ISO/IEC 25010 - Eficiencia'" :description="val.desc" />
                          <div class="flex justify-between text-[11px] md:text-xs font-bold uppercase tracking-widest mb-2">
                              <span class="text-slate-500 dark:text-slate-400">{{ val.label }}</span>
                              <span class="dark:text-white">{{ val.display || val.value.toFixed(1) + '%' }}</span>
                          </div>
                          <div class="w-full bg-slate-100 dark:bg-slate-800 h-2 rounded-full overflow-hidden">
                              <div class="bg-amber-500 h-full rounded-full transition-all duration-1000" :style="{ width: val.percent + '%' }"></div>
                          </div>
                      </div>
                  </div>
              </div>
            </div>
          </template>

          <!-- Security Alerts Widget -->
          <div class="glass p-6 md:p-10 rounded-[2.5rem] border-2 transition-colors overflow-hidden" 
               :class="filteredAlerts.length > 0 ? 'border-rose-100 dark:border-rose-900/30' : 'border-white dark:border-slate-900'">
            <div class="flex flex-col sm:flex-row sm:items-center justify-between mb-8 gap-4">
              <div class="flex items-center gap-4">
                <div class="p-3 rounded-2xl shadow-inner transition-colors duration-500" :class="filteredAlerts.length > 0 ? 'bg-rose-100 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400' : 'bg-slate-50 dark:bg-slate-900 text-slate-400'">
                  <ShieldAlert v-if="filteredAlerts.length > 0" class="w-6 h-6" :class="{'animate-pulse': !showArchived}" />
                  <Archive v-else-if="showArchived" class="w-6 h-6" />
                  <ShieldCheck v-else class="w-6 h-6" />
                </div>
                <div>
                  <h3 class="text-lg md:text-xl font-black dark:text-white tracking-widest uppercase">
                    {{ showArchived ? 'Historial de Amenazas' : 'Seguridad & Firewall' }}
                  </h3>
                  <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest">
                    {{ showArchived ? 'Registros de incidentes procesados y mitigados' : (filteredAlerts.length > 0 ? 'Amenazas bloqueadas en tiempo real' : 'Perímetro de red asegurado') }}
                  </p>
                </div>
              </div>
              <div class="flex gap-2">
                 <span v-if="filteredAlerts.length > 0" class="w-fit px-4 py-2" :class="showArchived ? 'bg-slate-200 dark:bg-slate-800 text-slate-600 dark:text-slate-400 rounded-xl text-[10px] font-black uppercase tracking-widest' : 'bg-rose-600 text-white rounded-xl text-[10px] font-black uppercase tracking-widest'">
                    {{ showArchived ? 'Estado: Archivado' : 'Estado: Alerta Crítica' }}
                 </span>
                 <span v-else class="w-fit px-4 py-2 bg-emerald-500 text-white rounded-xl text-[10px] font-black uppercase tracking-widest">Protegido</span>
              </div>
            </div>

            <div v-if="filteredAlerts.length === 0" class="py-16 flex flex-col items-center gap-4 border-2 border-dashed border-slate-100 dark:border-slate-800 rounded-3xl transition-colors">
              <div class="w-16 h-16 rounded-full bg-slate-50 dark:bg-slate-900 flex items-center justify-center">
                 <ShieldCheck class="w-8 h-8 text-slate-200 dark:text-slate-800" />
              </div>
              <p class="text-slate-400 dark:text-slate-600 font-black uppercase tracking-widest text-[10px]">Sin {{ showArchived ? 'archivos' : 'incidentes' }} registrados</p>
            </div>
            
            <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              <div v-for="alert in filteredAlerts" :key="alert.id" 
                   class="flex flex-col p-5 bg-slate-50/50 dark:bg-slate-900/40 border border-slate-100 dark:border-slate-800 rounded-3xl group hover:border-primary-500/50 transition-all cursor-pointer relative overflow-hidden shadow-sm hover:shadow-xl hover:shadow-primary-500/10 active:scale-[0.98]"
                   @click="selectedAlert = alert">
                
                <div class="flex items-center justify-between mb-4">
                  <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-xl bg-white dark:bg-slate-800 flex items-center justify-center text-rose-500 font-black text-[10px] shadow-sm">IP</div>
                    <div>
                      <p class="text-xs font-black text-slate-800 dark:text-white truncate">{{ alert.resource_id }}</p>
                      <p class="text-[9px] text-slate-500 dark:text-slate-400 font-mono tracking-tighter">{{ formatTime(alert.timestamp) }}</p>
                    </div>
                  </div>
                  <div class="p-1.5 rounded-lg bg-orange-50 dark:bg-yellow-900/20 text-orange-600 dark:text-yellow-500">
                    <ShieldAlert class="w-4 h-4" />
                  </div>
                </div>

                <div class="bg-white dark:bg-slate-800/50 p-3 rounded-2xl mb-4 h-16 overflow-hidden">
                   <p class="text-[9px] font-mono text-slate-500 dark:text-slate-400 break-all leading-relaxed">
                     {{ alert.resource }}
                   </p>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-[9px] font-black uppercase tracking-widest text-primary-500">Detalles</span>
                  <div class="flex gap-2">
                    <button 
                      v-if="!showArchived"
                      @click.stop="archiveAlert(alert)"
                      class="p-2 rounded-xl bg-slate-950 dark:bg-white text-white dark:text-slate-950 hover:scale-110 transition-all shadow-lg active:scale-95"
                      title="Archivar"
                    >
                      <Archive class="w-4 h-4" />
                    </button>
                    <button 
                      v-else
                      @click.stop="unarchiveAlert(alert)"
                      class="p-2 rounded-xl bg-emerald-500 text-white hover:scale-110 transition-all shadow-lg active:scale-95"
                      title="Desarchivar"
                    >
                      <RotateCcw class="w-4 h-4" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Activity Log Section (RESTORING FULL TABLE) -->
          <div v-if="!showArchived" class="glass p-6 md:p-10 rounded-[2.5rem] space-y-6">
            <h3 class="text-lg md:text-xl font-black dark:text-white flex items-center gap-2">
              <Activity class="text-primary-600" />
              Logs de Auditoría
            </h3>
            <div class="overflow-x-auto scrollbar-hide">
              <table class="w-full text-left">
                <thead>
                  <tr class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                    <th class="pb-4 px-2">Identidad</th>
                    <th class="pb-4 px-2">Método</th>
                    <th class="pb-4 px-2">Recurso</th>
                    <th class="pb-4 px-2">Latencia</th>
                    <th class="pb-4 px-2 text-right">Timestamp</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                  <tr v-for="log in metrics.recent_logs" :key="log.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors">
                    <td class="py-3 px-2">
                      <div class="flex items-center gap-2">
                        <div class="w-6 h-6 bg-slate-100 dark:bg-slate-800 rounded-lg flex items-center justify-center text-[9px] font-black text-slate-500">
                          {{ log.user_email?.[0]?.toUpperCase() || '?' }}
                        </div>
                        <span class="text-[11px] font-bold text-slate-700 dark:text-slate-300 truncate max-w-[120px]">{{ log.user_email }}</span>
                      </div>
                    </td>
                    <td class="py-3 px-2">
                      <span class="px-2 py-0.5 rounded-md text-[9px] font-black tracking-widest uppercase" :class="getActionClass(log.action)">
                        {{ log.action }}
                      </span>
                    </td>
                    <td class="py-3 px-2">
                        <span class="text-[10px] text-slate-500 dark:text-slate-400 font-mono truncate max-w-[150px] inline-block">{{ log.resource }}</span>
                    </td>
                    <td class="py-3 px-2">
                      <span class="text-[9px] font-black px-2 py-0.5 rounded-lg border dark:border-slate-700" :class="getLatencyClass(log.duration_ms)">
                        {{ log.duration_ms || 0 }}ms
                      </span>
                    </td>
                    <td class="py-3 px-2 text-right">
                      <span class="text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-tighter">{{ formatTime(log.timestamp) }}</span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Detailed Modal -->
    <div v-if="selectedAlert" class="fixed inset-0 z-[100] flex items-center justify-center p-4 md:p-6 bg-slate-950/90 backdrop-blur-xl animate-in fade-in duration-300">
      <div class="glass max-w-2xl w-full p-0 rounded-[2.5rem] overflow-hidden shadow-2xl animate-in zoom-in-95 duration-200 border border-slate-800">
        <div class="bg-rose-600 p-8 text-white relative">
          <button @click="selectedAlert = null" class="absolute top-6 right-6 p-2 hover:bg-white/20 rounded-xl transition-colors">
            <X class="w-6 h-6" />
          </button>
          <div class="flex items-center gap-4">
            <div class="p-3 bg-white/20 rounded-2xl backdrop-blur-md">
              <ShieldAlert class="w-8 h-8" />
            </div>
            <div>
              <h3 class="text-2xl font-black">Análisis de Amenaza</h3>
              <p class="text-rose-100 font-bold uppercase tracking-widest text-[10px]">Identificador: {{ selectedAlert.id }}</p>
            </div>
          </div>
        </div>

        <div class="p-8 space-y-8 bg-white dark:bg-slate-950 transition-colors duration-500">
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Origen / Atacante</p>
              <div class="bg-slate-50 dark:bg-slate-900 p-4 rounded-2xl flex items-center gap-3 border border-slate-100 dark:border-slate-800">
                <Globe class="w-5 h-5 text-indigo-500" />
                <span class="font-mono font-black text-slate-700 dark:text-white text-base md:text-lg">{{ selectedAlert.resource_id }}</span>
              </div>
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Fecha de Bloqueo</p>
              <div class="bg-slate-50 dark:bg-slate-900 p-4 rounded-2xl flex items-center gap-3 border border-slate-100 dark:border-slate-800">
                <Clock class="w-5 h-5 text-indigo-500" />
                <span class="font-black text-slate-700 dark:text-white text-xs uppercase">{{ formatDate(selectedAlert.timestamp) }} - {{ formatTime(selectedAlert.timestamp) }}</span>
              </div>
            </div>
          </div>

          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Payload Inyectado</p>
            <div class="bg-rose-50 dark:bg-rose-950/20 p-4 rounded-2xl border border-rose-100 dark:border-rose-900/30 break-all overflow-y-auto max-h-[140px] font-mono text-[10px] text-rose-600 dark:text-rose-400 leading-relaxed">
              {{ selectedAlert.resource }}
            </div>
          </div>

          <div class="bg-amber-50 dark:bg-amber-900/20 border border-amber-100 dark:border-amber-900/30 p-5 rounded-3xl flex gap-4">
             <div class="p-3 bg-white dark:bg-slate-800 rounded-2xl shadow-sm h-fit">
               <ShieldCheck class="w-6 h-6 text-amber-500" />
             </div>
             <div>
               <p class="font-black text-amber-800 dark:text-amber-400 text-xs uppercase tracking-widest">Análisis del WAF</p>
               <p class="text-amber-700/80 dark:text-amber-500/80 text-[10px] md:text-xs mt-1 leading-relaxed">
                 El sistema detectó un intento de inyección de código SQL malicioso. El ataque fue neutralizado con éxito antes de llegar a la capa de datos.
               </p>
             </div>
          </div>

          <div v-if="!showArchived" class="flex gap-3">
             <button @click="selectedAlert = null" class="flex-1 py-4 border-2 border-slate-100 dark:border-slate-800 rounded-2xl font-black text-xs uppercase tracking-widest text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-900 transition-colors">Volver</button>
             <button @click="archiveAlert(selectedAlert); selectedAlert = null" class="flex-[2] bg-slate-950 dark:bg-white text-white dark:text-slate-950 py-4 font-black uppercase tracking-widest text-xs rounded-2xl hover:scale-105 active:scale-95 transition-all shadow-xl dark:shadow-none">Archivar & Confirmar Revisión</button>
          </div>
          <button v-else @click="selectedAlert = null" class="w-full py-4 bg-primary-600 text-white font-black uppercase tracking-widest text-xs rounded-2xl hover:scale-105 active:scale-95 transition-all">Reporte Revisado</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, h, computed, watch } from 'vue';
import Navbar from '../components/Navbar.vue';
import api from '../api/axios';
import { useAuthStore } from '../store/auth';
import { 
  Loader2, ShieldCheck, Zap, Activity, ShieldAlert, 
  X, Archive, RotateCcw, Globe, Clock, Info 
} from 'lucide-vue-next';
import Chart from 'chart.js/auto';

const archivedAlerts = ref(JSON.parse(localStorage.getItem('archived_alerts') || '[]'));
const showArchived = ref(false);

const MetricTooltip = {
  props: ['title', 'iso', 'description'],
  setup(props) {
    return () => h('div', {
      class: 'absolute -top-2 left-0 right-0 z-50 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 pointer-events-none'
    }, [
      h('div', {
        class: 'absolute bottom-full left-1/2 -translate-x-1/2 mb-2 w-72 p-4 bg-slate-900 text-white rounded-xl shadow-2xl'
      }, [
        h('p', { class: 'font-black uppercase tracking-widest text-[10px] text-primary-400 mb-1' }, props.title),
        h('p', { class: 'text-slate-300 text-[10px] leading-relaxed font-bold' }, props.description),
        h('div', { class: 'absolute top-full left-1/2 -translate-x-1/2 border-8 border-transparent border-t-slate-900' })
      ])
    ]);
  }
};

const metrics = ref({
    total_users: 0, total_tasks: 0, total_requests: 0,
    avg_latency: 0, availability: 0,
    reliability: { maturity: 0, fault_tolerance: 0, recoverability: 0 },
    performance: { resource_utilization: 0, capacity: 0 },
    latency_history: [], recent_logs: [], security_alerts: []
});

const filteredAlerts = computed(() => {
    if (showArchived.value) {
        return archivedAlerts.value;
    }
    const archivedIds = new Set(archivedAlerts.value.map(a => a.id));
    return (metrics.value.security_alerts || []).filter(a => !archivedIds.has(a.id));
});

const archiveAlert = (alert) => {
    if (!archivedAlerts.value.find(a => a.id === alert.id)) {
        archivedAlerts.value.push(alert);
        localStorage.setItem('archived_alerts', JSON.stringify(archivedAlerts.value));
    }
};

const unarchiveAlert = (alert) => {
    archivedAlerts.value = archivedAlerts.value.filter(a => a.id !== alert.id);
    localStorage.setItem('archived_alerts', JSON.stringify(archivedAlerts.value));
};

const reliabilityMetrics = computed(() => ({
    avail: { label: 'Disponibilidad', title: 'Disponibilidad', value: metrics.value.availability || 100, desc: 'Fiabilidad del servicio.' },
    matu: { label: 'Madurez', title: 'Madurez', value: metrics.value.reliability?.maturity || 0, desc: 'Estabilidad del software.' },
    fault: { label: 'Tolerancia', title: 'Tolerancia a Fallos', value: metrics.value.reliability?.fault_tolerance || 0, desc: 'Capacidad ante errores.' }
}));

const performanceMetrics = computed(() => ({
    lat: { label: 'Latencia', title: 'Latencia Media', value: metrics.value.avg_latency || 0, percent: Math.min((metrics.value.avg_latency || 0) / 10, 100), display: (metrics.value.avg_latency || 0).toFixed(1) + 'ms', desc: 'Comportamiento temporal.' },
    res: { label: 'Recursos', title: 'Uso de Memoria', value: metrics.value.performance?.resource_utilization || 0, percent: metrics.value.performance?.resource_utilization || 0, desc: 'Eficiencia de hardware.' },
    cap: { label: 'Capacidad', title: 'Capacidad', value: metrics.value.performance?.capacity || 0, percent: metrics.value.performance?.capacity || 0, desc: 'Saturación del sistema.' }
}));

const loading = ref(true);
const selectedAlert = ref(null);
const latencyChartCanvas = ref(null);
let chartInstance = null;

const initChart = () => {
    if (!latencyChartCanvas.value) return;
    const ctx = latencyChartCanvas.value.getContext('2d');
    const isDarkMode = document.documentElement.classList.contains('dark');
    
    const history = [...(metrics.value.latency_history || [])].sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp));
    const labels = history.map(p => {
        const d = new Date(p.timestamp);
        return `${d.getHours()}:00`;
    });
    const data = history.map(p => p.latency);

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

const observer = new MutationObserver(() => {
    initChart();
});

const getHealthLabel = (v) => v >= 99 ? 'Optimized' : v >= 95 ? 'Stable' : 'Critical';
const getPerformanceLabel = (v) => v <= 50 ? 'Peak' : v <= 150 ? 'Steady' : 'High Load';
const getLatencyClass = (ms) => ms <= 50 ? 'text-emerald-500' : ms <= 150 ? 'text-amber-500' : 'text-rose-500';

const getActionClass = (a) => {
    switch(a) {
        case 'SECURITY_ALERT': return 'bg-rose-500 text-white';
        case 'CREATE': case 'POST': return 'text-emerald-500 bg-emerald-500/10';
        case 'DELETE': return 'text-rose-500 bg-rose-500/10';
        default: return 'text-indigo-500 bg-indigo-500/10';
    }
};

const formatDate = (s) => s ? new Date(s).toLocaleDateString() : '';
const formatTime = (s) => s ? new Date(s).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : '';

const fetchMetrics = async () => {
    try {
        const response = await api.get('/monitoring/');
        metrics.value = response.data;
        await nextTick();
        initChart();
    } catch (err) {
        console.error('Metrics Error:', err);
    } finally {
        loading.value = false;
    }
};

let refreshInterval = null;

onMounted(() => {
    fetchMetrics();
    observer.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] });
    refreshInterval = setInterval(() => {
        const authStore = useAuthStore();
        if (authStore.isAuthenticated && authStore.isAdmin) fetchMetrics();
    }, 10000); 
});

onUnmounted(() => {
    if (refreshInterval) clearInterval(refreshInterval);
    observer.disconnect();
});
</script>
