<template>
    <div class="h-full w-full overflow-y-auto rounded-3xl">
      <table class="table-fixed w-full border-collapse">
        <thead>
          <tr class="sticky top-0 h-12 divide-x divide-zinc-500 bg-zinc-700">
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
            class="h-10"
          >
            <td class="truncate px-4">{{ mail.subject }}</td>
            <td class="truncate px-4">{{ mail.from }}</td>
            <td class="truncate px-4">{{ mail.to }}</td>
          </tr>
        </tbody>
      </table>
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
    },
    computed: {
      mails() {
        return store.state.mails;
      },
    },
  };
  </script>
  