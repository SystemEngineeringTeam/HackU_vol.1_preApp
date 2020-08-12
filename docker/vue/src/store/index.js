import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

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
    users: ["秦", "福田", "外山", "相畑"],
    post: {
      title: "",
      deadlineDate: null,
      deadlineTime: null,
      users: [],
    },
    update: {
      id: null,
      title: "",
      deadlineDate: null,
      deadlineTime: null,
      users: [],
    },
    formFlag: false
  },
  mutations: {
    setTasks(state, tasks) {
      state.tasks = tasks;
    },
    removeTask(state, index){
      state.tasks.splice(index,1);
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
    setPostUser(state, users) {
      state.post.users = users;
    },
    setUpdateID(state, id) {
      state.update.id = id;
    },
    setUpdateTitle(state, title) {
      state.update.title = title;
    },
    setUpdateDeadlineDate(state, date) {
      state.update.deadlineDate = date;
    },
    setUpdateDeadlineTime(state, time) {
      state.update.deadlineTime = time;
    },
    setUpdateUser(state, users) {
      state.update.users = users;
    },
    setFormFlag(state, formFlag){
      state.formFlag = formFlag;
    },
  },
  actions: {
    async setTasks(context) {
      await axios
        .get(process.env.VUE_APP_URL_TASKS)
        .then((res) => context.commit("setTasks", res.data));
    },
    async setUsers(context) {
      await axios
        .get(process.env.VUE_APP_URL_USERS)
        .then((res) => context.commit("setUsers", res.data));
    },
  },
  modules: {},
});
