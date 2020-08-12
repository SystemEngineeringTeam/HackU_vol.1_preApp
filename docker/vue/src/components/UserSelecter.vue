<template>
  <v-row>
    <v-col cols="12" sm="6">
      <v-select
        v-model="users"
        :items="$store.state.users"
        :menu-props="{ maxHeight: '400' }"
        label="Select"
        multiple
        hint="Pick your favorite states"
        persistent-hint
      ></v-select>
    </v-col>
  </v-row>
</template>

<script>
export default {
  data: () => ({}),
  computed: {
    users: {
      get() {
        return this.$store.state.post.users;
      },
      set(value) {
        this.$store.commit("setPostUser", value);
      }
    }
  }
};
</script>