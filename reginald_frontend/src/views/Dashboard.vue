<template>
  <div class="flex-1 flex flex-col h-screen overflow-hidden">
    <Navbar />
    
    <main class="flex-1 overflow-y-auto p-4 md:p-8 lg:p-12">
      <div class="max-w-[1600px] mx-auto space-y-6 md:space-y-10">
        <!-- Header Section -->
        <header class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 md:gap-6">
          <div>
            <h2 class="text-3xl md:text-4xl font-black text-slate-900 dark:text-white tracking-tight">Mis Tareas</h2>
            <p class="text-slate-500 dark:text-slate-400 mt-1 text-sm md:text-base">Gestiona tus actividades diarias y objetivos estratégicos</p>
          </div>
          
          <button 
            v-if="auth.isAdmin"
            @click="showAddModal = true"
            class="bg-primary-600 hover:bg-primary-700 text-white font-bold px-5 py-3.5 md:px-8 md:py-4 rounded-2xl shadow-xl shadow-primary-200 dark:shadow-none transition-all flex items-center justify-center gap-2 group active:scale-95"
          >
            <Plus class="w-5 h-5 group-hover:rotate-90 transition-transform" />
            <span class="text-sm md:text-base">Nueva Tarea</span>
          </button>
        </header>

        <!-- Stats Grid - More compact on mobile -->
        <div class="grid grid-cols-2 md:grid-cols-3 gap-3 md:gap-6">
          <div class="glass p-4 md:p-6 rounded-3xl flex items-center gap-3 md:gap-4 transition-transform hover:scale-[1.02]">
            <div class="bg-blue-100 dark:bg-blue-900/30 p-2.5 md:p-3 rounded-2xl">
              <ClipboardList class="text-blue-600 dark:text-blue-400 w-5 h-5 md:w-6 md:h-6" />
            </div>
            <div>
              <p class="text-xl md:text-2xl font-black dark:text-white">{{ tasks.length }}</p>
              <p class="text-[10px] text-slate-500 dark:text-slate-400 uppercase font-black tracking-tighter">Totales</p>
            </div>
          </div>
          <div class="glass p-4 md:p-6 rounded-3xl flex items-center gap-3 md:gap-4 transition-transform hover:scale-[1.02]">
            <div class="bg-amber-100 dark:bg-amber-900/30 p-2.5 md:p-3 rounded-2xl">
              <Clock class="text-amber-600 dark:text-amber-400 w-5 h-5 md:w-6 md:h-6" />
            </div>
            <div>
              <p class="text-xl md:text-2xl font-black dark:text-white">{{ pendingTasks.length }}</p>
              <p class="text-[10px] text-slate-500 dark:text-slate-400 uppercase font-black tracking-tighter">Pendientes</p>
            </div>
          </div>
          <div class="glass p-4 md:p-6 rounded-3xl flex items-center gap-3 md:gap-4 transition-transform hover:scale-[1.02] col-span-2 md:col-span-1">
            <div class="bg-emerald-100 dark:bg-emerald-900/30 p-2.5 md:p-3 rounded-2xl">
              <CheckCircle2 class="text-emerald-600 dark:text-emerald-400 w-5 h-5 md:w-6 md:h-6" />
            </div>
            <div>
              <p class="text-xl md:text-2xl font-black dark:text-white">{{ completedTasks.length }}</p>
              <p class="text-[10px] text-slate-500 dark:text-slate-400 uppercase font-black tracking-tighter">Completadas</p>
            </div>
          </div>
        </div>

        <!-- Tasks List -->
        <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
          <Loader2 class="w-10 h-10 animate-spin text-primary-600" />
          <p class="text-slate-400 font-medium">Sincronizando tareas...</p>
        </div>

        <div v-else-if="tasks.length === 0" class="glass rounded-3xl p-12 md:p-20 flex flex-col items-center text-center border-dashed border-2">
            <div class="bg-slate-100 dark:bg-slate-800 p-6 rounded-full mb-6 text-slate-400">
                <Box class="w-12 h-12" />
            </div>
            <h3 class="text-xl font-bold text-slate-800 dark:text-white">Sin tareas activas</h3>
            <p class="text-slate-500 dark:text-slate-400 mt-2 max-w-xs text-sm">Organiza tu día agregando una nueva tarea en el botón superior.</p>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6 pb-12">
          <div 
            v-for="task in tasks" 
            :key="task.id"
            class="glass hover:border-primary-400/50 dark:hover:border-primary-500/50 transition-all group overflow-hidden flex flex-col rounded-3xl relative"
          >
            <div class="p-5 md:p-6 flex-1">
              <div class="flex items-start justify-between mb-4">
                <span 
                  class="px-2.5 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest"
                  :class="task.status === 'done' ? 'bg-emerald-100 dark:bg-emerald-950/50 text-emerald-700 dark:text-emerald-400' : 'bg-amber-100 dark:bg-amber-950/50 text-amber-700 dark:text-amber-400'"
                >
                  {{ task.status === 'done' ? 'Listo' : 'Haciendo' }}
                </span>
                <div class="flex flex-col items-end">
                  <span class="text-[10px] text-slate-400 dark:text-slate-500 font-bold uppercase tracking-tighter">{{ formatDate(task.created_at) }}</span>
                  <span v-if="task.assigned_to" class="text-[9px] text-primary-600 dark:text-primary-400 font-black uppercase mt-1">
                    {{ task.assigned_to.split('@')[0] }}
                  </span>
                </div>
              </div>
              
              <div v-if="editingTaskId === task.id" class="space-y-3">
                <input 
                  v-model="editTitle"
                  class="w-full px-3 py-2 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-primary-500 dark:text-white"
                  @keyup.enter="saveEdit(task)"
                />
                <div class="flex gap-2">
                  <button @click="saveEdit(task)" class="text-[10px] font-black uppercase tracking-widest text-emerald-600 dark:text-emerald-400">Guardar</button>
                  <button @click="editingTaskId = null" class="text-[10px] font-black uppercase tracking-widest text-slate-400">Cancelar</button>
                </div>
              </div>
              <h4 v-else class="text-base md:text-lg font-bold text-slate-800 dark:text-white group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors leading-snug">
                {{ task.title }}
              </h4>
            </div>

            <div class="p-4 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between transition-all md:opacity-0 md:translate-y-2 group-hover:opacity-100 group-hover:translate-y-0">
              <div class="flex gap-4">
                <button 
                  @click="requestStatusToggle(task)"
                  class="text-xs font-bold text-primary-600 dark:text-primary-400 hover:scale-110 transition-transform flex items-center gap-1.5"
                  title="Cambiar estado"
                >
                  <component :is="task.status === 'done' ? RotateCcw : CheckCircle" class="w-4.5 h-4.5 md:w-5 md:h-5" />
                </button>
                <button 
                  @click="startEdit(task)"
                  class="text-xs font-bold text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 hover:scale-110 transition-transform"
                  title="Editar"
                >
                  <Pencil class="w-4.5 h-4.5 md:w-5 md:h-5" />
                </button>
              </div>
              
              <button 
                @click="deleteTask(task.id)"
                class="text-xs font-bold text-rose-500 hover:text-rose-700 hover:scale-110 transition-all"
                title="Eliminar"
              >
                <Trash2 class="w-4.5 h-4.5 md:w-5 md:h-5" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Modals (Add, Confirm) updated for dark mode -->
    <div v-if="showAddModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-white/80 dark:bg-slate-950/80 backdrop-blur-md animate-in fade-in duration-300">
      <div class="glass max-w-md w-full p-8 rounded-[2.5rem] shadow-2xl animate-in zoom-in-95 duration-200">
        <h3 class="text-2xl font-black mb-6 dark:text-white tracking-tighter">Nueva Tarea</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1 mb-2">Descripción</label>
            <input 
              v-model="newTaskTitle"
              type="text"
              class="w-full px-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl mb-4 focus:ring-2 focus:ring-primary-600 outline-none dark:text-white transition-all"
              placeholder="Ej: Revisar logs del servidor"
              @keyup.enter="addTask"
            />
          </div>

          <div v-if="auth.isAdmin">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1 mb-2">Asignar Responsable</label>
            <select 
              v-model="assignedTo"
              class="w-full px-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl focus:ring-2 focus:ring-primary-600 outline-none appearance-none cursor-pointer dark:text-white transition-all"
            >
              <option value="">Mi perfil</option>
              <option v-for="user in usersList" :key="user.email" :value="user.email">
                {{ user.name }}
              </option>
            </select>
          </div>
        </div>

        <div class="flex gap-3 mt-8">
          <button @click="showAddModal = false" class="flex-1 py-4 text-slate-400 font-bold hover:bg-slate-100 dark:hover:bg-slate-800 rounded-2xl transition-colors">Cancelar</button>
          <button @click="addTask" class="flex-1 bg-primary-600 text-white font-bold py-4 rounded-2xl shadow-xl shadow-primary-200 dark:shadow-none active:scale-95 transition-all">Crear Tarea</button>
        </div>
      </div>
    </div>

    <!-- Confirm Modal -->
    <div v-if="showConfirmModal" class="fixed inset-0 z-[110] flex items-center justify-center p-6 bg-white/80 backdrop-blur-lg">
      <div class="glass max-w-sm w-full p-10 rounded-[2.5rem] text-center space-y-6 animate-in zoom-in-90 duration-300">
        <div class="w-20 h-20 bg-emerald-100 dark:bg-emerald-900/30 rounded-full flex items-center justify-center mx-auto text-emerald-600 dark:text-emerald-400">
           <CheckCircle2 class="w-10 h-10" />
        </div>
        <div>
          <h3 class="text-2xl font-black text-slate-900 dark:text-white leading-tight">¿Tarea finalizada?</h3>
          <p class="text-slate-500 dark:text-slate-400 mt-3 text-sm font-medium">Esto registrará la actividad como completada.</p>
        </div>
        <div class="flex flex-col gap-3 pt-2">
          <button @click="confirmCompletion" class="w-full bg-emerald-500 hover:bg-emerald-600 text-white font-black py-4 rounded-2xl shadow-xl shadow-emerald-200 dark:shadow-none transition-all active:scale-95">Sí, completar</button>
          <button @click="showConfirmModal = false" class="w-full py-3 text-slate-400 font-bold hover:text-slate-600 dark:hover:text-slate-300 transition-colors">Aún no</button>
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
import { 
    Plus, ClipboardList, Clock, CheckCircle2, Box, 
    Loader2, CheckCircle, RotateCcw, Trash2, Pencil 
} from 'lucide-vue-next';

