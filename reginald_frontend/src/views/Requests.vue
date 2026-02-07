<template>
  <div class="flex-1 flex flex-col h-screen overflow-hidden">
    <Navbar />
    
    <main class="flex-1 overflow-y-auto p-4 md:p-8 lg:p-12">
      <div class="max-w-[1600px] mx-auto space-y-12">
        <!-- Main Section: Active Requests -->
        <section class="space-y-6 md:space-y-10">
          <header class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 md:gap-6">
            <div>
              <h2 class="text-3xl md:text-4xl font-black text-slate-900 dark:text-white tracking-tight">Solicitudes Activas</h2>
              <p class="text-slate-500 dark:text-slate-400 mt-1 text-sm md:text-base">
                {{ auth.isAdmin ? 'Gestión centralizada de requerimientos de usuario' : 'Monitorea el progreso de tus requerimientos' }}
              </p>
            </div>
            
            <button 
              v-if="!auth.isAdmin"
              @click="showAddModal = true"
              class="bg-indigo-600 hover:bg-indigo-700 text-white font-bold px-6 py-4 rounded-2xl shadow-xl shadow-indigo-200 dark:shadow-none transition-all flex items-center justify-center gap-2 group active:scale-95"
            >
              <Send class="w-5 h-5 group-hover:translate-x-1 group-hover:-translate-y-1 transition-transform" />
              <span>Nueva Solicitud</span>
            </button>
          </header>

          <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
            <Loader2 class="w-10 h-10 animate-spin text-indigo-600" />
            <p class="text-slate-400 font-medium">Sincronizando solicitudes...</p>
          </div>

          <div v-else-if="activeRequests.length === 0" class="glass rounded-3xl p-12 md:p-20 flex flex-col items-center text-center border-dashed border-2">
              <div class="bg-indigo-50 dark:bg-indigo-900/30 p-6 rounded-full mb-6 text-indigo-400">
                  <Inbox class="w-12 h-12" />
              </div>
              <h3 class="text-xl font-bold text-slate-800 dark:text-white">Sin solicitudes pendientes</h3>
              <p class="text-slate-500 dark:text-slate-400 mt-2 max-w-xs text-sm">Todo está al día en tu bandeja de entrada.</p>
          </div>

          <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6">
            <div 
              v-for="req in activeRequests" 
              :key="req.id"
              class="glass hover:border-indigo-400/50 dark:hover:border-indigo-500/50 transition-all group overflow-hidden flex flex-col rounded-3xl"
            >
              <div class="p-5 md:p-6 flex-1">
                <div class="flex items-start justify-between mb-4">
                  <span 
                    class="px-2.5 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest"
                    :class="getStatusClass(req.status)"
                  >
                    {{ formatStatus(req.status) }}
                  </span>
                  <span class="text-[10px] text-slate-400 dark:text-slate-500 font-bold uppercase tracking-tighter">{{ formatDate(req.created_at) }}</span>
                </div>
                <h4 class="text-base md:text-lg font-bold text-slate-800 dark:text-white group-hover:text-indigo-600 dark:group-hover:text-indigo-400 transition-colors leading-snug">
                  {{ req.title }}
                </h4>
                <p class="text-xs md:text-sm text-slate-500 dark:text-slate-400 mt-3 line-clamp-3 leading-relaxed">{{ req.description }}</p>
                
                <div class="mt-5 pt-5 border-t border-slate-100 dark:border-slate-800 flex items-center gap-3">
                  <div class="w-7 h-7 rounded-lg bg-indigo-100 dark:bg-indigo-900/30 flex items-center justify-center text-[10px] font-black text-indigo-600 dark:text-indigo-400">
                      {{ req.created_by.charAt(0).toUpperCase() }}
                  </div>
                  <span class="text-[10px] text-slate-400 dark:text-slate-500 font-bold truncate">{{ req.created_by }}</span>
                </div>
              </div>

              <div v-if="auth.isAdmin && req.status === 'pending'" class="bg-slate-50/50 dark:bg-slate-900/50 p-4 border-t border-slate-100 dark:border-slate-800 flex items-center gap-2">
                <button 
                  @click="updateRequestStatus(req.id, 'approved')"
                  class="flex-1 bg-emerald-500 hover:bg-emerald-600 text-white text-xs font-black uppercase tracking-widest py-3 rounded-xl transition-all active:scale-95"
                >
                  Aprobar
                </button>
                <button 
                  @click="updateRequestStatus(req.id, 'rejected')"
                  class="flex-1 bg-white dark:bg-slate-800 hover:bg-rose-50 dark:hover:bg-rose-900/10 text-rose-600 border border-rose-100 dark:border-rose-900/30 text-xs font-black uppercase tracking-widest py-3 rounded-xl transition-all active:scale-95"
                >
                  Rechazar
                </button>
              </div>
            </div>
          </div>
        </section>

        <!-- History Section -->
        <section v-if="historyRequests.length > 0" class="space-y-6 md:space-y-10 pt-10 border-t border-slate-200 dark:border-slate-800">
          <header>
            <h3 class="text-2xl md:text-3xl font-black text-slate-800 dark:text-white tracking-tighter">Historial de Solicitudes</h3>
            <p class="text-slate-500 dark:text-slate-400 mt-1 text-sm">Registro histórico de requerimientos procesados</p>
          </header>

          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6 pb-20">
            <div 
              v-for="req in historyRequests" 
              :key="req.id"
              class="glass opacity-75 hover:opacity-100 grayscale-[0.5] hover:grayscale-0 transition-all group overflow-hidden flex flex-col rounded-[2rem] border border-slate-100 dark:border-slate-800 shadow-sm"
            >
              <div class="p-6 flex-1">
                <div class="flex items-start justify-between mb-4">
                  <span 
                    class="px-2.5 py-1 rounded-lg text-[8px] font-black uppercase tracking-widest"
                    :class="getStatusClass(req.status)"
                  >
                    {{ formatStatus(req.status) }}
                  </span>
                  <div class="flex items-center gap-3">
                    <span class="text-[9px] text-slate-400 font-bold">{{ formatDate(req.created_at) }}</span>
                    <button 
                      v-if="auth.isAdmin"
                      @click="deleteRequest(req.id)"
                      class="text-rose-400 hover:text-rose-600 transition-all hover:scale-125 p-1"
                      title="Eliminar del historial"
                    >
                      <Trash2 class="w-3.5 h-3.5" />
                    </button>
                  </div>
                </div>
                <h4 class="text-sm md:text-base font-bold text-slate-700 dark:text-slate-200 leading-snug">
                  {{ req.title }}
                </h4>
                
                <div class="mt-4 pt-4 border-t border-slate-50 dark:border-slate-800/50 flex items-center justify-between">
                  <div class="flex items-center gap-2">
                    <div class="w-6 h-6 rounded-lg bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-[9px] font-black text-slate-400">
                        {{ req.created_by.charAt(0).toUpperCase() }}
                    </div>
                    <span class="text-[9px] text-slate-400 font-medium truncate max-w-[80px]">{{ req.created_by }}</span>
                  </div>
                  <span v-if="req.handled_by" class="text-[8px] text-slate-500 font-bold uppercase italic opacity-60">ADMIN: {{ req.handled_by.split('@')[0] }}</span>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>

    <!-- Create Request Modal -->
    <div v-if="showAddModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-white/80 dark:bg-slate-950/80 backdrop-blur-md animate-in fade-in duration-300">
      <div class="glass max-w-md w-full p-8 rounded-[2.5rem] shadow-2xl animate-in zoom-in-95 duration-200">
        <h3 class="text-2xl font-black mb-6 dark:text-white tracking-tighter">Nueva Solicitud</h3>
        <div class="space-y-5">
            <div>
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1 mb-2 block">Asunto / Título</label>
                <input 
                    v-model="newRequest.title"
                    type="text"
                    class="w-full px-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl focus:ring-2 focus:ring-indigo-600 outline-none transition-all dark:text-white"
                    placeholder="Ej: Acceso a servidor de producción"
                />
            </div>
            <div>
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1 mb-2 block">Justificación / Detalles</label>
                <textarea 
                    v-model="newRequest.description"
                    rows="4"
                    class="w-full px-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl focus:ring-2 focus:ring-indigo-600 outline-none resize-none transition-all dark:text-white text-sm"
                    placeholder="Describe los motivos técnicos de tu solicitud..."
                ></textarea>
            </div>
        </div>
        <div class="flex gap-3 mt-8">
          <button @click="showAddModal = false" class="flex-1 py-4 text-slate-400 font-bold hover:bg-slate-100 dark:hover:bg-slate-800 rounded-2xl transition-colors">Cancelar</button>
          <button @click="createRequest" class="flex-1 bg-indigo-600 text-white font-bold py-4 rounded-2xl shadow-xl shadow-indigo-200 dark:shadow-none active:scale-95 transition-all">Enviar Solicitud</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import Navbar from '../components/Navbar.vue';
