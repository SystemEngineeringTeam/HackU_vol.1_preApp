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
    removeTask(state, index) {
      state.tasks.splice(index, 1);
    },
    updateTask(state, task){
      let index = state.tasks.findIndex(element => element.id === task.id);
      state.tasks.splice(index,1,task);
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