const auth = useAuthStore();
const tasks = ref([]);
const usersList = ref([]);
const loading = ref(true);
const showAddModal = ref(false);
const showConfirmModal = ref(false);
const taskToConfirm = ref(null);
const newTaskTitle = ref('');
const assignedTo = ref('');

const editingTaskId = ref(null);
const editTitle = ref('');

const pendingTasks = computed(() => tasks.value.filter(t => t.status !== 'done'));
const completedTasks = computed(() => tasks.value.filter(t => t.status === 'done'));

const fetchTasks = async () => {
    loading.value = true;
    try {
        const response = await api.get('/tasks/');
        tasks.value = Array.isArray(response.data) ? response.data : [];
    } catch (err) {
        console.error('Error al cargar tareas:', err);
        tasks.value = [];
    } finally {
        loading.value = false;
    }
};

const fetchUsers = async () => {
    if (!auth.isAdmin) return;
    try {
        const response = await api.get('/users/');
        usersList.value = response.data;
    } catch (err) {
        console.error('Error al cargar usuarios:', err);
    }
};

const addTask = async () => {
    if (!newTaskTitle.value.trim()) return;
    try {
        await api.post('/tasks/', { 
            title: newTaskTitle.value,
            created_by: auth.user?.email || 'anonymous',
            assigned_to: assignedTo.value || auth.user?.email
        });
        newTaskTitle.value = '';
        assignedTo.value = '';
        showAddModal.value = false;
        fetchTasks();
    } catch (err) {
        console.error('Error al crear tarea:', err);
    }
};

