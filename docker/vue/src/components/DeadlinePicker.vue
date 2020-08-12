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
            prepend-icon="event"
            readonly
            v-bind="attrs"
            v-on="on"
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
            prepend-icon="access_time"
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
  </v-row>
</template>

<script>
export default {
  //props: ['time'],
  data() {
    return {
      date: new Date().toISOString().substr(0, 10),
      datePick: false,
      time: null,
      timePick: false,
    };
  },
  methods: {},
};
</script>
