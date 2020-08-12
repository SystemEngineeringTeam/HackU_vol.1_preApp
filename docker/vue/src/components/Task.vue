<template>
  <v-card>
    <v-card-title class="justify-center">
      {{ task.title }}
    </v-card-title>
    <v-card-subtitle>
      {{ task.deadline }}
    </v-card-subtitle>
    <v-card-text>
      <span v-for="(user, i) in task.users" :key="i">
        {{ user }}
      </span>
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn>更新</v-btn>
      <v-btn>削除</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ["task"],

  data: () => ({}),
};
</script>