const deleteTask = async (id) => {
    if (!confirm('¿Seguro que deseas eliminar esta tarea?')) return;
    try {
        await api.delete(`/tasks/?id=${id}`);
        fetchTasks();
    } catch (err) {
        console.error('Error al eliminar tarea:', err);
    }
};

const startEdit = (task) => {
    editingTaskId.value = task.id;
    editTitle.value = task.title;
};

const saveEdit = async (task) => {
    if (!editTitle.value.trim()) return;
    try {
        await api.patch('/tasks/', {
            id: task.id,
            title: editTitle.value,
            status: task.status
        });
        editingTaskId.value = null;
        fetchTasks();
    } catch (err) {
        console.error('Error al editar tarea:', err);
    }
};

const requestStatusToggle = (task) => {
    if (task.status === 'done') {
        toggleStatus(task);
    } else {
        taskToConfirm.value = task;
        showConfirmModal.value = true;
    }
};

const confirmCompletion = () => {
    if (taskToConfirm.value) {
        toggleStatus(taskToConfirm.value);
    }
    showConfirmModal.value = false;
    taskToConfirm.value = null;
};

const toggleStatus = async (task) => {
    const newStatus = task.status === 'done' ? 'pending' : 'done';
    try {
        await api.patch('/tasks/', {
            id: task.id,
            status: newStatus
        });
        fetchTasks();
    } catch (err) {
        console.error('Error al actualizar estado:', err);
    }
};

const formatDate = (dateString) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('es-ES', {
        day: 'numeric',
        month: 'short'
    });
};

onMounted(() => {
    fetchTasks();
    fetchUsers();
});
</script>
