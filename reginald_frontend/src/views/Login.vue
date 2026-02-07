<template>
  <div class="min-h-screen flex flex-col items-center justify-center p-6 relative overflow-hidden transition-colors duration-500">
    <!-- Theme Toggle at top right -->
    <div class="absolute top-6 right-6 z-20">
      <button 
        @click="toggleDarkMode" 
        class="p-3 rounded-2xl bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 hover:scale-110 transition-all border border-slate-200 dark:border-slate-700 shadow-sm"
      >
        <Sun v-if="isDark" class="w-5 h-5" />
        <Moon v-else class="w-5 h-5" />
      </button>
    </div>

    <!-- Abstract Background -->
    <div class="absolute -top-24 -left-24 w-96 h-96 bg-primary-200 dark:bg-primary-900/20 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"></div>
    <div class="absolute -bottom-24 -right-24 w-96 h-96 bg-blue-200 dark:bg-indigo-900/20 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob animation-delay-2000"></div>

    <div class="glass max-w-md w-full p-8 md:p-10 rounded-[2.5rem] relative z-10 shadow-2xl transition-all border border-slate-200/50 dark:border-slate-800">
      <div class="flex items-center gap-3 text-center justify-center mb-10" >
        <img src="/LogoBlanco.svg" alt="Reginald OS" class="h-8 md:h-10 hidden dark:block">
        <img src="/LogoAzul.svg" alt="Reginald OS" class="h-8 md:h-10 block dark:hidden">
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div v-if="route.query.registered" class="bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400 p-4 rounded-2xl text-[11px] font-black uppercase tracking-widest border border-emerald-100 dark:border-emerald-900/30 flex items-center gap-3">
          <CheckCircle2 class="w-5 h-5 flex-shrink-0" />
          <span>¡Cuenta activada! Ya puedes entrar.</span>
        </div>

        <div>
          <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2 ml-1">Email Corporativo</label>
          <div class="relative group">
            <Mail class="absolute left-4 top-3.5 w-5 h-5 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
            <input 
              v-model="email"
              type="email" 
              required
              class="w-full pl-12 pr-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl focus:ring-2 focus:ring-primary-500 focus:border-transparent outline-none transition-all dark:text-white"
              placeholder="user@reginald.com"
            />
          </div>
        </div>

        <div>
          <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2 ml-1">Contraseña</label>
          <div class="relative group">
            <Lock class="absolute left-4 top-3.5 w-5 h-5 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
            <input 
              v-model="password"
              type="password" 
              required
              class="w-full pl-12 pr-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl focus:ring-2 focus:ring-primary-500 focus:border-transparent outline-none transition-all dark:text-white"
              placeholder="••••••••"
            />
          </div>
        </div>

        <div v-if="auth.error" class="bg-rose-50 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400 p-4 rounded-2xl text-xs font-bold flex items-center gap-3 border border-rose-100 dark:border-rose-900/30">
          <AlertCircle class="w-5 h-5 flex-shrink-0" />
          <span>{{ auth.error }}</span>
        </div>

        <button 
          type="submit"
          :disabled="auth.loading"
          class="w-full bg-slate-900 dark:bg-primary-600 hover:bg-black dark:hover:bg-primary-700 disabled:opacity-50 text-white font-black py-4 rounded-2xl shadow-xl shadow-slate-200 dark:shadow-none transition-all active:scale-95 flex items-center justify-center gap-2 uppercase tracking-widest text-sm"
        >
          <Loader2 v-if="auth.loading" class="w-5 h-5 animate-spin" />
          <span v-else>Entrar al Sistema</span>
        </button>

        <p class="text-center text-slate-500 dark:text-slate-400 text-xs mt-4 font-bold">
          ¿Sin credenciales? 
          <router-link to="/register" class="text-primary-600 dark:text-primary-400 font-black hover:underline cursor-pointer">Crea un perfil</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../store/auth';
import { useRouter, useRoute } from 'vue-router';
import { ShieldCheck, Mail, Lock, AlertCircle, Loader2, CheckCircle2, Sun, Moon } from 'lucide-vue-next';

const email = ref('');
const password = ref('');
const auth = useAuthStore();
const router = useRouter();
const route = useRoute();
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

const handleSubmit = async () => {
  const success = await auth.login(email.value, password.value);
  if (success) {
    router.push('/');
  }
};
</script>
