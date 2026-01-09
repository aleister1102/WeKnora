import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import "./assets/fonts.css";
import TDesign from "tdesign-vue-next";
// 引入组件库的少量全局样式变量
import "tdesign-vue-next/es/style/index.css";
import "@/assets/theme/theme.css";
import i18n from "./i18n";

const app = createApp(App);

// Global error handler to catch and log errors without blocking UI
app.config.errorHandler = (err, instance, info) => {
  console.error('Vue Error:', err);
  console.error('Component:', instance);
  console.error('Info:', info);
  // Don't re-throw - allow the app to continue
};

// Global warning handler (development only)
app.config.warnHandler = (msg, instance, trace) => {
  console.warn('Vue Warning:', msg);
  console.warn('Trace:', trace);
};

app.use(TDesign);
app.use(createPinia());
app.use(router);
app.use(i18n);

app.mount("#app");
