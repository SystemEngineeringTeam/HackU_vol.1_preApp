<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
      <DeadlinePicker :state="this.$store.state.post" :stateStr="stateStr"/>
      <UserSelecter :state="this.$store.state.post" :stateStr="stateStr"/>
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
