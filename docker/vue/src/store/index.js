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
        users: ["秦", "福田"],
      },
      {
        id: 1,
        title: "ペットの散歩",
        deadline: "2020-08-24 15:00",
        users: ["福田", "相畑", "外山"],
      },
      {
        id: 2,
        title: "スカイダイビング",
        deadline: "2020-08-11 15:00",
        users: ["相畑"],
      },
    ],
    users: [
      "秦",
      "福田",
      "外山",
      "相畑"
    ],
    userBool: [],
    post: {
      title: "",
      deadlineDate: null,
      deadlineTime: null,
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
    setUserBool(state, userBool){
      state.post.userBool = userBool;
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
