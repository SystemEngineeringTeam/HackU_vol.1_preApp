<template>
  <v-row>
    <v-col cols="12" sm="6">
      <v-select
        v-model="postUsers"
        :items="users"
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
  props: ['state','stateStr'],

  data: () => ({}),
  computed: {
    users(){
      return this.$store.state.users.map(user => user.name);
    },
    postUsers: {
      get() {
        return this.state.users;
      },
      set(value) {
        this.$store.commit("set" + this.stateStr + "User", value);
      }
    }
  }
};
</script>