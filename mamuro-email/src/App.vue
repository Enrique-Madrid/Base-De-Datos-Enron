<script setup>
import FoundEmails from './components/MainFrame/FoundEmails.vue';
import MailView from './components/MainFrame/MailView.vue';
import Bar from './components/SideBar/Bar.vue';
</script>

<template>
  <div class="h-screen w-screen bg-zinc-900 text-white">
    <div class="h-full flex flex-col">
      <Bar></Bar>
      <div class="h-full overflow-hidden flex flex-row m-8 gap-4 ">
        <div class="w-2/3">
          <FoundEmails></FoundEmails>
        </div>
        <div class="w-1/3">
          <MailView></MailView>
        </div>
      </div>
    </div>
  </div>
</template>
