<template>
  <div class="absolute -top-2 left-0 right-0 z-50 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 pointer-events-none">
    <div class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 w-72 p-4 bg-slate-900 text-white rounded-xl shadow-2xl">
      <p class="font-black uppercase tracking-widest text-[10px] text-primary-400 mb-1">{{ title }}</p>
      <p v-if="iso" class="text-[8px] font-black text-slate-500 uppercase tracking-tighter mb-1">{{ iso }}</p>
      <p class="text-slate-300 text-[10px] leading-relaxed font-bold">{{ description }}</p>
      <div class="absolute top-full left-1/2 -translate-x-1/2 border-8 border-transparent border-t-slate-900"></div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: String,
  iso: String,
  description: String
});
</script>
