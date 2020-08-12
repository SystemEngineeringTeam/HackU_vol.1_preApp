<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <PostTaskForm />
      </v-col>
      <v-col cols="12" v-for="(task, i) in this.$store.state.tasks" :key="i">
        <Task :task="task" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Task from "../components/Task";
import PostTaskForm from "../components/PostTaskForm";

export default {
  name: "ToDo",

  components: {
    Task,
    PostTaskForm,
  },

  data: () => ({}),
  created() {
    this.$store.dispatch("setTasks")
    this.$store.dispatch("setUsers")
  },
};
</script>
