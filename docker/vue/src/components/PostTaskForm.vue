<template>
  <v-card>
    <v-card-text>
      <v-text-field v-model="title" label="title" required></v-text-field>
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="test">
        作成
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    title: "こんにちは",
  }),
  methods:{
    test(){
      console.log(this.title);
      this.axios
      .get("https://api.coindesk.com/v1/bpi/currentprice.json")
      .then((response) => console.log(response));
    }
  }
};
</script>
