<template>
  <v-card>
    <v-card-title class="justify-center">
      タスク
    </v-card-title>
    <v-card-subtitle>
      8/13 21:00
    </v-card-subtitle>
    <v-card-text>
      <span v-for="(user, i) in users" :key="i">
        {{user}}
      </span>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "Task",

  data: () => ({
    users: ["けんしん","福田","あいはた","とやま"]
  }),
};
</script>
