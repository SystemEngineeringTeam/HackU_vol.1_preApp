<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <PostTaskForm v-if="!this.$store.state.formFlag"/>
        <UpdateTaskForm v-if="this.$store.state.formFlag"/> 
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
import UpdateTaskForm from "../components/UpdateTaskForm";

export default {
  name: "ToDo",

  components: {
    Task,
    PostTaskForm,
    UpdateTaskForm
  },

  data: () => ({}),
  created() {
    this.$store.dispatch("setTasks")
  },
};
</script>
