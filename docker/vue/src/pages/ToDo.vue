<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
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
        name: "string",
        deadline: "string",
        users: ["string"],
      },
      {
        id: 1,
        name: "string",
        deadline: "string",
        users: ["string"],
      },
      {
        id: 2,
        name: "string",
        deadline: "string",
        users: ["string"],
      },
    ],
  }),
};
</script>
