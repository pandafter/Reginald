<template>
  <div class="min-h-screen flex flex-col items-center justify-center p-6 relative overflow-hidden transition-colors duration-500 font-inter">
    <!-- Theme Toggle -->
    <div class="absolute top-6 right-6 z-20">
      <button 
        @click="toggleDarkMode" 
        class="p-3 rounded-2xl bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 hover:scale-110 transition-all border border-slate-200 dark:border-slate-700 shadow-sm"
      >
        <Sun v-if="isDark" class="w-5 h-5" />
        <Moon v-else class="w-5 h-5" />
      </button>
    </div>

    <!-- Animated background blobs -->
    <div class="absolute -top-24 -left-24 w-96 h-96 bg-primary-200/50 dark:bg-primary-900/20 rounded-full blur-3xl animate-pulse"></div>
    <div class="absolute -bottom-24 -right-24 w-96 h-96 bg-indigo-200/50 dark:bg-indigo-900/20 rounded-full blur-3xl animate-pulse delay-700"></div>

    <div class="glass w-full max-w-lg p-8 md:p-12 rounded-[2.5rem] relative z-10 shadow-2xl animate-in fade-in zoom-in duration-500 border border-slate-200/50 dark:border-slate-800">
      <div class="text-center mb-8 md:mb-10">
        <h1 class="text-3xl md:text-4xl font-black bg-gradient-to-r from-primary-600 to-indigo-600 dark:from-primary-400 dark:to-indigo-400 bg-clip-text text-transparent tracking-tighter mb-2">REGINALD</h1>
        <p class="text-slate-500 dark:text-slate-400 font-bold uppercase tracking-widest text-[10px]">Crea tu perfil de operación</p>
      </div>

      <form @submit.prevent="handleRegister" class="space-y-4 md:space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="space-y-2">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Nombre</label>
            <div class="relative group">
              <User class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-primary-600 transition-colors w-5 h-5" />
              <input 
                v-model="form.name"
                type="text" 
                required
                class="w-full pl-12 pr-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl outline-none focus:ring-2 focus:ring-primary-500 transition-all dark:text-white"
                placeholder="Nombre"
              />
            </div>
          </div>
          <div class="space-y-2">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Email</label>
            <div class="relative group">
              <Mail class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-primary-600 transition-colors w-5 h-5" />
              <input 
                v-model="form.email"
                type="email" 
                required
                class="w-full pl-12 pr-4 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl outline-none focus:ring-2 focus:ring-primary-500 transition-all dark:text-white"
                placeholder="ID Corporativo"
              />
            </div>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Contraseña de Seguridad</label>
          <div class="relative group">
            <Lock class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-primary-600 transition-colors w-5 h-5" />
            <input 
              v-model="form.password"
              :type="showPassword ? 'text' : 'password'"
              required
              class="w-full pl-12 pr-12 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl outline-none focus:ring-2 focus:ring-primary-500 transition-all dark:text-white"
              placeholder="••••••••"
            />
            <button 
              type="button"
              @click="showPassword = !showPassword"
              class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors"
            >
              <component :is="showPassword ? EyeOff : Eye" class="w-5 h-5" />
            </button>
          </div>
          
          <div class="px-2 pt-2 space-y-2">
            <div class="flex gap-1.5">
              <div v-for="i in 4" :key="i" 
                   class="h-1 flex-1 rounded-full transition-all duration-500"
                   :class="strength >= i ? strengthColor : 'bg-slate-200 dark:bg-slate-800 border border-slate-200/50 dark:border-transparent'">
              </div>
            </div>
            <div class="flex justify-between items-center">
              <p class="text-[9px] font-black uppercase tracking-widest transition-colors duration-500" :class="strengthTextClass">
                Nivel: {{ strengthLabel }}
              </p>
            </div>
            <ul class="grid grid-cols-2 gap-x-4 gap-y-1.5 mt-2">
              <li v-for="req in passwordRequirements" :key="req.label" 
                  class="flex items-center gap-1.5 text-[9px] font-bold transition-colors"
                  :class="req.valid ? 'text-emerald-600 dark:text-emerald-400' : 'text-slate-400 dark:text-slate-600'">
                <CheckCircle2 v-if="req.valid" class="w-2.5 h-2.5" />
                <Circle v-else class="w-2.5 h-2.5" />
                {{ req.label }}
              </li>
            </ul>
          </div>
        </div>

        <button 
          type="submit" 
          :disabled="loading || strength < 3"
          class="w-full bg-slate-900 dark:bg-primary-600 text-white font-black py-4 rounded-2xl shadow-xl shadow-slate-200 dark:shadow-none hover:bg-black dark:hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all active:scale-[0.98] mt-4 flex items-center justify-center gap-2 group uppercase tracking-widest text-xs"
        >
          <Loader2 v-if="loading" class="w-5 h-5 animate-spin" />
          <span v-else>Generar Cuenta</span>
          <ArrowRight v-if="!loading" class="w-4 h-4 group-hover:translate-x-1 transition-transform" />
        </button>

        <div v-if="error" class="bg-rose-50 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400 p-4 rounded-2xl text-[11px] font-bold border border-rose-100 dark:border-rose-900/30">
          {{ error }}
        </div>

        <p class="text-center text-slate-500 dark:text-slate-400 text-xs mt-8 font-bold">
          ¿Ya eres usuario? 
          <router-link to="/login" class="text-primary-600 dark:text-primary-400 font-black hover:underline">Accede aquí</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api/axios';
import { 
  User, Mail, Lock, Eye, EyeOff, CheckCircle2, Circle, 
  ArrowRight, Loader2, Sun, Moon 
} from 'lucide-vue-next';

const router = useRouter();
const loading = ref(false);
const error = ref(null);
const showPassword = ref(false);
const isDark = ref(false);

const form = reactive({
  name: '',
  email: '',
  password: ''
});

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

const passwordRequirements = computed(() => [
  { label: '8+ chars', valid: form.password.length >= 8 },
  { label: 'Mayúscula', valid: /[A-Z]/.test(form.password) },
  { label: 'Número', valid: /[0-9]/.test(form.password) },
  { label: 'Especial', valid: /[^A-Za-z0-9]/.test(form.password) }
]);

const strength = computed(() => {
  return passwordRequirements.value.filter(req => req.valid).length;
});

const strengthLabel = computed(() => {
  if (form.password.length === 0) return '-';
  if (strength.value === 1) return 'CRÍTICO';
  if (strength.value === 2) return 'DÉBIL';
  if (strength.value === 3) return 'SEGURO';
  if (strength.value === 4) return 'ÓPTIMO';
  return '';
});

const strengthColor = computed(() => {
  if (strength.value <= 1) return 'bg-rose-500';
  if (strength.value === 2) return 'bg-amber-500';
  if (strength.value === 3) return 'bg-blue-500';
  return 'bg-emerald-500';
});

const strengthTextClass = computed(() => {
  if (strength.value <= 1) return 'text-rose-500';
  if (strength.value === 2) return 'text-amber-500';
  if (strength.value === 3) return 'text-blue-500';
  return 'text-emerald-500';
});

const handleRegister = async () => {
  if (strength.value < 3) {
    error.value = 'Contraseña no cumple requisitos de seguridad.';
    return;
  }
  
  loading.value = true;
  error.value = null;
  
  try {
    await api.post('/users/register', form);
    router.push({ path: '/login', query: { registered: 'true' } });
  } catch (err) {
    error.value = err.response?.data || 'Error crítico en el registro.';
  } finally {
    loading.value = false;
  }
};
</script>
