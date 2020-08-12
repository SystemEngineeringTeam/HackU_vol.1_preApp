import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    tasks: [],
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
    formFlag: false,
  },
  mutations: {
    setTasks(state, tasks) {
      state.tasks = tasks;
    },
    removeTask(state, index) {
      state.tasks.splice(index, 1);
    },
    updateTask(state, task) {
      let index = state.tasks.findIndex((element) => element.id === task.id);
      state.tasks.splice(index, 1, task);
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
    setFormFlag(state, formFlag) {
      state.formFlag = formFlag;
    },
  },
  actions: {
    async setTasks(context) {
      await axios
        .get(process.env.VUE_APP_URL_TASKS)
        .then((res) => {
          res.data.forEach((task, index, array) => {
            array[index].deadline = task.deadline.substr(0, 16);
            console.log(task.deadline.substr(0, 16));
          })
          context.commit("setTasks", res.data);
        })
    },
    async setUsers(context) {
      await axios
        .get(process.env.VUE_APP_URL_USERS)
        .then((res) => context.commit("setUsers", res.data));
    },
    postAllReset(context) { 
      context.commit("setPostTitle","");
      context.commit("setPostDeadlineDate",null);
      context.commit("setPostDeadlineTime",null);
      context.commit("setPostUser",[]);
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
      await axios.post(process.env.VUE_APP_URL_TASKS, post_json).then((res) => {
        if (res.status == 200) {
          let deadline =
            context.state.post.deadlineDate +
            " " +
            context.state.post.deadlineTime;
          let task = {
            id: res.data.id, //ここのIDはバックエンドから取得する
            title: context.state.post.title,
            deadline: deadline,
            users: context.state.post.users,
          };
          let tasks = context.state.tasks;
          tasks.push(task);
          context.commit("setTasks", tasks);
          context.dispatch("postAllReset")
        }
      });
    },
    updateAllReset(context){
      context.commit("setUpdateTitle","");
      context.commit("setUpdateDeadlineDate",null);
      context.commit("setUpdateDeadlineTime",null);
      context.commit("setUpdateUser",[]);
    },
    async updateTask(context){
      const update = context.state.update;
      let deadline;
      if (update.deadlineDate && update.deadlineTime) {
        deadline = update.deadlineDate + " " + update.deadlineTime + ":00";
      } else if (update.deadlineDate && !update.deadlineTime) {
        deadline = update.deadlineDate + " " + "23:59:59";
      } else if (!update.deadlineDate && update.deadlineTime) {
        deadline =
          new Date().toISOString().substr(0, 10) +
          " " +
          update.deadlineTime +
          ":00";
      } else {
        deadline = null;
      }
      const update_json = {
        title: context.state.update.title,
        deadline: deadline,
        users: context.state.update.users,
      };
      console.log(update_json);
      await axios.put(process.env.VUE_APP_URL_TASKS + context.state.update.id, update_json).then((res) => {
        if (res.status == 200) {
          let deadline =
            context.state.update.deadlineDate +
            " " +
            context.state.update.deadlineTime;
          let task = {
            id: res.data.id, //ここのIDはバックエンドから取得する
            title: context.state.update.title,
            deadline: deadline,
            users: context.state.update.users,
          };
          context.commit("updateTask", task);
          context.dispatch("updateAllReset")
          context.state.formFlag = false;
        }
      });
    },
    async deleteTask(context, id) {
      await axios.delete(process.env.VUE_APP_URL_TASKS + id).then((res) => {
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
