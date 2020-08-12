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
    async postTask(context) {
      const post = context.state.post;
      let deadline;
      if (post.deadlineDate && post.deadlineTime) {
        deadline = post.deadlineDate + " " + post.deadlineTime + ":00";
      } else if (post.deadlineDate && !post.deadlineTime) {
        deadline = post.deadlineDate + " " + "23:59:59";
      } else if (!post.deadlineDate && post.deadlineTime) {
        deadline =
          new Date().toISOString().substr(0, 10) +
          " " +
          post.deadlineTime +
          ":00";
      } else {
        deadline = null;
      }
      const post_json = {
        title: context.state.post.title,
        deadline: deadline,
        users: context.state.post.users,
      };
      console.log(post_json);
      await axios
        .post(process.env.VUE_APP_URL_TASKS, post_json)
        .then((res) => context.commit("setUsers", res.data));
    },
  },
  modules: {},
});
