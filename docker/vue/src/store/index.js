import Vue from "vue";
import Vuex from "vuex";
import axios from 'axios';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    tasks: [
      {
        id: 0,
        title: "醤油をかける",
        deadline: "2020-08-24 09:00",
        users_id: [0, 1],
      },
      {
        id: 1,
        title: "ペットの散歩",
        deadline: "2020-08-24 15:00",
        users_id: [2, 3, 1],
      },
      {
        id: 2,
        title: "スカイダイビング",
        deadline: "2020-08-11 15:00",
        users_id: [1, 2],
      },
    ],
    users: [],
    post: {
      title: "",
      deadlineDate: "",
      deadlineTime: "",
      users: [],
    },
  },
  mutations: {
    setTasks(state, tasks) {
      state.tasks = tasks;
    },
    setUsers(state, users) {
      state.users = users;
    },
    setPostTitle(state, title) {
      state.post.title = title;
    },
    setPostDeadlineDate(state, date) {
      state.post.deadlineDate = date;
    },
    setPostDeadlineTime(state, time) {
      state.post.deadlineTime = time;
    },
    pushPostUsers(state, user) {
      state.post.users.push(user);
    },
    splicePostUser(state, index) {
      state.post.users.slice(index, 1);
    },
  },
  actions: {
    async setTasks(context) {
      await axios
        .get("http://localhost:3000")
        .then((res) => context.setTasks(res.data));
    },
    async setUsers(context) {
      await axios
        .get("http://localhost:3000")
        .then((res) => context.setUsers(res.data));
    },
  },
  modules: {},
});
