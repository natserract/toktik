import { defineCustomElement as VueDefineCustomElement, h, createApp, getCurrentInstance, Component } from 'vue'

export const defineCustomRootElement = (component: Component, { plugins = [] as any[] } = {}) =>
  VueDefineCustomElement({
    render: () => h(component),
    setup() {
      const app = createApp(component)

      plugins.forEach(app.use)

      const inst = getCurrentInstance()
      if (inst) {
        Object.assign(inst.appContext, app._context)
        Object.assign(inst.appContext.provides, app._context.provides)
      }
    },
  })
