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
    title: "こんにちは",
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
  },
};
</script>
