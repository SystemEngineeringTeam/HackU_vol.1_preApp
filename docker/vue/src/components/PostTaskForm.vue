<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
      <DeadlinePicker :state="this.$store.state.post" :stateStr="stateStr"/>
      <UserSelecter :state="this.$store.state.post" :stateStr="stateStr"/>
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="post">
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
    test() {
      console.log(this.title);
      this.axios
        .get("https://api.coindesk.com/v1/bpi/currentprice.json")
        .then((response) => console.log(response));
    },
    allReset() { 
      this.$store.commit("setPostTitle","");
      this.$store.commit("setPostDeadlineDate",null);
      this.$store.commit("setPostDeadlineTime",null);
      this.$store.commit("setPostUser",[]);
    },
    post() {
      //このメソッドが実行された時点でAPIを用いた通信をしなければならない
      let deadline = this.$store.state.post.deadlineDate + " " + this.$store.state.post.deadlineTime
      let task = {
        id: 12345,//ここのIDは後ほどちゃんと実装する
        title: this.$store.state.post.title,
        deadline: deadline,
        users: this.$store.state.post.users
      };
      let tasks = this.$store.state.tasks;
      tasks.push(task);
      this.$store.commit("setTasks",tasks);
      this.allReset();
      // console.log(this.$store.state.tasks);
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
