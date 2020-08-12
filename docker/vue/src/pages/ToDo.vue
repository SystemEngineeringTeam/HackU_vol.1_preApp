<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <PostTaskForm />
      </v-col>
      <v-col cols="12" v-for="(task, i) in this.$store.state.tasks" :key="i">
        <Task :task="task" :users="findUser(task.users_id)" />
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

  data: () => ({
    sample: [
    ],
    users: [
      { id: 0, name: "けんしん" },
      { id: 1, name: "福田" },
      { id: 2, name: "あいはた" },
      { id: 3, name: "とやま" },
    ],
  }),

  methods: {
    findUser: function(taskUser) {
      let re = [];
      for (var i = 0; i < taskUser.length; i++) {
        re.push(this.users.find((element) => element.id == taskUser[i]).name);
      }
      return re;
    },
  },
  created() {
    this.axios
      .get("http://localhost:80/")
      .then((response) => console.log(response));
  },
};
</script>
