<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <PostTaskForm />
      </v-col>
      <v-col cols="12" v-for="(task, i) in sample" :key="i">
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
      {
        id: 0,
        title: "醤油を買う",
        deadline: "2020-08-24 09:00",
        users_id: [0, 1],
      },
      {
        id: 1,
        title: "ペットの散歩",
        deadline: "2020-08-24 15:00",
        users_id: [2, 3, 1],
      },
      {
        id: 2,
        title: "スカイダイビング",
        deadline: "2020-08-11 15:00",
        users_id: [1, 2],
      },
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
