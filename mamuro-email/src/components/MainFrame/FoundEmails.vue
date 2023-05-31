<template>
    <div class="h-full w-full flex flex-col">
      <div class="overflow-y-auto rounded-3xl">
        <table class="table-fixed w-full border-collapse">
        <thead>
          <tr class="sticky top-0 h-8 divide-x divide-zinc-500 bg-zinc-700">
            <th>Subject</th>
            <th>From</th>
            <th>To</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-zinc-500 text-sm">
          <tr
            v-for="(mail, index) in mails"
            :key="mail.id"
            :class="rowClass(index)"
            @click="changeView(mail.id)"
            class="h-8 cursor-pointer"
          >
            <td class="truncate px-4">{{ mail.subject }}</td>
            <td class="truncate px-4">{{ mail.from }}</td>
            <td class="truncate px-4">{{ mail.to }}</td>
          </tr>
        </tbody>
      </table>
      </div>
      <div class="my-4 flex">
        <div class="flex flex-row w-16 p-2 bg-zinc-500 rounded-lg justify-center">
          <button @click="changePage(-1)" class="flex border-r border-zinc-300">
            <span class="material-symbols-outlined self-center">navigate_before</span>
          </button>
          <button @click="changePage(1)" class="flex">
            <span class="material-symbols-outlined self-center">navigate_next</span>
          </button>
        </div>
          <h1 class="mx-4 self-center">Page {{ actual_page }} of {{ total_pages }} - {{ total_mails }} Total Mails</h1>
      </div>
    </div>
  </template>
  
  <script>
  import store from '../../store/mails';
  
  export default {
    name: 'FoundEmails',
    methods: {
      changeView(id) {
        store.dispatch('changeMailView', this.mails[id]);
      },
      rowClass(index) {
        return index % 2 === 0 ? 'bg-zinc-900' : 'bg-zinc-950';
      },
      changePage(page) {
        store.dispatch('changeActualPage', page);
      },
    },
    computed: {
      mails() {
        return store.state.mails;
      },
      total_mails() {
        return store.state.total_mails;
      },
      actual_page() {
        return store.state.actual_page;
      },
      total_pages() {
        return store.state.total_pages;
      },
    },
  };
  </script>
  