<template>
  <v-row>
    <v-col cols="5">
      <v-menu
        v-model="datePick"
        :close-on-content-click="false"
        :nudge-right="40"
        transition="scale-transition"
        offset-y
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date"
            label="Picker without buttons"
            readonly
            v-bind="attrs"
            v-on="on"
            @click="dateInitial"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="date"
          @input="datePick = false"
          no-title
          scrollable
        ></v-date-picker>
      </v-menu>
    </v-col>
    <v-col cols="5">
      <v-menu
        ref="menu"
        v-model="timePick"
        :close-on-content-click="false"
        :nudge-right="40"
        :return-value.sync="time"
        transition="scale-transition"
        offset-y
        max-width="290px"
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="time"
            label="Picker in menu"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-time-picker
          v-if="timePick"
          v-model="time"
          full-width
          @click:minute="$refs.menu.save(time)"
        ></v-time-picker>
      </v-menu>
    </v-col>
    <v-col cols="5">
      <v-btn @click="resetDeadline">リセット</v-btn>
    </v-col>
  </v-row>
</template>

<script>
export default {
  //props: ['time','date'],

  data() {
    return {
      //date: new Date().toISOString().substr(0, 10),
      datePick: false,
      time: null,
      timePick: false,
    };
  },
  computed: {
    date: {
      get() {
        return this.$store.state.post.deadlineDate;
      },
      set(value) {
        this.$store.commit("setPostDeadlineDate", value);
      },
    },
  },
  methods: {
    resetDeadline: function() {
      this.date = null;
      this.time = null;
    },

    dateInitial: function() {
      if (this.date === null) {
        this.date = new Date().toISOString().substr(0, 10);
      }
    },
  },
};
</script>
