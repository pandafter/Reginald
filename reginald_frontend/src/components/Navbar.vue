<template>
  <nav class="glass sticky top-0 z-50 px-4 md:px-8 py-3 flex items-center justify-between border-b dark:border-slate-800">
    <div class="flex items-center gap-6 md:gap-12">
      <div class="flex items-center gap-3">
        <img src="/LogoBlanco.svg" alt="Reginald OS" class="h-8 md:h-10 hidden dark:block">
        <img src="/LogoAzul.svg" alt="Reginald OS" class="h-8 md:h-10 block dark:hidden">
      </div>
      
      <nav class="hidden lg:flex items-center gap-1 bg-slate-100/50 dark:bg-slate-800/50 p-1 rounded-2xl border border-slate-200/50 dark:border-slate-700/50">
        <router-link to="/dashboard" class="nav-link" active-class="nav-link-active">
          <LayoutDashboard class="w-4 h-4" /> Tareas
        </router-link>
        <router-link to="/requests" class="nav-link" active-class="nav-link-active">
          <Layers class="w-4 h-4" /> Solicitudes
        </router-link>
        <router-link v-if="auth.isAdmin" to="/users" class="nav-link" active-class="nav-link-active">
          <Users class="w-4 h-4" /> Usuarios
        </router-link>
        <router-link v-if="auth.isAdmin" to="/monitoring" class="nav-link" active-class="nav-link-active">
          <Activity class="w-4 h-4" /> Monitoreo
        </router-link>
      </nav>
    </div>

    <div class="flex items-center gap-3 md:gap-6" v-if="auth.isAuthenticated">
      <!-- Theme Toggle -->
      <button 
        @click="toggleDarkMode" 
        class="p-2.5 rounded-xl bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 hover:scale-110 transition-all border border-slate-200 dark:border-slate-700"
      >
        <Sun v-if="isDark" class="w-5 h-5" />
        <Moon v-else class="w-5 h-5" />
      </button>

      <div class="hidden md:flex flex-col items-end">
        <p class="text-sm font-bold text-slate-800 dark:text-slate-200">{{ auth.user?.name }}</p>
        <span class="text-[10px] font-black uppercase tracking-widest text-primary-600 dark:text-primary-400">{{ auth.user?.role }}</span>
      </div>
      
      <button 
        @click="handleLogout"
        class="p-2.5 md:px-4 md:py-2.5 rounded-xl bg-rose-50 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400 font-bold text-sm flex items-center gap-2 hover:bg-rose-600 hover:text-white transition-all border border-rose-100 dark:border-rose-900/30 shadow-sm"
      >
        <LogOut class="w-4 h-4" />
        <span class="hidden md:block">Salir</span>
      </button>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../store/auth';
import { useRouter } from 'vue-router';
import { 
  LayoutDashboard, LogOut, ShieldAlert, Layers, 
  Users, Activity, Moon, Sun 
} from 'lucide-vue-next';

const auth = useAuthStore();
const router = useRouter();
const isDark = ref(false);

const toggleDarkMode = () => {
  isDark.value = !isDark.value;
  if (isDark.value) {
    document.documentElement.classList.add('dark');
    localStorage.setItem('theme', 'dark');
  } else {
    document.documentElement.classList.remove('dark');
    localStorage.setItem('theme', 'light');
  }
};

onMounted(() => {
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    isDark.value = true;
    document.documentElement.classList.add('dark');
  }
});

const handleLogout = () => {
  auth.logout();
  router.push('/login');
};
</script>

<style scoped>
.nav-link {
  @apply flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-bold text-slate-600 dark:text-slate-400 transition-all hover:bg-white dark:hover:bg-slate-700 hover:text-primary-600;
}

.nav-link-active {
  @apply bg-white dark:bg-slate-700 text-primary-600 dark:text-primary-400 shadow-sm scale-105;
}
</style>