import api from '../api/axios';
import { useAuthStore } from '../store/auth';
import { Send, Loader2, Inbox, Trash2 } from 'lucide-vue-next';

const auth = useAuthStore();
const requests = ref([]);
const loading = ref(true);
const showAddModal = ref(false);

const newRequest = ref({
    title: '',
    description: ''
});

const activeRequests = computed(() => 
    requests.value.filter(r => r.status === 'pending' || r.status === 'review')
);

const historyRequests = computed(() => 
    requests.value.filter(r => r.status === 'approved' || r.status === 'rejected')
);

const fetchRequests = async () => {
    loading.value = true;
    try {
        const response = await api.get('/requests/');
        requests.value = Array.isArray(response.data) ? response.data : [];
    } catch (err) {
        console.error('Error al cargar solicitudes:', err);
    } finally {
        loading.value = false;
    }
};

const createRequest = async () => {
    if (!newRequest.value.title.trim()) return;
    try {
        await api.post('/requests/', newRequest.value);
        newRequest.value = { title: '', description: '' };
        showAddModal.value = false;
        fetchRequests();
    } catch (err) {
        console.error('Error al crear solicitud:', err);
    }
};

const updateRequestStatus = async (id, status) => {
    try {
        await api.patch('/requests/', { id, status });
        fetchRequests();
    } catch (err) {
        console.error('Error al procesar solicitud:', err);
    }
};

const deleteRequest = async (id) => {
    if (!confirm('¿Seguro que deseas eliminar esta solicitud del historial?')) return;
    try {
        await api.delete(`/requests/?id=${id}`);
        fetchRequests();
    } catch (err) {
        console.error('Error al eliminar solicitud:', err);
    }
};

const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('es-ES', {
        day: 'numeric',
        month: 'short',
        year: 'numeric'
    });
};

const formatStatus = (status) => {
    const map = {
        'pending': 'Pendiente',
        'review': 'En Revisión',
        'approved': 'Aprobada',
        'rejected': 'Rechazada'
    };
    return map[status] || status;
};

const getStatusClass = (status) => {
    const map = {
        'pending': 'bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-400',
        'review': 'bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400',
        'approved': 'bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400',
        'rejected': 'bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-400'
    };
    return map[status] || 'bg-slate-100 text-slate-700';
};

onMounted(fetchRequests);
</script>
