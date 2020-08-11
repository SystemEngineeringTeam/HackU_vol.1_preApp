<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12" v-for="(task,i) in sample" v-bind:key="i">
        <Task />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Task from "../components/Task";

export default {
  name: "ToDo",

  components: {
    Task,
  },

  data: () => ({
    sample: [
      {
        id: 0,
        title: "醤油を買う",
        deadline: "2020-08-24 09:00",
        users: [0, 1],
      },
      {
        id: 1,
        title: "ペットの散歩",
        deadline: "2020-08-24 15:00",
        users: [2, 3],
      },
      {
        id: 2,
        title: "スカイダイビング",
        deadline: "2020-08-11 15:00",
        users: [1],
      },
    ],
    users: [{ id: 0, name: "けんしん" }, { id: 1, name: "福田"}, { id: 2, name:"あいはた"}, { id: 3, name:"とやま"}],
  }),
};
</script>
