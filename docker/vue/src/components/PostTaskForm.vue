<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
      <DeadlinePicker />
      <UserSelecter />
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="test">
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
  }),
  methods: {
    test() {
      console.log(this.title);
      this.axios
        .get("https://api.coindesk.com/v1/bpi/currentprice.json")
        .then((response) => console.log(response));
    },
    allReset() { 
      this.$store.state.post.title = "";
      this.$store.state.post.deadlineDate = null;
      this.$store.state.post.deadlineTime = null;
      this.$store.state.post.users = [];
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
