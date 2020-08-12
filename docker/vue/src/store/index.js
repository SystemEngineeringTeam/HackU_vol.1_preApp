import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    tasks: [
    ],
    users: [],
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
    removeTask(state, index) {
      state.tasks.splice(index, 1);
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
    async deleteTask(context, id) {
      await axios
        .delete(process.env.VUE_APP_URL_TASKS + id)
        .then((res) => {
          if (res.status == 200) {
            const index = context.state.tasks.findIndex(
              (element) => element.id === id
            );
            context.commit("removeTask", index);
          }
        });
    },
  },
  modules: {},
});
