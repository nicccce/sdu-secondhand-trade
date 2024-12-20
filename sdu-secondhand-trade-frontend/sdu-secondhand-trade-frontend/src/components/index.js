import ImageView from './ImageView/index.vue'

export const componentPlugin = {
    install(app) {
        app.component('ImageView', ImageView)
    }
}