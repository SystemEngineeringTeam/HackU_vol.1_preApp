<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
      <DeadlinePicker :state="this.$store.state.update" :stateStr="stateStr" />
      <UserSelecter :state="this.$store.state.update" :stateStr="stateStr" />
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="doUpdate">更新</v-btn>
      <v-btn @click="allReset">リセット</v-btn>
      <v-btn @click="cancel">キャンセル</v-btn>
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
    stateStr: "Update"
  }),
  methods: {
    // test() {
    //   console.log(this.title);
    //   this.axios
    //     .get("https://api.coindesk.com/v1/bpi/currentprice.json")
    //     .then(response => console.log(response));
    // },
    allReset() {
        this.$store.dispatch("updateAllReset");
    },
    cancel() {
      this.$store.dispatch("updateAllReset");
      this.$store.commit("setFormFlag", false);
    },
    doUpdate() {
      this.$store.dispatch("updateTask");
    }
  },
  computed: {
    title: {
      get() {
        return this.$store.state.update.title;
      },
      set(value) {
        this.$store.commit("setUpdateTitle", value);
      }
    }
  }
};
</script>
