<template>
  <div class="flex-1 flex flex-col h-screen overflow-hidden">
    <Navbar />
    
    <main class="flex-1 overflow-y-auto p-4 md:p-8 lg:p-12">
      <div class="max-w-[1400px] mx-auto space-y-6 md:space-y-10">
        <header class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <h2 class="text-3xl md:text-4xl font-black text-slate-900 dark:text-white tracking-tight">Directorio de Usuarios</h2>
            <p class="text-slate-500 dark:text-slate-400 mt-1 text-sm md:text-base">Administración central de identidades y privilegios</p>
          </div>
          
          <div class="glass px-5 py-3 rounded-2xl flex items-center gap-3 border border-slate-200 dark:border-slate-800">
            <UsersIcon class="text-primary-600 dark:text-primary-400 w-5 h-5" />
            <span class="font-black text-slate-700 dark:text-slate-200 text-sm tracking-tight">{{ users.length }} Cuentas de Usuario</span>
          </div>
        </header>

        <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
          <Loader2 class="w-10 h-10 animate-spin text-primary-600" />
          <p class="text-slate-400 font-black uppercase tracking-widest text-[10px]">Cargando base de datos...</p>
        </div>

        <div v-else class="glass overflow-hidden rounded-[2rem] border border-white/20 dark:border-slate-800 shadow-2xl dark:shadow-none">
          <div class="overflow-x-auto">
            <table class="w-full text-left border-separate border-spacing-0">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-900/50">
                  <th class="px-6 py-5 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Nombre Completo</th>
                  <th class="px-6 py-5 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800 hidden md:table-cell">Correo Electrónico</th>
                  <th class="px-6 py-5 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Privilegios</th>
                  <th class="px-6 py-5 text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800 text-right hidden lg:table-cell">UUID de Usuario</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-100 dark:divide-slate-800/50">
                <tr v-for="user in users" :key="user.id" class="hover:bg-white/40 dark:hover:bg-slate-800/20 transition-all group">
                  <td class="px-6 py-4 md:py-5">
                    <div class="flex items-center gap-4">
                      <div class="w-10 h-10 md:w-12 md:h-12 rounded-xl bg-gradient-to-br from-primary-500/10 to-indigo-500/10 dark:from-primary-500/20 dark:to-indigo-500/20 border border-primary-500/20 flex items-center justify-center font-black text-primary-600 dark:text-primary-400 shadow-sm text-base md:text-lg">
                        {{ user.name.charAt(0).toUpperCase() }}
                      </div>
                      <div class="flex flex-col">
                        <span class="font-black text-slate-800 dark:text-white leading-tight">{{ user.name }}</span>
                        <span class="text-[10px] text-slate-400 md:hidden mt-0.5">{{ user.email }}</span>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 text-xs font-bold text-slate-600 dark:text-slate-400 hidden md:table-cell">{{ user.email }}</td>
                  <td class="px-6 py-4">
                    <span 
                      class="px-3 py-1.5 rounded-lg text-[9px] font-black uppercase tracking-widest shadow-sm"
                      :class="user.role === 'admin' ? 'bg-primary-600 text-white' : 'bg-slate-200 dark:bg-slate-800 text-slate-600 dark:text-slate-400'"
                    >
                      {{ user.role }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-right text-[10px] font-mono text-slate-300 dark:text-slate-600 uppercase tracking-tighter hidden lg:table-cell">{{ user.id }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Navbar from '../components/Navbar.vue';
import api from '../api/axios';
import { Users as UsersIcon, Loader2 } from 'lucide-vue-next';

const users = ref([]);
const loading = ref(true);

const fetchUsers = async () => {
    loading.value = true;
    try {
        const response = await api.get('/users/');
        users.value = Array.isArray(response.data) ? response.data : [];
    } catch (err) {
        console.error('Error al cargar usuarios:', err);
    } finally {
        loading.value = false;
    }
};

onMounted(fetchUsers);
</script>
