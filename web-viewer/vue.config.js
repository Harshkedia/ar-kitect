module.exports = {
  publicPath: "/",
  outputDir: "./dist/",
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "~@/sass/main.scss";`
      }
    }
  },
  chainWebpack: config => {
    ["vue-modules", "vue", "normal-modules", "normal"].forEach(match => {
      config.module
        .rule("sass")
        .oneOf(match)
        .use("sass-loader")
        .tap(opt => Object.assign(opt, { prependData: `@import '~@/sass/main.scss'` }));
    });

    config.module
      .rule("raw")
      .test(/\.obj$/)
      .use("raw-loader")
      .loader("raw-loader")
      .end();
  },

  configureWebpack: {
    stats: {
      // warnings: false
    },
    entry: {},
    optimization: {
      runtimeChunk: "single",
      splitChunks: {
        chunks: "all",
        maxInitialRequests: Infinity,
        minSize: 0,
        cacheGroups: {
          vendor: {
            test: /[\\/]node_modules[\\/]/,
            name(module) {
              const packageName = module.context.match(/[\\/]node_modules[\\/](.*?)([\\/]|$)/)[1];
              return `npm.${packageName.replace("@", "")}`;
            }
          }
        }
      }
    },
    devtool: false
  }
};
