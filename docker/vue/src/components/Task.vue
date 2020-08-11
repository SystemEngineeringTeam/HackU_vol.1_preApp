<template>
  <v-card>
    <v-card-title class="justify-center">
      {{ task.title }}
    </v-card-title>
    <v-card-subtitle>
      {{ task.deadline }}
    </v-card-subtitle>
    <v-card-text>
      <!-- <span v-for="(userId, i) in task.users" :key="i">
        {{ users.find((user) => user.id == userId).name }}
      </span> -->
      <span v-for="(user, i) in users" :key="i">
        {{ user }}
      </span>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: ["task", "users"],

  data: () => ({}),
};
</script>
