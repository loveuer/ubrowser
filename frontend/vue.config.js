// const path = require('path');

module.exports = {
    // Project deployment base
    // By default we assume your app will be deployed at the root of a domain,
    // e.g. https://www.my-app.com/
    // If your app is deployed at a sub-path, you will need to specify that
    // sub-path here. For example, if your app is deployed at
    // https://www.foobar.com/my-app/
    // then change this to '/my-app/'
    publicPath: '/',

    // where to output built files
    // 打包后的输出目录
    outputDir: 'dist',

    // A directory (relative to outputDir) to nest generated static assets (js, css, img, fonts) under.
    assetsDir: 'static',

    // whether to use eslint-loader for lint on save.
    // 保存时是不是用eslint-loader 来lint 代码
    lintOnSave: false,

    // use the full build with in-browser compiler?
    // https://vuejs.org/v2/guide/installation.html#Runtime-Compiler-vs-Runtime-only
    // 使用runtime-only 还是 in-browser compiller
    // compiler: true,

    // tweak internal webpack configuration.
    // see https://github.com/vuejs/vue-cli/blob/dev/docs/webpack.md
    // webpack 配置~
    chainWebpack: () => {},
    configureWebpack: {
        optimization: {
            splitChunks: false
        },
        output: {
            filename: "static/js/desktop_[name].js",
            chunkFilename: 'static/js/desktop_[name].js',
        },
    },
    // vue-loader options
    // https://vue-loader.vuejs.org/en/options.html
    // vue-loader 配置
    // vueLoader: {},

    // generate sourceMap for production build?
    // 生产环境的sourceMap 要不要？
    // 就是生成.js.map文件,用来定位js哪里出错了的
    productionSourceMap: false,

    // CSS related options
    css: {
        // extract CSS in components into a single CSS file (only in production)
        extract: false,

        // enable CSS source maps?
        sourceMap: false,

        // pass custom options to pre-processor loaders. e.g. to pass options to
        // sass-loader, use { sass: { ... } }
        loaderOptions: {},

        // Enable CSS modules for all css / pre-processor files.
        // This option does not affect *.vue files.
        // 用不用 css Modules 啊？
        modules: false
    },

    // use thread-loader for babel & TS in production build
    // enabled by default if the machine has more than 1 cores
    // 使用多线程否？
    parallel: require('os').cpus().length > 1,

    // split vendors using autoDLLPlugin?
    // can also be an explicit Array of dependencies to include in the DLL chunk.
    // See https://github.com/vuejs/vue-cli/blob/dev/docs/cli-service.md#dll-mode
    // 用不用 autoDLLPlugin，厉害了
    // dll: false,

    // options for the PWA plugin.
    // see https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-pwa
    // pwa 相关
    pwa: {},

    // configure webpack-dev-server behavior
    // Webpack dev server
    devServer: {
        open: false,
        host: '0.0.0.0',
        port: 9119,
        https: false,
        hotOnly: false,
        // proxy: 'http://127.0.0.1:8000/', // 配置跨域处理,只有一个代理
        proxy: {
            '/api': {
                // target: 'http://192.168.31.99:8111',
                target: 'http://localhost:8111',
                // target: '<url>',
                ws: true,
                changeOrigin: true
            },
            '/static': {
                target: 'http://localhost:8111',
            }
        }, // 配置多个代理
        // before: app => {}
    }
}
