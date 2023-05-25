<template>
    <div class="h-full w-full bg-zinc-700 rounded-3xl p-6">
        <div class="h-full flex flex-col gap-2">
            <div class="flex flex-row gap-2">
                <h1 class="font-bold">From:</h1>
                <p>{{ mail.from }}</p>
            </div>
            <div class="flex flex-row gap-2">
                <h1 class="font-bold">To:</h1>
                <p>{{ mail.to }}</p>
            </div>
            <div class="flex flex-row gap-2">
                <h1 class="font-bold">Category:</h1>
                <p>{{ mail.category }}</p>
            </div>
            <div class="flex flex-row gap-2">
                <h1 class="font-bold">Subject:</h1>
                <p>{{ mail.subject }}</p>
            </div>
            <h1 class="overflow-y-auto whitespace-pre-line">{{ mail.body }}</h1>
        </div>
    </div>
</template>

<script>
import store from '../../store/mails';

export default {
    name: 'MailView',
    data() {
        return {
            mail: store.state.mail
        }
    }
}
</script>