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
      <v-btn @click="updateTask">更新</v-btn>
      <v-btn @click="deleteTask(task.id)">削除</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ["task"],

  data: () => ({}),

  methods: {
    updateTask: function(){
      this.$store.commit("setUpdateID",this.task.id);
      this.$store.commit("setUpdateTitle",this.task.title);
      let str = this.task.deadline.split(' ');
      this.$store.commit("setUpdateDeadlineDate",str[0]);
      this.$store.commit("setUpdateDeadlineTime",str[1]);
      this.$store.commit("setUpdateUser",this.task.users);
      this.$store.commit("setFormFlag",true);
    },
    deleteTask: function(id){
      this.$store.dispatch("deleteTask", id)
    }
  },
};
</script>
