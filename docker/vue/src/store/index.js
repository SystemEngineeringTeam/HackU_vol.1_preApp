import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    tasks: [],
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
    splicePostUsers(state, index) {
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
