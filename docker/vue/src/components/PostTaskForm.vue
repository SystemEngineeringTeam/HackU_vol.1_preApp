<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
      <DeadlinePicker :state="this.$store.state.post" :stateStr="stateStr"/>
      <UserSelecter :state="this.$store.state.post" :stateStr="stateStr"/>
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="postTask">
        作成
      </v-btn>
      <v-btn @click="allReset">
        リセット
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import DeadlinePicker from "../components/DeadlinePicker";
import UserSelecter from "../components/UserSelecter";

export default {
  components: {
    DeadlinePicker,
    UserSelecter
  },
  data: () => ({
    picker: "",
    datePick: false,
    stateStr: "Post"
  }),
  methods: {
    postTask() {
      this.$store.dispatch("postTask");
    },
    allReset(){
      this.$store.dispatch("postAllReset");
    }
  },
  computed: {
    title: {
      get() {
        return this.$store.state.post.title;
      },
      set(value) {
        this.$store.commit("setPostTitle", value);
      },
    }
  }
};
</script>
